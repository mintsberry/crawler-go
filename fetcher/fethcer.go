package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeChannel = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-timeChannel
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML,"+
		"like Gecko) Chrome/98.0.4758.82 Safari/537.36 Edg/98.0.1108.51")
	request.Header.Add("Cookie", "Cookie: ll=\"118172\"; bid=1scVzdhCQeA; Hm_lvt_16a14f3002af32bf3a75dfe352478639=1640618163; gr_user_id=d98f890e-5c56-4667-9114-9e8ec1d3058d; viewed=\"23008813_27615777\"; douban-fav-remind=1; __utmz=30149280.1642439292.2.2.utmcsr=cn.bing.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmc=30149280; dbcl2=\"188798333:4pAdWMkgQVs\"; ck=wLxj; push_noty_num=0; push_doumail_num=0; __utmv=30149280.18879; _pk_id.100001.4cf6=e89c674d6453cbb7.1640618163.2.1645270008.1640618182.; _pk_ses.100001.4cf6=*; __utma=30149280.770649054.1641727641.1645264104.1645270009.5; __utmb=30149280.0.10.1645270009; __utma=223695111.1816329907.1645270009.1645270009.1645270009.1; __utmb=223695111.0.10.1645270009; __utmc=223695111; __utmz=223695111.1645270009.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)")
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
