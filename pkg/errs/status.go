package errs

// ErrStatus ...
type ErrStatus uint8

// List of status codes
const (
	StatusOther      ErrStatus = iota // Unclassified error or nil error
	StatusUnexpected                  // Unexpected error
	StatusConflict                    // Action cannot be performed.
	StatusInternal                    // Internal error.
	StatusNotFound                    // Entity does not exist.
	StatusInvalid                     // validation failed
	StatusIO                          // I/O error
)

var statusText = map[ErrStatus]string{
	StatusOther:      "Other",
	StatusUnexpected: "Unexpected",
	StatusConflict:   "Conflict",
	StatusInternal:   "Internal",
	StatusNotFound:   "Not Found",
	StatusInvalid:    "Invalid",
	StatusIO:         "IO",
}

func (c ErrStatus) String() string {
	return statusText[c]
}
