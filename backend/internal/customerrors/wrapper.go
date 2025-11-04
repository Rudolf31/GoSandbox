package customerrors

type Wrapper struct {
	Error       error  `json:"error"`
	Description string `json:"description"`
}
