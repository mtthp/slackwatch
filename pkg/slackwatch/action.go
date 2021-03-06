package slackwatch

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// An Action 's execute method is called when an interesting message is
// received if armed.
type Action interface {
	Execute(Message)
}

// URLAction specifies an HTTP request to make on Alert. If Body is provided,
// an HTTP post is made, otherwise, an HTTP get.
type URLAction struct {
	URL         string
	Body        string
	ContentType string // defaults to application/octet-stream
}

// Execute is called to make the HTTP request
func (u URLAction) Execute(m Message) {
	var res *http.Response
	var err error
	if u.Body != "" {
		ct := u.ContentType
		if ct == "" {
			ct = "application/octet-stream"
		}
		res, err = http.Post(u.URL, ct, strings.NewReader(u.Body))
	} else {
		res, err = http.Get(u.URL)
	}
	if err != nil {
		logrus.Errorf("Error requesting %s: %v", u.URL, err)
	}
	if res != nil {
		_ = res.Body.Close()
	}
}

// CommandAction specifies a command to execute on Alert
type CommandAction struct {
	Command string
	Args    string
}

// Execute is called to run the command.
func (c CommandAction) Execute(m Message) {
	cmd := exec.Command(c.Command, c.Args)
	err := cmd.Run()
	if err != nil {
		logrus.Errorf("Error running %s: %v", c.Command, err)
	}
}
