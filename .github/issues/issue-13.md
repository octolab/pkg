---
code: null
id: MDU6SXNzdWU2MDI5NTU5MzA=
databaseId: 602955930
number: 13
url: https://github.com/octolab/pkg/issues/13
title: "review signals for trap"
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
createdAt: 2020-04-20T05:37:51Z
updatedAt: 2020-04-20T18:48:03Z
lastEditedAt: null
closedAt: 2020-04-20T18:48:03Z
---

# review signals for trap

`os.Interrupt, syscall.SIGINT, syscall.SIGTERM` -> `syscall.SIGINT, syscall.SIGTERM` cause os.Interrupt is an alias for syscall.SIGINT

<!-- 2020-04-20T18:48Z https://github.com/octolab/pkg/issues/13#issuecomment-616741447
won't do, will save compatible with plan9
-->
