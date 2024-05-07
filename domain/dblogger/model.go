package dblogger

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DbLoggerData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Status      string             `bson:"status" json:"status"`
	Metadata    interface{}        `bson:"metadata" json:"metadata"`
	CreatedAt   string             `bson:"createdAt" json:"createdAt"`
}

type InfoRegister struct {
	StatusCode  string          `json:"statusCode" bson:"statusCode"`
	ResponseREG ConfirmResponse `json:"responseREG" bson:"responseREG"`
}

type RegistrationConsumerRequest struct {
	ConfirmRequest ConfirmRequest `json:"confirmRequest" bson:"confirmRequest"`
	StudentID      int            `json:"studentId" bson:"studentId"`
	StudentCode    string         `json:"studentCode" bson:"studentCode"`
	UserID         string         `json:"userId" bson:"userId"`
	Email          string         `json:"email" bson:"email"`
	Sub            string         `json:"sub" bson:"sub"`
}

type ConfirmRequest struct {
	Course []ConfirmCourseRequest `json:"course" bson:"course"`
	Ip     string                 `json:"ip" bson:"ip"`
}

type ConfirmCourseRequest struct {
	ClassID       int64   `json:"classid" bson:"classid"`
	CourseID      int64   `json:"courseid" bson:"courseid"`
	CourseCode    string  `json:"coursecode" bson:"coursecode"`
	Section       string  `json:"section" bson:"section"`
	Action        int64   `json:"action" bson:"action"`
	CreditAttempt float32 `json:"creditattempt" bson:"creditattempt"`
	GradeMode     string  `json:"grademode" bson:"grademode"`
}

type ConfirmResponse struct {
	StudentID   *string                 `json:"studentId" bson:"studentId"`
	ErrorID     int64                   `json:"errorid" bson:"errorId"`
	ErrorString string                  `json:"errorstring" bson:"errorString"`
	Course      []ConfirmCourseResponse `json:"course" bson:"course"`
	Ip          string                  `json:"ip" bson:"ip"`
}

type ConfirmCourseResponse struct {
	ClassID       int64   `json:"classid" bson:"classId"`
	Action        int64   `json:"action" bson:"action"`
	CreditAttempt float32 `json:"creditattempt" bson:"creditAttempt"`
	GradeMode     string  `json:"grademode" bson:"gradeMode"`
	ErrorID       int64   `json:"errorid" bson:"errorId"`
	ErrorString   string  `json:"errorstring" bson:"errorString"`
}
