package gh

import (
	"encoding/json"

	"github.com/cli/go-gh"
)

type Label struct {
	Id          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

func CreateLabel(owner string, repo string, labelName string) (*Label, error) {
	args := CreateLabelApiArgs(owner, repo, labelName)

	stdout, _, err := gh.Exec(args...)
	if err != nil {
		return nil, err
	}

	label := &Label{}
	err = json.Unmarshal(stdout.Bytes(), label)

	if err != nil {
		return nil, err
	}

	return label, nil
}

func GetLabel(owner string, repo string, labelName string) (*Label, error) {
	args := GetLabelApiArgs(owner, repo, labelName)

	stdout, _, err := gh.Exec(args...)
	if err != nil {
		return nil, err
	}

	label := &Label{}
	err = json.Unmarshal(stdout.Bytes(), label)

	if err != nil {
		return nil, err
	}

	return label, nil
}
