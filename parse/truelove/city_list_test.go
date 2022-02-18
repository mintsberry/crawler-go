package truelove

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//fetch, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	file, err := ioutil.ReadFile("city_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", fetch)

	list := ParseCityList(file)

	const resultSize = 408

	if len(list.Request) != resultSize {
		t.Errorf("result should have %d, but had %d", resultSize, len(list.Request))
	}

	for _, item := range list.Items {
		fmt.Printf("Get city: %s\n", item)
	}
}
