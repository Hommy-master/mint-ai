package util

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	out, errOut, err := Exec("netstat", "-na")
	if err != nil {
		t.Errorf("exec failed: %+v, errOut: %s", err, errOut)
		return
	}

	fmt.Printf("out: %s, outErr: %s\n", out, errOut)
}
