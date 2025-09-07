package geocaching

type SearchTerms struct {
	Latitude      float32
	Longitude     float32
	RadiusMeters  int
	AreaName      string
	IgnorePremium bool
}

type APIConfig struct {
	// The URL of the Geocaching API.
	GeocachingAPIURL string
	HTTPProxyURL     string
	UnThrottle       bool // Should we disable rate-limiting for this API?
}
