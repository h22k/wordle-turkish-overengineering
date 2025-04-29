package validator

type InputValidator interface {
	Validate(v any) error
}
