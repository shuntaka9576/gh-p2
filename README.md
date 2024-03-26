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

### Ls Issue

*Quick start*

List an issue in project.

```bash
# To specify an organization, use -o instead of -u.
gh p2 ls -u "ownerName" -p "project"  |\
  jq -r '["type", "number", "title", "status", "url", "point", "labels"], (.items[] |[.type, .number, .title, .singleSelectValues.Status, .url, .numberValues.Point, (.labels | join(", "))]) | @csv'
```

Stdout
```csv
"type","number","title","status","url"
"ISSUE",48,"title2","In Progress","https://github.com/shuntaka9576/kanban-test/issues/48"
"DRAFT_ISSUE",0,"","In Progress",""
"ISSUE",49,"Fix bug",,"https://github.com/shuntaka9576/kanban-test/issues/49"
"DRAFT_ISSUE",0,"",,""
"DRAFT_ISSUE",0,"",,""
"ISSUE",50,"aaa[]",,"https://github.com/shuntaka9576/kanban-test/issues/50"
"ISSUE",51,"aaa",,"https://github.com/shuntaka9576/kanban-test/issues/51"
"ISSUE",52,"aaa",,"https://github.com/shuntaka9576/kanban-test/issues/52"
"ISSUE",58,"test",,"https://github.com/shuntaka9576/kanban-test/issues/58"
"ISSUE",89,"Fix bug","Todo","https://github.com/shuntaka9576/kanban-test/issues/89"
"ISSUE",90,"Fix bug","Todo","https://github.com/shuntaka9576/kanban-test/issues/90"
"DRAFT_ISSUE",0,"",,""
```

*Details*

|flag|short|required|default|
|---|---|---|---|
|--user or --org|-u or -o|true|""
|--project-title|-p|true|""


## Special Thanks

* https://github.com/yusukebe/gh-markdown-preview
