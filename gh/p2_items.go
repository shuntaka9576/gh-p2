package gh

import (
	"github.com/cli/go-gh"
)

type GetProjectItemsParams struct {
	ProjectId string
	Cursor    *string
}

func GetProjectItems(params *GetProjectItemsParams) (*[]byte, error) {
	ghql := "query=" + GetProjectItemsQuery(params.ProjectId, params.Cursor)
	args := append(graphqlArgs, ghql)

	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		return nil, err
	}

	bytes := stdOut.Bytes()

	return &bytes, nil
}
