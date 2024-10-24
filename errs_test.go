package errs

import (
	"testing"
)

type erTest struct{}

func (e *erTest) Error() string {
	return ""
}

func getETest() error {
	return &erTest{}
}

func TESTNilError(t *testing.T) {
	err := E(Op("123"), getETest(), Unauthorized)
	if err != nil {
		t.Error("NOT NIL")
	}
}
