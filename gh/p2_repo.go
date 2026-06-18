package gh

type GetRepoParams struct {
	Owner string
	Repo  string
}

func GetRepo(params *GetRepoParams) (*[]byte, error) {
	ghql := "query=" + GetRepoQuery(params.Owner, params.Repo)
	args := append(graphqlArgs, ghql)
	bytes, err := execGh(args...)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}
