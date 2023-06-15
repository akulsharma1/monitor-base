package resources

type Message struct {
	Store             string                        `json:"store"`
	PolymarketMessage []PolyMarketNewListingAPIResp `json:"polymarket_message"`
}

type PolyMarketNewListingAPIResp struct {
	Store            string
	ID               string      `json:"id"`
	Ticker           string      `json:"ticker"`
	Slug             string      `json:"slug"`
	Title            string      `json:"title"`
	Subtitle         interface{} `json:"subtitle"`
	Description      string      `json:"description"`
	ResolutionSource string      `json:"resolutionSource"`
	StartDate        interface{} `json:"startDate"`
	CreationDate     string      `json:"creationDate"`
	EndDate          string      `json:"endDate"`
	Image            string      `json:"image"`
	Icon             string      `json:"icon"`
	FeaturedImage    interface{} `json:"featuredImage"`
	Active           bool        `json:"active"`
	Closed           bool        `json:"closed"`
	Archived         bool        `json:"archived"`
	New              bool        `json:"new"`
	Featured         bool        `json:"featured"`
	Restricted       bool        `json:"restricted"`
	Liquidity        int         `json:"liquidity"`
	Volume           float64     `json:"volume"`
	Volume24Hr       float64     `json:"volume24hr"`
	OpenInterest     int         `json:"openInterest"`
	SortBy           string      `json:"sortBy"`
	Markets          []struct {
		ID            string      `json:"id"`
		Slug          string      `json:"slug"`
		EndDate       string      `json:"endDate"`
		Liquidity     string      `json:"liquidity"`
		StartDate     interface{} `json:"startDate"`
		Outcomes      []string    `json:"outcomes"`
		OutcomePrices []string    `json:"outcomePrices"`
		Volume        string      `json:"volume"`
		Active        bool        `json:"active"`
		FormatType    interface{} `json:"formatType"`
		Closed        bool        `json:"closed"`
		ClosedTime    interface{} `json:"closedTime"`
		New           bool        `json:"new"`
		Featured      bool        `json:"featured"`
		Archived      bool        `json:"archived"`
		Restricted    bool        `json:"restricted"`
		MarketGroup   interface{} `json:"marketGroup"`
		EndDateIso    string      `json:"endDateIso"`
		StartDateIso  interface{} `json:"startDateIso"`
		UmaEndDateIso string      `json:"umaEndDateIso"`
		Question      string      `json:"question"`
		Typename      string      `json:"__typename"`
	} `json:"markets"`
	Categories []struct {
		ID             string `json:"id"`
		Label          string `json:"label"`
		ParentCategory string `json:"parentCategory"`
		Slug           string `json:"slug"`
		Typename       string `json:"__typename"`
	} `json:"categories"`
	Typename string `json:"__typename"`
}