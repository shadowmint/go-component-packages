package packages

// ErrLoadFailed is a generic error raised when loading a template list fails.
type ErrLoadFailed struct{
	Errors []error
}
