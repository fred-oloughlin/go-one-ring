package oneringclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Response struct {
	Total  int
	Limit  int
	Offset int
	Page   int
	Pages  int
}

type BookResponse struct {
	Response
	Books []Book `json:"docs"`
}

type ChapterResponse struct {
	Response
	Chapters []Chapter `json:"docs"`
}

type Book struct {
	Id   string `json:"_id"`
	Name string
}

type Chapter struct {
	Id          string `json:"_id"`
	ChapterName string
}

func getResponse(endpoint string) ([]byte, error) {

	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println("ERROR | Request failed to: ", endpoint)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("ERROR | Request failed to: ", endpoint)
		log.Println("ERROR | With message: ", resp.Body)
		return nil, errors.New(resp.Status)
	}

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR | Unable to parse response from: ", endpoint)
		return nil, err
	}

	return body, nil
}

func GetBookResponse(endpoint string) (BookResponse, error) {
	body, err := getResponse(endpoint)
	if err != nil {
		return BookResponse{}, err
	}

	var Response BookResponse
	json.Unmarshal(body, &Response)
	if len(Response.Books) == 0 {
		log.Println("ERROR | ")
		return BookResponse{}, errors.New("Could not parse content from: " + endpoint)
	}

	return Response, nil
}

func GetChapterResponse(endpoint string) (ChapterResponse, error) {
	body, err := getResponse(endpoint)
	if err != nil {
		return ChapterResponse{}, err
	}

	var Response ChapterResponse
	json.Unmarshal(body, &Response)
	if len(Response.Chapters) == 0 {
		log.Println("ERROR | ")
		return ChapterResponse{}, errors.New("Could not parse content from: " + endpoint)
	}

	return Response, nil
}

func BuildURL(path_items ...string) string {
	var base string = "https://the-one-api.dev/v2"
	endpoint, _ := url.JoinPath(base, path_items...)

	return endpoint
}

func BuildPrefix(index int) string {
	chnoCharCount := len(strconv.Itoa(index))
	spacesNeeded := 4 - chnoCharCount
	prefixSpaces := strings.Repeat(" ", spacesNeeded)
	prefix := strconv.Itoa(index) + prefixSpaces + "-" + " "

	return prefix
}
