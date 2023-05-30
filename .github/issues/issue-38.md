---
code: null
id: MDU6SXNzdWU4MDYwODI4Mjk=
databaseId: 806082829
number: 38
url: https://github.com/octolab/pkg/issues/38
title: "add async.DelayedJob"
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
createdAt: 2021-02-11T05:08:49Z
updatedAt: 2023-01-09T13:03:35Z
lastEditedAt: null
closedAt: null
---

# add async.DelayedJob

```go
job := async.NewDelayedJob(make([]func(), 0, 2))

if cond1 {
  job.Do(...)
}

if cond2 {
  job.Do(...)
}

job.Wait()
```

<!-- 2021-02-11T07:05Z https://github.com/octolab/pkg/issues/38#issuecomment-777242477
motivation: if cond1 is false, then no need to create a goroutine
-->
