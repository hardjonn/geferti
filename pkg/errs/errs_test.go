package errs

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoArgumentsShouldPanic(t *testing.T) {
	defer panicRecover(t, "E() did not panic with not arguments")

	_ = E()
}

func TestBadArgumentShouldPanic(t *testing.T) {
	defer panicRecover(t, "E() did not panic with bad argument")

	_ = E(0)
}

func TestNilShouldReturnNoError(t *testing.T) {
	e := E(nil)

	assert.Equal(t, e.Error(), "no error", "nil should return 'no error'")
}

func TestStringShouldBeConvertedToError(t *testing.T) {
	errText := "error"
	e := E(errText)
	r := errors.New(errText)

	assert.Equal(t, e.(*Error).Err, r, "string should be handled properly and be converted to error")
}

func TestErrorShouldBeHandledAsError(t *testing.T) {
	err := errors.New("error")
	e := E(err)

	assert.Equal(t, e.Error(), err.Error(), "error message should match")
	assert.Equal(t, e.(*Error).Err, err, "error type should match")
}

func TestNestedError(t *testing.T) {
	err := errors.New("error")
	e1 := E(err)
	e2 := E(e1)
	e3 := E(e2)

	assert.Equal(t, e1.Error(), err.Error(), "level 1 error message should match")
	assert.Equal(t, e2.Error(), e1.Error(), "level 2 error message should match")
	assert.Equal(t, e3.Error(), e2.Error(), "level 3 error message should match")
}

func TestLastArgumentHasPriority(t *testing.T) {
	err := errors.New("error")
	oErr := errors.New("overridden error")
	e1 := E(err)
	e2 := E(e1, oErr)

	assert.Equal(t, e2.Error(), oErr.Error(), "last argument has priority and should override the previous one")
}

type statusTest struct {
	err    error
	status ErrStatus
}

func TestStatuses(t *testing.T) {

	e1 := E("nested error should return the first/initial status", StatusInvalid)
	e2 := E(e1, StatusUnexpected)
	e3 := E(e2, StatusNotFound)

	statusTests := []statusTest{
		{E(nil), StatusOther},
		{nil, StatusOther},
		{E("conflict error -- status has been setup manually", StatusConflict), StatusConflict},
		{E("internal error -- status not specified"), StatusInternal},
		{E("internal error with explicit Other status", StatusOther), StatusInternal},
		{e3, StatusInvalid},
	}

	for _, test := range statusTests {
		assert.Equal(t, Status(test.err), test.status, fmt.Sprintf("error [%v] should have status [%s]", test.err, test.status))
	}
}

func TestNestedOps(t *testing.T) {
	e1 := E(Op("op1"), "error")
	e2 := E(Op("op2"), e1)
	e3 := E(Op("op3"), e2)

	actualOps := Ops(e3)
	expectedOps := []Op{"op1", "op2", "op3"}

	assert.Equal(t, actualOps, expectedOps, "nested ops should start from the first one")
}

func TestNestedOpsOverriding(t *testing.T) {
	e1 := E(Op("op1"), "error 1")
	e2 := E(Op("op2"), e1, "error 2")

	actualOps := Ops(e2)
	expectedOps := []Op{"op2"}

	assert.Equal(t, actualOps, expectedOps, "recent extra error should override previous ops")
}

type detailsTest struct {
	err     error
	details []byte
}

func TestDetails(t *testing.T) {
	// detailsTests := []detailsTest{
	// 	{}
	// }

	e1 := E(Op("op1"), "error invalid", StatusInvalid)
	e2 := E(Op("op2"), e1)

	actualDetails := Details(e2)
	expectedDetails := &ErrDetails{
		Ops:    Ops(e2),
		Code:   Status(e2),
		Status: Status(e2).String(),
	}

	rawDetails := Raw(e2)
	expectedRawDetails, _ := json.Marshal(expectedDetails)

	assert.Equal(t, actualDetails, expectedDetails, "error details should match")
	assert.Equal(t, rawDetails, expectedRawDetails, "error raw details should match")
}

func panicRecover(t *testing.T, message string) {
	err := recover()
	if err == nil {
		t.Fatal(message)
	}
}
