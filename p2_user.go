package ghp2

import (
	"encoding/json"

	"github.com/shuntaka9576/gh-p2/gh"
)

type UserGhRes struct {
	Data map[string]struct {
		Id    string `json:"id"`
		Login string `json:"login"`
	} `json:"data"`
}

type UserRes struct {
	Users []User
}

type User struct {
	Id    string
	Login string
}

func (c *Client) GetUsers(users []string) (*UserRes, error) {
	res, err := gh.GetUser(users)
	if err != nil {
		return nil, err
	}

	parsed := &UserGhRes{}
	json.Unmarshal(*res, parsed)

	userRes := &UserRes{Users: []User{}}
	for _, user := range parsed.Data {
		userRes.Users = append(userRes.Users, User{
			Id:    user.Id,
			Login: user.Login,
		})

	}

	return userRes, nil
}
