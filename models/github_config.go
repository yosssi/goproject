package models

import (
	"github.com/yosssi/gogithub"
	"github.com/yosssi/goproject/consts"
	"github.com/yosssi/goproject/utils"
)

// GitHubConfig represents a GitHub config.
type GitHubConfig struct {
	AccessToken string `yaml:"access_token"`
}

// GitHubClient generates a GitHub client and returns it.
func (g *GitHubConfig) GitHubClient() *gogithub.Client {
	return &gogithub.Client{AccessToken: g.AccessToken}
}

// NewGitHubConfig parses a yaml file, generates a GitHubConfig and returns it.
func NewGitHubConfig() (*GitHubConfig, error) {
	config := &GitHubConfig{}
	if err := utils.YamlUnmarshal(consts.GitHubConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
