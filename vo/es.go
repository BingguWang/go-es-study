package vo

type EsJsonResult struct {
	ScrollId string `json:"_scroll_id"`
	Took     int    `json:"took" mapstructure:"took"`
	TimedOut bool   `json:"time_out" mapstructure:"time_out"`
	Shards   *Shard `json:"_shards" mapstructure:"_shards"`
	Hits     *Hits  `json:"hits" mapstructure:"hits"`
}
type Shard struct {
	Total      int `json:"total" mapstructure:"total"`
	Successful int `json:"successful" mapstructure:"successful"`
	Skipped    int `json:"skipped" mapstructure:"skipped"`
	Failed     int `json:"failed" mapstructure:"failed"`
}
type Hits struct {
	Total    *Total `json:"total"  mapstructure:"total"`
	MaxScore int    `json:"maxScore"  mapstructure:"maxScore"`
	Hits     []*Hit `json:"hits" mapstructure:"hits"`
}
type Total struct {
	Value    int    `json:"value" mapstructure:"value"`
	Relation string `json:"relation"  mapstructure:"relation"`
}
type Hit struct {
	Index  string                 `json:"_index" mapstructure:"_index"`
	Id     string                 `json:"_id" mapstructure:"_id"`
	Score  float64                `json:"_score" mapstructure:"_score"`
	Source map[string]interface{} `json:"_source" mapstructure:"_source"`
}
