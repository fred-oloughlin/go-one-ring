package ring

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR | Unable to parse response from: ", endpoint)
		return nil, err
	}

	return body, nil
}

func GetBookResponse() (BookResponse, error) {
	endpoint, err := BuildURL("book")
	if err != nil {
		return BookResponse{}, err
	}
	
	body, err := getResponse(endpoint)
	if err != nil {
		return BookResponse{}, err
	}

	var response BookResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return BookResponse{}, err
	}
	if len(response.Books) == 0 {
		log.Println("ERROR | ")
		return BookResponse{}, errors.New("Could not parse content from: " + endpoint)
	}

	return response, nil
}

func GetChapterResponse(book_id string) (ChapterResponse, error) {
	endpoint, err := BuildURL("book", book_id, "chapter")
	if err != nil {
		return ChapterResponse{}, err
	}
	
	body, err := getResponse(endpoint)
	if err != nil {
		return ChapterResponse{}, err
	}

	var response ChapterResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ChapterResponse{}, err
	}
	if len(response.Chapters) == 0 {
		log.Println("ERROR | ")
		return ChapterResponse{}, errors.New("Could not parse content from: " + endpoint)
	}

	return response, nil
}

func BuildURL(path_items ...string) (string, error) {
	var base string = "https://the-one-api.dev/v2"
	endpoint, err := url.JoinPath(base, path_items...)

	return endpoint, err
}

func BuildPrefix(index int) string {
	chapterNumberCharCount := len(strconv.Itoa(index))
	spacesNeeded := 4 - chapterNumberCharCount
	prefixSpaces := strings.Repeat(" ", spacesNeeded)
	
	return fmt.Sprintf("%d%s- ", index, prefixSpaces)
}
