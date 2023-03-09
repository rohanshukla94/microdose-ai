package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

// Either the input to the API is the description itself
// Or the input to the API can be a website link ( more user friendly )

var feedUrls = []string{"https://wellnessmama.com/feed/", "https://articles.mercola.com/sites/articles/rss.aspx", "https://www.precisionnutrition.com/blog/feed", "https://www.abundanceandhealth.co.uk/en/blog/rss/posts?lang=en", "https://www.healthy-magazine.co.uk/feed/", "https://healthychild.com/feed/", "https://youmustgethealthy.com/feed/", "https://www.naturalwellness.com/nwupdate/feed/", "https://www.behealthynow.co.uk/feed/", "https://info.nihadc.com/integrative-health-blog/rss.xml"}

func ParseRssFeeds() error {
	fmt.Println("links", feedUrls)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrls[0])
	fmt.Println(feed)
	return nil
}
