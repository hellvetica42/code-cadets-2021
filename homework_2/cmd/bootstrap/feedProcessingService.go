package bootstrap

import "code-cadets-2021/lecture_2/06_offerfeed/internal/domain/services"

func FeedProcessingService(queue services.Queue, feed ...services.Feed) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feed, queue)
}
