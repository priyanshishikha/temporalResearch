package main

import (
	greeting "greeting"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	//connection to temporal clustef
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	//defer will close the connection at the end of the execution
	defer c.Close()

	// greeting-tasks: is to specify the task queue name for the worker to fetch from
	w := worker.New(c, "greeting-tasks", worker.Options{})
	//Register the workflow
	w.RegisterWorkflow(greeting.GreetSomeone)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
