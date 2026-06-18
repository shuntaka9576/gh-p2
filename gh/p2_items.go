package gh

type GetProjectItemsParams struct {
	ProjectId string
	Cursor    *string
}

func GetProjectItems(params *GetProjectItemsParams) (*[]byte, error) {
	ghql := "query=" + GetProjectItemsQuery(params.ProjectId, params.Cursor)
	args := append(graphqlArgs, ghql)

	bytes, err := execGh(args...)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}
