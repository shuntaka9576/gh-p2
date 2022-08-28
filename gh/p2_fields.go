package gh

import (
	"github.com/cli/go-gh"
)

type GetProjectFiledsParams struct {
	ProjectId string
}

func GetProjectFields(params *GetProjectFiledsParams) (*[]byte, error) {
	ghql := "query=" + GetProjectFieldsQuery(params.ProjectId)
	args := append(graphqlArgs, ghql)

	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		return nil, err
	}

	bytes := stdOut.Bytes()

	return &bytes, nil
}
