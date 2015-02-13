package other

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.theatlantic.com/international/archive/2014/12/do-police-body-cameras-work-ferguson/383323/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Lets just print out the body real quick to save out sanity
	robots, err := ioutil.ReadAll(resp.Body)
	body := string(robots)
	//fmt.Println(body)

	z := html.NewTokenizer(strings.NewReader(body))

	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			log.Fatalln(z.Err())
			return
		case html.TextToken:
			if len(z.Token().Data) > 0 && depth > 0 {
				fmt.Println(z.Token().Data)
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			stn := string(tn)
			if len(tn) > 0 && stn != "meta" && !strings.Contains(stn, "script") {
				if tt == html.StartTagToken {
					depth++
				}
				fmt.Println(strings.Repeat("  ", depth) + stn)
				if tt == html.EndTagToken {
					depth--
				}
			}
		}
	}
}
