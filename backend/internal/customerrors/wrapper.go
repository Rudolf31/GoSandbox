package customerrors

type Wrapper struct {
	//TODO: Remove id from Wrapper
	Error       error `json:"error"`
	ID          int
	Description string `json:"description"`
}
