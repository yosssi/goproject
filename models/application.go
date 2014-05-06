package models

import (
	"github.com/yosssi/goelasticsearch"
	"github.com/yosssi/gogithub"
	"github.com/yosssi/gologger"
)

// Application represents an application context.
type Application struct {
	LoggerConfig        *LoggerConfig
	ElasticsearchConfig *ElasticsearchConfig
	GitHubConfig        *GitHubConfig
	Logger              *gologger.Logger
	ElasticsearchClient *goelasticsearch.Client
	GitHubClient        *gogithub.Client
}

// NewApplication generates an Application and returns it.
func NewApplication() (*Application, error) {
	loggerConfig, err := NewLoggerConfig()
	if err != nil {
		return nil, err
	}

	elasticsearchConfig, err := NewElasticsearchConfig()
	if err != nil {
		return nil, err
	}

	githubConfig, err := NewGitHubConfig()
	if err != nil {
		return nil, err
	}

	logger := loggerConfig.Logger()

	elasticsearchClient := elasticsearchConfig.ElasticsearchClient()

	githubClient := githubConfig.GitHubClient()

	app := &Application{
		LoggerConfig:        loggerConfig,
		ElasticsearchConfig: elasticsearchConfig,
		GitHubConfig:        githubConfig,
		Logger:              logger,
		ElasticsearchClient: elasticsearchClient,
		GitHubClient:        githubClient,
	}

	return app, nil
}
