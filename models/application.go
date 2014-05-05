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
	Logger              *gologger.Logger
	GitHubClient        *gogithub.Client
	ElasticsearchClient *goelasticsearch.Client
}
