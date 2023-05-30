---
code: null
id: MDU6SXNzdWU4MTI1Njc0MzY=
databaseId: 812567436
number: 45
url: https://github.com/octolab/pkg/issues/45
title: "improve io.Discard for defer"
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
createdAt: 2021-02-20T09:51:12Z
updatedAt: 2023-01-09T13:03:34Z
lastEditedAt: null
closedAt: null
---

# improve io.Discard for defer

if I use io.Discard with defer I have bad dev experience:

```go
resp, err := http.Do(req)
if err {...}
defer safe.Close(io.Discard(resp.Body), ...) // now I can't read from body

json.NewDecoder(resp.Body).Decode() // woops
```

<!-- 2021-02-20T09:56Z https://github.com/octolab/pkg/issues/45#issuecomment-782598113
and write docs with tests
-->
