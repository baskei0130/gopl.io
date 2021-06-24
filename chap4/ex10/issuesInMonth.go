// 検索後に一致したGitHubイシューの表を表示
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/chap4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	AMonthAgo := time.Now().Add(-24 * 31 * time.Hour)
	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		if item.CreatedAt.After(AMonthAgo) {
			fmt.Printf("#%-5d %9.9s %.55s ",
				item.Number, item.User.Login, item.Title)
			fmt.Println(item.CreatedAt)
		}
	}
}
