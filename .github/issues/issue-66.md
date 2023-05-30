---
code: null
id: I_kwDODBUd6M5KuCoJ
databaseId: 1253583369
number: 66
url: https://github.com/octolab/pkg/issues/66
title: "env: dump to dotenv"
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
createdAt: 2022-05-31T09:21:35Z
updatedAt: 2023-01-09T13:03:22Z
lastEditedAt: null
closedAt: null
---

# env: dump to dotenv

Instead of https://github.com/octomation/maintainer/blob/dbfaab2b34ba9d2363d3e665b20c8bb5783581b1/internal/config/tool_test.go#L44-L46

```go
vars.Dump(writer, [grouper])
```

`grouper` allows to sort and split them by prefix, e.g.

```dotenv
GITHUB_VAR1=...
GITHUB_VAR2=...

BITBUCKET_VAR1=...
...
```
