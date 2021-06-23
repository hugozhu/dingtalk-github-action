package main

import (
	"flag"
	"io/ioutil"

	dingtalk "github.com/hugozhu/godingtalk"
	githubactions "github.com/sethvargo/go-githubactions"
)

var img *bool
var f *string

func init() {
	f = flag.String("f", "", "File Path")
	flag.Parse()
}

func main() {
	title := githubactions.GetInput("title")
	msgtype := githubactions.GetInput("type")
	token := githubactions.GetInput("token")
	file := githubactions.GetInput("file")

	if title == "" {
		githubactions.Fatalf("missing 'title'")
		return
	}

	if token == "" {
		githubactions.Fatalf("missing 'token'")
		return
	}

	githubactions.AddMask(title)

	message := ""
	if file != "" {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			githubactions.Fatalf("faild to read file %v", err)
		}
		message = string(data)
	}

	c := dingtalk.NewDingTalkClient("", "")

	var err error

	if msgtype != "text" {
		_, err = c.SendRobotTextMessage(token, title+"\n\n"+message)
	} else {
		_, err = c.SendRobotMarkdownMessage(token, title, "# "+title+"\n\n"+message)
	}

	if err != nil {
		githubactions.Fatalf("failed to send dingtalk message, %v", err)
	}
}
