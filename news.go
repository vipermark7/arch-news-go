package main

import (
  "fmt"
  "github.com/mmcdole/gofeed"
  "html"
)


func main() {
	var parsedNews *gofeed.Feed = parseNews()
  fmt.Println(html.UnescapeString(parsedNews.Items[0].Description))
}

func parseNews() *gofeed.Feed {

  ARCH_NEWS_URL :=  "https://archlinux.org/feeds/news/"
  fp := gofeed.NewParser()
  parsed, err := fp.ParseURL(ARCH_NEWS_URL)

  if err != nil {
    fmt.Println(err)
  }
  return parsed;
}


