package workers

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/yosssi/goproject/consts"
	"github.com/yosssi/goproject/models"
)

// GitHubSearcher represents a GitHub searcher.
type GitHubSearcher struct {
	app     *models.Application
	wg      *sync.WaitGroup
	indexer *GitHubIndexer
}

// NewGitHubSearcher generates a GitHub worker and returns it.
func NewGitHubSearcher(app *models.Application, wg *sync.WaitGroup) *GitHubSearcher {
	wg.Add(1)
	return &GitHubSearcher{app: app, wg: wg, indexer: NewGitHubIndexer(app)}
}

// Run runs the main task.
func (g *GitHubSearcher) Run() {
	g.Search()
}

// Search searches Go projects from GitHub.
func (g *GitHubSearcher) Search() {
	g.app.Logger.Infof("[%s] Search starts.", reflect.TypeOf(g))

	// Invoke an indexer.
	go g.indexer.Run()

	for page := 1; ; page++ {
		result, err := g.app.GitHubClient.SearchRepositories(fmt.Sprintf(consts.QueryUpdate, page))
		if err != nil {
			log.Fatal(err)
		}

		l := len(result.Items)

		g.app.Logger.Debugf("[%s] page: %d, items: %d", reflect.TypeOf(g), page, l)

		if l == 0 {
			break
		}

		for _, repository := range result.Items {
			g.indexer.repositoryC <- repository
		}

		time.Sleep(5 * time.Second)
	}

	g.wg.Done()
}
