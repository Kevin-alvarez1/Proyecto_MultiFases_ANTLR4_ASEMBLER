package visitor

type EvaluationError struct {
	Message string
	Line    int
	Column  int
}

func (e *EvaluationError) Error() string {
	return e.Message
}
