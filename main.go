package main

import (
	"flag"
	"regexp"

	dingtalk "github.com/hugozhu/godingtalk"
	githubactions "github.com/sethvargo/go-githubactions"
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
	msg := githubactions.GetInput("msg")
	corpid := githubactions.GetInput("corpid")
	corpsecret := githubactions.GetInput("corpsecret")
	token := githubactions.GetInput("token")

	if msg == "" || corpid == "" || corpsecret == "" || token == "" {
		githubactions.Fatalf("missing one of the inputs: 'msg', 'corpid', 'corpsecret', 'token'")
		return
	}

	githubactions.AddMask(msg)

	c := dingtalk.NewDingTalkClient(corpid, corpsecret)
	c.RefreshAccessToken()
	c.SendRobotMarkdownMessage(token, stripeMarkdown(msg), msg)
}
