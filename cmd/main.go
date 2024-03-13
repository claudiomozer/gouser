package main

import "github.com/claudiomozer/gouser/internal/app/http"

func main() {
	httpContainer := http.Start()
	defer httpContainer.CloseContainer()
}
