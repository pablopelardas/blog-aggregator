package scraper

import (
	"context"
	"fmt"
	"internal/api"
	"internal/database"
	"sync"
	"time"

	"github.com/google/uuid"
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
					pubDate, err := time.Parse(time.RFC1123Z, item.PubDate) 
					if err != nil {
						fmt.Printf("Error parsing date: %s", err)
						continue
					}
					_,err = config.DB.CreatePost(context.TODO(), database.CreatePostParams{
						FeedID: uuid.NullUUID{UUID: feed.ID, Valid: true},
						Title: item.Title,
						Url: item.Link,
						Description: item.Description,
						PublishedAt: pubDate,
						ID:          uuid.New(),
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					})
					if err != nil {
						// if error include duplicate key, ignore
						if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
							continue
						}
						fmt.Printf("Error creating post: %s", err)
					}
				}
			}(feed)
		}
		wg.Wait()
		time.Sleep(60 * time.Second)
	}

}