package gh

func GetUser(users []string) (*[]byte, error) {
	ghql := "query=" + GetUserQuery(users)
	args := append(graphqlArgs, ghql)
	bytes, err := execGh(args...)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}
