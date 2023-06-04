package main

import (
  "fmt"
  "github.com/mmcdole/gofeed"
  "html"
  "strings"
)


func main() {
  var parsedNews *gofeed.Feed = parseNews()
  for _, post := range parsedNews.Items {
    printNewsItem(post)
  }
}

func printNewsItem(item *gofeed.Item) {
  fmt.Println(html.UnescapeString(item.Title))
  for _, word := range strings.Split(item.Description, " ") {
    fmt.Print(html.UnescapeString(word) + " ")
  }
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


