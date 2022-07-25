package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func checkUrlAndSaveLocal(urlString string) {
	resp, err := http.Get(urlString)
	if err != nil {
		fmt.Println("server nahi chal raha bhaiyon aur behno", urlString, err)
	} else {
		if resp.StatusCode != 200 {
			fmt.Println("server 2xx nahi de raha bhaiyon aur behno", urlString, resp.Status)
		} else {
			fileName := "./savedHtmlPages/" + strings.Split(urlString, "//")[1] + ".txt"
			fmt.Println("recording response for", urlString, " in file:", fileName)
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			err = ioutil.WriteFile(fileName, bodyBytes, 0664)
			if err != nil {
				fmt.Println("apan response file mein nahi likh paye, bhaiyon aur behno", fileName, err)
			}
		}
	}
}

func main() {

	urls := []string{"https://sdfasdf.casdfasd", "https://google.com", "https://google.com/sdfasdfg.html", "https://fast.com"}

	start := time.Now()
	defer fmt.Println("lapsed time", time.Since(start))
	for _, v := range urls {
		checkUrlAndSaveLocal(v)
	}

}
