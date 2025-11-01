package customerrors

type Wrapper struct {
	Error       error `json:"error"`
	ID          int
	Description string `json:"description"`
}
