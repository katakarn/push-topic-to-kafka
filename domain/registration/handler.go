package registration

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"net/http"
	"strconv"
	"testKafka/httputil"
	"testKafka/producer"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type insertRecordFunc func(coll CollectionName, data *SubmissionData) error

func (ir insertRecordFunc) InsertRecord(coll CollectionName, data *SubmissionData) error {
	return ir(coll, data)
}

type getLatestFunc func(int, string, int, int) (SubmissionData, int64, error)

func (gl getLatestFunc) GetLatest(studentID int, collectionName string, acadYear int, semester int) (SubmissionData, int64, error) {
	return gl(studentID, collectionName, acadYear, semester)
}

type submitFunc func(ConfirmRequest, string) (ConfirmResponse, error)

func (sf submitFunc) Submit(requestBody ConfirmRequest, token string) (ConfirmResponse, error) {
	return sf(requestBody, token)
}

type getEnrollSemesterFunc func(context.Context, int) (*Semester, error)

func (fn getEnrollSemesterFunc) getEnrollSemester(ctx context.Context, scheduleGroupID int) (*Semester, error) {
	return fn(ctx, scheduleGroupID)
}

type getTokenFunc func(TokenRequest) (TokenResponse, error)

func (gt getTokenFunc) getToken(request TokenRequest) (TokenResponse, error) {
	return gt(request)
}

// SubmitRegistrationHandler godoc
//
//	@tags			Register
//	@summary		Submit registration student
//	@description	Submit registration student
//	@id				Submit registration student
//	@security		Bearer
//	@produce		json
//	@param			RequestBody	body						ConfirmRequest	true	"Submit registration student"
//	@success		200			httputil.SuccessResponse	200				"Registration submitted successfully."
//	@failure		400			{object}					httputil.HTTPBadRequestErrors
//	@failure		500			{object}					httputil.HTTPError
//	@router			/submit [post]
func SubmitRegistrationHandler(p *producer.Producer, insert insertRecordFunc, ges getEnrollSemesterFunc, gl getLatestFunc, gt getTokenFunc, sm submitFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		studentIDValue, exist := ctx.Get("studentId")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "Missing required parameter: studentID",
			})
			return
		}
		studentCodeValue, exist := ctx.Get("studentCode")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "studentCode not found",
			})
			return
		}
		userIdValue, exist := ctx.Get("userId")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "userId not found",
			})
			return
		}
		emailValue, exist := ctx.Get("email")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "email not found",
			})
			return
		}
		subValue, exist := ctx.Get("sub")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "sub not found",
			})
			return
		}
		studentCode := studentCodeValue.(string)
		studentID := studentIDValue.(int)
		userId := userIdValue.(string)
		email := emailValue.(string)
		sub := subValue.(string)

		// Validate JSON request body
		var request ConfirmRequest
		err := ctx.ShouldBindJSON(&request)

		if err != nil {
			var errorString string
			var fieldRequired string
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
				for i, fe := range validationErrors {
					errorString += fe.Field()
					if i < len(validationErrors)-1 {
						errorString += ", "
					}
					fieldRequired = httputil.GetErrorMsg(fe)
				}
				errorString += ": " + fieldRequired
			} else {
				errorString = "No body input"
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: errorString,
			})
			return
		}

		if len(request.Course) == 0 {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "'course' must have at least one element",
			})
			return
		}

		scheduleGroupIDStr, exist := ctx.Get("scheduleGroupId")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "scheduleGroupId not found",
			})
			return
		}
		scheduleGroupID := scheduleGroupIDStr.(int)

		logrus.WithFields(logrus.Fields{
			"data":        request,
			"studentCode": studentCode,
			"studentID":   studentID,
		}).Info("Submit-Registration")

		current, _ := ges.getEnrollSemester(ctx, scheduleGroupID)
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "current year or current semester not found",
			})
			return

		}
		_, scount, err := gl.GetLatest(studentID, string(SubmissionCollection), current.AcadYear, current.Semester)
		if err != nil && err != mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
				Code:    httputil.InternalServerError,
				Message: err.Error(),
			})
			return
		}
		_, ccount, err := gl.GetLatest(studentID, string(ConfirmationCollection), current.AcadYear, current.Semester)
		if err != nil && err != mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
				Code:    httputil.InternalServerError,
				Message: err.Error(),
			})
			return
		}
		reqReg := RegistrationConsumerRequest{
			StudentID:      studentID,
			StudentCode:    studentCode,
			UserID:         userId,
			Email:          email,
			ConfirmRequest: request,
			Sub:            sub,
		}

		id := primitive.NewObjectID()
		now := time.Now()
		sequence := int(scount) + 1
		switch sequence {
		case 0, 1:
			data := &SubmissionData{
				ID:              id,
				Name:            string(TopicRegisSubmit),
				Description:     "registration with status process.",
				Status:          SubmissionStatusProcess,
				Request:         reqReg,
				AcadYear:        current.AcadYear,
				Semester:        current.Semester,
				SequenceSubmit:  int(scount) + 1,
				SequenceConfirm: int(ccount),
				CreatedAt:       now,
				UpdateAt:        now,
			}
			err = insert.InsertRecord(SubmissionCollection, data)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "เกิดข้อผิดพลาดในการบันทึกข้อมูล",
				})
				return
			}

			msg, err := json.Marshal(&data)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
					Code:    httputil.BadRequest,
					Message: "Error marshaling JSON",
				})
				return
			}

			p.ProduceMessage(string(TopicRegisSubmit), msg)
			log.Print("Send topic regis submit to kafka")
		default:
			token, err := gt.getToken(TokenRequest{
				StudentID: strconv.Itoa(studentID),
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "Err get token: " + err.Error(),
				})
				return
			}
			submit, err := sm.Submit(request, token.Token)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "Err submit: " + err.Error(),
				})
				return
			}
			var status StatusEnum
			var desc string
			if submit.ErrorID == 0 {
				status = SubmissionStatusSuccess
				desc = "registration with status success."
			} else {
				status = SubmissionStatusFailed
				desc = "registration with status failed."
			}
			data := &SubmissionData{
				ID:              id,
				Name:            string(TopicRegisSubmit),
				Description:     desc,
				Status:          status,
				Request:         reqReg,
				AcadYear:        current.AcadYear,
				Semester:        current.Semester,
				SequenceSubmit:  int(scount) + 1,
				SequenceConfirm: int(ccount),
				CreatedAt:       now,
				UpdateAt:        now,
				Response: &InfoRegister{
					StatusCode:  strconv.Itoa(http.StatusOK),
					ResponseREG: submit,
				},
			}

			err = insert.InsertRecord(SubmissionCollection, data)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "เกิดข้อผิดพลาดในการบันทึกข้อมูล",
				})
				return
			}

			msg, err := json.Marshal(&data)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
					Code:    httputil.BadRequest,
					Message: "Error marshaling JSON",
				})
				return
			}
			p.ProduceMessage(string(TopicNotifySubmit), msg)
			log.Print("Send topic notify submit to kafka")

			ctx.JSON(http.StatusOK, submit.MapSubmitStudentRegistration())
			return
		}
		ctx.JSON(http.StatusOK, SubmitStudentRegistration{
			Code:    "-1",
			Message: "Send Topic Regis Submit to Kafka",
		})
	}
}

