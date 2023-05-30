---
code: null
id: MDU6SXNzdWU2NDc0NDc3MjM=
databaseId: 647447723
number: 21
url: https://github.com/octolab/pkg/issues/21
title: "add sync.Times as extended sync.Once"
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
createdAt: 2020-06-29T14:58:48Z
updatedAt: 2020-06-29T15:26:49Z
lastEditedAt: null
closedAt: 2020-06-29T15:26:49Z
---

# add sync.Times as extended sync.Once

```go
sync.Times(100).Do(func () {
   ...
})
```
