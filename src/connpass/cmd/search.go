package cmd

import (
	"fmt"
	"github.com/TsujiTakuya55/go-connpass/connpass"
)

func Search(c *Connpass, keyWord string, count string) error {

	clinet := connpass.NewClient()
	connpass, _, err := clinet.Keyword.Get(keyWord)

	if err != nil {
		return err
	}

	for _, v := range *connpass.Events {
		fmt.Println("タイトル : ", *v.Title)
		fmt.Println("キャッチ : ", *v.Catch)
		fmt.Println("開催会場 : ", *v.Place)
		fmt.Println("URL : ", *v.EventUrl)
		fmt.Println("開催日時 : ", *v.StartedAt)
		fmt.Println("定員 : ", *v.Limit)
		fmt.Println("参加者数 : ", *v.Accepted)
		fmt.Println("補欠者数 : ", *v.Waiting)
		fmt.Println("==================================")
	}
	return nil
}
