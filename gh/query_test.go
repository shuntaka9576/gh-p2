package gh

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateLabelApiArgs(t *testing.T) {
	wants := []string{
		"api",
		"--method",
		"POST",
		"/repos/repoName/labelName/labels",
		"-f",
		"name=sample",
	}

	gots := CreateLabelApiArgs("repoName", "labelName", "sample")

	if d := cmp.Diff(gots[:6], wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}

	if gots[len(gots)-1] == "" {
		t.Errorf("error color is empty: %s", gots[len(gots)-1])
	}
}
