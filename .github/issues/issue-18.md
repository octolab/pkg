---
code: null
id: MDU6SXNzdWU2NDYxNjI2NTA=
databaseId: 646162650
number: 18
url: https://github.com/octolab/pkg/issues/18
title: "add Merge for errors classificator"
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
createdAt: 2020-06-26T10:19:58Z
updatedAt: 2020-06-26T11:45:20Z
lastEditedAt: null
closedAt: 2020-06-26T11:45:20Z
---

# add Merge for errors classificator

```go
local := packageX.Classificator.Merge(packageY.Classificator, packageZ.Classificator)
if !local.Consistent() {
  panic()
}
```
