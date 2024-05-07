package dblogger

type RegistrationLogType string

const (
	OpenPaymentWebview       RegistrationLogType = "Open Payment Webview"
	RegistrationConfirmation RegistrationLogType = "Registration Confirmation"
)

type RegistrationLog struct {
	StudentCode     string              `json:"studentCode" bson:"studentCode"`
	StudentID       int                 `json:"studentId" bson:"studentId"`
	FacultyID       int                 `json:"facultyId"`
	LevelID         int                 `json:"levelId"`
	ScheduleGroupID int                 `json:"scheduleGroupId"`
	StudentYear     int                 `json:"studentYear"`
	DepartmentID    int                 `json:"departmentId"`
	RegisterDate    string              `json:"registerDate" bson:"registerDate"`
	Courses         []Course            `json:"courses" bson:"courses"`
	Year            string              `json:"year" bson:"year"`
	Semester        string              `json:"semester" bson:"semester"`
	Ip              string              `json:"ip" bson:"ip"`
	Type            RegistrationLogType `json:"type" bson:"type"`
	DateTime        string              `json:"dateTime" bson:"dateTime"`
	Result          interface{}         `json:"result" bson:"result"`
}

type Course struct {
	ClassID       int64   `json:"classid" bson:"classId"`
	CourseID      int64   `json:"courseid" bson:"courseId"`
	CourseCode    string  `json:"coursecode" bson:"courseCode"`
	Section       string  `json:"section" bson:"section"`
	Action        int64   `json:"action" bson:"action"`
	CreditAttempt float32 `json:"creditattempt" bson:"creditAttempt"`
	GradeMode     string  `json:"grademode" bson:"gradeMode"`
}
