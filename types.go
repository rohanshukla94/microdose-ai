package main

type CsvData struct {
	ID              string
	ProductLink     string
	Prompt          string
	MetaDescription string
	HindiTranslate  string
	Image           string
}

type GptMetaDesc struct {
	Description string
}

type HealthRSSFeeds struct {
	ID          string
	Link        string
	Title       string
	Description string
}

type ProductsDesc struct {
	ID   string
	Link string
}
