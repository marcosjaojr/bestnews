package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"math/big"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	best_ids := getBestIds()

	fmt.Println("\nPress 'Enter' to next news >>>\n")
	for i := range best_ids {
		news := getNewsDetail(best_ids[i])
		fmt.Println(i+1,"#",news["title"])
		fmt.Println(" "," ",news["url"])
		fmt.Println("----------------------------------------------------")
  	bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	fmt.Printf("That's all folks...\n")

}

// Make simple http request
func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return data
}

// Get ids of the best stories
func getBestIds() []big.Int{
	best_ids := getData("https://hacker-news.firebaseio.com/v0/topstories.json")

	ids := make([]big.Int,0)
  json.Unmarshal(best_ids, &ids)

	return ids;
}

// Get detail of a specific story
func getNewsDetail(id big.Int) map[string]interface{} {

	news := getData("https://hacker-news.firebaseio.com/v0/item/"+ id.String() +".json?type=story")

	detail := map[string]interface{}{}
	json.Unmarshal(news, &detail)

	return detail
}
