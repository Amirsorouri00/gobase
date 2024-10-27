package service

import "strings"

// ErrClass is the error class.
type ErrClass int

// List of error classes.
const (
	EUnknown    ErrClass = iota // Unknown error
	EFile                       // File error
	EDB                         // Database error
	ENetwork                    // Network error
	EBadArg                     // Bad argument
	EAccess                     // Access denied
	ENotFound                   // Not found
	ETimeout                    // Operation timed out
	EConflict                   // Conflict
	EValidation                 // Validation
)

var errCLasses = map[ErrClass]string{
	EUnknown:    "unknown",
	EFile:       "file",
	EDB:         "db",
	ENetwork:    "network",
	EBadArg:     "badarg",
	EAccess:     "access",
	ENotFound:   "notfound",
	ETimeout:    "timeout",
	EConflict:   "conflict",
	EValidation: "validation",
}

// String returns the error class name.
func (e ErrClass) String() string {
	if name, ok := errCLasses[e]; ok {
		return name
	}
	return "unknown"
}

// Error is an error type that can be used to return errors from services.
type Error struct {
	Service string   `json:"service"` // Service name.
	Message string   `json:"message"` // Error message.
	Cause   error    `json:"cause"`   // Underlying error.
	Class   ErrClass `json:"class"`   // Error class.
	IsTemp  bool     `json:"is_temp"` // Is the error temporary?
	Meta    any      `json:"meta"`    // Additional data.
}

// Error returns the full error message.
func (e *Error) Error() string {
	var sb strings.Builder
	if e.Service != "" {
		sb.WriteString(e.Service + ": ")
	}
	sb.WriteString(e.Message)
	if e.Cause != nil {
		sb.WriteString("(" + e.Cause.Error() + ")")
	}
	sb.WriteString(" [" + e.Class.String() + "]")
	if e.IsTemp {
		sb.WriteString(" [temp]")
	}
	return sb.String()
}

// Unwrap returns the underlying error.
func (e *Error) Unwrap() error {
	return e.Cause
}

// IsBadArg returns true if the error is a bad argument error.
func IsBadArg(err error) bool {
	se, ok := err.(*Error)
	return ok && se.Class == EBadArg
}

// IsBadArg returns true if the error is a bad argument error.
func IsValidation(err error) bool {
	se, ok := err.(*Error)
	return ok && se.Class == EValidation
}

// IsAccess returns true if the error is an access denied error.
func IsAccess(err error) bool {
	se, ok := err.(*Error)
	return ok && se.Class == EAccess
}

// IsNotFound returns true if the error is a not found error.
func IsNotFound(err error) bool {
	se, ok := err.(*Error)
	return ok && se.Class == ENotFound
}

// IsConflict returns true if the error is a conflict error.
func IsConflict(err error) bool {
	se, ok := err.(*Error)
	return ok && se.Class == EConflict
}
