---
code: null
id: MDU6SXNzdWU2NDIxMjUxNzg=
databaseId: 642125178
number: 17
url: https://github.com/octolab/pkg/issues/17
title: "errors classifier"
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
createdAt: 2020-06-19T17:29:13Z
updatedAt: 2020-06-20T07:24:46Z
lastEditedAt: null
closedAt: 2020-06-20T07:24:46Z
---

# errors classifier

```go
classifier := errors.Classifier{
  "some": []error{x, y, z}
}
class := classifier.Classify(err)
metrics.Count(class).Send()
```
