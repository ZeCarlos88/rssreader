package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Descriptiom string `xml:"description"`
	Link        string `xml:"link"`
}

func main() {
	url := "http://feeds.itjobs.pt/feed/emprego?q=golang+developer"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao tentar receber RSS", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao tentar ler a resposta", err)
	}
	var feed Feed
	err = xml.Unmarshal(body, &feed)

	if err != nil {
		fmt.Println("Erro ao tentar fazer parser da resposta", err)
	}

	fmt.Println("title:", feed.Channel.Title)
	fmt.Println("Description:", feed.Channel.Description)

	for _, item := range feed.Channel.Items {

		fmt.Println("Title:", item.Title)
		fmt.Println("Description:", item.Descriptiom)
		fmt.Println("Link:", item.Link)
		fmt.Println("*****************************")
	}

}
