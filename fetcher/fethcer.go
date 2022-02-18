package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML,"+
		"like Gecko) Chrome/98.0.4758.82 Safari/537.36 Edg/98.0.1108.51")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	//e := determineEncoding(resp.Body)
	//reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(resp.Body)
}
