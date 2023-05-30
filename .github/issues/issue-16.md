---
code: null
id: MDU6SXNzdWU2NDAzNzY3OTU=
databaseId: 640376795
number: 16
url: https://github.com/octolab/pkg/issues/16
title: "misplaced compiler directive"
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
createdAt: 2020-06-17T11:45:21Z
updatedAt: 2020-06-17T20:10:43Z
lastEditedAt: null
closedAt: 2020-06-17T20:10:43Z
---

# misplaced compiler directive

https://travis-ci.com/github/octolab/pkg/jobs/350339344


```
# go.octolab.org/runtime_test [go.octolab.org/runtime.test]
262runtime/caller_test.go:225:4: misplaced compiler directive
263runtime/caller_test.go:229:4: misplaced compiler directive
```

<!-- 2020-06-17T20:10Z https://github.com/octolab/pkg/issues/16#issuecomment-645597913
fixed https://travis-ci.com/github/octolab/pkg/jobs/350583936
-->
