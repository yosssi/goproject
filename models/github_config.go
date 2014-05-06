package models

import (
	"github.com/yosssi/gogithub"
	"github.com/yosssi/goproject/consts"
	"github.com/yosssi/goproject/utils"
)

// GitHubConfig represents a GitHub config.
type GitHubConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// GitHubClient generates a GitHub client and returns it.
func (g *GitHubConfig) GitHubClient() *gogithub.Client {
	return gogithub.NewClient(g.ClientID, g.ClientSecret)
}

// NewGitHubConfig parses a yaml file, generates a GitHubConfig and returns it.
func NewGitHubConfig() (*GitHubConfig, error) {
	config := &GitHubConfig{}
	if err := utils.YamlUnmarshal(consts.GitHubConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
