---
code: null
id: MDU6SXNzdWU1Nzc0OTY5OTU=
databaseId: 577496995
number: 4
url: https://github.com/octolab/pkg/issues/4
title: "timeout/deadline correction"
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
createdAt: 2020-03-08T12:58:38Z
updatedAt: 2020-03-23T19:30:31Z
lastEditedAt: null
closedAt: 2020-03-11T05:17:41Z
---

# timeout/deadline correction

`time.TimeoutCorrection(http.Header(req.Header).Timeout(time.Second), 10*time.Millisecond)`

<!-- 2020-03-08T12:59Z https://github.com/octolab/pkg/issues/4#issuecomment-596203097
it may be worth fixing the interface or package structure:

`http.Header(req.Header).Timeout(time.Second)` -> `extended.Header(req.Header).Timeout(time.Second)` || `config.FromHeader(req.Header).Timeout()`
-->
