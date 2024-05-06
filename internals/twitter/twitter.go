package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/models"
	"github.com/josephelias94/tweet-deleter/internals/validator"
)

type Client struct {
	AuthorizedClient *http.Client
	user             models.User
}

func buildUrl(url, wildcard, value string) string {
	return strings.ReplaceAll(url, wildcard, value)
}

func parseJson[T any](response []byte, v *T) error {
	if json.Valid(response) == false {
		return errors.New(constants.ERROR_TW_JSON_INVALID + "Provided response: " + string(response))
	}

	if err := json.Unmarshal(response, &v); err != nil {
		return errors.New(constants.ERROR_TW_JSON_CONVERTING + "ErrorMessage: " + err.Error())
	}

	if err := validator.ValidateFields(string(response), v); err != nil {
		return err
	}

	return nil
}

func (c *Client) makeRequest(method, url string, body io.Reader) (int, []byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return 0, nil, errors.New(constants.ERROR_TW_REQUEST_BUILDING + "ErrorMessage: " + err.Error())
	}

	response, err := c.AuthorizedClient.Do(request)
	if err != nil {
		return 0, nil, errors.New(constants.ERROR_TW_REQUEST_DURING + "ErrorMessage: " + err.Error())
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, nil, errors.New(constants.ERROR_TW_REQUEST_RESPONSE + "ErrorMessage: " + err.Error())
	}

	return response.StatusCode, responseBody, nil
}

func (c *Client) SetUser(username string) {
	url := buildUrl(constants.GET_USER_BY_USERNAME, ":username", username)

	statusCode, response, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if statusCode != 200 {
		log.Fatalf("%vResponse: \"%v\"", constants.ERROR_TW_S_USER_FAILED_STATUS_CODE, string(response))
	}

	body := models.GetUserResponse{}

	if err := parseJson(response, &body); err != nil {
		log.Fatal(err)
	}

	c.user = body.Data
}

func (c *Client) GetLikedTweets() []models.Tweet {
	if c.user.Id == "" {
		log.Fatal(constants.ERROR_TW_G_L_TWEETS_USER_UNSET)
	}

	url := buildUrl(constants.GET_LIKED_TWEETS_BY_USER, ":id", c.user.Id)

	statusCode, response, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if statusCode != 200 {
		log.Fatalf("%vResponse: \"%v\"", constants.ERROR_TW_G_L_TWEETS_FAILED_STATUS_CODE, string(response))
	}

	body := models.GetLikedTweetsResponse{}

	if err := parseJson(response, &body); err != nil {
		log.Fatal(err)
	}

	return body.Data
}

func (c *Client) GetTweets() []models.Tweet {
	if c.user.Id == "" {
		log.Fatal(constants.ERROR_TW_G_TWEETS_USER_UNSET)
	}

	url := buildUrl(constants.GET_TWEETS_BY_USER, ":id", c.user.Id)

	statusCode, response, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if statusCode != 200 {
		log.Fatalf("%vResponse: \"%v\"", constants.ERROR_TW_G_TWEETS_FAILED_STATUS_CODE, string(response))
	}

	body := models.GetTweetsResponse{}

	if err := parseJson(response, &body); err != nil {
		log.Fatal(err)
	}

	return body.Data
}

func (c *Client) DeleteTweet(id string) (bool, error) {
	url := buildUrl(constants.DELETE_TWEET, ":id", id)

	statusCode, response, err := c.makeRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}

	if statusCode != 200 {
		body := models.FailedRequest{}

		if err := parseJson(response, &body); err != nil {
			return false, err
		}

		message := fmt.Sprintf("%v ErrorTitle: %v | ErrorDetail: %v",
			constants.ERROR_TW_D_TWEETS_FAILED_STATUS_CODE, body.Title, body.Detail)

		if body.Status != nil {
			message = fmt.Sprintf("%v | ErrorStatus: %v", message, body.Status)
		}

		if body.Errors != nil {
			message = fmt.Sprintf("%v | Errors: \"%v\"", message, body.Errors)
		}

		return false, errors.New(message)
	}

	body := models.DeleteTweetResponse{}

	if err := parseJson(response, &body); err != nil {
		return false, err
	}

	return body.Data.Deleted, nil
}

func (c *Client) DeleteLikedTweet(tweetId string) (bool, error) {
	url := buildUrl(constants.DELETE_LIKED_TWEET, ":id", c.user.Id)
	url = buildUrl(url, ":tweet_id", tweetId)

	statusCode, response, err := c.makeRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}

	if statusCode != 200 {
		body := models.FailedRequest{}

		if err := parseJson(response, &body); err != nil {
			return false, err
		}

		message := fmt.Sprintf("%v ErrorTitle: %v | ErrorDetail: %v",
			constants.ERROR_TW_D_L_TWEETS_FAILED_STATUS_CODE, body.Title, body.Detail)

		if body.Status != nil {
			message = fmt.Sprintf("%v | ErrorStatus: %v", message, body.Status)
		}

		if body.Errors != nil {
			message = fmt.Sprintf("%v | Errors: \"%v\"", message, body.Errors)
		}

		return false, errors.New(message)
	}

	body := models.DeleteLikedTweetResponse{}

	if err := parseJson(response, &body); err != nil {
		return false, err
	}

	return !body.Data.Liked, nil
}
