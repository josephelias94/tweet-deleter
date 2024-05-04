package twitter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/models"
	"github.com/josephelias94/tweet-deleter/internals/validator"
)

type Client struct {
	Token string
	User  models.User
}

func buildUrl(url, wildcard, value string) string {
	return strings.ReplaceAll(url, wildcard, value)
}

func parseJson[T any](response []byte, v *T) {
	if json.Valid(response) == false {
		log.Fatalf("%v Provided response: \"%v\"", constants.ERROR_TW_JSON_INVALID, response)
	}

	if err := json.Unmarshal(response, &v); err != nil {
		log.Fatalf("%v Message: \"%v\"", constants.ERROR_TW_JSON_CONVERTING, err)
	}

	validator.ValidateFields(string(response), v)
}

func (c *Client) makeRequest(method, url string, body io.Reader) (int, []byte, error) {
	client := &http.Client{}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("%v Message: \"%v\"", constants.ERROR_TW_REQUEST_BUILDING, url)
	}

	request.Header.Add("Authorization", "Bearer "+c.Token)

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("%v Message: \"%v\"", constants.ERROR_TW_REQUEST_DURING, err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%v Message: \"%v\"", constants.ERROR_TW_REQUEST_RESPONSE, err)
	}

	return response.StatusCode, responseBody, nil
}

func (c *Client) SetUser(username string) {
	url := buildUrl(constants.GET_USER_BY_USERNAME, ":username", username)
	statusCode, response, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(constants.ERROR_TW_USER_REQUEST)
	}

	if statusCode != 200 {
		log.Fatalf("%v Response: \"%v\"", constants.ERROR_TW_USER_FAILED_STATUS_CODE, string(response))
	}

	body := models.GetUserResponse{}

	parseJson(response, &body)

	c.User = body.Data
}

func (c *Client) GetTweets() []models.Tweet {
	if c.User.Id == "" {
		log.Fatal(constants.ERROR_TW_TWEETS_USER_UNSET)
	}

	url := buildUrl(constants.GET_TWEETS_BY_USER, ":id", c.User.Id)

	statusCode, response, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(constants.ERROR_TW_TWEETS_REQUEST)
	}

	if statusCode != 200 {
		log.Fatalf("%v Response: \"%v\"", constants.ERROR_TW_TWEETS_FAILED_STATUS_CODE, string(response))
	}

	body := models.GetTweetsResponse{}

	parseJson(response, &body)

	return body.Data
}

func (c *Client) DeleteTweet(id string) bool {
	url := buildUrl(constants.DELETE_TWEET, ":id", id)

	statusCode, response, err := c.makeRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatalf("%v Id: %v", constants.ERROR_TW_T_DELETE_REQUEST, id)
	}

	if statusCode != 200 {
		log.Fatalf("%v Id: %v | Response: \"%v\"", constants.ERROR_TW_T_DELETE_FAILED_STATUS_CODE, id, string(response))
	}

	body := models.DeleteTweetResponse{}

	parseJson(response, &body)

	return body.Data.Deleted
}
