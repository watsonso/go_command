package main

import "fmt"

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類義語検索に失敗しました: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%qに類義語はありませんでした\n")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
