package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

type Category struct {
	LessThanAMonth []*github.Issue
	LessThanAYear  []*github.Issue
	MoreThanAYear  []*github.Issue
}

func categorize(items []*github.Issue) *Category {
	category := new(Category)
	for _, i := range items {
		now := time.Now()
		if i.CreatedAt.Year() < now.Year() {
			category.MoreThanAYear = append(category.MoreThanAYear, i)
		} else {
			if i.CreatedAt.Month() < now.Month() {
				category.LessThanAYear = append(category.LessThanAYear, i)
			} else {
				category.LessThanAMonth = append(category.LessThanAMonth, i)
			}
		}
	}
	return category
}

func printIssues(issues []*github.Issue) {
	for _, item := range issues {
		fmt.Printf("#%-5d %v %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n",
	//		item.Number, item.User.Login, item.Title)
	//}
	cat := categorize(result.Items)
	fmt.Printf("Less than a month: %d issues\n", len(cat.LessThanAMonth))
	printIssues(cat.LessThanAMonth)
	fmt.Printf("Less than a year: %d issues\n", len(cat.LessThanAYear))
	printIssues(cat.LessThanAYear)
	fmt.Printf("More than a year: %d issues\n", len(cat.MoreThanAYear))
	printIssues(cat.MoreThanAYear)

}
