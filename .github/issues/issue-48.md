---
code: null
id: MDU6SXNzdWU4MTI2ODIwNzc=
databaseId: 812682077
number: 48
url: https://github.com/octolab/pkg/issues/48
title: "add ReadCloser to testing"
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
createdAt: 2021-02-20T18:44:55Z
updatedAt: 2023-01-09T13:03:32Z
lastEditedAt: null
closedAt: null
---

# add ReadCloser to testing

```go
RepeatableReadCloser(
    helper.ReadCloser("file:testdata/some.data"),
    bytes.NewBuffer(make([]byte, 0, 5)),
)
```
