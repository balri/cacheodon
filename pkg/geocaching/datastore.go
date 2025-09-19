package geocaching

type Attribute int

const (
	Dogs                  Attribute = 1
	AccessOrParkingFee    Attribute = 2
	ClimbingGear          Attribute = 3
	Boat                  Attribute = 4
	ScubaGear             Attribute = 5
	RecommendedForKids    Attribute = 6
	TakesLessThanAnHour   Attribute = 7
	ScenicView            Attribute = 8
	SignificantHike       Attribute = 9
	DifficultClimbing     Attribute = 10
	MayRequireWading      Attribute = 11
	MayRequireSwimming    Attribute = 12
	AvailableAtAllTimes   Attribute = 13
	RecommendedAtNight    Attribute = 14
	AvailableDuringWinter Attribute = 15
	CactiNearby           Attribute = 16
	PoisonPlants          Attribute = 17
	DangerousAnimals      Attribute = 18
	Ticks                 Attribute = 19
	AbandonedMines        Attribute = 20
	CliffsFallingRocks    Attribute = 21
	Hunting               Attribute = 22
	DangerousArea         Attribute = 23
	WheelchairAccessible  Attribute = 24
	ParkingAvailable      Attribute = 25
	PublicTransportation  Attribute = 26
	DrinkingWaterNearby   Attribute = 27
	PublicRestroomsNearby Attribute = 28
	TelephoneNearby       Attribute = 29
	PicnicTablesNearby    Attribute = 30
	CampingAvailable      Attribute = 31
	Bicycles              Attribute = 32
	Motorcycles           Attribute = 33
	Quads                 Attribute = 34
	OffRoadVehicles       Attribute = 35
	Snowmobiles           Attribute = 36
	Horses                Attribute = 37
	Campfires             Attribute = 38
	Thorns                Attribute = 39
	StealthRequired       Attribute = 40
	StrollerAccessible    Attribute = 41
	NeedsMaintenance      Attribute = 42
	WatchForLivestock     Attribute = 43
	FlashlightRequired    Attribute = 44
	LostAndFoundTour      Attribute = 45
	TruckDriverRV         Attribute = 46
	FieldPuzzle           Attribute = 47
	UVLightRequired       Attribute = 48
	Snowshoes             Attribute = 49
	CrossCountrySkis      Attribute = 50
	SpecialToolRequired   Attribute = 51
	NightCache            Attribute = 52
	ParkAndGrab           Attribute = 53
	AbandonedStructure    Attribute = 54
	ShortHike             Attribute = 55
	MediumHike            Attribute = 56
	LongHike              Attribute = 57
	FuelNearby            Attribute = 58
	FoodNearby            Attribute = 59
	WirelessBeacon        Attribute = 60
	PartnershipCache      Attribute = 61
	SeasonalAccess        Attribute = 62
	TouristFriendly       Attribute = 63
	TreeClimbing          Attribute = 64
	FrontYard             Attribute = 65
	TeamworkRequired      Attribute = 66
	GeoTour               Attribute = 67
	BonusCache            Attribute = 69
	PowerTrail            Attribute = 70
	ChallengeCache        Attribute = 71
	SolutionChecker       Attribute = 72
)

type CacheSize int

const (
	NotChosen   CacheSize = 1
	Micro       CacheSize = 2
	Regular     CacheSize = 3
	Large       CacheSize = 4
	VirtualSize CacheSize = 5
	Other       CacheSize = 6
	Small       CacheSize = 8
)

type CacheType int

const (
	Traditional    CacheType = 2
	Multi          CacheType = 3
	Virtual        CacheType = 4
	Letterbox      CacheType = 5
	Event          CacheType = 6
	Unknown        CacheType = 8
	APE            CacheType = 9
	Webcam         CacheType = 11
	Locationless   CacheType = 12
	CITO           CacheType = 13
	Earthcache     CacheType = 137
	Mega           CacheType = 453
	GPSMaze        CacheType = 1304
	Wherigo        CacheType = 1858
	CommunityEvent CacheType = 3653
	HQCache        CacheType = 3773
	HQCelebration  CacheType = 3774
	BlockParty     CacheType = 4738
	Giga           CacheType = 7005
)

