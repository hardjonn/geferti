package errs

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Error main error struct.
type Error struct {
	Op     Op
	Status ErrStatus
	Err    error
}

// ErrDetails stands for logging in JSON format.
type ErrDetails struct {
	Ops    []Op      `json:"ops"`
	Code   ErrStatus `json:"code"`
	Status string    `json:"status"`
}

// Op stands for operation.
type Op string

// E constructs the error
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errs.E with no arguments")
	}
	e := &Error{}

	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case ErrStatus:
			e.Status = arg
		case *Error:
			copy := *arg
			e.Err = &copy
		case error:
			e.Err = arg
		case string:
			e.Err = errors.New(arg)
		case nil:
			e.Err = nil
		default:
			panic("bad call to errs.E handler")
		}
	}

	return e
}

func (e *Error) Error() string {
	b := new(bytes.Buffer)

	if e.Err != nil {
		b.WriteString(e.Err.Error())
	}

	if b.Len() == 0 {
		return "no error"
	}

	return b.String()
}

// Ops combines all the ops into one slice starting from the first one.
func Ops(err error) []Op {
	res := []Op{}

	for e, ok := err.(*Error); ok; e, ok = e.Err.(*Error) {
		if e.Op != "" {
			res = append([]Op{e.Op}, res...)
		}
	}

	return res
}

// Status finds the first error code.
func Status(err error) ErrStatus {
	if err == nil || err.(*Error).Err == nil {
		return StatusOther
	}

	status := StatusInternal

	for e, ok := err.(*Error); ok; e, ok = e.Err.(*Error) {
		if e.Status != StatusOther {
			status = e.Status
		}
	}

	return status
}

// Details generates JSON output.
func Details(err error) *ErrDetails {
	status := Status(err)

	return &ErrDetails{
		Ops:    Ops(err),
		Code:   status,
		Status: status.String(),
	}
}

// Raw returns raw json.
func Raw(err error) []byte {
	details := Details(err)

	r, _ := json.Marshal(details)
	return r
}
