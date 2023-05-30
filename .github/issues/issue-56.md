---
code: null
id: MDU6SXNzdWU5NDAwODI0Njk=
databaseId: 940082469
number: 56
url: https://github.com/octolab/pkg/issues/56
title: "testing: env: refactor SetEnvs"
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
createdAt: 2021-07-08T17:39:48Z
updatedAt: 2023-01-09T13:03:28Z
lastEditedAt: null
closedAt: null
---

# testing: env: refactor SetEnvs

Setenv will be available at go 1.17: https://tip.golang.org/pkg/testing/#T.Setenv

https://github.com/octolab/pkg/blob/d8517c0f3c4d86584cf71bf8bdae434cfcefc4fc/testing/env.go#L46-L48

<!-- 2022-05-20T19:28Z https://github.com/octolab/pkg/issues/56#issuecomment-1133248647
also, `T.Cleanup` is present, so I can extend for example `testing.SetEnv` to use it, see https://github.com/golang/mock/commit/6d816de489c18a7e9a8fbd2aa5bb2dd75f2bbc86
-->
