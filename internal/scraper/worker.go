package scraper

import (
	"context"
	"fmt"
	"internal/api"
	"internal/database"
	"sync"
	"time"
)

func ScraperWorker() {

	fmt.Println("Scraper worker started")
	// Get 10 feeds every 60 seconds
	for {
		var wg sync.WaitGroup
		config := api.GetAPIConfig()
		feeds, err := config.DB.GetNextFeedsToFetch(context.TODO(), 1)
		if err != nil {
			fmt.Println("Error getting feeds")
		}
		for _, feed := range feeds {
			wg.Add(1)
			go func(feed database.Feed) {
				defer wg.Done()
				fmt.Println("Fetching feed: ", feed.Url)
				rss := FetchFeeds(feed.Url)
				config.DB.MarkFeedAsFetched(context.TODO(), database.MarkFeedAsFetchedParams{
					ID: feed.ID,
				})
				for _, item := range rss.Channel.Item {
					fmt.Println("Post: ", item.Title)
				}
			}(feed)
		}
		wg.Wait()
		time.Sleep(60 * time.Second)
	}

}