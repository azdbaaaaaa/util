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

func (c *GitlabClient) CreateMergeRequest(gitUrl, title, sourceBranch, targetBranch string) (mr *gitlab.MergeRequest, resp *gitlab.Response, err error) {
	project, err := c.getFullProject(gitUrl)
	if err != nil {
		return
	}
	mr, resp, err = c.client.MergeRequests.CreateMergeRequest(project, &gitlab.CreateMergeRequestOptions{
		Title:        &title,
		SourceBranch: &sourceBranch,
		TargetBranch: &targetBranch,
	})
	if err != nil {
		return
	}
	return
}

func (c *GitlabClient) GetMergeRequest(gitUrl string, mrIID int) (mr *gitlab.MergeRequest, resp *gitlab.Response, err error) {
	project, err := c.getFullProject(gitUrl)
	if err != nil {
		return
	}
	mr, resp, err = c.client.MergeRequests.GetMergeRequest(project, mrIID, &gitlab.GetMergeRequestsOptions{})
	if err != nil {
		return
	}
	return
}

func (c *GitlabClient) GetMergeRequestChange(gitUrl string, mrIID int) (mr *gitlab.MergeRequest, resp *gitlab.Response, err error) {
	project, err := c.getFullProject(gitUrl)
	if err != nil {
		return
	}
	mr, resp, err = c.client.MergeRequests.GetMergeRequestChanges(project, mrIID, &gitlab.GetMergeRequestChangesOptions{})
	if err != nil {
		return
	}
	return
}

func (c *GitlabClient) CreateMergeRequestNote(gitUrl string, mrIID int, body *string) (note *gitlab.Note, resp *gitlab.Response, err error){
	project, err := c.getFullProject(gitUrl)
	if err != nil {
		return
	}
	note, resp, err = c.client.Notes.CreateMergeRequestNote(project, mrIID, &gitlab.CreateMergeRequestNoteOptions{Body: body})
	if err != nil {
		return
	}
	return
}

func (c *GitlabClient) ListMergeRequestNote(gitUrl string, mrIID int) (notes []*gitlab.Note, resp *gitlab.Response, err error){
	project, err := c.getFullProject(gitUrl)
	if err != nil {
		return
	}
	notes, resp, err = c.client.Notes.ListMergeRequestNotes(project, mrIID, &gitlab.ListMergeRequestNotesOptions{})
	if err != nil {
		return
	}
	return
}