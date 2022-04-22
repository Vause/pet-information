package shared

type Pet interface {
	GetName() string
}

// Represents the Image section
type Image struct {
	Height int64  `json:"height"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

// Represents the pet height section
type Height struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

// Represents the pet weight section
type Weight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}
