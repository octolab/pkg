---
code: null
id: I_kwDODBUd6M5BF5oh
databaseId: 1092065825
number: 57
url: https://github.com/octolab/pkg/issues/57
title: "throw instead panic"
issueType: null
labels: []
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2022-01-02T17:38:36Z
updatedAt: 2023-01-09T13:03:27Z
lastEditedAt: null
closedAt: null
---

# throw instead panic

```go
errors.Throw(err error)
# vs
panic(any interface{})
```
