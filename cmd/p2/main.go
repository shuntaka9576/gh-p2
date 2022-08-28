package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	ghp2 "github.com/shuntaka9576/gh-p2"
	"github.com/shuntaka9576/gh-p2/cli"
)

type Globals struct {
	OrgName  string          `short:"o" name:"org" help:"Specify an organization name."`
	UserName string          `short:"u" name:"user" help:"Specify a user name."`
	Version  cli.VersionFlag `short:"v" name:"version" help:"print the version."`
}

var CLI struct {
	Globals
	Create struct {
		ProjectTitle string   `short:"p" required:"" name:"project-title" help:"Specify the title of ProjectV2."`
		Title        string   `short:"t" required:"" name:"title" help:"Specify issue title."`
		Repo         string   `short:"r" name:"repo" help:"Specify the repository name. Owner name is not required. This flag is not available when creating draft issues."`
		Fields       []string `short:"f" name:"fields" help:"Specify ProjectV2 custom fields in the format {keyName}:{valueName}. e.g. Status:Todo, Point:3, date:2022-08-29. See https://docs.github.com/ja/graphql/reference/input-objects#projectv2fieldvalue for data types. Iteration is not currently supported."`
		Labels       []string `short:"l" name:"labels" help:"Specify the label name to be set for the issue. If a label with the target name does not exist, a new one will be created with a random color. This flag is not available when creating draft issues."`
		Draft        bool     `short:"d" name:"draft" help:"Due to GitHub specifications, the --label and --repo options cannot be used together."`
		Assignees    []string `short:"a" name:"assignees" help:"Specify the GitHub account ID to be assigned."`
	} `cmd:"" help:"Option to create an issue or draft issue directly in Project V2."`
}

func main() {
	kontext := kong.Parse(&CLI,
		kong.Name("p2"),
	)

	client, err := ghp2.InitClient(&ghp2.InitParams{
		OrgName:  &CLI.OrgName,
		UserName: &CLI.UserName,
	})

	if errors.Is(err, ghp2.ErrorInvalidSpecifyClientType) {
		fmt.Fprintf(os.Stderr, "Please set either --org(-o) or --user(-u) flag. Specify whether the GitHub Project you want to retrieve belongs to organization or user.")

		os.Exit(1)
	}

	c := cli.Cmd{
		Client: *client,
	}

	res, err := c.Client.ListProject()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)

		os.Exit(1)
	}

	var projectId string
	for _, project := range res.Projects() {
		if project.Title == CLI.Create.ProjectTitle {
			projectId = project.Id
		}
	}

	if projectId == "" {
		fmt.Fprintf(os.Stderr, "Not found project name: %s\n", CLI.Create.ProjectTitle)
		if len(res.Projects()) > 0 {
			fmt.Fprintf(os.Stderr, "The following project names are available and can be specified in the project-title(-p) flag.\n")
			for _, project := range res.Projects() {
				fmt.Fprintf(os.Stderr, "  * %s\n", project.Title)
			}
		} else {
			fmt.Fprintf(os.Stderr, "There are no ProjectV2 resources available for this organization or user.\n")
		}

		os.Exit(1)
	}

	switch kontext.Command() {
	case "create":
		err = c.Create(&cli.CreateParamas{
			ProjectId: projectId,
			Title:     CLI.Create.Title,
			Repo:      CLI.Create.Repo,
			Fields:    CLI.Create.Fields,
			Labels:    CLI.Create.Labels,
			Draft:     CLI.Create.Draft,
			Assignees: CLI.Create.Assignees,
		})
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)

		os.Exit(1)
	}

	os.Exit(0)
}
