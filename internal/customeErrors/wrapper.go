package customeerrors

type Wrapper struct {
	Error       error
	ID          int
	Description string
}
