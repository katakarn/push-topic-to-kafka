package registration

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusEnum string

const (
	SubmissionStatusSuccess StatusEnum = "SUCCESS"
	SubmissionStatusFailed  StatusEnum = "FAILED"
	SubmissionStatusRetry   StatusEnum = "RETRY"
	SubmissionStatusProcess StatusEnum = "PROCESS"
)

type CollectionName string

const (
	SubmissionCollection   CollectionName = "submissions"
	ConfirmationCollection CollectionName = "confirmations"
)

type SubmissionData struct {
	ID              primitive.ObjectID          `bson:"_id,omitempty" json:"id,omitempty"`
	Name            string                      `bson:"name" json:"name"`
	Description     string                      `bson:"description" json:"description"`
	Status          StatusEnum                  `bson:"status" json:"status"`
	Attempt         int                         `bson:"attempt" json:"attempt"`
	Request         RegistrationConsumerRequest `bson:"request" json:"request"`
	Response        *InfoRegister               `bson:"response" json:"response"`
	SequenceSubmit  int                         `bson:"sequenceSubmit" json:"sequenceSubmit"`
	SequenceConfirm int                         `bson:"sequenceConfirm" json:"sequenceConfirm"`
	AcadYear        int                         `bson:"acadYear" json:"acadYear"`
	Semester        int                         `bson:"semester" json:"semester"`
	CreatedAt       time.Time                   `bson:"createdAt" json:"createdAt"`
	UpdateAt        time.Time                   `bson:"updatedAt" json:"updatedAt"`
}
type InfoRegister struct {
	StatusCode  string          `json:"statusCode" bson:"statusCode"`
	ResponseREG ConfirmResponse `json:"responseREG" bson:"responseREG"`
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
type RegistrationConsumerRequest struct {
	ConfirmRequest ConfirmRequest `json:"confirmRequest" bson:"confirmRequest"`
	StudentID      int            `json:"studentId" bson:"studentId"`
	StudentCode    string         `json:"studentCode" bson:"studentCode"`
	UserID         string         `json:"userId" bson:"userId"`
	Email          string         `json:"email" bson:"email"`
	Sub            string         `json:"sub" bson:"sub"`
}

type TopicsEnum string

const (
	TopicRegisSubmit   TopicsEnum = "SUBMIT_REGISTRATION"
	TopicRegisConfirm  TopicsEnum = "CONFIRM_REGISTRATION"
	TopicNotifySubmit  TopicsEnum = "NOTIFY_SUBMISSION_REGISTER"
	TopicNotifyConfirm TopicsEnum = "NOTIFY_CONFIRMATION_REGISTER"
)

type ScheduleItem struct {
	tableName           struct{} `pg:"app_registration_schedule_item"`
	ScheduleGroupID     int      `pg:"schedulegroupid"`
	ScheduleCode        int      `pg:"schedulecode"`
	ScheduleCodeName    string   `pg:"schedulecodename"`
	ScheduleCodeNameEng string   `pg:"schedulecodenameeng"`
	AcadYear            int      `pg:"acadyear"`
	Semester            int      `pg:"semester"`
	DateFrom            string   `pg:"datefrom"`
	DateTo              string   `pg:"dateto"`
}

type ScheduleGroup struct {
	tableName            struct{} `pg:"app_registration_schedule_group"`
	ScheduleGroupID      int      `pg:"schedulegroupid,pk"`
	ScheduleGroupName    string   `pg:"schedulegroupname"`
	Schedulegroupnameeng string   `pg:"schedulegroupnameeng"`
	CurrentAcadyear      int      `pg:"currentacadyear"`
	CurrentSemester      int      `pg:"currentsemester"`
	EnrollAcadyear       int      `pg:"enrollacadyear"`
	EnrollSemester       int      `pg:"enrollsemester"`
}

type Semester struct {
	AcadYear int `db:"ACADYEAR" json:"acadYear"`
	Semester int `db:"SEMESTER" json:"semester"`
}

type SuggestionCoursesRequest struct {
	AcadYear string `json:"acadYear" binding:"required"`
	Semester string `json:"semester" binding:"required"`
}

type ConfirmationPeriodREG struct {
	DateFrom     *time.Time `pg:"datefrom" json:"dateFrom"`
	DateTo       *time.Time `pg:"dateto" json:"dateTo"`
	ScheduleCode int        `pg:"schedulecode" json:"scheduleCode"`
	InPeriod     int        `pg:"inperiod" json:"inPeriod"`
}

func (c *ConfirmationPeriodREG) GetInPeriod() bool {
	if c.InPeriod == 1 {
		return true
	}
	return false
}

func (c *ConfirmationPeriodREG) GetDateFrom() *string {
	var date string
	if c.DateFrom != nil {
		date = c.DateFrom.Format("2006-01-02T15:04:05.000Z")
	}
	return &date
}

func (c *ConfirmationPeriodREG) GetDateTo() *string {
	var date string
	if c.DateTo != nil {
		date = c.DateTo.Format("2006-01-02T15:04:05.000Z")
	}
	return &date
}

type Debt struct {
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

type RegistrationStudentToken struct {
	tableName struct{}  `pg:"registration_student_token"`
	ID        int       `pg:",pk,type:serial"`
	StudentID string    `pg:"student_id,type:varchar(255),notnull,unique"`
	Token     string    `pg:"token,type:text,notnull"`
	ExpiredAt time.Time `pg:"expired_at,type:timestamp"`
}
type CoursesSection struct {
	AcadYear int          `json:"acadYear"`
	Semester int          `json:"semester"`
	Course   CourseDetail `json:"course"`
	Sections SectionInfo  `json:"sections"`
}

type CourseDetail struct {
	CourseID          int    `pg:"courseid" json:"courseId"`
	CourseName        string `pg:"coursename" json:"courseName"`
	CourseNameEng     string `pg:"coursenameeng" json:"courseNameEng"`
	CourseCode        string `pg:"coursecode" json:"courseCode"`
	RevisionCode      string `pg:"revisioncode" json:"-"`
	FacultyID         int    `pg:"facultyid" json:"facultyId"`
	FacultyName       string `pg:"facultyname" json:"facultyName"`
	FacultyNameEng    string `pg:"facultynameeng" json:"facultyNameEng"`
	LevelID           int    `pg:"levelid" json:"levelId"`
	LevelName         string `pg:"levelname" json:"levelName"`
	LevelNameEng      string `pg:"levelnameeng" json:"levelNameEng"`
	CourseTypeName    string `pg:"coursetypename" json:"courseTypeName"`
	CourseTypeNameEng string `pg:"coursetypenameeng" json:"courseTypeNameEng"`
	CourseUnit        string `pg:"courseunit" json:"courseUnit"`
	Condition         string `pg:"condition" json:"condition"`
	CreditTotal       int    `pg:"credittotal" json:"credittotal"`
	GradeMode         string `pg:"grademode" json:"grademode"`
}

type SectionInfo struct {
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

type SeatQuantity struct {
	ClassID  int `pg:"classid" json:"-"`
	Total    int `pg:"totalseat" json:"total"`
	Reserved int `pg:"enrollseat" json:"reserved"`
}

type Instructors struct {
	ClassID         int    `pg:"classid" json:"-"`
	InstructorID    int    `pg:"officerid" json:"instructorId"`
	ExtendPrefix    string `pg:"extendprefix" json:"extendPrefix"`
	ExtendPrefixEng string `pg:"extendprefixeng" json:"extendPrefixEng"`
	Prefix          string `pg:"prefixname" json:"prefix"`
	PrefixEng       string `pg:"prefixnameeng" json:"prefixEng"`
	FirstName       string `pg:"officername" json:"firstName"`
	FirstNameEng    string `pg:"officernameeng" json:"firstNameEng"`
	LastName        string `pg:"officersurname" json:"lastName"`
	LastNameEng     string `pg:"officersurnameeng" json:"lastNameEng"`
}
type SectionSchedules struct {
	ClassID     int    `pg:"classid" json:"-"`
	StartTime   string `pg:"timefrom" json:"startTime"`
	EndTime     string `pg:"timeto" json:"endTime"`
	WeekDay     int    `pg:"weekday" json:"weekDay"`
	WeekDayName string `pg:"weekdaynameeng" json:"weekDayName"`
	Location    string `pg:"buildingname" json:"location"`
	Room        string `pg:"roomname" json:"-"`
}

type RegistrationStatusEnum string

const (
	NoRegistrationHistory RegistrationStatusEnum = "NONE"
	SubmittingProcess     RegistrationStatusEnum = "SUBMITTING_PROCESS"
	SubmitSuccess         RegistrationStatusEnum = "SUBMIT_SUCCESS"
	SubmitFailed          RegistrationStatusEnum = "SUBMIT_FAILED"
	ConfirmingProcess     RegistrationStatusEnum = "CONFIRMING_PROCESS"
	ConfirmSuccess        RegistrationStatusEnum = "CONFIRM_SUCCESS"
	ConfirmFailed         RegistrationStatusEnum = "CONFIRM_FAILED"
)

type EnrollmentHistoryQueryRes struct {
	CourseID       int64   `db:"COURSEID" json:"courseId"`
	CourseCode     string  `db:"COURSECODE" json:"courseCode"`
	CourseUnit     string  `db:"COURSEUNIT" json:"courseUnit"`
	CourseName     string  `db:"COURSENAME" json:"courseName"`
	CourseNameEng  string  `db:"COURSENAMEENG" json:"courseNameEng"`
	Section        string  `db:"SECTION" json:"section"`
	CreditAttempt  *int64  `db:"CREDITTOTAL" json:"creditattempt"`
	GradeMode      *string `db:"GRADEMODE" json:"grademode"`
	ClassID        int64   `db:"CLASSID" json:"classId"`
	FacultyName    string  `db:"FACULTYNAME" json:"facultyName"`
	FacultyNameEng string  `db:"FACULTYNAMEENG" json:"facultyNameEng"`
}

type MyProgramCourse struct {
	ProgramID      int     `db:"PROGRAMID" json:"programId"`
	ProgramName    *string `db:"PROGRAMNAME" json:"programName"`
	ProgramNameEng *string `db:"PROGRAMNAMEENG" json:"programNameEng"`
	FacultyID      *string `db:"FACULTYID" json:"facultyId"`
	FacultyName    *string `db:"FACULTYNAME" json:"facultyName"`
	FacultyNameEng *string `db:"FACULTYNAMEENG" json:"facultyNameEng"`
	LevelID        *string `db:"LEVELID" json:"levelId"`
	LevelName      *string `db:"LEVELNAME" json:"levelName"`
	LevelNameEng   *string `db:"LEVELNAMEENG" json:"levelNameEng"`
	CourseCode     *string `db:"COURSECODE" json:"courseCode"`
	CourseName     *string `db:"COURSENAME" json:"courseName"`
	CourseNameEng  *string `db:"COURSENAMEENG" json:"courseNameEng"`
	CourseUnit     *string `db:"COURSEUNIT" json:"courseUnit" `
	CourseTypeName *string `db:"COURSETYPENAME" json:"courseTypeName"`
}

type MyProgramOnCurrentSemesterInfo struct {
	ProgramID      int    `db:"PROGRAMID" json:"programId"`
	FacultyName    string `db:"FACULTYNAME" json:"facultyName"`
	FacultyNameEng string `db:"FACULTYNAMEENG" json:"facultyNameEng"`
	FacultyID      string `db:"FACULTYID" json:"facultyId"`
	CampusID       int    `db:"CAMPUSID" json:"campusId"`
}

type CourseDetailTimetable struct {
	Weekday        *int       `db:"WEEKDAY" json:"weekday"`
	WeekDayName    *string    `db:"WEEKDAYNAME" json:"weekDayName"`
	WeekDayNameEng *string    `db:"WEEKDAYNAMEENG" json:"weekDayNameEng"`
	Start          *string `db:"TIMEFROM" json:"start"`
	End            *string `db:"TIMETO" json:"end"`
	Location       string     `db:"-" json:"location"` // location หาจากไหน
	ClassID        string     `db:"CLASSID"  json:"-"`
	RoomCode       *string    `db:"ROOMCODE" json:"roomCode"`
	BuildingCode   *string    `db:"BUILDINGCODE" json:"buildingCode"`
}

func TakeOutNullString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

type CourseDetailLecturer struct {
	OfficerId         int     `db:"OFFICERID" json:"officerId"`
	OfficerName       *string `db:"OFFICERNAME" json:"officerName"`
	OfficerNameEng    *string `db:"OFFICERNAMEENG" json:"officerNameEng"`
	OfficerSurName    *string `db:"OFFICERSURNAME" json:"officerSurName"`
	OfficerSurNameEng *string `db:"OFFICERSURNAMEENG" json:"officerSurNameEng"`
	PrefixName        *string `db:"PREFIXNAME" json:"prefixName"`
	PrefixNameEng     *string `db:"PREFIXNAMEENG" json:"prefixNameEng"`
	ClassID           string  `db:"CLASSID"  json:"-"`
}

type OfficerNameFromdatabase struct {
	PrefixName        *string `db:"PREFIXNAME" json:"prefixName"`
	OfficerName       *string `db:"OFFICERNAME" json:"officerName"`
	OfficerSurname    *string `db:"OFFICERSURNAME" json:"officerSurname"`
	PrefixNameEng     *string `db:"PREFIXNAMEENG" json:"prefixNameEng"`
	OfficerNameEng    *string `db:"OFFICERNAMEENG" json:"officerNameEng"`
	OfficerSurnameEng *string `db:"OFFICERSURNAMEENG" json:"officerSurnameEng"`
}

type OfficerName struct {
	ExtendPrefix    string `json:"extendPrefix"`
	ExtendPrefixEng string `json:"extendPrefixEng"`
	Prefix          string `json:"prefix"`
	PrefixEng       string `json:"prefixEng"`
	FirstName       string `json:"firstName"`
	FirstNameEng    string `json:"firstNameEng"`
	LastName        string `json:"lastName"`
	LastNameEng     string `json:"lastNameEng"`
}
type NewsREG struct {
	WebMsgID        int        `pg:"webmsgid"`
	TargetType      string     `pg:"targettype"`
	TargetTypeName  string     `pg:"targettypename"`
	WebTitle        *string    `pg:"webtitle"`
	WebMsg          *string    `pg:"webmsg"`
	WEBTitleENG     *string    `pg:"webtitleeng"`
	WebMsgENG       *string    `pg:"webmsgeng"`
	DateFrom        *time.Time `pg:"datefrom"`
	DateTo          *time.Time `pg:"dateto"`
	ClassID         *int       `pg:"classid"`
	StudentID       *int       `pg:"studentid"`
	FacultyID       *int       `pg:"facultyid"`
	LevelID         *int       `pg:"levelid"`
	Sender          *string    `pg:"sender"`
	SenderENG       *string    `pg:"sendereng"`
	CreateDateTime  *time.Time `pg:"createdatetime"`
	Priority        int        `pg:"priority"`
	ImageFileName   *string    `pg:"imagefilename"`
	ProgramID       *int       `pg:"programid"`
	AdmitAcadYear   *int       `pg:"admitacadyear"`
	CreateOfficerID *int       `pg:"createofficerid"`
}

type NewsType string
type AnnouncerID string
type CategoryID string
type NewsStatus string

type NewsItem struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NewsId          string             `bson:"newsId" json:"newsId"`
	Type            NewsType           `bson:"type" json:"type"`
	Title           string             `bson:"title" json:"title"`
	Content         string             `bson:"content" json:"content"`
	ImageBanner     string             `bson:"imageBanner" json:"imageBanner"`
	ImageBannerName *string            `bson:"imageBannerName" json:"imageBannerName"`
	Images          []string           `bson:"images" json:"images"`
	PublishDate     time.Time          `bson:"publishDate,omitempty" json:"publishDate,omitempty"`
	UnpublishDate   *time.Time         `bson:"unpublishDate,omitempty" json:"unpublishDate,omitempty"`
	CategoryID      CategoryID         `bson:"categoryId" json:"categoryId"`
	AnnouncerID     AnnouncerID        `bson:"announcerId" json:"announcerId"`
	Campus          []string           `bson:"campus" json:"campus"`
	Status          NewsStatus         `bson:"-" json:"status,omitempty"`
	IsDeleted       bool               `bson:"isDeleted" json:"isDeleted"`
	CreateDate      time.Time          `bson:"createDate,omitempty" json:"createDate,omitempty"`
	UpdateDate      time.Time          `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	CreateBy        *string            `bson:"createBy" json:"createBy"`
	UpdateBy        string             `bson:"updateBy" json:"updateBy"`
}

func (n *NewsItem) ToNewsResponse() NewsResponse {
	imageBannerName := ""
	if n.ImageBannerName != nil {
		imageBannerName = *n.ImageBannerName
	}
	createBy := ""
	if n.CreateBy != nil {
		createBy = *n.CreateBy
	}

	return NewsResponse{
		ID:              n.ID,
		NewsId:          n.NewsId,
		Type:            string(n.Type),
		Title:           n.Title,
		Content:         n.Content,
		ImageBanner:     n.ImageBanner,
		ImageBannerName: imageBannerName,
		Images:          n.Images,
		PublishDate:     n.PublishDate.Format("2006-01-02T15:04:05.000Z"),
		UnpublishDate:   n.UnpublishDate.Format("2006-01-02T15:04:05.000Z"),
		CategoryID:      string(n.CategoryID),
		AnnouncerID:     string(n.AnnouncerID),
		Campus:          n.Campus,
		Status:          string(n.Status),
		IsDeleted:       n.IsDeleted,
		CreateDate:      n.CreateDate.Format("2006-01-02T15:04:05.000Z"),
		UpdateDate:      n.UpdateDate.Format("2006-01-02T15:04:05.000Z"),
		CreateBy:        createBy,
		UpdateBy:        n.UpdateBy,
	}
}
