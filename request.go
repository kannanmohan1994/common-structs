package commonstructs

type FetchUserDetailsRequest struct {
	UserID string
}

type VerifyUserRequest struct {
	UserID string
}

type NotifyUserRequest struct {
	PhoneNumber string
	Steps       []string
}

type WaitDriverRequest struct {
	UserID       string
	UserLocation string
}

type MatchDriverRequest struct {
	UserID       string
	UserLocation string
}

type AssignDriverWorkflowRequest struct {
	UserID   string
	Location string
}
