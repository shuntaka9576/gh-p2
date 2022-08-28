package ghp2

import (
	"github.com/shuntaka9576/gh-p2/gh"
)

type Client struct {
	ClientType gh.ClientType
	Name       string
}

type InitParams struct {
	OrgName  *string
	UserName *string
}

func InitClient(params *InitParams) (*Client, error) {
	if (*params.OrgName != "" && *params.UserName != "") || (*params.OrgName == "" && *params.UserName == "") {
		return nil, ErrorInvalidSpecifyClientType
	}

	if *params.OrgName != "" {
		return &Client{
			ClientType: gh.ORGANIZATION,
			Name:       *params.OrgName,
		}, nil
	}

	if *params.UserName != "" {
		return &Client{
			ClientType: gh.USER,
			Name:       *params.UserName,
		}, nil
	}

	return nil, ErrorInvalidSpecifyClientType
}
