package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func readLines(fpath string) []string {
	content, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatalf("Can't find %s", fpath)
	}
	lines := strings.Split(strings.Trim(string(content), "\n"), "\n")
	return lines
}

func seedFromFile(fpath string, rkey string, db *redis.Client) {
	seeds := readLines(fpath)

	pipe := db.Pipeline()
	for _, seed := range seeds {
		pipe.SAdd(rkey, seed)
	}
	_, err := pipe.Exec()
	if err != nil {
		log.Println("Could not write to redis!")
		log.Fatal(err)
	}

	log.Printf("Added %d shouts from %s to the database at %s\n", len(seeds), fpath, rkey)
}

func removeFromFile(fpath string, rkey string, db *redis.Client) {
	seeds := readLines(fpath)

	pipe := db.Pipeline()
	for _, seed := range seeds {
		pipe.SRem(rkey, seed)
	}
	_, err := pipe.Exec()
	if err != nil {
		log.Println("Could not write to redis!")
		log.Fatal(err)
	}

	log.Printf("Removed %d shouts from %s to the database at %s\n", len(seeds), fpath, rkey)
}

func main() {
	loaded := godotenv.Load("../../.env")
	if loaded != nil {
		log.Println("No .env file found; using defaults")
	}

	prefix, found := os.LookupEnv("REDIS_KEY")
	if !found {
		prefix = "LB"
	}

	var rkey string
	rkey = fmt.Sprintf("%s:YELLS", prefix)

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPort = os.Getenv("REDIS_PORT")
	var redisPassword = os.Getenv("REDIS_PASSWORD")
	log.Printf("using redis @ %s to store our data", redisHost)
	db := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	removeFromFile("SYSTEMANTICS", rkey, db)
	removeFromFile("STAR_FIGHTING", rkey, db)
	seedFromFile("SEEDS", rkey, db)
	rkey = fmt.Sprintf("%s:CATS", prefix)
	seedFromFile("CATS", rkey, db)
	rkey = fmt.Sprintf("%s:SW", prefix)
	seedFromFile("STAR_FIGHTING", rkey, db)
}
