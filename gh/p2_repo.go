package gh

import (
	"github.com/cli/go-gh"
)

type GetRepoParams struct {
	Owner string
	Repo  string
}

func GetRepo(params *GetRepoParams) (*[]byte, error) {
	ghql := "query=" + GetRepoQuery(params.Owner, params.Repo)
	args := append(graphqlArgs, ghql)
	stdOut, _, err := gh.Exec(args...)

	if err != nil {
		return nil, err
	}

	byte := stdOut.Bytes()

	return &byte, nil
}
