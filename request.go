package commonstructs

type FetchUserDetailsRequest struct {
	UserID string
}

type VerifyUserRequest struct {
	UserID string
}

type NotifyUserRequest struct {
	UserID string
	Steps  StepsStore
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
