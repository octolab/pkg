---
code: null
id: MDU6SXNzdWU3MzU5MTU5MzE=
databaseId: 735915931
number: 28
url: https://github.com/octolab/pkg/issues/28
title: "helper for testing"
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
createdAt: 2020-11-04T08:30:07Z
updatedAt: 2020-11-11T16:28:01Z
lastEditedAt: null
closedAt: 2020-11-11T16:28:01Z
---

# helper for testing

```go

	ctx := context.Background()
	deadline, has := t.Deadline()
	if has {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, deadline.Add(50*time.Millisecond))
		defer cancel()
	}

```

<!-- 2020-11-04T10:45Z https://github.com/octolab/pkg/issues/28#issuecomment-721657340
```go
func ContextWithDeadline(
	ctx context.Context,
	src interface{ Deadline() (time.Time, bool) },
) (context.Context, context.CancelFunc) {
	deadline, has := src.Deadline()
	if has {
		return context.WithDeadline(ctx, deadline.Add(50*time.Millisecond))
	}
	return ctx, func() {}
}
```
-->
