package ghp2

import (
	"encoding/json"

	"github.com/shuntaka9576/gh-p2/gh"
)

type GetRepoRes struct {
	Data struct {
		Repository struct {
			Id string `json:"id"`
		} `json:"repository"`
	} `json:"data"`
}
type GetRepoParams struct {
	Owner string
	Repo  string
}

func (c *Client) GetRepo(params GetRepoParams) (*GetRepoRes, error) {
	ghRes, err := gh.GetRepo(&gh.GetRepoParams{
		Owner: c.Name,
		Repo:  params.Repo,
	})

	if err != nil {
		return nil, err
	}

	res := &GetRepoRes{}

	err = json.Unmarshal(*ghRes, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
