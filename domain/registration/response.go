package registration

import (
	"testKafka/client"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Info struct {
	Data     DetailRegistration `json:"data"`
	Messsage string             `json:"messsage"`
	Status   string             `json:"status"`
}

type DetailRegistration struct {
	Courses          []Course `json:"courses"`
	RegistrationDate string   `json:"registrationDate"`
	Semester         string   `json:"semester"`
	StudentID        int      `json:"studentID"`
	Year             string   `json:"year"`
}

type Course struct {
	CourseCode string `json:"courseCode"`
	CourseID   int    `json:"courseID"`
	Section    string `json:"section"`
}

type EventResponse struct {
	ScheduleCode int    `json:"scheduleCode"`
	Title        string `json:"title"`
	TitleEng     string `json:"titleEng"`
	AcadYear     int    `json:"acadYear"`
	Semester     int    `json:"semester"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
}

type RegistrationPeriodResponse struct {
	ConfirmStartDate                 string `json:"confirmStartDate"`
	ConfirmEndDate                   string `json:"confirmEndDate"`
	PaymentStartDate                 string `json:"paymentStartDate"`
	PaymentEndDate                   string `json:"paymentEndDate"`
	IsConfirmationRegistrationPeriod bool   `json:"isConfirmationRegistrationPeriod"`
	IsPaymentPeriod                  bool   `json:"isPaymentPeriod"`
	IsAddDeleteExchangePeriod        bool   `json:"isAddDeleteExchangePeriod"`
	AddDeleteExchangeStartDate       string `json:"addDeleteExchangeStartDate"`
	AddDeleteExchangeEndDate         string `json:"addDeleteExchangeEndDate"`
}

type TokenResponse struct {
	Token string `json:"token" bson:"token"`
}

type GetDebtResponse struct {
	FeeId         int64   `json:"feeId"`
	FeeName       string  `json:"feeName"`
	FeeNameEng    string  `json:"feeNameEng"`
	CourseCode    string  `json:"courseCode"`
	CourseName    string  `json:"courseName"`
	CourseNameEng string  `json:"courseNameEng"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
	VoucherName   string  `json:"voucherName"`
}
type DebtResponse struct {
	AcadYear string `json:"acadYear"`
	Semester string `json:"semester"`
	Debt     []Debt `json:"debt"`
	Paid     []Debt `json:"paid"`
	Process  []Debt `json:"process"`
}

type GetCourseSection struct {
	AcadYear     int           `json:"acadYear"`
	Semester     int           `json:"semester"`
	Course       CourseDetail  `json:"course"`
	SectionsInfo []SectionInfo `json:"sections"`
}

type GetCoursesResponse struct {
	AcadYear        int64             `json:"acadYear"`
	Semester        int64             `json:"semester"`
	CoursesResponse []CoursesResponse `json:"courses"`
}

type Courses struct {
	CourseID                int64  `pg:"courseid"`
	CourseName              string `pg:"coursename"`
	CourseNameEng           string `pg:"coursenameeng"`
	CourseCode              string `pg:"coursecode"`
	RevisionCode            string `pg:"revisioncode" `
	FacultyID               int64  `pg:"facultyid" `
	FacultyName             string `pg:"facultyname" `
	FacultyNameEng          string `pg:"facultynameeng" `
	LevelID                 int64  `pg:"levelid" `
	LevelName               string `pg:"levelname" `
	LevelNameEng            string `pg:"levelnameeng" `
	CourseTypeName          string `pg:"coursetypename" `
	CourseTypeNameEng       string `pg:"coursetypenameeng" `
	CourseUnit              string `pg:"courseunit" `
	CreditTotal             int    `pg:"credittotal" `
	GradeMode               string `pg:"grademode" `
	IsReserveSeatStr        string `pg:"reserve" `
	IsInProgramStructureStr string `pg:"prostruc" `
}

type CoursesResponse struct {
	CourseID             int64  `json:"courseId"`
	CourseName           string `json:"courseName"`
	CourseNameEng        string `json:"courseNameEng"`
	CourseCode           string `json:"courseCode"`
	RevisionCode         string `json:"-"`
	FacultyID            int64  `json:"facultyId"`
	FacultyName          string `json:"facultyName"`
	FacultyNameEng       string `json:"facultyNameEng"`
	LevelID              int64  `json:"levelId"`
	LevelName            string `json:"levelName"`
	LevelNameEng         string `json:"levelNameEng"`
	CourseTypeName       string `json:"courseTypeName"`
	CourseTypeNameEng    string `json:"courseTypeNameEng"`
	CourseUnit           string `json:"courseUnit"`
	CreditTotal          int    `json:"creditTotal"`
	GradeMode            string `json:"gradeMode"`
	IsReserveSeat        bool   `json:"isReserveSeat"`
	IsInProgramStructure bool   `json:"isInProgramStructure"`
}

type GetCoursesSuggestionResponse struct {
	AcadYear int64                       `json:"acadYear"`
	Semester int64                       `json:"semester"`
	Courses  []CoursesSuggestoinResponse `json:"courses"`
}

type CoursesSuggestion struct {
	CourseID                int64  `pg:"courseid"`
	CourseName              string `pg:"coursename"`
	CourseNameEng           string `pg:"coursenameeng"`
	CourseCode              string `pg:"coursecode"`
	RevisionCode            string `pg:"revisioncode" `
	FacultyID               int64  `pg:"facultyid" `
	FacultyName             string `pg:"facultyname" `
	FacultyNameEng          string `pg:"facultynameeng" `
	Section                 string `pg:"section"`
	CourseUnit              string `pg:"courseunit" `
	CreditTotal             int    `pg:"credittotal" `
	GradeMode               string `pg:"grademode" `
	IsReserveSeatStr        string `pg:"reserve" `
	IsInProgramStructureStr string `pg:"prostruc" `
}

type CoursesSuggestoinResponse struct {
	CourseID             int64  `json:"courseId"`
	CourseName           string `json:"courseName"`
	CourseNameEng        string `json:"courseNameEng"`
	CourseCode           string `json:"courseCode"`
	RevisionCode         string `json:"-"`
	FacultyID            int64  `json:"facultyId"`
	FacultyName          string `json:"facultyName"`
	FacultyNameEng       string `json:"facultyNameEng"`
	Section              string `json:"section"`
	CourseUnit           string `json:"courseUnit"`
	CreditTotal          int    `json:"creditTotal"`
	GradeMode            string `json:"gradeMode"`
	IsReserveSeat        bool   `json:"isReserveSeat"`
	IsInProgramStructure bool   `json:"isInProgramStructure"`
}

type GetPreRegistrationHistoryResponse struct {
	ClassID              int64                      `json:"classId"`
	CourseID             int64                      `json:"courseId"`
	CourseCode           string                     `json:"courseCode"`
	FacultyID            int64                      `json:"facultyId"`
	FacultyName          string                     `json:"facultyName"`
	FacultyNameEng       string                     `json:"facultyNameEng"`
	LevelID              int64                      `json:"levelId"`
	LevelName            string                     `json:"levelName"`
	LevelNameEng         string                     `json:"levelNameEng"`
	CourseUnit           string                     `json:"courseUnit"`
	CreditTotal          int64                      `json:"credittotal"`
	CourseName           string                     `json:"courseName"`
	CourseNameEng        string                     `json:"courseNameEng"`
	CourseTypeName       string                     `json:"courseTypeName"`
	CourseTypeNameEng    string                     `json:"courseTypeNameEng"`
	Condition            string                     `json:"condition"`
	Section              string                     `json:"section"`
	Action               int64                      `json:"action"`
	CreditAttempt        int64                      `json:"creditattempt"`
	GradeMode            string                     `json:"grademode"`
	Instructors          []client.InstructorHistory `json:"instructors"`
	Schedules            []client.SchedulesHistory  `json:"schedules"`
	IsReserveSeat        bool                       `json:"isReserveSeat"`
	IsInProgramStructure bool                       `json:"isInProgramStructure"`
}

type GetRegistrationHistoryResponse struct {
	StudentID        string                                 `json:"studentId"`
	Semester         int                                    `json:"semester"`
	AcadYear         int                                    `json:"acadYear"`
	Sequence         int64                                  `json:"sequence"`
	RegistrationDate string                                 `json:"registrationDate"`
	Course           []GetRegistrationHistoryCourseResponse `json:"course"`
}

type GetRegistrationHistoryCourseResponse struct {
	CourseID      int64  `json:"courseId"`
	CourseCode    string `json:"courseCode"`
	CourseUnit    string `json:"courseUnit"`
	CourseName    string `json:"courseName"`
	CourseNameEng string `json:"courseNameEng"`
	Section       string `json:"section"`
	Action        int64  `json:"action"`
	CreditAttempt int64  `json:"creditattempt"`
	GradeMode     string `json:"grademode"`
}

type RegistrationStatus interface {
	isRegistrationStatus()
}

func (s RegistrationStatusEnum) isRegistrationStatus() {}

type RegistrationStatusResponse struct {
	RegistrationStatus RegistrationStatus `json:"registrationStatus"`
}

type Tracking struct {
	RegistrationStatus RegistrationStatus `json:"registrationStatus"`
	SequenceSubmit     int                `json:"sequenceSubmit"`
	SequenceConfirm    int                `json:"sequenceConfirm"`
	ErrorID            *int64             `json:"errorid"`
	ErrorString        string             `json:"errorString"`
	Courses            []CourseTracking   `json:"courses"`
}

type CourseTracking struct {
	Course       CourseTrackingDetail `json:"course"`
	SectionsInfo []SectionTracking    `json:"sections"`
	ErrorID      int64               `json:"errorid"`
	ErrorString  string               `json:"errorString"`
}

type CourseTrackingDetail struct {
	Action        *int64  `json:"action"`
	CreditAttempt float32 `json:"creditattempt"`
	CourseDetail
}

type SubmitStudentRegistration struct {
	Code        string         `json:"code" bson:"code"`
	Message     string         `json:"message" bson:"message"`
	ErrorObject SubmitResponse `json:"errorObject" bson:"errorObject"`
}

type SubmitResponse struct {
	StudentID   string                  `json:"studentId" bson:"studentId"`
	ErrorID     *int64                  `json:"errorid" bson:"errorId"`
	ErrorString string                  `json:"errorstring" bson:"errorString"`
	Course      []ConfirmCourseResponse `json:"course" bson:"course"`
	Ip          string                  `json:"ip" bson:"ip"`
}

// Map ConfirmResponse to SubmitStudentRegistration
func (cr ConfirmResponse) MapSubmitStudentRegistration() SubmitStudentRegistration {
	submitResponse := SubmitResponse{
		StudentID:   *cr.StudentID,
		ErrorID:     &cr.ErrorID,
		ErrorString: cr.ErrorString,
		Course:      cr.Course,
		Ip:          cr.Ip,
	}
	return SubmitStudentRegistration{
		Code:        strconv.Itoa(int(cr.ErrorID)),
		Message:     cr.ErrorString,
		ErrorObject: submitResponse,
	}
}

type SectionTracking struct {
	ClassID                 int                `pg:"classid" json:"classId"`
	Section                 string             `pg:"section" json:"section"`
	SeatInfo                SeatQuantity       `json:"seatInfo"`
	Instructors             []Instructors      `json:"instructors"`
	SectionSchedules        []SectionSchedules `json:"schedules"`
	Remark                  string             `pg:"classnote" json:"remark"`
	IsReserveSeatStr        string             `pg:"reserve" json:"-"`
	IsInProgramStructureStr string             `pg:"prostruc" json:"-"`
	IsReserveSeat           bool               `json:"isReserveSeat"`
	IsInProgramStructure    bool               `json:"isInProgramStructure"`
}

type GetRegistrationStatusResponse struct {
	Status StatusEnum    `json:"status"`
	Data   *InfoRegister `json:"data"`
}

func mapCoursesResponse(courses []Courses, currentAcadYear string, currentSemester string) GetCoursesResponse {
	courseList := []CoursesResponse{}
	for _, course := range courses {
		IsReserveSeat, _ := strconv.ParseBool(course.IsReserveSeatStr)
		IsInProgramStructure, _ := strconv.ParseBool(course.IsInProgramStructureStr)
		courseList = append(courseList, CoursesResponse{
			CourseID:             course.CourseID,
			CourseName:           course.CourseName,
			CourseNameEng:        course.CourseNameEng,
			CourseCode:           course.CourseCode + "-" + course.RevisionCode,
			FacultyID:            course.FacultyID,
			FacultyName:          course.FacultyName,
			FacultyNameEng:       course.FacultyNameEng,
			LevelID:              course.LevelID,
			LevelName:            course.LevelName,
			LevelNameEng:         course.LevelNameEng,
			CourseTypeName:       course.CourseTypeName,
			CourseTypeNameEng:    course.CourseTypeNameEng,
			CourseUnit:           course.CourseUnit,
			CreditTotal:          course.CreditTotal,
			GradeMode:            course.GradeMode,
			IsReserveSeat:        IsReserveSeat,
			IsInProgramStructure: IsInProgramStructure,
		})
	}
	AcadYear, _ := strconv.ParseInt(currentAcadYear, 10, 64)
	Semester, _ := strconv.ParseInt(currentSemester, 10, 64)
	response := GetCoursesResponse{
		AcadYear:        AcadYear,
		Semester:        Semester,
		CoursesResponse: courseList,
	}
	return response
}

func mapCoursesSuggestionResponse(coursesSuggestion []CoursesSuggestion, currentAcadYear string, currentSemester string) GetCoursesSuggestionResponse {
	courseList := []CoursesSuggestoinResponse{}
	for _, course := range coursesSuggestion {
		IsReserveSeat, _ := strconv.ParseBool(course.IsReserveSeatStr)
		IsInProgramStructure, _ := strconv.ParseBool(course.IsInProgramStructureStr)
		courseList = append(courseList, CoursesSuggestoinResponse{
			CourseID:             course.CourseID,
			CourseName:           course.CourseName,
			CourseNameEng:        course.CourseNameEng,
			CourseCode:           course.CourseCode + "-" + course.RevisionCode,
			FacultyID:            course.FacultyID,
			FacultyName:          course.FacultyName,
			FacultyNameEng:       course.FacultyNameEng,
			Section:              course.Section,
			CourseUnit:           course.CourseUnit,
			CreditTotal:          course.CreditTotal,
			GradeMode:            course.GradeMode,
			IsReserveSeat:        IsReserveSeat,
			IsInProgramStructure: IsInProgramStructure,
		})
	}
	AcadYear, _ := strconv.ParseInt(currentAcadYear, 10, 64)
	Semester, _ := strconv.ParseInt(currentSemester, 10, 64)
	response := GetCoursesSuggestionResponse{
		AcadYear: AcadYear,
		Semester: Semester,
		Courses:  courseList,
	}
	return response
}

type CourseSectionDetail struct {
	Section           string                        `pg:"section" json:"section"`
	CourseID          int                           `pg:"courseid" json:"courseId"`
	CourseName        string                        `pg:"coursename" json:"courseName"`
	CourseNameEng     string                        `pg:"coursenameeng" json:"courseNameEng"`
	CourseCode        string                        `pg:"coursecode" json:"courseCode"`
	RevisionCode      string                        `pg:"revisioncode" json:"-"`
	CourseDescription string                        `pg:"description1" json:"courseDescription"`
	Programs          []Program                     `json:"programs"`
	MidTermSchedules  []CourseSectionDetailSchedule `json:"midTermSchedules"`
	FinalSchedules    []CourseSectionDetailSchedule `json:"finalSchedules"`
	ReserveFor        []ReserveInfo                 `json:"reserveFor"`
}

type Program struct {
	FacultyID         int    `pg:"facultyid" json:"facultyId"`
	FacultyName       string `pg:"facultyname" json:"facultyName"`
	FacultyNameEng    string `pg:"facultynameeng" json:"facultyNameEng"`
	DepartmentID      int    `pg:"departmentid" json:"departmentId"`
	DepartmentName    string `pg:"departmentname" json:"departmentName"`
	DepartmentNameEng string `pg:"departmentnameeng" json:"departmentNameEng"`
}

type CourseSectionDetailSchedule struct {
	Date      string `pg:"examdate" json:"date"`
	StartTime string `pg:"examtimefrom" json:"startTime"`
	EndTime   string `pg:"examtimeto" json:"endTime"`
	Location  string `json:"location"`
	Room      string `pg:"roomname" json:"-"`
	Building  string `pg:"buildingname" json:"-"`
}

type ReserveInfo struct {
	CampusID          int          `pg:"campusid" json:"campusId"`
	CampusName        string       `pg:"campusname" json:"campusName"`
	CampusNameEng     string       `pg:"campusnameeng" json:"campusNameEng"`
	FacultyID         int          `pg:"facultyid" json:"facultyId"`
	FacultyName       string       `pg:"facultyname" json:"facultyName"`
	FacultyNameEng    string       `pg:"facultynameeng" json:"facultyNameEng"`
	DepartmentID      int          `pg:"departmentid" json:"departmentId"`
	DepartmentName    string       `pg:"departmentname" json:"departmentName"`
	DepartmentNameEng string       `pg:"departmentnameeng" json:"departmentNameEng"`
	ReserveInfo       string       `pg:"reserveinfo" json:"reserveInfo"`
	SeatInfo          SeatQuantity `json:"seatInfo"`
	Total             int          `pg:"reserveseat" json:"-"`
	Reserved          int          `pg:"enrollseat" json:"-"`
}

type GetEnrollmentHistoryResponse struct {
	AcadYear          string              `json:"acadYear"`
	Semester          string              `json:"semester"`
	EnrollmentHistory []EnrollmentHistory `json:"enrollmentHistory"`
}

type EnrollmentHistory struct {
	CourseID      int64   `json:"courseId"`
	CourseCode    string  `json:"courseCode"`
	CourseUnit    string  `json:"courseUnit"`
	CourseName    string  `json:"courseName"`
	CourseNameEng string  `json:"courseNameEng"`
	Section       string  `json:"section"`
	CreditAttempt *int64  `json:"creditattempt"`
	GradeMode     *string `json:"grademode"`
	ClassID       int64   `json:"classId"`

	Instructors          []Instructors      `json:"instructors"`
	SectionSchedules     []SectionSchedules `json:"schedules"`
	IsReserveSeat        bool               `json:"isReserveSeat"`
	IsInProgramStructure bool               `json:"isInProgramStructure"`
	FacultyName          string             `json:"facultyName"`
	FacultyNameEng       string             `json:"facultyNameEng"`
}

type NewsResponse struct {
	ID              primitive.ObjectID `json:"id,omitempty"`
	NewsId          string             `json:"newsId"`
	Type            string             `json:"type"`
	Title           string             `json:"title"`
	Content         string             `json:"content"`
	ImageBanner     string             `json:"imageBanner"`
	ImageBannerName string             `json:"imageBannerName"`
	Images          []string           `json:"images"`
	PublishDate     string             `json:"publishDate,omitempty"`
	UnpublishDate   string             `json:"unpublishDate,omitempty"`
	CategoryID      string             `json:"categoryId"`
	AnnouncerID     string             `json:"announcerId"`
	Campus          []string           `json:"campus"`
	Status          string             `json:"status,omitempty"`
	IsDeleted       bool               `json:"isDeleted"`
	CreateDate      string             `json:"createDate,omitempty"`
	UpdateDate      string             `json:"updateDate,omitempty"`
	CreateBy        string             `json:"createBy"`
	UpdateBy        string             `json:"updateBy"`
}