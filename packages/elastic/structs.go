package elastic

type ElasticAggregationResponse struct {
	Aggregations ElasticAggregation
}

type ElasticAggregation struct {
	FullName ElasticAllBuckets `json:"full_name"`
}

type ElasticAllBuckets struct {
	Buckets []ElasticBucket
}

type ElasticBucket struct {
	Key   string
	Count int `json:"doc_count"`
}

type AllHostnames struct {
	Hostnames []Hostname
}

type Hostname struct {
	Name string
}
