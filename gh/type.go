package gh

type ClientType = string

const (
	ORGANIZATION ClientType = "organization"
	USER         ClientType = "user"
)

type PROJECT_V2_DATA_TYPE = string

const (
	ASSIGNEES            PROJECT_V2_DATA_TYPE = "ASSIGNEES"
	DATE                 PROJECT_V2_DATA_TYPE = "DATE"
	ITERATION            PROJECT_V2_DATA_TYPE = "ITERATION"
	LABELS               PROJECT_V2_DATA_TYPE = "LABELS"
	LINKED_PULL_REQUESTS PROJECT_V2_DATA_TYPE = "LINKED_PULL_REQUESTS"
	MILESTONE            PROJECT_V2_DATA_TYPE = "MILESTONE"
	NUMBER               PROJECT_V2_DATA_TYPE = "NUMBER"
	REPOSITORY           PROJECT_V2_DATA_TYPE = "REPOSITORY"
	REVIEWERS            PROJECT_V2_DATA_TYPE = "REVIEWERS"
	SINGLE_SELECT        PROJECT_V2_DATA_TYPE = "SINGLE_SELECT"
	TEXT                 PROJECT_V2_DATA_TYPE = "TEXT"
	TITLE                PROJECT_V2_DATA_TYPE = "TITLE"
	TRACKS               PROJECT_V2_DATA_TYPE = "TRACKS"
)

type ITEM_TYPE = string

const (
	ISSUE       ITEM_TYPE = "ISSUE"
	DRAFT_ISSUE ITEM_TYPE = "DRAFT_ISSUE"
)
