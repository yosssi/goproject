// package main provides main processes
package main

import (
	"log"
	"sync"

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

	// Generate a WaitGroup.
	wg := &sync.WaitGroup{}

	// Invoke workers.
	for _, worker := range workers.All(app, wg) {
		go worker.Run()
	}

	// Wait until all workers end their tasks.
	wg.Wait()
}
