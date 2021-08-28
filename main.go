package main

import (
	"example.com/database/crawl_data/crawler"
	"example.com/database/todb"
)

func main() {
	db := todb.ConnectDB()
	crawler.Crawl(db)

	// api.InitiallizeServer()

}
