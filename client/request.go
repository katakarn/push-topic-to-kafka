package client

type GenerateTokenRequest struct {
	StudentCode string `json:"studentcode"`
	StudentID   string `json:"studentid"`
	Key         string `json:"key"`
}

type GenerateTokenRequestV2 struct {
	StudentCode string `json:"studentcode"`
	Key         string `json:"key"`
}

type GetPreRegistrationHistoryRequest struct {
	AcadYear string `json:"acadYear"`
	Semester string `json:"semester"`
}

type GetRegistrationHistoryRequest struct {
	AcadYear string `json:"acadYear"`
	Semester string `json:"semester"`
}