type confirmFunc func(ConfirmRequest, string) (ConfirmResponse, error)

func (cf confirmFunc) Confirm(requestBody ConfirmRequest, token string) (ConfirmResponse, error) {
	return cf(requestBody, token)
}

// ConfirmRegistrationHandler godoc
//
//	@tags			Register
//	@summary		Confirm registration student
//	@description	Confirm registration student
//	@id				Confirm registration student
//	@security		Bearer
//	@produce		json
//	@param			RequestBody	body						ConfirmRequest	true	"Confirm registration student"
//	@success		200			httputil.SuccessResponse	200				"Registration confirmed successfully."
//	@failure		400			{object}					httputil.HTTPBadRequestErrors
//	@failure		500			{object}					httputil.HTTPError
//	@router			/confirm [post]
func ConfirmRegistrationHandler(p *producer.Producer, insert insertRecordFunc, ges getEnrollSemesterFunc, gl getLatestFunc, gt getTokenFunc, cf confirmFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		studentIDValue, exist := ctx.Get("studentId")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "Missing required parameter: studentID",
			})
			return
		}
		studentCodeValue, exist := ctx.Get("studentCode")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "studentCode not found",
			})
			return
		}
		userIdValue, exist := ctx.Get("userId")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "userId not found",
			})
			return
		}
		emailValue, exist := ctx.Get("email")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "email not found",
			})
			return
		}
		subValue, exist := ctx.Get("sub")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "sub not found",
			})
			return
		}
		studentCode := studentCodeValue.(string)
		studentID := studentIDValue.(int)
		userId := userIdValue.(string)
		email := emailValue.(string)
		sub := subValue.(string)

		// Validate JSON request body
		var request ConfirmRequest
		err := ctx.ShouldBindJSON(&request)

		if err != nil {
			var errorString string
			var fieldRequired string
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
				for i, fe := range validationErrors {
					errorString += fe.Field()
					if i < len(validationErrors)-1 {
						errorString += ", "
					}
					fieldRequired = httputil.GetErrorMsg(fe)
				}
				errorString += ": " + fieldRequired
			} else {
				errorString = "No body input"
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: errorString,
			})
			return
		}

		if len(request.Course) == 0 {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "'course' must have at least one element",
			})
			return
		}

		scheduleGroupIDStr, exist := ctx.Get("scheduleGroupId")
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "scheduleGroupId not found",
			})
			return
		}
		scheduleGroupID := scheduleGroupIDStr.(int)

		logrus.WithFields(logrus.Fields{
			"data":        request,
			"studentCode": studentCode,
			"studentID":   studentID,
		}).Info("Confirm-Registration")

		current, _ := ges.getEnrollSemester(ctx, scheduleGroupID)
		if !exist {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "current year or current semester not found",
			})
			return

		}
		_, scount, err := gl.GetLatest(studentID, string(SubmissionCollection), current.AcadYear, current.Semester)
		if err != nil && err != mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
				Code:    httputil.InternalServerError,
				Message: err.Error(),
			})
			return
		}
		_, ccount, err := gl.GetLatest(studentID, string(ConfirmationCollection), current.AcadYear, current.Semester)
		if err != nil && err != mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
				Code:    httputil.InternalServerError,
				Message: err.Error(),
			})
			return
		}
		reqReg := RegistrationConsumerRequest{
			StudentID:      studentID,
			StudentCode:    studentCode,
			UserID:         userId,
			Email:          email,
			ConfirmRequest: request,
			Sub:            sub,
		}

		id := primitive.NewObjectID()
		now := time.Now()
		sequence := int(ccount) + 1
		switch sequence {
		case 0, 1:
			data := SubmissionData{
				ID:              id,
				Name:            string(TopicRegisConfirm),
				Description:     "registration with status process.",
				Status:          SubmissionStatusProcess,
				Request:         reqReg,
				AcadYear:        current.AcadYear,
				Semester:        current.Semester,
				SequenceSubmit:  int(scount),
				SequenceConfirm: int(ccount) + 1,
				CreatedAt:       now,
				UpdateAt:        now,
			}
			err = insert.InsertRecord(ConfirmationCollection, &data)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "เกิดข้อผิดพลาดในการบันทึกข้อมูล",
				})
				return
			}

			msg, err := json.Marshal(&data)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
					Code:    httputil.BadRequest,
					Message: "Error marshaling JSON",
				})
				return
			}

			p.ProduceMessage(string(TopicRegisConfirm), msg)
			log.Print("Send topic regis confirm to kafka")
		default:
			token, err := gt.getToken(TokenRequest{
				StudentID: strconv.Itoa(studentID),
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "Err get token: " + err.Error(),
				})
				return
			}
			confirm, err := cf.Confirm(request, token.Token)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "Err confirm: " + err.Error(),
				})
				return
			}
			var status StatusEnum
			var desc string
			if confirm.ErrorID == 0 {
				status = SubmissionStatusSuccess
				desc = "registration with status success."
			} else {
				status = SubmissionStatusFailed
				desc = "registration with status failed."
			}
			data := SubmissionData{
				ID:              id,
				Name:            string(TopicRegisConfirm),
				Description:     desc,
				Status:          status,
				Request:         reqReg,
				AcadYear:        current.AcadYear,
				Semester:        current.Semester,
				SequenceSubmit:  int(scount),
				SequenceConfirm: int(ccount) + 1,
				CreatedAt:       now,
				UpdateAt:        now,
				Response: &InfoRegister{
					StatusCode:  strconv.Itoa(http.StatusOK),
					ResponseREG: confirm,
				},
			}
			err = insert.InsertRecord(ConfirmationCollection, &data)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, httputil.HTTPError{
					Code:    httputil.InternalServerError,
					Message: "เกิดข้อผิดพลาดในการบันทึกข้อมูล",
				})
				return
			}
			msg, err := json.Marshal(&data)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
					Code:    httputil.BadRequest,
					Message: "Error marshaling JSON",
				})
				return
			}
			p.ProduceMessage(string(TopicNotifyConfirm), msg)
			log.Print("Send topic notify confirm to kafka")

			ctx.JSON(http.StatusOK, confirm.MapSubmitStudentRegistration())
			return
		}
		ctx.JSON(http.StatusOK, SubmitStudentRegistration{
			Code:    "-1",
			Message: "Send Topic Regis Confirm to Kafka",
		})
	}
}

// PushKafkaHandler
func PushKafkaHandler(p *producer.Producer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqReg := RegistrationConsumerRequest{
			StudentID:      630710645,
			StudentCode:    "630710645",
			UserID:         "6ad099b0-aa6e-4cb0-8c9a-1ab6d3f80a7c",
			Email:          "MADMUD_C@su.ac.th",
			ConfirmRequest: ConfirmRequest{},
			Sub:            "mobile",
		}
		data := SubmissionData{
			ID:              primitive.NewObjectID(),
			Name:            string(TopicRegisConfirm),
			Description:     "registration with status process.",
			Status:          SubmissionStatusProcess,
			Request:         reqReg,
			AcadYear:        2566,
			Semester:        2,
			SequenceSubmit:  1,
			SequenceConfirm: 2,
			CreatedAt:       time.Now(),
			UpdateAt:        time.Now(),
		}

		msg, err := json.Marshal(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httputil.HTTPError{
				Code:    httputil.BadRequest,
				Message: "Error marshaling JSON",
			})
			return
		}

		p.ProduceMessage(string(TopicNotifySubmit), msg)
		log.Print("Send topic regis confirm to kafka")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Send topic regis submit to kafka",
		})
	}
}
