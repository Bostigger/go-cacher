package main

import (
	"fmt"
	"github.com/bostigger/go-cacher/controllers"
)

func main() {
	fmt.Printf("GO CACHING <LRU>")
	cache := controllers.NewCache()
	for _, movies := range []string{"Avatar", "Ben 10", "Blacklist", "Silo", "Money Heist", "Blacklist", "Merlin", "Seeker"} {
		cache.Check(movies)
		cache.Display()
	}
}
