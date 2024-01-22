package main

import (
	"bufio"
	"fmt"
	"log"
	"main/src/thesaurus"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	apiKey := os.Getenv("bth_api_key")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%qに類語はありませんでした\n", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}