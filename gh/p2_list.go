package gh

import "github.com/cli/go-gh"

type ListProjectParams struct {
	ClientType ClientType
	Name       string
}

func ListProject(params *ListProjectParams) (*[]byte, error) {
	ghql := "query=" + GetListQuery(params.ClientType, params.Name)
	args := append(graphqlArgs, ghql)
	stdOut, _, err := gh.Exec(args...)

	if err != nil {
		return nil, err
	}

	bytes := stdOut.Bytes()

	return &bytes, nil
}
