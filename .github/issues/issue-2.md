---
code: null
id: MDU6SXNzdWU1MzgwMTg3MTE=
databaseId: 538018711
number: 2
url: https://github.com/octolab/pkg/issues/2
title: "safe.Background"
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
createdAt: 2019-12-15T08:16:05Z
updatedAt: 2020-03-23T19:30:30Z
lastEditedAt: null
closedAt: 2020-03-14T18:08:25Z
---

# safe.Background

```go
	wg :=&sync.WaitGroup{}
	wg.Add(len(data))
	for range data {
		go safe.Background(wg, func() { handler.ServeHTTP(rec, &http.Request{}) }, 1)
	}
```
