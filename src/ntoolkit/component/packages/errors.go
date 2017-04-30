package packages

// ErrLoadFailed is a generic error raised when loading a template list fails.
type ErrLoadFailed struct{
	Errors []error
}

// ErrBadFile is an error to wrap a file load error with details
type ErrBadFile struct {
}
