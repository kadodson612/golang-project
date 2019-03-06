package main

import (
    "fmt"
    "strings"
    "io/ioutil"

    "github.com/nlopes/slack"
)

var YAML_FILE string

type function func([]string,*slack.MessageEvent,*slack.RTM)

//MAIN FUNCTION -- instantiate bot
func main() {

    YAML_FILE = "squad.yaml"

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
                command := strings.HasPrefix(text, "*")
                debug(text)
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

func call(f function, tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {
    f(tokens, ev, rtm)
}

func check_messages(text string, ev *slack.MessageEvent, rtm *slack.RTM) {

    tokens := strings.Fields(text)
    cmd := strings.Replace(tokens[0], "*", "", -1)

    // dictionary associating a command with a function
    commands := map[string]function {
        "gif": get_gif,
        "speak": speak,
        "joke": joke,
        "insult": insult,
        "show": show,
        "aka": aka,
        "addfriend": add_friend,
        "addphrase": add_phrase,
        "addjoke": add_joke,
        "addinsult": add_insult,
        "addalias": add_alias,
        "rmfriend": remove_friend,
        "rmphrase": remove_phrase,
        "rmjoke": remove_joke,
        "rminsult": remove_insult,
        "rmalias": remove_alias}

    _, ok := commands[cmd]

    if ok {
        call(commands[cmd], tokens, ev, rtm)
    } else {
        send_message("Not a command, dood", ev, rtm)
    }

}
