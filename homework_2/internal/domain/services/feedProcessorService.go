package services

import (
	"context"
	"log"
	"sync"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
)

type FeedProcessorService struct {
	feeds  []Feed
	queue Queue
}

func NewFeedProcessorService(
	feeds []Feed,
	queue Queue,
) *FeedProcessorService {
	return &FeedProcessorService{
		feeds:  feeds,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	var updates []chan models.Odd

	for _, feed := range f.feeds {
		updates = append(updates, feed.GetUpdates())
	}
	source := f.queue.GetSource()

	defer close(source)
	defer log.Printf("shutting down %s", f)

	wg := &sync.WaitGroup{}
	wg.Add(len(updates))

	for _, ch := range updates {
		go func(c chan models.Odd) {
			defer wg.Done()

			for update := range c {
				update.Coefficient *= 2
				source <- update
			}
		}(ch)
	}

	wg.Wait()

	return nil
}

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
