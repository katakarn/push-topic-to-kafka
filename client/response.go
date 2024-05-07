package client

type GenerateTokenResponse struct {
	Token string `json:"token" bson:"token"`
}

type GetPreRegistrationHistoryResponse struct {
	ClassID              int64               `json:"classId"`
	CourseID             int64               `json:"courseId"`
	CourseCode           string              `json:"courseCode"`
	FacultyID            int64               `json:"facultyId"`
	FacultyName          string              `json:"facultyName"`
	FacultyNameEng       string              `json:"facultyNameEng"`
	LevelID              int64               `json:"levelId"`
	LevelName            string              `json:"levelName"`
	LevelNameEng         string              `json:"levelNameEng"`
	CourseUnit           string              `json:"courseUnit"`
	CreditTotal          int64               `json:"credittotal"`
	CourseName           string              `json:"courseName"`
	CourseNameEng        string              `json:"courseNameEng"`
	CourseTypeName       string              `json:"courseTypeName"`
	CourseTypeNameEng    string              `json:"courseTypeNameEng"`
	Condition            string              `json:"condition"`
	Section              int                 `json:"section"`
	Action               int64               `json:"action"`
	CreditAttempt        int64               `json:"creditattempt"`
	GradeMode            string              `json:"grademode"`
	Instructors          []InstructorHistory `json:"instructors"`
	Schedules            []SchedulesHistory  `json:"schedules"`
	IsReserveSeat        bool                `json:"isReserveSeat"`
	IsInProgramStructure bool                `json:"isInProgramStructure"`
}

type InstructorHistory struct {
	InstructorID     int    `json:"instructorId"`
	ExtendPrefix     string `json:"extendPrefix"`
	ExtenddPrefixEng string `json:"extendPrefixEng"`
	Prefix           string `json:"prefix"`
	PrefixEng        string `json:"prefixEng"`
	FirstName        string `json:"firstName"`
	FirstNameEng     string `json:"firstNameEng"`
	LastName         string `json:"lastName"`
	LastNameEng      string `json:"lastNameEng"`
}
type SchedulesHistory struct {
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	WeekDay     int    `json:"weekDay"`
	WeekDayName string `json:"weekDayName"`
	Location    string `json:"location"`
}

type GetRegistrationHistoryCourseResponse struct {
	CourseID      int64  `json:"courseId"`
	CourseCode    string `json:"courseCode"`
	CourseUnit    string `json:"courseUnit"`
	CourseName    string `json:"courseName"`
	CourseNameEng string `json:"courseNameEng"`
	Section       int    `json:"section"`
	Action        int64  `json:"action"`
	CreditAttempt int64  `json:"creditattempt"`
	GradeMode     string `json:"grademode"`
}

type GetRegistrationHistoryResponse struct {
	StudentID        string                                 `json:"studentId"`
	Semester         int                                    `json:"semester"`
	AcadYear         int                                    `json:"acadYear"`
	Sequence         int64                                  `json:"sequence"`
	RegistrationDate string                                 `json:"registrationDate"`
	Course           []GetRegistrationHistoryCourseResponse `json:"course"`
}
