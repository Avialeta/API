package main

import (
	"fmt"

	"github.com/fellah/stop"

	"github.com/avialeta/api/api"
	"github.com/avialeta/api/job"
)

func main() {
	go api.Serve()
	go job.Scheduler()

	// Get stop signal
	<-stop.Ch

	fmt.Println("done")
}
