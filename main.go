package main

import (
	_ "go_apis/conn"
	routers "go_apis/router"
)

func main() {
	router := routers.Routers()
	router.Run(":9090")
}

