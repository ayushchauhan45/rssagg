package main

import (
	"context"
	"log"
	"time"
	"sync"

	"github.com/ayushchauhan_45/rssagg/internal/database"
)



func startScrapping(
	db *database.Queries,
	concurrency int, 
	timeBetweenRequest time.Duration,
){
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
    
	ticker:= time.NewTicker(timeBetweenRequest)
	
	for ; ; <-ticker.C {
		feeds,err:= db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err!=nil{
			log.Println("Error fetching feeds:",err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _,feed := range feeds{
			wg.Add(1)

			go srcapFeed(db,wg,feed)
		}
        wg.Wait()
	}

}

func srcapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed){
	defer wg.Done()
	_,err:= db.MarkFeedAsFetched(context.Background(),feed.ID)

	if err!=nil{
		log.Println("Error marking feed as fetched:",err)
	}
	rssFeed,err:= urlToFeed(feed.Url)
	if err!=nil{
		log.Println("Error fetching feed:",err)
		return
	}

	for _, item := range rssFeed.Channel.Items{
		log.Println("Found post",item, "on feed",feed.Name)
	}
	log.Printf("Feed %s collected ,%v posts found",feed.Name,len(rssFeed.Channel.Items))
}