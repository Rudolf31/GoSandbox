package customeerrors

type Wrapper struct {
	Error       error `json:"error"`
	ID          int
	Description string `json:"description"`
}
