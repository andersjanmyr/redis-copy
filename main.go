package main

import (
	"log"
	"os"

	"gopkg.in/redis.v3"
)

func NewClient(url string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func copyData(keys []string, from *redis.Client, to *redis.Client) {
	for _, key := range keys {
		dump, err := from.Dump(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		_, err = to.Restore(key, 0, dump).Result()
		if err != nil {
			log.Println("Warning", err.Error(), "skipping")
		}
	}
}

func main() {
	log.Printf("Copying data from: '%s', to: '%s'\n", os.Args[1], os.Args[2])
	from := NewClient(os.Args[1])
	to := NewClient(os.Args[2])
	var cursor int64
	var keys []string
	var err error

	for {
		scanCmd := from.Scan(cursor, "*", 100)
		cursor, keys, err = scanCmd.Result()
		if err != nil {
			log.Fatal(err)
		}
		copyData(keys, from, to)
		if cursor == 0 {
			break
		}
	}

}
