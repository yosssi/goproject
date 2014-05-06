package searchers

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/yosssi/goproject/models"
)

const (
	workerName  = "searchers.GitHub"
	queryUpdate = "q=language:go&sort=update&order=asc&page=%d"
)

// GitHub represents a GitHub searcher.
type GitHub struct {
	app *models.Application
	wg  *sync.WaitGroup
}

// NewGitHub generates a GitHub worker and returns it.
func NewGitHub(app *models.Application, wg *sync.WaitGroup) *GitHub {
	wg.Add(1)
	return &GitHub{app: app, wg: wg}
}

// Run runs the main task.
func (g *GitHub) Run() {
	g.Search()
}

// Search searches Go projects from GitHub.
func (g *GitHub) Search() {
	g.app.Logger.Infof("[%s] Search starts.", workerName)

	for page := 1; ; page++ {
		result, err := g.app.GitHubClient.SearchRepositories(fmt.Sprintf(queryUpdate, page))
		if err != nil {
			log.Fatal(err)
		}

		l := len(result.Items)

		g.app.Logger.Debugf("[%s] page: %d, items: %d,", workerName, page, l)

		if l == 0 {
			break
		}

		time.Sleep(5 * time.Second)
	}

	g.wg.Done()
}
