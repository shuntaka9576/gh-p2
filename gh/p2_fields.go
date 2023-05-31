package gh

import (
	"github.com/cli/go-gh"
)

type GetProjectFieldsParams struct {
	ProjectId string
}

func GetProjectFields(params *GetProjectFieldsParams) (*[]byte, error) {
	ghql := "query=" + GetProjectFieldsQuery(params.ProjectId)
	args := append(graphqlArgs, ghql)

	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		return nil, err
	}

	bytes := stdOut.Bytes()

	return &bytes, nil
}
