package main

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/todo/internal"
)

func main() {
	drone.New(internal.New).Boot()
}
