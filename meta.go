package main

// var links = []string{"https://lovedepot.com/featured", "https://lovedepot.com/featured/page/2", "https://lovedepot.com/featured/page/3", "https://lovedepot.com/best-sellers/", "https://lovedepot.com/best-sellers/page/2/", "https://lovedepot.com/new-arrivals", "https://lovedepot.com/new-arrivals/page/2/", "https://lovedepot.com/new-arrivals/page/3/", "https://lovedepot.com/new-arrivals/page/4/", "https://lovedepot.com/new-arrivals/page/5/", "https://lovedepot.com/new-arrivals/page/6/"}

//func CrawlForProductLinks(productLinks *[]string) {
//
//	c := colly.NewCollector()
//
//	crawledLinks := make([]ProductLinks, 0)
//
//	c.OnHTML("div.desk", func(e *colly.HTMLElement) {
//		t := e.ChildAttr("a", "href")
//		item := ProductLinks{}
//		item.ID = uuid.NewString()
//		item.Link = t
//
//		crawledLinks = append(crawledLinks, item)
//	})
//
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL.String())
//	})
//
//	for _, item := range *productLinks {
//
//		err := c.Visit(item)
//		if err != nil {
//			return
//		}
//	}
//
//	fmt.Println(crawledLinks)
//	payload, err := getProductLinksAndCrawl(&crawledLinks)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	err = GenerateMetaDescription(payload)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func GenerateMetaDescription(content *[]ProductsDesc) error {
//
//	csvData := make([]CsvData, 0)
//
//	item := CsvData{}
//
//	var wg = &sync.WaitGroup{}
//	for i, result := range *content {
//		wg.Add(5)
//
//		result := result
//		go func(i int) {
//			fmt.Println(i)
//			resp, _ := MakeRequest(result.Description, result.Link)
//
//			item.ID = resp.ID
//			item.ProductLink = resp.ProductLink
//			item.MetaDescription = resp.MetaDescription
//			item.Prompt = resp.Prompt
//
//			csvData = append(csvData, resp)
//
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//
//	err := WriteToCsv(csvData)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func getProductLinksAndCrawl(crawledLinks *[]ProductLinks) (*[]ProductsDesc, error) {
//
//	c := colly.NewCollector()
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL.String())
//	})
//	crawledContent := make([]ProductsDesc, 0)
//
//	c.OnHTML("span[data-sheets-value]", func(e *colly.HTMLElement) {
//
//		item := ProductsDesc{}
//
//		item.ID = uuid.NewString()
//		item.Link = e.Request.URL.Path
//		item.Description = "Write meta description for the following product description: " + e.Text
//		crawledContent = append(crawledContent, item)
//
//	})
//
//	for _, data := range *crawledLinks {
//		c.Visit(data.Link)
//	}
//
//	return &crawledContent, nil
//}
