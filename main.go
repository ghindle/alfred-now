package main

import (
	"fmt"
	"time"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func run() {
	query := wf.Args()[0]

	now := time.Now()

	addOption(now.Format(time.RFC3339), "RFC 3339")
	addOption(now.Format(time.RFC822), "RFC 822")
	addOption(now.Format(time.RFC1123), "RFC 1123")
	addOption(fmt.Sprintf("%d", now.Unix()), "Unix seconds")
	addOption(fmt.Sprintf("%d", now.UnixMilli()), "Unix milliseconds")
	addOption(fmt.Sprintf("%d", now.UnixMicro()), "Unix microseconds")

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("Unknown format", "")
	wf.SendFeedback()
}

func addOption(value, formatName string) {
	wf.NewItem(formatName).
		Subtitle(value).
		Valid(true).
		Arg(value)
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
