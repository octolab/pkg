---
code: null
id: MDU6SXNzdWU2ODU2MjIyODE=
databaseId: 685622281
number: 22
url: https://github.com/octolab/pkg/issues/22
title: "add sync.Setenvs(...)"
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
createdAt: 2020-08-25T16:24:11Z
updatedAt: 2020-08-26T20:49:21Z
lastEditedAt: null
closedAt: 2020-08-26T20:49:21Z
---

# add sync.Setenvs(...)

```
release, err := sync.Setenvs("A", "b", "C", "d")
...
defer release()
```
