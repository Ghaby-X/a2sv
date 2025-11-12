package main

import (
	"fmt"

	"github.com/Ghaby-X/task_manager/router"
)

func main() {
	r := router.GetTaskRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("Failed to run error")
	}
}
