package main

import "fmt"

func main() {

	link := "https://www.1mg.com/drugs/augmentin-625-duo-tablet-138629"
	question := "As a SEO senior marketer, create summary for the page https://www.1mg.com/drugs/augmentin-625-duo-tablet-138629. Also provide a Hindi translation of the same without sentence construction errors that you can avoid by using literal translation"

	csv, err := MakeRequest(question, link)
	//fmt.Println(csv)
	if err != nil {
		fmt.Print(err)
	}
	err = WriteToCsv(csv)
	//
	if err != nil {
		fmt.Print(err)
	}

	//error = ParseRssFeeds()
	//if error != nil {
	//	fmt.Print(error)
	//
	//}
}
