package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/models"
)

type Client struct {
	Token string
	User  models.User
}

func buildUrl(url, wildcard, value string) string {
	return strings.ReplaceAll(url, wildcard, value)
}

func parseJson[T any](response []byte, v *T) {
	if err := json.Unmarshal(response, &v); err != nil {
		log.Fatalf("twitter: Error converting json. Message: \"%v\"", err)
	}
}

func (c *Client) makeRequest(method, url string, body io.Reader) (int, []byte, error) {
	client := &http.Client{}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Printf("twitter: Error while building a request. Message: \"%v\"", url)
		return 0, nil, err
	}

	request.Header.Add("Authorization", "Bearer "+c.Token)

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("twitter: Error during a request. Message: \"%v\"", err)
		return 0, nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("twitter: Error getting the response. Message: \"%v\"", err)
		return 0, nil, err
	}

	return response.StatusCode, responseBody, nil
}

func (c *Client) SetUser(username string) {
	url := buildUrl(constants.GET_USER_BY_USERNAME, ":username", username)
	statusCode, response, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("tweet: Unable to make request to set user")
	}

	if statusCode != 200 {
		log.Fatalf("tweet: Request to get user failed. Response: %v", string(response))
	}

	body := models.GetUserResponse{}

	parseJson(response, &body)

	c.User = body.Data
}

func (c *Client) GetTweets() []models.Tweet {
	if c.User.Id == "" {
		log.Fatal("twitter: User is not set")
	}

	url := buildUrl(constants.GET_TWEETS_BY_USER, ":id", c.User.Id)

	statusCode, response, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("tweet: Unable to make request to get tweets")
	}

	if statusCode != 200 {
		log.Fatalf("tweet: Request to get tweets failed. Response: %v", string(response))
	}

	body := models.GetTweetsResponse{}

	parseJson(response, &body)

	return body.Data
}

func (c *Client) DeleteTweet(id string) bool {
	url := buildUrl(constants.DELETE_TWEET, ":id", id)

	statusCode, response, err := c.makeRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatalf("tweet: Unable to make request to delete a tweet. Id: %v", id)
	}

	if statusCode != 200 {
		log.Fatalf("tweet: Request to delete a tweet failed. Id: %v. Response: %v", id, string(response))
	}

	body := models.DeleteTweetResponse{}

	parseJson(response, &body)

	return body.Data.Deleted
}
