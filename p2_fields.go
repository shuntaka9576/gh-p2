package ghp2

import (
	"encoding/json"

	"github.com/shuntaka9576/gh-p2/gh"
)

type GetProjectFieldsParams struct {
	ProjectId string
}

type GetProjectFieldsGhRes struct {
	Data struct {
		Node struct {
			Fields struct {
				Nodes []struct {
					Id       string `json:"id"`
					Name     string `json:"name"`
					DataType string `json:"dataType"`
					Options  []struct {
						Id   string `json:"id"`
						Name string `json:"name"`
					} `json:"options"`
				} `json:"nodes"`
			} `json:"fields"`
		} `json:"node"`
	} `json:"data"`
}

type GetProjectFieldsRes struct {
	Fields []Filed
}

type Filed struct {
	Id       string
	DataType gh.PROJECT_V2_DATA_TYPE
	Name     string
	Options  []struct {
		Id   string
		Name string
	}
}

func (c *Client) GetProjectFields(params *GetProjectFieldsParams) (*GetProjectFieldsRes, error) {
	payload, err := gh.GetProjectFields(&gh.GetProjectFieldsParams{
		ProjectId: params.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	parsed := &GetProjectFieldsGhRes{}
	err = json.Unmarshal(*payload, parsed)
	if err != nil {
		return nil, err
	}

	res := &GetProjectFieldsRes{}

	for _, node := range parsed.Data.Node.Fields.Nodes {
		filed := Filed{
			Id:       node.Id,
			DataType: node.DataType,
			Name:     node.Name,
		}

		if len(node.Options) > 0 {
			filed.Options = []struct {
				Id   string
				Name string
			}(node.Options)
		}

		res.Fields = append(res.Fields, filed)
	}

	return res, nil
}
