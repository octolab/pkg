---
code: null
id: I_kwDODBUd6M5BpnOt
databaseId: 1101427629
number: 58
url: https://github.com/octolab/pkg/issues/58
title: "extend errors.Is"
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
createdAt: 2022-01-13T09:25:30Z
updatedAt: 2023-01-09T13:03:27Z
lastEditedAt: null
closedAt: null
---

# extend errors.Is

```go
	if err != nil {
		switch {
		case errors.Is(err, bitbucket.ErrEventNotFound):
			fallthrough
		case errors.Is(err, bitbucket.ErrEventNotSpecifiedToParse):
			fallthrough
		case errors.Is(err, bitbucket.ErrInvalidHTTPMethod):
			fallthrough
		case errors.Is(err, bitbucket.ErrParsingPayload):
			code = http.StatusBadRequest
```

to

```go
	if err != nil {
		switch {
		case errors.Is(err, bitbucket.ErrEventNotFound, bitbucket.ErrEventNotSpecifiedToParse, ...):
			code = http.StatusBadRequest
```

<!-- 2022-01-13T09:25Z https://github.com/octolab/pkg/issues/58#issuecomment-1011950471
maybe IsOneOf
-->
