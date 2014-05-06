// package main provides main processes.
package main

import (
	"fmt"
	"log"

	"github.com/yosssi/goproject/models"
)

// main executes main processes.
func main() {
	app, err := models.NewApplication()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", app.GitHubClient)
}
