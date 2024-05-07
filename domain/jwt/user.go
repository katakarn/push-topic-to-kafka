package jwt

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

type UserClaim struct {
	ID     uuid.UUID `json:"id"`
	UserID string    `json:"userId"`
	Email  string    `json:"email"`
	REGStudentEssentialIds
	Sub string `json:"sub"`
	jwt.StandardClaims
}

type REGStudentEssentialIds struct {
	StudentID       int    `json:"studentId"`
	StudentCode     string `json:"studentCode"`
	FacultyID       int    `json:"facultyId"`
	LevelID         int    `json:"levelId"`
	ScheduleGroupID int    `json:"scheduleGroupId"`
	StudentYear     int    `json:"studentYear"`
	DepartmentID    int    `json:"departmentId"`
}

func GetClaims(tokenString string) (*UserClaim, error) {
	parser := jwt.Parser{}
	var claim UserClaim
	token, _, err := parser.ParseUnverified(
		tokenString,
		&claim,
	)

	if err != nil {
		return nil, err
	}

	_, ok := token.Claims.(*UserClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	return &claim, nil
}
