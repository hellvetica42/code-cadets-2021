package services

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"context"
)

type FeedProcessorService struct {
	feed Feed
	queue Queue
}

func NewFeedProcessorService(feed Feed, queue Queue) *FeedProcessorService {
	// it should receive "Feed" & "Queue" interfaces through constructor
	return &FeedProcessorService{
		feed: feed,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	updates := f.feed.GetUpdates()
	// - get source channel from queue interface
	source := f.queue.GetSource()

	defer close(source)
	// // repeatedly:
	// - range over updates channel
	for up := range updates {
		// - multiply each odd with 2
		up.Coefficient *= 2
		// - send it to source channel
		source <- up
	}
	return nil
	//
	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel
}

// define feed interface here
type Feed interface {
	GetUpdates() chan models.Odd
}

// define queue interface here
type Queue interface {
	GetSource() chan models.Odd
}
