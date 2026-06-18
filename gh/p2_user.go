package gh

import (
	"github.com/cli/go-gh/v2"
)

func GetUser(users []string) (*[]byte, error) {
	ghql := "query=" + GetUserQuery(users)
	args := append(graphqlArgs, ghql)
	stdOut, _, err := gh.Exec(args...)

	if err != nil {
		return nil, err
	}

	bytes := stdOut.Bytes()

	return &bytes, nil
}
