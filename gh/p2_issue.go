package gh

import (
	"github.com/cli/go-gh"
)

type CreateDraftIssueParams struct {
	ProjectId  string
	Title      string
	AssiginIds []string
}

func CreateDraftIssue(params *CreateDraftIssueParams) (*[]byte, error) {
	ghql := "query=" + GetDraftIssueMutation(params.ProjectId, params.Title, "", params.AssiginIds)
	args := append(graphqlArgs, ghql)
	stdOut, _, err := gh.Exec(args...)

	if err != nil {
		return nil, err
	}

	byte := stdOut.Bytes()

	return &byte, nil
}

type CreateIssueParams struct {
	RepositoryId string
	Title        string
	Body         string
	LabelIds     []string
	AssigneeIds  []string
}

func CreateIssue(params *CreateIssueParams) (*[]byte, error) {
	ghql := "query=" + IssueMutation(params.RepositoryId, params.Title, params.Body, params.AssigneeIds, params.LabelIds)
	args := append(graphqlArgs, ghql)
	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		return nil, err
	}

	byte := stdOut.Bytes()

	return &byte, nil

}

type AddItemParams struct {
	ProjectId string
	ContentId string
}

func AddItem(params *AddItemParams) (*[]byte, error) {
	ghql := "query=" + AddItemMutation(params.ProjectId, params.ContentId)
	args := append(graphqlArgs, ghql)
	stdOut, _, err := gh.Exec(args...)

	if err != nil {
		return nil, err
	}

	bytes := stdOut.Bytes()

	return &bytes, nil
}

type UpdateItemParams struct {
	ProjectId string
	ItemId    string
	FieldId   string
	ValueType PROJECT_V2_DATA_TYPE
	Value     any
}

func UpdateItem(params *UpdateItemParams) (*[]byte, error) {
	ghql := "query=" + UpdateItemMutation(params.ProjectId, params.ItemId, params.FieldId, params.ValueType, params.Value)
	args := append(graphqlArgs, ghql)
	stdOut, _, err := gh.Exec(args...)

	if err != nil {
		return nil, err
	}

	bytes := stdOut.Bytes()

	return &bytes, nil
}