type Difficulty float32

const (
	Difficulty1   Difficulty = 1.0
	Difficulty1_5 Difficulty = 1.5
	Difficulty2   Difficulty = 2.0
	Difficulty2_5 Difficulty = 2.5
	Difficulty3   Difficulty = 3.0
	Difficulty3_5 Difficulty = 3.5
	Difficulty4   Difficulty = 4.0
	Difficulty4_5 Difficulty = 4.5
	Difficulty5   Difficulty = 5.0
)

type OriginType string

const (
	City           OriginType = "city"
	Region         OriginType = "region"
	Country        OriginType = "country"
	GeocacheOrigin OriginType = "geocache"
	Coords         OriginType = "coords"
	Query          OriginType = "query"
)

type SortType string

const (
	Distance       SortType = "distance"
	GeocacheName   SortType = "geocacheName"
	FavouritePoint SortType = "favoritePoint"
	ContainerSize  SortType = "containerSize"
	DifficultySort SortType = "difficulty"
	TerrainSort    SortType = "terrain"
	TrackableCount SortType = "trackableCount"
	FoundDate      SortType = "foundDate"
	PlacedDate     SortType = "placedDate"
)

type Terrain float32

const (
	Terrain1   Terrain = 1.0
	Terrain1_5 Terrain = 1.5
	Terrain2   Terrain = 2.0
	Terrain2_5 Terrain = 2.5
	Terrain3   Terrain = 3.0
	Terrain3_5 Terrain = 3.5
	Terrain4   Terrain = 4.0
	Terrain4_5 Terrain = 4.5
	Terrain5   Terrain = 5.0
)

type SearchTerms struct {
	Latitude       float32      `json:"lat,omitempty"`
	Longitude      float32      `json:"lon,omitempty"`
	RadiusMeters   int          `json:"r,omitempty"`
	IgnorePremium  bool         // deprecated: use ShowPremium
	ShowPremium    *bool        `json:"sp,omitempty"`
	ShowDisabled   *bool        `json:"sd,omitempty"`
	ShowArchived   *bool        `json:"sa,omitempty"`
	SearchTerm     string       `json:"st,omitempty"`
	SortAsc        *bool        `json:"asc,omitempty"`
	OriginType     OriginType   `json:"op,omitempty"`
	OriginID       string       `json:"oid,omitempty"`
	HideOwned      *bool        `json:"ho,omitempty"`
	HideFound      *bool        `json:"hf,omitempty"`
	FillGrid       []string     `json:"m,omitempty"`
	NotFoundBy     []string     `json:"nfb,omitempty"`
	CacheName      string       `json:"cn,omitempty"`
	CacheSize      []CacheSize  `json:"cs,omitempty"`
	CacheType      []CacheType  `json:"ct,omitempty"`
	Difficulty     []Difficulty `json:"d,omitempty"`
	Terrain        []Terrain    `json:"t,omitempty"`
	FoundAfter     string       `json:"fad,omitempty"`
	FoundBefore    string       `json:"fbd,omitempty"`
	FoundEnd       string       `json:"fed,omitempty"`
	FoundStart     string       `json:"fsd,omitempty"`
	FoundOn        string       `json:"fod,omitempty"`
	PlacedAfter    string       `json:"pad,omitempty"`
	PlacedBefore   string       `json:"pbd,omitempty"`
	PlacedEnd      string       `json:"ped,omitempty"`
	PlacedStart    string       `json:"psd,omitempty"`
	PlacedOn       string       `json:"pod,omitempty"`
	FoundBy        []string     `json:"fb,omitempty"`
	HiddenBy       string       `json:"hb,omitempty"`
	Attributes     []Attribute  `json:"att,omitempty"`
	Sort           SortType     `json:"sort,omitempty"`
	Corrected      *bool        `json:"cc,omitempty"`
	FavouriteCount int          `json:"fp,omitempty"`
	PersonalNote   *bool        `json:"pn,omitempty"`
}

type APIConfig struct {
	// The URL of the Geocaching API.
	GeocachingAPIURL string
	HTTPProxyURL     string
	UnThrottle       bool // Should we disable rate-limiting for this API?
}
