package middleware

import (
	"errors"
	"net/http"
	"testKafka/domain/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserMobileAuthorizes() gin.HandlerFunc {
	return func(context *gin.Context) {
		clientToken := context.Request.Header.Get("Authorization")
		if clientToken == "" {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": errors.New("no authorization header provided").Error(),
			})
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": errors.New("incorrect format of authorization token").Error(),
			})
			return
		}

		token, err := jwt.GetClaims(clientToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		context.Set("token", clientToken)
		context.Set("userId", token.UserID)
		context.Set("email", token.Email)
		context.Set("scheduleGroupId", token.ScheduleGroupID)
		context.Set("studentId", token.StudentID)
		context.Set("studentCode", token.StudentCode)
		context.Set("levelId", token.LevelID)
		context.Set("facultyId", token.FacultyID)
		context.Set("departmentId", token.DepartmentID)
		context.Set("studentYear", token.StudentYear)
		context.Set("sub", token.Sub)

		context.Next()
	}
}
