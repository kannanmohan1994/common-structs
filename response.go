package commonstructs

type FetchUserDetailsResponse struct {
	UserID string
}

type VerifyUserResponse struct {
	UserID string
}

type NotifyUserResponse struct{}

type MatchDriverResponse struct{}

type WaitDriverResponse struct {
	Accept bool
}

type AssignDriverWorkflowResponse struct {
	UserID string
}
