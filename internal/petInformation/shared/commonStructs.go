package shared

type Image struct {
	Height int64  `json:"height"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type Height struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

type Weight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}