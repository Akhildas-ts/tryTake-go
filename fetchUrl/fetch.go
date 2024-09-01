package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// fetch html content with the url 


// * frist we will create a client for request 
// * we got the response from the request 
// * we just read the repsonse with ioutil (package),there we get body part 
// * if we need to go with more with the html content we can use the goquery (used for web scraping and html manuplating)

func main() {

	url := "https://github.com/Arundas666"

	fmt.Println(fetchHtmlContent(url))

}

func fetchHtmlContent(url string) (string, error) {

	// here we are creating a client for requesting
	// with a time , after the time if there is no response the request will canncel

	client := &http.Client{

		Timeout: 10 * time.Second,
	}

	// getting the response

	resp, err := client.Get(url)

	if err != nil {
		return "", errors.New("faild to fetch url   ")
	}

	// closing the response body after done work
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return "", errors.New("get an error from fetching data ")
	}

	// go query from a package is used to web Scrab and manipulate the html content
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))

	if err != nil {

		return "", errors.New("error from scrabing html page ")

	}

	fmt.Println(doc)

	// emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// var emails []string

	// doc.Find("body").Each(func(i int, s *goquery.Selection){

	//     text := s.Text()
	//     matches := emailRegex.FindAllString(text,-1)

	//     for _,email := range matches{

	//         emails = append(emails, email)
	//     }

	// })

	response := string(body)

	fmt.Println("response:", response)

	return response, nil

}
