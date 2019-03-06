package main

import (
    "fmt"
    "strings"
    "io/ioutil"

    "github.com/nlopes/slack"
)

var YAML_FILE string

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

func check_messages(text string, ev *slack.MessageEvent, rtm *slack.RTM) {

    tokens := strings.Fields(text)
    cmd := strings.Replace(tokens[0], "*", "", -1)

    // TODO: make a dictionary associating a cmd with a function and its parameters

    switch cmd {

    case "gif":             // sends a gif
        get_gif(tokens, ev, rtm)
    case "speak":           // sends a catchphrase
        speak(tokens, ev, rtm)
    case "joke":            // sends an inside joke
        joke(tokens, ev, rtm)
    case "insult":          // sends an insult
        insult(tokens, ev, rtm)
    case "show":            // prints out all the items of a given category (joke, friend, etc.)
        show(tokens, ev, rtm)
    case "aka":             // prints out all a friend's aliases
        aka(tokens, ev, rtm)

    case "addfriend":       // add a friend
        add_friend(tokens, ev, rtm)
    case "addphrase":       // add a catchphrase
        add_phrase(tokens, ev, rtm)
    case "addjoke":         // add an inside joke
        add_joke(tokens, ev, rtm)
    case "addinsult":       // add an insult
        add_insult(tokens, ev, rtm)
    case "addalias":        // add an alias
        add_alias(tokens, ev, rtm)

    case "rmfriend":        // remove a friend
        remove_friend(tokens, ev, rtm)
    case "rmphrase":        // remove a catchphrase
        remove_phrase(tokens, ev, rtm)
    case "rmjoke":          // remove an inside joke
        remove_joke(tokens, ev, rtm)
    case "rminsult":        // remove an insult
        remove_insult(tokens, ev, rtm)
    case "rmalias":         // remove an alias
        remove_alias(tokens, ev, rtm)

    default:
        send_message("Not a command dood", ev, rtm)

    }

}


