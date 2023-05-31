package ghp2

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shuntaka9576/gh-p2/gh"
)

type CreateIssueParams struct {
	ProjectId string
	Title     string
	Body      string
	Draft     bool
	Repo      string
	Fields    []CreateFiled
	Assignees []string
	Labels    []string
}

type CreateFiled struct {
	Id       string
	Name     string
	DataType gh.PROJECT_V2_DATA_TYPE
	Value    any
}

type CreateDraftIssueRes struct {
	Data struct {
		AddProjectV2DraftIssue struct {
			ProjectItem struct {
				Id string `json:"id"`
			} `json:"projectItem"`
		} `json:"addProjectV2DraftIssue"`
	} `json:"data"`
}

type CreateIssueRes struct {
	Data struct {
		CreateIssue struct {
			Issue struct {
				Id string `json:"id"`
			} `json:"issue"`
		} `json:"createIssue"`
	} `json:"data"`
}

type AddItemRes struct {
	Data struct {
		AddProjectV2ItemById struct {
			Item struct {
				Id string `json:"id"`
			}
		} `json:"addProjectV2ItemById"`
	} `json:"data"`
}

func (c *Client) CreateIssue(params *CreateIssueParams) (err error) {
	var assiginIds []string
	if len(params.Assignees) > 0 {
		users, err := c.GetUsers(params.Assignees)
		if err != nil {
			return err
		}

		for _, user := range users.Users {
			assiginIds = append(assiginIds, user.Id)
		}
	}

	itemId := ""
	if params.Draft {
		res, err := gh.CreateDraftIssue(&gh.CreateDraftIssueParams{
			ProjectId:  params.ProjectId,
			Title:      params.Title,
			AssiginIds: assiginIds,
		})
		if err != nil {
			return err
		}

		draftIssueResult := &CreateDraftIssueRes{}
		err = json.Unmarshal(*res, draftIssueResult)
		if err != nil {
			return err
		}
		itemId = draftIssueResult.Data.AddProjectV2DraftIssue.ProjectItem.Id
	} else {
		var labelIds []string
		for _, label := range params.Labels {
			createdLabel, err := gh.CreateLabel(c.Name, params.Repo, label)
			if err != nil {
				// TODO check status code
				fmt.Fprintf(os.Stderr, "skip create label %s (already exits)\n", label)
				getLabel, err := gh.GetLabel(c.Name, params.Repo, label)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to get label %s\n", label)
					return err
				}
				labelIds = append(labelIds, getLabel.NodeID)
			} else {
				labelIds = append(labelIds, createdLabel.NodeID)
			}
		}
		repository, err := c.GetRepo(GetRepoParams{
			Owner: c.Name,
			Repo:  params.Repo,
		})
		if err != nil {
			return err
		}

		createIssueResult, err := gh.CreateIssue(&gh.CreateIssueParams{
			RepositoryId: repository.Data.Repository.Id,
			Title:        params.Title,
			Body:         params.Body,
			AssigneeIds:  assiginIds,
			LabelIds:     labelIds,
		})
		if err != nil {
			return err
		}

		createIssueRes := &CreateIssueRes{}
		err = json.Unmarshal(*createIssueResult, createIssueRes)
		if err != nil {
			return err
		}

		contentId := createIssueRes.Data.CreateIssue.Issue.Id

		addResult, err := gh.AddItem(&gh.AddItemParams{
			ProjectId: params.ProjectId,
			ContentId: contentId,
		})
		if err != nil {
			return err
		}

		addItemRes := &AddItemRes{}
		err = json.Unmarshal(*addResult, addItemRes)
		if err != nil {
			return err
		}

		itemId = addItemRes.Data.AddProjectV2ItemById.Item.Id
	}

	// project item apply fields
	for _, field := range params.Fields {
		_, err := gh.UpdateItem(&gh.UpdateItemParams{
			ProjectId: params.ProjectId,
			ItemId:    itemId,
			FieldId:   field.Id,
			ValueType: field.DataType,
			Value:     field.Value,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
