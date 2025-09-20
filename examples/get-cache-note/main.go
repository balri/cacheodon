package main

import (
	"fmt"
	"os"

	"github.com/balri/cacheodon/pkg/geocaching"
	log "github.com/sirupsen/logrus"
)

func BoolPtr(b bool) *bool {
	return &b
}

func main() {
	config := geocaching.APIConfig{
		GeocachingAPIURL: "https://www.geocaching.com",
	}

	client, err := geocaching.NewGeocachingAPI(config)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Auth(
		os.Getenv("GEOCACHING_CLIENT_ID"),
		os.Getenv("GEOCACHING_CLIENT_SECRET"),
	)
	if err != nil {
		log.Fatal(err)
	}

	cache := geocaching.Geocache{
		Code: "GC7VKAP",
	}

	note, err := client.GetCacheNoteForGeocache(cache)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cache Note for %s (%s):\n%s\n", cache.Code, cache.Name, note)
}
