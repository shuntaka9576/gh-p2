package ghp2

import (
	"encoding/json"

	"github.com/shuntaka9576/gh-p2/gh"
)

type GetProjectItemsParams struct {
	ProjectId string
}

type GetProjectItemsGhRes struct {
	Data struct {
		Node struct {
			Title string `json:"title"`
			Items struct {
				PageInfo struct {
					HasNextPage bool   `json:"hasNextPage"`
					EndCursor   string `json:"endCursor"`
				} `json:"pageInfo"`
				Nodes []struct {
					ID         string `json:"id"`
					CreatedAt  string `json:"createdAt"`
					UpdatedAt  string `json:"updatedAt"`
					IsArchived bool   `json:"isArchived"`
					Content    struct {
						TypeName string `json:"__typename,omitempty"`
						Number   int    `json:"number,omitempty"`
						Title    string `json:"title,omitempty"`
						State    string `json:"state,omitempty"`
						URL      string `json:"url,omitempty"`
						Body     string `json:"body,omitempty"`
					} `json:"content"`
					FieldValues struct {
						Nodes []struct {
							TypeName string `json:"__typename,omitempty"`
							Name     string `json:"name,omitempty"`
							Filed    struct {
								Name string `json:"name,omitempty"`
							} `json:"field"`
						} `json:"nodes"`
					} `json:"fieldValues"`
				} `json:"nodes"`
			} `json:"items"`
		} `json:"node"`
	} `json:"data"`
}

type GetProjectItemsRes struct {
	Title      string      `json:"title"`
	IssueItems []IssueItem `json:"items"`
}

type IssueItem struct {
	Title              string            `json:"title"`
	SingleSelectValues map[string]string `json:"singleFiledValues"`
	ItemType           gh.ITEM_TYPE      `json:"type"`
	Body               string            `json:"body"`
	URL                string            `json:"url"`
	Number             int               `json:"number"`
}

func (c *Client) GetProjectItems(params *GetProjectItemsParams) (*GetProjectItemsRes, error) {
	var cursor *string
	res := &GetProjectItemsRes{}

	for {
		payload, err := gh.GetProjectItems(&gh.GetProjectItemsParams{
			ProjectId: params.ProjectId,
			Cursor:    cursor,
		})
		if err != nil {
			return nil, err
		}

		parsed := &GetProjectItemsGhRes{}
		err = json.Unmarshal(*payload, parsed)
		if err != nil {
			return nil, err
		}

		if res.Title == "" { // Set the title only once
			res.Title = parsed.Data.Node.Title
		}

		for _, node := range parsed.Data.Node.Items.Nodes {
			if node.IsArchived {
				continue
			}

			item := IssueItem{
				Title:  node.Content.Title,
				Body:   node.Content.Body,
				URL:    node.Content.URL,
				Number: node.Content.Number,
			}

			switch node.Content.TypeName {
			case "DraftIssue":
				item.ItemType = gh.DRAFT_ISSUE
			case "Issue":
				item.ItemType = gh.ISSUE
			default:
				return nil, err
			}

			singleSelectValues := map[string]string{}
			for _, fieldValue := range node.FieldValues.Nodes {
				if fieldValue.TypeName == "ProjectV2ItemFieldSingleSelectValue" {
					singleSelectValues[fieldValue.Filed.Name] = fieldValue.Name
				}
			}
			item.SingleSelectValues = singleSelectValues

			res.IssueItems = append(res.IssueItems, item)
		}

		if !parsed.Data.Node.Items.PageInfo.HasNextPage {
			break
		}
		cursor = &parsed.Data.Node.Items.PageInfo.EndCursor
	}

	return res, nil
}
