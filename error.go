package dotenv

// Implements error interface for more specific package errors
type DotError struct {
	s string
}

// Returns error description
func (e *DotError) Error() string {
	return e.s
}

// Creates new environment error
func NewError(s string) *DotError {
	return &DotError{s}
}

// Line parsing error. Length of split is too short
var LPR_SHORT = NewError("Error parsing line. Lengt of slice is too short")

// Line parsing error. Length of split is too long
var LPR_LONG = NewError("Error parsing line. Length of slice is too long")

// Line parsing error. Invalid line not in correct format
var LPR_INVALID_LINE = NewError("Error parsing line. Line is composed of special invalid characters")
