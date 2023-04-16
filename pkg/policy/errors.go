package policy

type EvaluationFailure struct {
	err error
}

func (e *EvaluationFailure) Error() string {
	return e.err.Error()
}

func ThrowError(err interface{ error }) error {
	return err
}
