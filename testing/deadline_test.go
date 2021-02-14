package testing_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	. "go.octolab.org/testing"
	"go.octolab.org/time"
)

func TestWithDeadline(t *testing.T) {
	root, cancel := context.WithTimeout(context.TODO(), 100*time.Millisecond)
	defer cancel()

	ctx, cancel := WithDeadline(context.TODO(), root, 50*time.Millisecond)
	defer cancel()

	infinite, cancel := WithDeadline(context.TODO(), context.TODO(), 10*time.Millisecond)
	defer cancel()

	test, cancel := WithDeadline(infinite, t, time.Millisecond)
	defer cancel()

	select {
	case <-root.Done():
		require.Fail(t, "root is not expected")
	case <-infinite.Done():
		require.Fail(t, "infinite is not expected")
	case <-test.Done():
		require.Fail(t, "test is not expected")
	case <-ctx.Done():
		// success
	}
}
