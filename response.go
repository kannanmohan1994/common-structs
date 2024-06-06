package commonstructs

type FetchUserDetailsResponse struct {
	UserID string
}

type VerifyUserResponse struct {
	UserID string
}

type Step struct {
	Progress string
	Error    error
}

type NotifyUserResponse struct{ Steps []Step }

type MatchDriverResponse struct{}

type WaitDriverResponse struct {
	Accept bool
}
