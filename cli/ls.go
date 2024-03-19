package cli

import (
	"encoding/json"
	"fmt"
	"os"

	ghp2 "github.com/shuntaka9576/gh-p2"
)

type LsParamas struct {
	ProjectId string
	Field     string
}

func (c *Cmd) Ls(params *LsParamas) error {
	res, err := c.Client.GetProjectItems(&ghp2.GetProjectItemsParams{
		ProjectId: params.ProjectId,
	})

	if err != nil {
		return err
	}

	json, err := json.Marshal(res)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%s\n", json)

	return nil
}
