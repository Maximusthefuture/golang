package scheduler

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	//"time"
)

func Schedule() {
	gocron.Every(5).Second().Do(Print)
	gocron.Start()
}

func Print() {
	fmt.Println("Helloaqeqweqwe")
}
