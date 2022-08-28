package gh

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateLabelApiArgs(t *testing.T) {
	wants := []string{"label", "create", "labelName", "--repo", "repoName", "--color"}
	gots := CreateLabelApiArgs("repoName", "labelName")

	if d := cmp.Diff(gots[:6], wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}

	if gots[len(gots)-1] == "" {
		t.Errorf("error color is empty: %s", gots[len(gots)-1])
	}
}
