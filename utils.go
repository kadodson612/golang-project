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

func delete_item(list []string, item string) []string {
    for i := range list {
        if list[i] == item {
            if i < len(list)-1 {
                list = append(list[:i], list[i+1:]...)
            } else {
                list = list[:i]
            }
            break
        }
    }

    return(list)
}
