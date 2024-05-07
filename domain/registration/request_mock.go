package registration

import (
	"context"
)

type RegistrationRequestMock struct {
	SubmitRegistrationError  error
	GetEventResult           []EventResponse
	GetEventError            error
	GetCoursesResult         GetCoursesResponse
	GetCoursesError          error
	GetEnrollSemesterResult  Semester
	GetEnrollSemesterError   error
	GetCoursesSectionsResult GetCourseSection
	GetCoursesSectionsError  error
}

func (r *RegistrationRequestMock) CheckAttemptForRegistration(ctx context.Context, studentID int) error {
	return r.SubmitRegistrationError
}

func (r *RegistrationRequestMock) GetEvent(ctx context.Context, schedulegroupID int, studentYear int) ([]EventResponse, error) {
	return r.GetEventResult, r.GetEventError
}

func (r *RegistrationRequestMock) GetCourses(studentID int, studentCode string, currentAcadYear string, currentSemester string, search string) (GetCoursesResponse, error) {
	return r.GetCoursesResult, r.GetCoursesError
}

func (r *RegistrationRequestMock) GetEnrollSemester(schedulegroupID int) (*Semester, error) {
	return &r.GetEnrollSemesterResult, r.SubmitRegistrationError
}

func (r *RegistrationRequestMock) GetSuggestionCourses(acadYear int64, semester int64) (GetCoursesResponse, error) {
	return r.GetCoursesResult, r.GetCoursesError
}

func (r *RegistrationRequestMock) GetCoursesSections(courseId string, search string, filter string, currentAcadYear int64, currentSemester int64, studentID int, studentCode string, facultyID int) (GetCourseSection, error) {
	return r.GetCoursesSectionsResult, r.GetCoursesSectionsError
}
