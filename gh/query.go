package gh

import (
	"fmt"
	"strings"
)

func GetUserQuery(users []string) string {
	var userQueries string
	for i, assignee := range users {
		userQuery := fmt.Sprintf(`user%d: user(login: "%s") {
  id
  login
}`, i, assignee)
		userQueries = userQueries + userQuery
	}

	query := fmt.Sprintf(`query {
%s
}`, userQueries)

	return query
}

func GetListQuery(clientType ClientType, name string) string {
	query := fmt.Sprintf(`query{
  %s(login: "%s") {
    projectsV2(first: 20) {
      nodes {
        id
        title
      }
    }
  }
}
`, clientType, name)

	return query
}

func GetProjectItemsQuery(projectId string, cursor *string) string {
	var afterClause string
	if cursor != nil {
		afterClause = fmt.Sprintf(`, after: "%s"`, *cursor)
	}

	query := fmt.Sprintf(`query{
  node(id: "%s") {
    ... on ProjectV2 {
      title
      items(first: 20%s) {
        pageInfo {
          hasNextPage
          endCursor
        }
        nodes {
          id
          createdAt
          updatedAt
          isArchived
          content {
            __typename
          }
          fieldValues(first: 20) {
            nodes {
              __typename
              ... on ProjectV2ItemFieldSingleSelectValue {
                name
                optionId
                field {
                  ... on ProjectV2SingleSelectField {
                    id
                    name
                  }
                }
              }
            }
          }
          content {
            ... on DraftIssue {
              title
              body
            }
            ... on Issue {
              number
              title
              state
              url
              body
              labels(first: 20) {
                nodes {
                  name
                }
              }
            }
          }
        }
      }
    }
  }
}`, projectId, afterClause)

	return query
}

func GetProjectFieldsQuery(projectId string) string {
	query := fmt.Sprintf(`query{
  node(id: "%s") {
    ... on ProjectV2 {
      fields(first: 20) {
        nodes {
          ... on ProjectV2Field {
            id
            name
            dataType
          }
          ... on ProjectV2IterationField {
            id
            name
            dataType
            configuration {
              iterations {
                startDate
                id
              }
            }
          }
          ... on ProjectV2SingleSelectField {
            id
            name
            dataType
            options {
              id
              name
            }
          }
        }
      }
    }
  }
}`, projectId)

	return query

}

func convertArrayString(array []string) (arrayString string) {
	arrayString = "["
	for i, string := range array {
		arrayString += fmt.Sprintf("\"%s\"", string)
		if i < len(array)-1 {
			arrayString += ","
		}
	}
	arrayString += "]"

	return arrayString
}

func GetDraftIssueMutation(projectId string, title string, body string, assigneeIds []string) string {
	assigneeIdsString := convertArrayString(assigneeIds)

	query := fmt.Sprintf(`mutation {
  addProjectV2DraftIssue(input: {projectId: "%s" title: "%s" body: "%s" assigneeIds: %s}) {
    projectItem {
      id
    }
  }
}`, projectId, title, body, assigneeIdsString)

	return query
}

func UpdateItemMutation(projectId string, itemId string, fieldId string, valueType PROJECT_V2_DATA_TYPE, value any) string {
	queryKey := ""
	applyValue := fmt.Sprintf("\"%s\"", value)
	switch valueType {
	case DATE:
		queryKey = "date"
	case NUMBER:
		queryKey = "number"
		applyValue = fmt.Sprintf("%s", value)
	case SINGLE_SELECT:
		queryKey = "singleSelectOptionId"
	case TEXT:
		queryKey = "text"
	case ITERATION:
		queryKey = "iterationId" // TODO
	default:
		panic("no supported error")
	}

	query := fmt.Sprintf(`mutation {
  updateProjectV2ItemFieldValue(
    input: {
      projectId: "%s"
      itemId: "%s"
      fieldId: "%s"
      value: {
        %s: %s
      }
    }
  ) {
    projectV2Item {
      id
    }
  }
}`, projectId, itemId, fieldId, queryKey, applyValue)

	return query
}

func GetRepoQuery(owner string, repo string) string {
	query := fmt.Sprintf(`query{
  repository(owner: "%s", name: "%s") {
    id
  }
}
`, owner, repo)

	return query

}

func IssueMutation(repositoryId string, title string, body string, assigneeIds []string, labelIds []string) string {
	assigneeIdsString := convertArrayString(assigneeIds)
	labelIdsString := convertArrayString(labelIds)
	body = strings.ReplaceAll(body, `"`, `\"`)

	query := fmt.Sprintf(`mutation {
		createIssue(input: {repositoryId: "%s" title: "%s" body: "%s" assigneeIds: %s, labelIds: %s}) {
    issue {
      id
    }
  }
}`, repositoryId, title, body, assigneeIdsString, labelIdsString)

	return query
}

func AddItemMutation(projectId string, contentId string) string {
	query := fmt.Sprintf(`mutation {
  addProjectV2ItemById(input: {projectId: "%s" contentId: "%s"}) {
    item {
      id
    }
  }
}`, projectId, contentId)

	return query
}
