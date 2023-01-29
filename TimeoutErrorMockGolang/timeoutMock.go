package main

type TimeoutError struct {
	err     string
	timeout bool
}

func (e *TimeoutError) Error() string {
	return e.err
}
func (e *TimeoutError) Timeout() bool {
	return e.timeout
}

func (e *TimeoutError) Temporary() bool {
	return true
}

func (e *TimeoutError) SetTimeout(timeout bool) {
	e.timeout = timeout
}

func (e *TimeoutError) SetError(err string) {
	e.err = err
}

func CreateTimeoutErrorMock() *TimeoutError {
	timeoutError := &TimeoutError{}
	timeoutError.SetTimeout(true)
	timeoutError.SetError("timeout error")
	return timeoutError
}
