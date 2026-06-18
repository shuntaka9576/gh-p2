package gh

type GetProjectFieldsParams struct {
	ProjectId string
}

func GetProjectFields(params *GetProjectFieldsParams) (*[]byte, error) {
	ghql := "query=" + GetProjectFieldsQuery(params.ProjectId)
	args := append(graphqlArgs, ghql)

	bytes, err := execGh(args...)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}
