package elasticsearch

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	elasticmodel "git.ahamove.com/searare-common/searare-kit/elasticsearch/models"

	ES "git.ahamove.com/searare-common/searare-kit/elasticsearch"
	"github.com/stretchr/testify/assert"
)

var ESConfig = ES.Config{
	Host:                  []string{"http://103.110.85.105:9200"},
	Username:              "elastic",
	Password:              "bWlVNrl77CMznG7ICZCI",
	TimeoutConnect:        120,
	MaxIdleConnsPerHost:   20,
	PoolSize:              20,
	RetryOnStatus:         []int{502, 503, 504},
	DisableRetry:          false,
	EnableRetryOnTimeout:  true,
	MaxRetries:            5,
	RetriesTimeout:        10,
	ResponseHeaderTimeout: 15,
	Index:                 "feed",
}

// Test connect ElasticSearch
func TestESConnection(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
}

// Test create index

func TestCreateIndex(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	//
	setting := `{"settings":{"index":{"number_of_shards":4,"number_of_replicas":1}}}`
	var body map[string]interface{}
	json.Unmarshal([]byte(setting), &body)

	indexName := "searare_feed_test"
	ES.ESConnector.CreateIndex(indexName, body)
}

// Test delete index

func TestDeleteIndex(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	//
	setting := `{"acknowledged":true}`
	var body map[string]interface{}
	json.Unmarshal([]byte(setting), &body)

	indexName := "searare_feed_test"
	ES.ESConnector.DeleteIndex(indexName, body)
}
func TestUpdateIndexSettings(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	//
	setting := `{"index" : {"number_of_replicas" : 6}}`
	indexName := "searare_feed_test"
	response, err := ES.ESConnector.UpdateIndexSettings(indexName, setting)
	assert.Equal(t, err, nil)
	fmt.Println(response, err)
}

// Test Document
func TestCreateOne(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	post := elasticmodel.Post{
		Content: "Hello Ahamove!",
		Tags:    []string{"greeting", "guide"},
		Stats: elasticmodel.PostStats{
			CommentCount: 10,
			ViewCount:    44,
			LikeCount:    32,
		},
		Attachments: []*elasticmodel.PostAttachment{
			{
				URL:     "http://ahamove.com/images/1.png",
				Content: "selfpng",
			},
			{
				URL:     "http://ahamove.com/images/1.png",
				Content: "selfpng",
			},
		},
		Status:           "OK",
		CreatedAt:        float64(time.Now().Unix()),
		UpdatedAt:        float64(time.Now().Unix()),
		InteractionScore: 12,
		Type:             elasticmodel.PostTypeFree,
		Place: &elasticmodel.PostPlace{
			Address:  "HCM",
			Name:     "Q10",
			Location: &elasticmodel.Location{Lat: 11, Lng: 33},
		},
		Categories: []string{"greeting"},
	}
	response, err := ES.ESConnector.CreateOne(index, "", post)
	assert.Equal(t, err, nil)
	t.Log(response, err)
}

func TestUpdateOne(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	id := "_ylJ_IEBTIi260cNkVZX"
	post := elasticmodel.Post{
		Content: "Hello Ahamove update!",
		Tags:    []string{"greeting", "guide"},
		Stats: elasticmodel.PostStats{
			CommentCount: 10,
			ViewCount:    44,
			LikeCount:    32,
		},
		Attachments: []*elasticmodel.PostAttachment{
			{
				URL:     "http://ahamove.com/images/1.png",
				Content: "selfpng",
			},
			{
				URL:     "http://ahamove.com/images/1.png",
				Content: "selfpng",
			},
		},
		Status:           "OK",
		CreatedAt:        float64(time.Now().Unix()),
		UpdatedAt:        float64(time.Now().Unix()),
		InteractionScore: 12,
		Type:             elasticmodel.PostTypeFree,
		Place: &elasticmodel.PostPlace{
			Address:  "HCM",
			Name:     "Q10",
			Location: &elasticmodel.Location{Lat: 11, Lng: 33},
		},
		Categories: []string{"greeting"},
	}
	response, err := ES.ESConnector.CreateOne(index, id, post)
	assert.Equal(t, err, nil)
	t.Log(response, err)
}
func TestFindAll(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	search := map[string]interface{}{
		"content": "Ahamove",
	}
	response, err := ES.ESConnector.FindAllMatch(index, search)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(response, err)
	assert.Equal(t, err, nil)
}

func TestFindOne(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	search := map[string]interface{}{
		"content": "Ahamove",
	}

	response, err := ES.ESConnector.FindOneMatch(index, search)
	if err != nil {
		fmt.Println(err)
		return
	}
	// convert to ESCreateDocResponse
	post := elasticmodel.Post{}
	err = json.Unmarshal(response.Source, &post)
	t.Log(post, err)
	assert.Equal(t, err, nil)
}
func TestFindPaging(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	search := map[string]interface{}{
		"content": "Ahamove",
	}
	response, err := ES.ESConnector.FindMatchPaging(index, search, 0, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(response, err)
	assert.Equal(t, err, nil)
}
func TestFindMatchRange(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	rangeSearch := ES.RangeSearch{
		Key: "created_at",
	}
	search := map[string]interface{}{
		"content": "Quang",
	}
	response, err := ES.ESConnector.FindMatchWithRange(index, search, rangeSearch)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(response, err)
	assert.Equal(t, err, nil)
}

func TestFindMatchRangePaging(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	rangeSearch := ES.RangeSearch{
		Key: "created_at",
	}
	search := map[string]interface{}{
		"content": "Quang",
	}
	offset := 0
	limit := 1
	response, err := ES.ESConnector.FindMatchWithRangePaging(index, search, offset, limit, rangeSearch)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(response, err)
	assert.Equal(t, err, nil)
}

func TestCountMatch(t *testing.T) {
	ES.ESConnector, _ = ES.NewElasticClient(ESConfig)
	ES.ESConnector.Ping()
	index := "searare_feed_test"
	search := map[string]interface{}{
		"content": "Welcome",
	}
	response, err := ES.ESConnector.CountMatchDocuments(index, search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
	t.Log(response, err)
	assert.Equal(t, err, nil)
}
