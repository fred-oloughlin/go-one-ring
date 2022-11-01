package main

import (
	"fmt"
	"log"
	ring "onering/oneringclient"
)

func main() {
	endpoint := ring.BuildURL("book")
	BookResponse, err := ring.GetBookResponse(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	for _, book := range BookResponse.Books {

		endpoint := ring.BuildURL("book", book.Id, "chapter")
		ChapterResponse, err := ring.GetChapterResponse(endpoint)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println()
		fmt.Println(book.Name)
		for i, chapter := range ChapterResponse.Chapters {
			prefix := ring.BuildPrefix(i)
			fmt.Println(prefix, chapter.ChapterName)
		}
	}

}
