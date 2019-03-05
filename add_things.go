package main

import (
    "fmt"
    "strings"

    "github.com/nlopes/slack"
    //"gopkg.in/yaml.v2"

)

func add_friend(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) != 2 {
        send_message("Usage: ?addfriend <name>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])

    // check if friend exists

    // add friend to yaml struct and write

    fmt.Println("friend:", name)

}

func add_phrase(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) != 3 {
        send_message("Usage: ?addphrase <name> <phrase>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    phrase := strings.ToLower(strings.Join(tokens[2:], " "))

    // make sure friend exists

    // check that phrase is not a duplicate 

    // add phrase to yaml struct and write

    fmt.Println("name:", name)
    fmt.Println("phrase:", phrase)

}

func add_joke(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) != 2 {
        send_message("Usage: ?addjoke <joke>", ev, rtm)
    }

    joke := strings.ToLower(strings.Join(tokens[1:], " "))

    // add joke to yaml struct and write

    fmt.Println("joke:", joke)

}
