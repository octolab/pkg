---
code: null
id: MDU6SXNzdWU4MjM4OTE2OTE=
databaseId: 823891691
number: 52
url: https://github.com/octolab/pkg/issues/52
title: "rethink SetEnvs from testing"
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
createdAt: 2021-03-07T11:13:59Z
updatedAt: 2023-01-09T13:03:30Z
lastEditedAt: null
closedAt: null
---

# rethink SetEnvs from testing

it has an important disadvantages: it disable parallel test execution and serialize them.

define interface instead:

```go
type Environment interface {
    Expand()
    Lookup()
    Get()
    Set()
    Unset()
    Clear()
}
```

and add basic implementation based on the `os` package.
