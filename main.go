// package main provides main processes.
package main

import (
	"log"

	"github.com/yosssi/goproject/models"
	"github.com/yosssi/goproject/workers"
)

// main executes main processes.
func main() {
	// Generate an application.
	app, err := models.NewApplication()
	if err != nil {
		log.Fatal(err)
	}
	app.Logger.Infof("An application was generated. [app: %+v]", app)

	// Invoke workers.
	wkrs := []workers.Worker{workers.NewGitHub(app)}
	for _, worker := range wkrs {
		worker.Search()
	}
}
