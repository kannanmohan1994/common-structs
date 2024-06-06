package commonstructs

type Step struct {
	Name     string
	Progress string
	Error    error
}
type StepsStore []Step

func (s StepsStore) StartNextStage() {
	for _, step := range s {
		if step.Progress == "not_started" {
			step.Progress = "in_progress"
			break
		}
	}
}

func (s StepsStore) FailPresentStage() {
	for _, step := range s {
		if step.Progress == "in_progress" {
			step.Progress = "failed"
			break
		}
	}
}

func (s StepsStore) CompletePresentStage() {
	for _, step := range s {
		if step.Progress == "in_progress" {
			step.Progress = "completed"
			break
		}
	}
}
