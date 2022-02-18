package regex

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"

func main() {
	compile := regexp.MustCompile("ccmouse@gmail.com")
	findString := compile.FindString(text)
	fmt.Println(findString)
}
