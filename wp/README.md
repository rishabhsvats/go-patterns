
# Compress video to make it ready for web format

~~~
$ go run app/main.go
~~~


# Idea
videoQueue chan VideoProcessing job
wp (VideoDispatcher) ---> VideoDispatcher {videoQueue(jobQueue), workerPool}
videoWorker {new jobQueue, workerpool} 

start
    videoWorker.workerPool <- videoWorker.jobQueue
    job := <-videoWorker.jobQueue
    videoWorker.processVideoJob(job.Video)

run 
    videoWorker.start()
    VideoDispatcher.dispatch

dispatch
    job := <-VideoDispatcher.jobQueue
    workerJobQueue := <-VideoDispatcher.WorkerPool
	workerJobQueue <- job

-----------
package main

import "fmt"
import "time"

func main() {

	// make the request chan chan that both go-routines will be given	
	requestChan := make(chan chan string)
	
	// start the goroutines
	go goroutineC(requestChan)
	go goroutineD(requestChan)
	
	// sleep for a second to let the goroutines complete 	
	time.Sleep(time.Second)
	
}

func goroutineC(requestChan chan chan string) {
	
	// make a new response chan
	responseChan := make(chan string)
	
	// send the responseChan to goRoutineD
	requestChan <- responseChan
	
	// read the response
	response := <- responseChan
	
	fmt.Printf("Response: %v\n", response)
	
}

func goroutineD(requestChan chan chan string) {
	
	// read the responseChan from the requestChan
	responseChan := <- requestChan
	
	// send a value down the responseChan
	responseChan <- "wassup!"

	
}