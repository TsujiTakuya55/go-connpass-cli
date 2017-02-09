package cmd

import (
	"encoding/json"
	"fmt"
	"connpass/utils"
	"github.com/pkg/errors"
	"net/http"
)

func List(c *Connpass, count string) {

	resp, err := utils.NewResponse(utils.BaseUrl + "?count=" + count + "&order=2")

	if err != nil {
		errors.Wrap(err, "NewResponse is failed")
		return
	}

	if resp.StatusCode != http.StatusOK {
		errors.Wrap(err, "StatusCode is funny")
		return
	}

	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		errors.Wrap(err, "decode is failed")
		return
	}

	for _, v := range c.Events {
		fmt.Println("タイトル : ", v.Title)
		fmt.Println("キャッチ : ", v.Catch)
		fmt.Println("開催会場 : ", v.Place)
		fmt.Println("URL : ", v.EventUrl)
		fmt.Println("開催日時 : ", v.StartedAt)
		fmt.Println("定員 : ", v.Limit)
		fmt.Println("参加者数 : ", v.Accepted)
		fmt.Println("補欠者数 : ", v.Waiting)
		fmt.Println("==================================")
	}
}
