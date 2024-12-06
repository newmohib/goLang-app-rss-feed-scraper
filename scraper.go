package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/newmohib/goLang-app-rss-feed-scraper/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %v duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Println("error fetching feeds: ", err)
			continue
		}

		// fetch feeds concurrently
		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Println("error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found Post:\n", "Titel : ", item.Title, "\n Name : ", feed.Name)

		// _, err := db.CreateItem(context.Background(), database.CreateItemParams{
		// 	ID:        uuid.New(),
		// 	CreatedAt: time.Now().UTC(),
		// 	UpdatedAt: time.Now().UTC(),
		// 	Title:     item.Title,
		// 	Link:      item.Link,
		// 	FeedID:    feed.ID,
		// })

		// if err != nil {
		// 	log.Println("error creating item: ",err)
		// 	return
		// }
	}
	log.Println("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
