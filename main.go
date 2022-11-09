package main

import (
	"fmt"
	"log"
	"onering/ring"
	"os"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	BookResponse, err := ring.GetBookResponse()
	if err != nil {
		log.Println(err)
		exitCode = 1
		return
	}

	for _, book := range BookResponse.Books {
		ChapterResponse, err := ring.GetChapterResponse(book.Id)
		if err != nil {
			log.Println(err)
			exitCode = 1
			return
		}

		fmt.Println()
		fmt.Println(book.Name)
		for i, chapter := range ChapterResponse.Chapters {
			prefix := ring.BuildPrefix(i)
			fmt.Println(prefix, chapter.ChapterName)
		}
	}
}
