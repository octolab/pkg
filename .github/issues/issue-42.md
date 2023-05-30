---
code: null
id: MDU6SXNzdWU4MDk4OTI1NTk=
databaseId: 809892559
number: 42
url: https://github.com/octolab/pkg/issues/42
title: "add MustReturn function"
issueType: null
labels:
  - "help wanted"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2021-02-17T06:18:27Z
updatedAt: 2021-02-17T06:26:04Z
lastEditedAt: null
closedAt: 2021-02-17T06:26:04Z
---

# add MustReturn function

```go
import "github.com/Masterminds/semver"

var (
	min := fn.MustReturn(semver.NewConstraint(">= 92.0")).(*semver.Constraints)
)
```
