---
code: null
id: I_kwDODBUd6M5DnUQY
databaseId: 1134380056
number: 59
url: https://github.com/octolab/pkg/issues/59
title: "add inverted safe.Return"
issueType: null
labels: []
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2022-02-12T18:05:02Z
updatedAt: 2022-02-13T11:49:33Z
lastEditedAt: null
closedAt: 2022-02-13T11:49:33Z
---

# add inverted safe.Return

see https://github.com/octolab/pkg/blob/d8517c0f3c4d86584cf71bf8bdae434cfcefc4fc/unsafe/suppress.go#L21-L31

<!-- 2022-02-13T11:49Z https://github.com/octolab/pkg/issues/59#issuecomment-1038049094
won't do. incorrect semantics, `unsafe.Return(some.Action()).(Type)` is more consistent
-->
