---
code: null
id: MDU6SXNzdWU2NDYxNzE0ODA=
databaseId: 646171480
number: 19
url: https://github.com/octolab/pkg/issues/19
title: "hide fallback under the hood"
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
createdAt: 2020-06-26T10:36:48Z
updatedAt: 2020-06-26T11:45:20Z
lastEditedAt: null
closedAt: 2020-06-26T11:45:20Z
---

# hide fallback under the hood

```go
classificator.Classify(err, fallback) -> classificator.Classify(err)
```
