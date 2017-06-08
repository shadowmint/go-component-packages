package packages

// ErrLoadFailed is a generic error raised when loading a template list fails.
type ErrLoadFailed struct {
	Errors []error
}

// ErrBadFile is an error to wrap a file load error with details
type ErrBadFile struct {
}

// ErrDuplicateName is raised if multiple objects in the workspace have a common name.
type ErrDuplicateName struct {
}

// ErrBadTemplate is raised when try to thaw a missing or broken template
type ErrBadTemplate struct {
}
