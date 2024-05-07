package registration

type ConfirmCourseRequest struct {
	ClassID       int64   `json:"classid" bson:"classid" binding:"required"`
	CourseID      int64   `json:"courseid" bson:"courseid" binding:"required"`
	CourseCode    string  `json:"coursecode" bson:"coursecode" binding:"required"`
	Section       string  `json:"section" bson:"section" binding:"required"`
	Action        *int64  `json:"action" bson:"action" binding:"required"`
	CreditAttempt float32 `json:"creditattempt" bson:"creditattempt" binding:"required"`
	GradeMode     string  `json:"grademode" bson:"grademode" binding:"required,oneof=GD SW"`
}

type ConfirmRequest struct {
	Course []ConfirmCourseRequest `json:"course" bson:"course" binding:"required,dive"`
	Ip     string                 `json:"ip" bson:"ip" binding:"required"`
}

type TokenRequest struct {
	StudentCode string `json:"studentcode"`
	StudentID   string `json:"studentid"`
	Key         string `json:"key"`
}
type GetTokenREGRequest struct {
	StudentCode string `json:"studentcode"`
	Key         string `json:"key"`
}

type GetDebtRequest struct {
	AcadYear string `json:"acadYear"`
	Semester string `json:"semester"`
}
