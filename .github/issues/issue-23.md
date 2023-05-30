---
code: null
id: MDU6SXNzdWU2ODYwNTYwODY=
databaseId: 686056086
number: 23
url: https://github.com/octolab/pkg/issues/23
title: "extend pointer package"
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
createdAt: 2020-08-26T06:42:45Z
updatedAt: 2020-08-26T06:48:40Z
lastEditedAt: 2020-08-26T06:48:20Z
closedAt: 2020-08-26T06:48:40Z
---

# extend pointer package

- intToInt32
- intToInt64
- int32ToInt
- int64ToInt

<!-- 2020-08-26T06:48Z https://github.com/octolab/pkg/issues/23#issuecomment-680693561
won't do, use `pointer.ToInt64(int64(pointer.ValueOfInt(intPtr)))` instead
-->
