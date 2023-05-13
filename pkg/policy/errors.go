package policy

type EvaluationFailure struct {
	err error
}

func (e *EvaluationFailure) Error() string {
	return e.err.Error()
}

// ThrowError returns the error type that is passed. It acts as a
// single consistent way of throwing errors from the pkg level.
// As long as the error type is defined before hand, this can be
// used to throw errors up to the calling code without discrimination
// around the error type.
func ThrowError(err interface{ error }) error {
	return err
}
