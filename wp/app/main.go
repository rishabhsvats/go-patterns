package main

import (
	"fmt"

	"github.com/rishabhsvats/go-patterns/wp/streamer"
)

func main() {
	const numJobs = 1
	const numWorkers = 2

	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	//Get a worker pool

	wp := streamer.New(videoQueue, numWorkers)

	//Start the worker pool

	wp.Run()
	fmt.Println("Worker pool started. Press enter to continue.")
	_, _ = fmt.Scanln()

	//Create 4 videos to send to the worker pool
	ops := &streamer.VideoOptions{
		SegmentDuration: 10,
		MaxRate1080p:    "1200k",
		MaxRate720p:     "600k",
		MaxRate480p:     "400k",
	}
	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "hls", notifyChan, ops)
	//Send the videos to the worker pool

	videoQueue <- streamer.VideoProcessingJob{Video: video}
	//Print out the results

	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i:", i, "msg:", msg)
	}
	fmt.Println("Done!")

}
