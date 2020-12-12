package main

import (
	"bufio"
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

func check(err error) {
	if err != nil {
		panic(err)
	}
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
	rkey := fmt.Sprintf("%s:YELLS", prefix)

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPort = os.Getenv("REDIS_PORT")
	var redisPassword = os.Getenv("REDIS_PASSWORD")
	log.Printf("using redis @ %s to store our data", redisHost)
	db := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	f, err := os.OpenFile("SAVED_LOUDS", os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	out := bufio.NewWriter(f)

	louds, err2 := db.SMembers(rkey).Result()
	check(err2)

	log.Printf("found %d louds to save\n", len(louds))

	for _, shout := range louds {
		out.WriteString(shout)
		out.WriteString("\n")
	}

	out.Flush()
}
