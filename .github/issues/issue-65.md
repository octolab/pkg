---
code: null
id: I_kwDODBUd6M5Kt0IT
databaseId: 1253523987
number: 65
url: https://github.com/octolab/pkg/issues/65
title: "env: add analogue of `os.Getenv` and `os.Setenv`"
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
createdAt: 2022-05-31T08:32:20Z
updatedAt: 2023-01-09T13:03:23Z
lastEditedAt: null
closedAt: null
---

# env: add analogue of `os.Getenv` and `os.Setenv`

```go
	vars := env.Environment{
		env.Must("GIT_REMOTE", "git@github.com:octomation/maintainer.git"),
		env.Must("GITHUB_TOKEN", "secret"),
	}
	vars.Get("GIT_REMOTE")
	vars.Set("GIT_REMOTE", "new value")
```
