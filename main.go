package main

import (
	"musicMod/model"
	"musicMod/router"
)

func main() {
	model.InitDb()
	r := router.InitRouter()
	r.Run(":8890")
}
