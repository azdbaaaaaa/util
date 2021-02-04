package util

import (
	"errors"
	"github.com/xanzy/go-gitlab"
	"strings"
)

type GitlabConfig struct {
	Uri   string `json:"uri"`
	Token string `json:"token"`
}

type GitlabClient struct {
	client *gitlab.Client
}

func NewGitlabClient(conf GitlabConfig) (gc *GitlabClient, err error) {
	client, err := gitlab.NewClient(conf.Token, gitlab.WithBaseURL(conf.Uri))
	if err != nil {
		Logger.Log.Errorf("初始化gitlab客户端错误：%v", err)
		return nil, err
	}
	return &GitlabClient{client: client}, nil
}

func (c *GitlabClient) getFullProject(gitUrl string) (string, error) {
	fullProject := strings.ReplaceAll(strings.Split(gitUrl, ":")[1], ".git", "")
	if fullProject == "" {
		return "", errors.New("错误的URL")
	}
	return fullProject, nil
}

func (c *GitlabClient) GetBranchList(gitUrl string, search string) ([]*gitlab.Branch, error) {
	project, err := c.getFullProject(gitUrl)
	if err != nil {
		return nil, err
	}
	branches, _, err := c.client.Branches.ListBranches(project, &gitlab.ListBranchesOptions{
		Search: &search,
	})
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func (c *GitlabClient) CreateMergeRequest(gitUrl, title, sourceBranch, targetBranch string) (mr *gitlab.MergeRequest, err error) {
	project, err := c.getFullProject(gitUrl)
	if err != nil {
		return nil, err
	}
	mr, _, err = c.client.MergeRequests.CreateMergeRequest(project, &gitlab.CreateMergeRequestOptions{
		Title:        &title,
		SourceBranch: &sourceBranch,
		TargetBranch: &targetBranch,
	})
	if err != nil {
		return
	}
	return
}
