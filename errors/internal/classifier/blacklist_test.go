package classifier_test

import (
	"errors"
	"testing"

	. "go.octolab.org/errors/internal/classifier"
)

func TestBlacklistClassifier_Classify(t *testing.T) {
	var (
		errInBlacklist    = errors.New("is in blacklist")
		errNotInBlacklist = errors.New("is not in blacklist")
	)
	list := BlacklistClassifier([]error{errInBlacklist})

	if list.Classify(nil) != Succeed {
		t.Error("succeed is expected")
	}

	if list.Classify(errNotInBlacklist) != Retry {
		t.Error("retry is expected")
	}

	if list.Classify(errInBlacklist) != Fail {
		t.Error("fail is expected")
	}
}
