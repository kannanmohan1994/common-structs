package commonstructs

type Step struct {
	Name     string
	Progress string
	Error    error
}
type StepsStore []Step

func (s StepsStore) StartNextStage() {
	for idx := range s {
		if s[idx].Progress == "in_progress" {
			s[idx].Progress = "completed"
		}
		if s[idx].Progress == "not_started" {
			s[idx].Progress = "in_progress"
			break
		}
	}
}

func (s StepsStore) FailPresentStage() {
	for idx := range s {
		if s[idx].Progress != "completed" {
			s[idx].Progress = "failed"
			break
		}
	}
}

func (s StepsStore) CompletePresentStage() {
	for idx := range s {
		if s[idx].Progress == "in_progress" {
			s[idx].Progress = "completed"
			break
		}
	}
}

func (s StepsStore) MoveToStep(step int) {
	for idx := range s {
		if idx < step {
			s[idx].Progress = "completed"
		} else {
			s[idx].Progress = "not_started"
		}
	}
}
