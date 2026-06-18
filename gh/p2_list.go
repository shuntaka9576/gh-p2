package gh

type ListProjectParams struct {
	ClientType ClientType
	Name       string
}

func ListProject(params *ListProjectParams) (*[]byte, error) {
	ghql := "query=" + GetListQuery(params.ClientType, params.Name)
	args := append(graphqlArgs, ghql)
	bytes, err := execGh(args...)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}
