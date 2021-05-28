package main

import (
	app "github.com/parkedwards/go-dashboard-api/pkg"
)

func main() {
	r := app.Init()
	app.Boot(r)
}
