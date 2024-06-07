package main

import (
	"github.com/dronestock/deb/internal"
	"github.com/dronestock/drone"
)

func main() {
	drone.New(internal.New).Boot()
}
