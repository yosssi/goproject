// Package workers provides workers.
package workers

import (
	"sync"

	"github.com/yosssi/goproject/models"
	"github.com/yosssi/goproject/workers/indexers"
	"github.com/yosssi/goproject/workers/searchers"
)

// Worker represents a worker interface.
type Worker interface {
	Run()
}

// All generates all workers and returns them.
func All(app *models.Application, wg *sync.WaitGroup) []Worker {
	return []Worker{
		searchers.NewGitHub(app, wg),
		indexers.NewGitHub(app, wg),
	}
}
