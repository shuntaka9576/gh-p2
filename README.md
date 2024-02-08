# gh-p2

GitHub CLI extension ProjectV2 utility ✨.

![gif](https://github.com/shuntaka9576/gh-p2/blob/main/doc/gif/p2.gif?raw=true)

## Installation

Add projectV2 socpes to GitHub CLI.
```bash
gh auth login --scopes 'project'
```

Install extension.
```bash
gh extension install shuntaka9576/gh-p2
```

Upgrade:

```bash
gh extension upgrade shuntaka9576/gh-p2
```

## Usage

### Global Flags

If ProjectV2 belongs to User, use --user(-u). if it belongs to Organization, use --org(-o) to specify the name.

### Create Issue

#### Draft issue

*Quick start*

Create and add a draft issue to ProjectV2.

```bash
gh p2 create -u "ownerName" -p "projectTitle" -t "Title" -f "Status:Todo" -d
```

*Details*

|flag|short|required|default|
|---|---|---|---|
|--user or --org|-u or -o|true|""
|--project-title|-p|true|""
|--title|-t|true|""
|--draft|-r|true|false
|--assignees|-a|false|[]
|--fields|-f|false|[]
|--body|-b|false|""
|--repo|-r|unavailable|""
|--labels|-l|unavailable|[]

```bash
gh p2 create \
  --user "shuntaka9576" \
  --project-title "testProject" \
  --title "Fix bug" \
  --draft \
  --assignees "shuntaka9576" \
  --fields "Status:Todo" \
  --fields "point:3" \
  --fields "deadline:2022-08-11" \
```
#### Issue

*Quick start*

Create an issue in `{ownerName}/{repositoryName}` repository and add it to ProjectV2.

```bash
gh p2 create -u "ownerName" -r "repositoryName" -p "projectTitle" -t "Title"
```

*Details*

|flag|short|required|default|
|---|---|---|---|
|--user or --org|-u or -o|true|""
|--repo|-r|true|""
|--project-title|-p|true|""
|--title|-t|true|""
|--draft|-r|false|false
|--assignees|-a|false|[]
|--labels|-l|false|[]
|--fields|-f|false|[]
|--body|-b|false|""

```bash
gh p2 create \
  --user "shuntaka9576" \
  --repo "repositoryName" \
  --project-tilte "testProject" \
  --title "Fix bug" \
  --assignees "shuntaka9576" \
  --labels "label1" \
  --labels "label2" \
  --fields "Status:Todo" \
  --fields "point:3" \
  --fields "deadline:2022-08-11"
```

### Show Project

#### List Issues with Custom Fields

Retrieve a list of issues linked to `Single Select Custom Fields` in GitHub Projects.

*Quick start*

```bash
gh p2 show -u "ownerName" -r "repositoryName" -p "projectTitle" -f "Status:TODO"
```

*Details*

|flag|short|required|default|
|---|---|---|---|
|--user or --org|-u or -o|true|""|
|--repo|-r|true|""|
|--project-title|-p|true|""|
|--filter|-f|true|""|

```bash
gh p2 show \
  --user "shuntaka9576" \
  --project-title "testProject" \
  --repo "repositoryName" \
  --filter "Status:TODO"
```

## Special Thanks

* https://github.com/yusukebe/gh-markdown-preview
