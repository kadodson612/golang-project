package main

import (
    "fmt"

    "github.com/nlopes/slack"
)


func send_message(msg string, ev *slack.MessageEvent, rtm *slack.RTM) {

    rtm.SendMessage(rtm.NewOutgoingMessage(msg, ev.Channel))

}

func debug(msg string) {
    fmt.Println("DEBUG:", msg)
}
