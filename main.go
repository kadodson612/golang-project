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
    cmd := strings.Replace(tokens[0], "?", "", -1)

    // TODO: make a dictionary associating a cmd with a function and its parameters

    switch cmd {

    case "gif":             // sends a gif
        get_gif(tokens, ev, rtm)
    case "speak":           // sends a catchphrase
        debug(cmd)
        //speak()
    case "joke":            // sends an inside joke
        debug(cmd)
        //joke()
    case "insult":          // sends an insult
        debug(cmd)
        //insult()
    case "deprecate":       // sends a deprecation
        debug(cmd)
        //deprecate()
    case "show":            // prints out all the items of a given category (joke, friend, etc.)
        debug(cmd)
        //show_items()

    case "addfriend":       // add a friend
        add_friend(tokens, ev, rtm)
    case "addphrase":       // add a catchphrase
        add_phrase(tokens, ev, rtm)
    case "addjoke":         // add an inside joke
        add_joke(tokens, ev, rtm)
    case "addinsult":       // add an insult
        debug(cmd)
        //add_insult(tokens, ev, rtm)
    case "adddeprecate":    // add a deprecation
        debug(cmd)
        //add_deprecate(tokens, ev, rtm)
    case "addalias":        // add an alias
        debug(cmd)
        //add_alias(tokens, ev, rtm)

    case "rmfriend":        // remove a friend
        debug(cmd)
        //remove_friend()
    case "rmphrase":        // remove a catchphrase
        debug(cmd)
        //remove_phrase()
    case "rmjoke":          // remove an inside joke
        debug(cmd)
        //remove_joke()
    case "rminsult":        // remove an insult
        debug(cmd)
        //remove_insult()
    case "rmdeprecate":     // remove a deprecation
        debug(cmd)
        //remove_deprecate()
    case "rmalias":         // remove an alias
        debug(cmd)
        //remove_alias()

    default:
        send_message("Not a command dood", ev, rtm)

    }

}

func send_message(msg string, ev *slack.MessageEvent, rtm *slack.RTM) {

    rtm.SendMessage(rtm.NewOutgoingMessage(msg, ev.Channel))

}

func debug(msg string) {
    fmt.Println("DEBUG:", msg)
}
