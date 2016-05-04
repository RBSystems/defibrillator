package elastic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetHostnames() (AllHostnames, error) {
	// Ask Elasticsearch for all hostnames via an aggregation
	var postBody = []byte(`{
  "aggs": {
    "full_name": {
      "terms": {
      "field": "device.hostname",
        "size": 0
      }
    }
  }
}`)

	req, err := http.NewRequest("POST", "http://search-byu-oit-av-metrics-ruenjnrqfuhghh7omvtmgcqe7m.us-west-1.es.amazonaws.com/events/_search", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return AllHostnames{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return AllHostnames{}, err
	}

	elasticResponse := ElasticAggregationResponse{}
	json.Unmarshal(body, &elasticResponse)

	hostnames := AllHostnames{}

	for i := range elasticResponse.Aggregations.FullName.Buckets {
		hostname := Hostname{Name: elasticResponse.Aggregations.FullName.Buckets[i].Key}

		hostnames.Hostnames = append(hostnames.Hostnames, hostname)
	}

	return hostnames, nil
}
