package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)
/**
* 返回response
*/
func getResponse (url string) *http.Response  {
	client := http.Client{}
	request,_ := http.NewRequest("GET",url,nil)
	request.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
	response,_ := client.Do(request)
	return response

}

func GetJokes(){
	response := getResponse("http://www.qiushibaike.com")
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil{
		log.Fatal(err)
	}
	doc.Find(".content").Each(func(i int, s *goquery.Selection){
		fmt.Println(s.Text())
	})
}

func main(){
	GetJokes()
}