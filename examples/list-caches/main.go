package main

import (
	"fmt"
	"log"
	"os"

	"github.com/balri/cacheodon/pkg/geocaching"
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

	searchTerms := geocaching.SearchTerms{
		Latitude:      -27.4698,
		Longitude:     153.0251,
		RadiusMeters:  25000,
		IgnorePremium: false,
		ShowDisabled:  BoolPtr(false),
		SortAsc:       BoolPtr(true),
		Sort:          "distance",
		OriginType:    "query",
		HideOwned:     BoolPtr(true),
		NotFoundBy:    []string{os.Getenv("GEOCACHING_CLIENT_ID")},
		CacheType: []geocaching.CacheType{
			geocaching.Traditional,
			geocaching.Multi,
			geocaching.Virtual,
			geocaching.Letterbox,
			geocaching.Unknown,
			geocaching.Webcam,
			geocaching.Earthcache,
			geocaching.Wherigo,
		},
	}

	caches, err := client.Search(searchTerms)
	if err != nil {
		log.Fatal(err)
	}

	for _, cache := range caches {
		fmt.Printf(
			"%s (%s) - %f,%f\n",
			cache.Code,
			cache.Name,
			cache.PostedCoordinates.Latitude,
			cache.PostedCoordinates.Longitude,
		)
	}
}
