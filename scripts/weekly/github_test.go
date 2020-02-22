package main

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGitHub_RecentWeeklyIssues(t *testing.T) {
	gh, err := NewGitHub()
	require.Nil(t, err)
	issues, err := gh.RecentWeeklyIssues(context.Background())
	require.Nil(t, err)
	for _, issue := range issues {
		t.Logf("%d %s %s", issue.GetNumber(), issue.GetTitle(), issue.GetCreatedAt())
	}
}
