package main

import (
	"awesomeProject1/demo/router"
)

func main() {
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
