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

func TestIsseMutation(t *testing.T) {
	cases := []struct {
		name string
		body string
		want string
	}{
		{
			name: "contain html",
			body: "sample\n<details>\n<summary>\nSample Code</summary>\n<pre>\n<code>\necho Hello World\n</code>\n</pre></details>",
			want: `mutation {
		createIssue(input: {repositoryId: "test" title: "issueTitle" body: "sample
<details>
<summary>
Sample Code</summary>
<pre>
<code>
echo Hello World
</code>
</pre></details>" assigneeIds: [], labelIds: []}) {
    issue {
      id
    }
  }
}`,
		},
		{
			name: "double quote",
			body: `"""`,
			want: `mutation {
		createIssue(input: {repositoryId: "test" title: "issueTitle" body: "\"\"\"" assigneeIds: [], labelIds: []}) {
    issue {
      id
    }
  }
}`,
		},
		{
			name: "contain double quote and html",
			body: `test<img src="" />`,
			want: `mutation {
		createIssue(input: {repositoryId: "test" title: "issueTitle" body: "test<img src=\"\" />" assigneeIds: [], labelIds: []}) {
    issue {
      id
    }
  }
}`,
		},
	}

	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := IssueMutation("test", "issueTitle", tt.body, []string{}, []string{})

			if got != tt.want {
				t.Errorf("want = %s, but got = %s", tt.want, got)
			}
		})
	}
}
