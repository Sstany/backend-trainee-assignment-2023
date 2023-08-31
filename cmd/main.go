package main

import "segmenty/app/core"

func main() {
	service := core.NewService("0.0.0.0", "8090")

	service.Start()
}
