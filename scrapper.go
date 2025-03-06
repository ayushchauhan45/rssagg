package main

import (
	"log"
	"time"

	"github.com/ayushchauhan_45/rssagg/internal/database"
)



func startScrapping(
	db *database.Queries,
	concurrency int, 
	timeBetweenRequest time.Duration,
){
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
    
	time.NewTicker(timeBetweenRequest)
	

}