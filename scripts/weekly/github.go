package main

import (
	"context"
	"strings"
	"time"

	"github.com/google/go-github/v29/github"
)

// GitHub is a wrapper for go-github
type GitHub struct {
	client *github.Client
}

func NewGitHub() (*GitHub, error) {
	client := github.NewClient(nil)
	return &GitHub{client: client}, nil
}

func IsWeeklyIssue(issue *github.Issue) bool {
	return strings.HasPrefix(issue.GetTitle(), "Weekly-")
}

func (g *GitHub) RecentWeeklyIssues(ctx context.Context) ([]*github.Issue, error) {
	issues, _, err := g.client.Issues.ListByRepo(ctx, "dyweb", "weekly", &github.IssueListByRepoOptions{
		State:     "all",
		Sort:      "created",
		Direction: "desc",
		Since:     time.Now().Add(-1 * 2 * 30 * 24 * time.Hour),
	})
	if err != nil {
		return nil, err
	}
	var filtered []*github.Issue
	for _, issue := range issues {
		if IsWeeklyIssue(issue) {
			filtered = append(filtered, issue)
		}
		//log.Infof("issue %d %d %s", issue.GetID(), issue.GetNumber(), issue.GetTitle())
	}
	return filtered, nil
}

func (g *GitHub) Issue(ctx context.Context, id int) (*github.Issue, error) {
	issue, _, err := g.client.Issues.Get(ctx, "dyweb", "weekly", id)
	if err != nil {
		return nil, err
	}
	return issue, nil
}
