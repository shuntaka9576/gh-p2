package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	ghp2 "github.com/shuntaka9576/gh-p2"
	"github.com/shuntaka9576/gh-p2/gh"
)

type CreateParamas struct {
	ProjectId        string
	SetProjectStatus string
	Title            string
	Body             string
	Repo             string
	Fields           []string
	Labels           []string
	Draft            bool
	Assignees        []string
}

func (c *Cmd) Create(params *CreateParamas) error {
	// check create draft issue flags
	if params.Draft {
		if params.Repo != "" || len(params.Labels) > 0 {
			fmt.Fprintf(os.Stderr, "If you are creating a Draft issue, you cannot use the --repo or --labels flag.\n")

			return errors.New("exec create error: error specifying flags when creating draft")
		}
	}

	// get project status item
	projectFields, err := c.Client.GetProjectFields(&ghp2.GetProjectFieldsParams{
		ProjectId: params.ProjectId,
	})
	if err != nil {
		return fmt.Errorf("exec create error: %s", err)
	}

	updateFields := []ghp2.CreateFiled{}
	for _, updateFiled := range params.Fields {
		updateFiledName := strings.Split(updateFiled, ":")[0]
		updateFiledValue := strings.Split(updateFiled, ":")[1]

		for _, filed := range projectFields.Fields {
			if updateFiledName == filed.Name {
				// TODO ghp2.ITERATION
				if filed.DataType == gh.SINGLE_SELECT {
					for _, option := range filed.Options {
						if updateFiledValue == option.Name {
							updateFields = append(updateFields, ghp2.CreateFiled{
								Id:       filed.Id,
								Name:     filed.Name,
								DataType: filed.DataType,
								Value:    option.Id,
							})
						}
					}
				} else {
					updateFields = append(updateFields, ghp2.CreateFiled{
						Id:       filed.Id,
						Name:     filed.Name,
						DataType: filed.DataType,
						Value:    updateFiledValue,
					})
				}
			}
		}
	}

	err = c.Client.CreateIssue(&ghp2.CreateIssueParams{
		ProjectId: params.ProjectId,
		Title:     params.Title,
		Body:      params.Body,
		Draft:     params.Draft,
		Fields:    updateFields,
		Repo:      params.Repo,
		Assignees: params.Assignees,
		Labels:    params.Labels,
	})

	if err != nil {
		return err
	}

	return nil
}
