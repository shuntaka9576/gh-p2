package ghp2

import (
	"encoding/json"

	"github.com/shuntaka9576/gh-p2/gh"
)

type ListProjectResItem struct {
	Id    string
	Title string
}

type ListProjectItemList = []ListProjectResItem

type ListProjectRes interface {
	Projects() ListProjectItemList
}

type ListProjectOrgRes struct {
	Data struct {
		Organization struct {
			ProjectsV2 struct {
				Nodes []struct {
					Id    string `json:"id"`
					Title string `json:"title"`
				} `json:"nodes"`
			} `json:"projectsV2"`
		} `json:"organization"`
	} `json:"data"`
}

type ListProjectUserRes struct {
	Data struct {
		User struct {
			ProjectsV2 struct {
				Nodes []struct {
					Id    string `json:"id"`
					Title string `json:"title"`
				} `json:"nodes"`
			} `json:"projectsV2"`
		} `json:"user"`
	} `json:"data"`
}

func (l ListProjectOrgRes) Projects() (res ListProjectItemList) {
	for _, pj := range l.Data.Organization.ProjectsV2.Nodes {
		res = append(res, ListProjectResItem{
			Id:    pj.Id,
			Title: pj.Title,
		})
	}

	return res
}

func (l ListProjectUserRes) Projects() (res ListProjectItemList) {
	for _, pj := range l.Data.User.ProjectsV2.Nodes {
		res = append(res, ListProjectResItem{
			Id:    pj.Id,
			Title: pj.Title,
		})
	}

	return res
}

func (c *Client) ListProject() (ListProjectRes, error) {
	res, err := gh.ListProject(&gh.ListProjectParams{
		ClientType: c.ClientType,
		Name:       c.Name,
	})
	if err != nil {
		return nil, err
	}

	var project ListProjectRes
	if c.ClientType == gh.ORGANIZATION {
		org := ListProjectOrgRes{}
		err = json.Unmarshal(*res, &org)
		if err != nil {
			return nil, err
		}
		project = org
	} else {
		user := ListProjectUserRes{}
		err = json.Unmarshal(*res, &user)
		if err != nil {
			return nil, err
		}
		project = user
	}

	return project, nil
}
