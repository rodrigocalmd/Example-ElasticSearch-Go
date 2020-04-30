package types

import "time"

type Document struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Score       float64   `json:"score"`
	Create      time.Time `json:"date"`
}

type ResponseES struct {
	ID            string   `json:"_id"`
	Index         string   `json:"_index"`
	PrimaryTerm   int64    `json:"_primary_term"`
	SeqNo         int64    `json:"_seq_no"`
	Shards        ShardsES `json:"_shards"`
	Type          string   `json:"_type"`
	Version       int64    `json:"_version"`
	ForcedRefresh bool     `json:"forced_refresh"`
	Result        string   `json:"result"`
}

type ShardsES struct {
	Failed     int64 `json:"failed"`
	Skipped    int64 `json:"skipped,omitempty"`
	Successful int64 `json:"successful"`
	Total      int64 `json:"total"`
}

type SearchResponseES struct {
	Shards   ShardsES `json:"_shards"`
	Hits     HitsES   `json:"hits"`
	TimedOut bool     `json:"timed_out"`
	Took     int64    `json:"took"`
}

type HitsES struct {
	Hits     []HitResponseES     `json:"hits"`
	MaxScore float64             `json:"max_score"`
	Total    HitsResponseTotalES `json:"total"`
}

type HitResponseES struct {
	ID     string   `json:"_id"`
	Index  string   `json:"_index"`
	Score  float64    `json:"_score"`
	Source Document `json:"_source"`
	Type   string   `json:"_type"`
}

type HitsResponseTotalES struct {
	Relation string `json:"relation"`
	Value    int64  `json:"value"`
}
