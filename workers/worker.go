// Package workers provides workers.
package workers

import (
	"sync"

	"github.com/yosssi/goproject/models"
)

// Worker represents a worker interface.
type Worker interface {
	Run()
}

// Searchers generates all searchers and returns them.
func Searchers(app *models.Application, wg *sync.WaitGroup) []Worker {
	return []Worker{
		NewGitHubSearcher(app, wg),
	}
}
