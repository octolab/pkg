---
code: null
id: MDU6SXNzdWU1ODI1NTIxNDA=
databaseId: 582552140
number: 7
url: https://github.com/octolab/pkg/issues/7
title: "new package sequence/generator"
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
createdAt: 2020-03-16T19:33:28Z
updatedAt: 2020-04-01T19:51:49Z
lastEditedAt: null
closedAt: 2020-04-01T19:51:49Z
---

# new package sequence/generator

The concept:

```go
sequence := new(integer.generator).Since(10).Rhythm(func (int64) int64).Until(10^5)

for range sequence.Source() { ... } // math.MaxInt64 by default

sequence.Next() bool

sequence.Current() int64

sequence.FastForward(int64) bool
```

```go
uid := new(uuid.generator)

uid.Next() string ?
```

```go
random := new(rand.generator)

random.Next() ?
```

<!-- 2020-03-21T18:53Z https://github.com/octolab/pkg/issues/7#issuecomment-602086854
research
- https://github.com/gofrs/uuid
- https://github.com/google/uuid
-->

<!-- 2020-04-01T19:51Z https://github.com/octolab/pkg/issues/7#issuecomment-607457125
implemented
-->
