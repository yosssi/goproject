package indexers

import (
	"sync"

	"github.com/yosssi/goproject/models"
)

const (
	workerName = "indexers.GitHub"
)

// GitHub represents a GitHub indexer.
type GitHub struct {
	app *models.Application
	wg  *sync.WaitGroup
}

// NewGitHub generates a GitHub indexer and returns it.
func NewGitHub(app *models.Application, wg *sync.WaitGroup) *GitHub {
	wg.Add(1)
	return &GitHub{app: app, wg: wg}
}

// Run runs the main task.
func (g *GitHub) Run() {
	g.Index()
}

// Index indexes Go projects into Elasticsearch.
func (g *GitHub) Index() {
	g.app.Logger.Infof("[%s] Index starts.", workerName)

	g.wg.Done()
}
