package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/redis.v3"
)

func NewClient(url string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func copyData(keys []string, from *redis.Client, to *redis.Client, force *bool) {
	for _, key := range keys {
		dump, err := from.Dump(key).Result()
		if err != nil {
			log.Fatal(err)
		}
		if *force && to.Exists(key).Val() {
			log.Println("Deleting existing key due to --force:", key)
			_, err = to.Del(key).Result()
			if err != nil {
				log.Fatal(err.Error())
			}
		}
		_, err = to.Restore(key, 0, dump).Result()
		if err != nil {
			log.Println("Warning", err.Error(), "skipping")
		}
	}
}

func main() {
	force := flag.Bool("force", false, "Overwrite existing keys")
	help := flag.Bool("help", false, "Show help text")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: redis-copy [options] <from> <to>\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	log.Printf("Copying data from: '%s', to: '%s'\n", flag.Args()[0], flag.Args()[1])
	from := NewClient(flag.Args()[0])
	to := NewClient(flag.Args()[1])
	var cursor int64
	var keys []string
	var err error

	for {
		scanCmd := from.Scan(cursor, "*", 100)
		cursor, keys, err = scanCmd.Result()
		if err != nil {
			log.Fatal(err)
		}
		copyData(keys, from, to, force)
		if cursor == 0 {
			break
		}
	}

}
