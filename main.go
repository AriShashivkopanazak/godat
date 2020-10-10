/*hello world

more lines

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Description"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{Title: "hello world", Desc: "hello world", Content: "hello world"},
	}
	fmt.Println("articles Endpoint hit")
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the homepage")
	fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("godat is now working")
	handleRequests()
}
