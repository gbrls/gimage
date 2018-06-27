package img

import (
	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func getURL(name string) string {
	
	uri, err := url.Parse("https://www.google.com/search?tbm=isch")
	if err != nil {
		panic(err)
	}
	uri.Query().Set("q", name)
	uri.RawQuery = uri.Query().Encode()
	return uri.String()
}

func downloadImage(url string, folder string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		return err
	}
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.Mkdir((folder), os.ModePerm)

	}
	f, err := os.Create(fmt.Sprintf("%v/%v.jpeg", folder, url[len("https://encrypted-tbn0.gstatic.com/images?q=tbn:"):]))
	if err != nil {
		return fmt.Errorf("error creating files: %v", err)
	}
	defer f.Close()

	return jpeg.Encode(f, img, nil)
}

//GetImagesURLs searches for 'name' in google and downloads its firsts results
func GetImagesURLs(name string, folder string) error {
	urls := []string{}

	c := colly.NewCollector()

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		urls = append(urls, e.Attr("src"))
	})

	c.Visit(getURL(name))

	for i := 1; i < len(urls); i++ {
		err := downloadImage(urls[i], folder)
		if err != nil {
			return err
		}
	}

	return nil
}
