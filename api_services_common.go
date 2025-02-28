package broadpeakio

// Misc types start
type GetAllSlotsInput struct {
	Offset     uint
	Limit      uint
	Categories []uint
	From       string
	To         string
}

type AdBreakInsertion struct {
	AdServer  *Identifiable `json:"adServer,omitempty"`
	GapFilter *Identifiable `json:"gapFilter,omitempty"`
}

type ServerSideAdTracking struct {
	Enable                          bool `json:"enable,omitempty"`
	CheckAdMediaSegmentAvailability bool `json:"checkAdMediaSegmentAvailability,omitempty"`
}

type LiveAdPreRoll struct {
	AdServer    *Identifiable `json:"adServer,omitempty"` //required
	MaxDuration uint          `json:"maxDuration,omitempty"`
	Offset      uint          `json:"offset,omitempty"`
}

type LiveAdReplacement struct {
	AdServer  *Identifiable `json:"adServer,omitempty"` //required
	GapFiller *Identifiable `json:"gapFiller,omitempty"`
	SpotAware SpotAware     `json:"spotAware,omitempty"`
}

type SpotAware struct {
	Mode string `json:"mode,omitempty"`
}

type VodAdInsertion struct {
	AdServer *Identifiable `json:"adServer,omitempty"` //required
}

type QueryParam struct {
	Name  string `json:"name"`
	Value string `json:"url"`
	Type  string `json:"type"`
}

type AdServer struct {
	Id              uint         `json:"id"`
	Name            string       `json:"name"`
	Url             string       `json:"url"`
	Type            string       `json:"type"`
	QueryParameters []QueryParam `json:"queryParameters"`
}

type Source struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Type   string `json:"type"`
	Origin Origin `json:"origin"`
}

type Origin struct {
	CustomHeaders []CustomHeader `json:"customHeaders"`
}

type CustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GapFiller struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

type TranscodingProfile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Id      uint   `json:"id"`
}

type LiveAdPreRollOutput struct {
	AdServer    AdServer `json:"adServer,omitempty"`
	MaxDuration uint     `json:"maxDuration"`
	Offset      uint     `json:"offset"`
}

type LiveAdReplacementOutput struct {
	AdServer  AdServer  `json:"adServer"`
	GapFiller GapFiller `json:"gapFiller"`
}

type AdBreakInsertionOutput struct {
	AdServer  AdServer  `json:"adServer"`
	GapFiller GapFiller `json:"gapFiller"`
}

type VodAdInsertionOutput struct {
	AdServer AdServer `json:"adServer"`
}

// Misc types end
