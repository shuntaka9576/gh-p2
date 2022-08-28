package ghp2

import (
	"encoding/json"

	"github.com/shuntaka9576/gh-p2/gh"
)

type GetProjectFieldsParams struct {
	ProjectId string
}

type GetProjectFiledsGhRes struct {
	Data struct {
		Node struct {
			Fileds struct {
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

type GetProjectFiledsRes struct {
	Fileds []Filed
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

func (c *Client) GetProjectFields(params *GetProjectFieldsParams) (*GetProjectFiledsRes, error) {
	payload, err := gh.GetProjectFields(&gh.GetProjectFiledsParams{
		ProjectId: params.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	parsed := &GetProjectFiledsGhRes{}
	err = json.Unmarshal(*payload, parsed)
	if err != nil {
		return nil, err
	}

	res := &GetProjectFiledsRes{}

	for _, node := range parsed.Data.Node.Fileds.Nodes {
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

		res.Fileds = append(res.Fileds, filed)
	}

	return res, nil
}
