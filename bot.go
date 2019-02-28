package main

import (
    "fmt"
    "regexp"
    "strings"
    "io/ioutil"

    "github.com/nlopes/slack"
)

//MAIN FUNCTION -- instantiate bot
func main() {

    // fetch token file
    data, err := ioutil.ReadFile("token")
    if err != nil {
        return
    }

    token := string(data)
    token = strings.TrimSuffix(token, "\n")

    api := slack.New(token)
    rtm := api.NewRTM()
    go rtm.ManageConnection()

Loop:
    for {
        select {

        case msg := <-rtm.IncomingEvents:

            switch ev := msg.Data.(type) {

            case *slack.MessageEvent:
                text := ev.Text
                text = strings.TrimSpace(text)
                text = strings.ToLower(text)

                check_messages(text, ev, rtm)

            case *slack.RTMError:
                fmt.Printf("Error: %s\n", ev.Error())

            case *slack.InvalidAuthEvent:
                fmt.Printf("Invalid credentials")
                break Loop

            default:
                // Take no action

            }
        }
    }

}

func check_messages(text string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check if the message string exactly matches "meow"
    info := rtm.GetInfo()
    matched, _ := regexp.MatchString("^meow$", text)

    if ev.User != info.User.ID && matched {
        rtm.SendMessage(rtm.NewOutgoingMessage("Is that Maru I hear?", ev.Channel))
    }

}
