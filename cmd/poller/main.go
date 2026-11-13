// Command poller polls Rutgers schedule data on an interval and logs
// section open/close events to Postgres.
package main

import (
	"log"
	"os"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set (copy .env.example to .env)")
	}
	log.Println("poller starting")
	// TODO(day 2): fetch, parse, detect changes, store events, loop.
}
