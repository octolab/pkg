---
code: null
id: MDU6SXNzdWU3NzgyMzcxNTI=
databaseId: 778237152
number: 32
url: https://github.com/octolab/pkg/issues/32
title: "extend safe package by finalizers"
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
createdAt: 2021-01-04T17:12:24Z
updatedAt: 2023-01-09T13:03:23Z
lastEditedAt: null
closedAt: null
---

# extend safe package by finalizers

https://golang.org/pkg/runtime/#SetFinalizer


```go
file, err := safe.OpenFile()
if err ...
// no call defer file.Close()
// panic in runtime
```

how it works

```go
Open() -> SetFinalizer

Close() -> UnsetFinalizer
```

<!-- 2021-01-05T15:39Z https://github.com/octolab/pkg/issues/32#issuecomment-754713857
also, the custom profiler can be usefull:

```gp
Open() -> profile.Add

Close() -> profile.Remove
```
-->
