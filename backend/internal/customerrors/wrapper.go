package customerrors

type Wrapper struct {
	Error        error  `json:"-"`
	ErrorMessage string `json:"error,omitempty"`
	Description  string `json:"description,omitempty"`
}

func NewErrorWrapper(err error, description string) *Wrapper {
	wrapper := &Wrapper{
		Error:       err,
		Description: description,
	}

	if err != nil {
		wrapper.ErrorMessage = err.Error()
	}

	return wrapper
}
