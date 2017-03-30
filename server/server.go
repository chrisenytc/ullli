package main

import (
	"github.com/chrisenytc/ullli/adapters"
	"github.com/chrisenytc/ullli/config"
	"github.com/chrisenytc/ullli/router"
)

func main() {
	// Load configs
	config.Load()

	// Load database
	adapters.LoadDatabase()

	// Load proxy server
	router.Load()
}
