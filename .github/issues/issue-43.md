---
code: null
id: MDU6SXNzdWU4MTE3MzA0OTA=
databaseId: 811730490
number: 43
url: https://github.com/octolab/pkg/issues/43
title: "discard before close"
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
createdAt: 2021-02-19T06:14:14Z
updatedAt: 2021-02-19T07:21:44Z
lastEditedAt: null
closedAt: 2021-02-19T07:21:44Z
---

# discard before close

`safe.Close(resp.Body, ...)` works, but not well. It will be better use something like this:

```go
defer safe.Close(io.Discard(resp.Body), ...)
```
