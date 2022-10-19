package main

import (
	"context"
	"greeting"
	"log"
	"os"

	"go.temporal.io/sdk/client"
)

// This code is to execute the workflow before this ensure a worker is up running by first running code of worker in different terminal
func main() {
	//Connection to temporal cluster
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()
	//configurations to execute the workflow
	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "greeting-tasks",
	}
	//ExecuteWorkflow will publish the task of executing the task of GreetSomeone in task queue and will let the process run in background
	//It will continue to the point of results are required and fetched of this ExecuteWorkflow in this case see below till the point of we.Get func
	//this way multiple process can be processed in an asynch manner (parallel fashion)
	//Note if there are multiple workers present the task of ExecuteWorkflow will be picked up by different worker in the background
	//and this worker can continue the rest of code if we have only 1 worker upon calling of Get Func this worker has to redirect itself to the task
	// and nothing special then same serial computing....... multiple worker is where magic happens
	// for having multiple workers no code change required just up more pods for workers to run on
	we, err := c.ExecuteWorkflow(context.Background(), options, greeting.GreetSomeone, os.Args[1])
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
