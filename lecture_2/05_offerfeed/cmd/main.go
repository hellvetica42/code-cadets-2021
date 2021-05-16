package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"
	"context"
	"time"
)

func main() {

	cntx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	offerFeed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()
	feedProcessor := services.NewFeedProcessorService(offerFeed, queue)

	go offerFeed.Start(cntx)
	go queue.Start(cntx)
	go feedProcessor.Start(cntx)

	time.Sleep(time.Second*5)
}
