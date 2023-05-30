---
code: null
id: I_kwDODBUd6M5JkE3d
databaseId: 1234193885
number: 63
url: https://github.com/octolab/pkg/issues/63
title: "safe: Writer functionality"
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
createdAt: 2022-05-12T15:59:45Z
updatedAt: 2023-01-09T13:03:24Z
lastEditedAt: null
closedAt: null
---

# safe: Writer functionality

opposite to https://github.com/octolab/pkg/blob/19e671e0420f209e9ee8c9294722eeb4d4834cf3/unsafe/suppress.go#L11 define `safe.Write(io.Write(...), count)` to check written length and no error.
