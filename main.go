package main

import (
	"flag"
	"os"
	"regexp"
	"fmt"
	dingtalk "github.com/hugozhu/godingtalk"
)

var img *bool
var f *string

func init() {
	f = flag.String("f", "", "File Path")
	flag.Parse()
}

func stripeMarkdown(str string) string {
	str = regexp.MustCompile("[*|#]+").ReplaceAllString(str, "")
	str = regexp.MustCompile("\\s+").ReplaceAllString(str, " ")
	str = regexp.MustCompile("^ ").ReplaceAllString(str, "")
	return str
}

func main() {
	c := dingtalk.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	c.RefreshAccessToken()
	c.SendRobotMarkdownMessage(os.Getenv("token"), stripeMarkdown(os.Args[1]), os.Args[1])	
	fmt.Println(os.Args[1])
}
