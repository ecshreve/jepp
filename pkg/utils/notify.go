package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func SendGotification(t, m string) {
	postURL := fmt.Sprintf("http://gotify.slab.lan/message?token=%s", os.Getenv("GOTIFY_TOKEN"))
	_, err := http.PostForm(postURL, url.Values{"message": {m}, "title": {t}})
	if err != nil {
		fmt.Println(err)
	}
}
