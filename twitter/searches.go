package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Tweets struct {
	Statuses []Tweet `json:"statuses"`
}

type SearchService struct {
	sling *sling.Sling
}

func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("search/"),
	}
}

type SearchTweetsParams struct {
	Query           string `url:"q,omitempty"`
	Geocode         int64  `url:"geocode,omitempty"`
	Lang            string `url:"lang,omitempty"`
	Local           string `url:"local,omitempty"`
	ResultType      string `url:result_type,omitempty`
	Count           int    `url:"count,omitempty"`
	Until           string `url:"until,omitempty"`
	SinceID         int64  `url:"since_id,omitempty"`
	MaxID           int64  `url:"max_id,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"`
	Callback        string `url:"callback,omitempty"`
}

func (s *SearchService) Tweets(params *SearchTweetsParams) (Tweets, *http.Response, error) {
	searchTweets := new(Tweets)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("tweets.json").QueryStruct(params).Receive(searchTweets, apiError)
	return *searchTweets, resp, relevantError(err, *apiError)
}
