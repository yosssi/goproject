package workers

import "github.com/yosssi/goproject/models"

// GitHub represents a GitHub worker.
type GitHub struct {
	app *models.Application
}

// NewGitHub generates a GitHub worker and returns it.
func NewGitHub(app *models.Application) *GitHub {
	return &GitHub{app: app}
}

// Search searches Go projects from GitHub and
// indexes them into Elasticsearch.
func (g *GitHub) Search() {
	g.app.Logger.Infof("GitHub.Search starts")
}
