package transpData

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

/*------------------------------------------------- DAILYCHECK -----
|  Function DAILYCHECK
|
|  Purpose: CHECKS DAILY FOR CHANGES IN THE PORTAL DATABASE AND
|      UPDATES APPLICATION DATABASE TO MATCH.
|
|  Parameters:
|      url (STRING) -- HTML ADDRESS TO WATCH FOR CHANGES.
|
|  Returns:  BOOLEAN INDICATING IF DATABASE WAS UPDATED.
*-------------------------------------------------------------------*/
func DailyCheck(url string) {
	transp, _ := http.Get(url)
	//body, _ := ioutil.ReadAll(transp.Body)

	//fmt.Println("HTML:\n", string(body))

	z := html.NewTokenizer(transp.Body)

	latestFile := ""

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				fmt.Println("Found an link!")
			}

			for _, a := range t.Attr {
				if a.Key == "href" {
					fmt.Println("Found href:", a.Val)
					latestFile = a.Val
					return
				}
			}
		}
	}

	fileUrl := url + latestFile

	if err := DownloadFile(latestFile, fileUrl); err != nil {
		panic(err)
	}
	defer transp.Body.Close()
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
