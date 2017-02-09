package cmd

type Connpass struct {
	ResultsReturned  int `json:"results_returned"`
	ResultsAvailable int `json:"results_available"`
	ResultsStart     int `json:"results_start"`
	Events           []Event `json:"events"`
}

type Event struct {
	EventId          int `json:"event_id"`
	Title            string `json:"title"`
	Catch            string `json:"catch"`
	Description      string `json:"description"`
	EventUrl         string `json:"event_url"`
	HashTag          string `json:"hash_tag"`
	StartedAt        string `json:"started_at"`
	EndedAt          string `json:"ended_at"`
	Limit            int `json:"limit"`
	EventType        string `json:"event_type"`
	Series                        `json:"series"`
	Address          string `json:"address"`
	Place            string `json:"place"`
	//Lat              float32 `json:lat`
	//Lon              float32 `json:lon`
	OwnerId          int `json:"owner_id"`
	OwnerNickname    string `json:"owner_nickname"`
	OwnerDisplayName string `json:"owner_display_name"`
	Accepted         int `json:"accepted"`
	Waiting          int `json:"waiting"`
	Updated_at       string `json:"updated_at"`
}

type Series struct {
	Id    int `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}