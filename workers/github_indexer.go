package workers

import (
	"fmt"
	"reflect"

	"github.com/yosssi/gogithub"
	"github.com/yosssi/goproject/consts"
	"github.com/yosssi/goproject/models"
)

// GitHubIndexer represents a GitHub indexer.
type GitHubIndexer struct {
	app         *models.Application
	repositoryC chan gogithub.Repository
}

// NewGitHubIndexer generates a GitHub indexer and returns it.
func NewGitHubIndexer(app *models.Application) *GitHubIndexer {
	return &GitHubIndexer{app: app, repositoryC: make(chan gogithub.Repository, 64)}
}

// Run runs the main task.
func (g *GitHubIndexer) Run() {
	g.Index()
}

// Index indexes Go projects into Elasticsearch.
func (g *GitHubIndexer) Index() {
	g.app.Logger.Infof("[%s] Index starts.", reflect.TypeOf(g))
	for repository := range g.repositoryC {
		g.app.Logger.Infof("[%s] Repository: %s", reflect.TypeOf(g), repository.FullName)
		result := make(map[string]interface{})
		g.app.ElasticsearchClient.Create(consts.IndexGoProject, consts.TypeGitHubRepository, repository, &result)
		fmt.Printf("%+v", result)
	}
}
