package main

import (
    "fmt"
    "regexp"
    "strings"
    "io/ioutil"
    "os"

    "github.com/nlopes/slack"
    "github.com/peterhellberg/giphy"
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

                // check if message is for bot
                command, _ := regexp.MatchString("^?.*", text)
                if command {
                    check_messages(text, ev, rtm)
                }

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

    tokens := strings.Fields(text)
    cmd := strings.Replace(tokens.get(0), "?", "", -1)

    switch cmd {

    case "gif":
        get_gif(tokens, ev, rtm)
    case "addfriend":
        add_friend()
    case "addphrase":
        add_phrase()
    case "addjoke":
        add_joke()
    case "rmfriend":
        remove_friend()
    case "rmphrase":
        remove_phrase()
    case "rmjoke":
        remove_joke()
    default:
        usage(ev, rtm)

    }

}

func usage(ev *slack.MessageEvent, rtm *slack.RTM) {

    rtm.SendMessage(rtm.NewOutgoingMessage("Sorry, I don't know what you're trying to do", ev.Channel))

}
