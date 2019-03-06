package main

import (
    "fmt"
    "strings"

    "github.com/nlopes/slack"
)

func add_friend(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) != 2 {
        send_message("Usage: ?addfriend <name>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    yaml := read_yaml("yaml/friends.yaml")

    // check if friend exists
    _, exists := yaml.Friends[name]

    if exists != true {

        // add friend to yaml struct and write
        yaml.Friends[name] = Friend{}
        write_yaml("yaml/friends.yaml", yaml)

    }

}

func add_phrase(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 3 {
        send_message("Usage: ?addphrase <name> <phrase>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    phrase := strings.ToLower(strings.Join(tokens[2:], " "))
    yaml := read_yaml("yaml/friends.yaml")

    _, exists := yaml.Friends[name]

    if exists == true {

        // add phrase to yaml struct and write
        friend := yaml.Friends[name]
        friend.Phrases = append(friend.Phrases, phrase)

        fmt.Println(friend.Phrases)

        yaml.Friends[name] = friend
        write_yaml("yaml/friends.yaml", yaml)
    }

    fmt.Println("PHRASE ADDED")

}

func add_joke(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 2 {
        send_message("Usage: ?addjoke <joke>", ev, rtm)
    }

    joke := strings.ToLower(strings.Join(tokens[1:], " "))

    // add joke to yaml struct and write

    fmt.Println("joke:", joke)

}

func add_insult(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 2 {
        send_message("Usage: ?addinsult <insult>", ev, rtm)
    }

}

func add_alias(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 3 {
        send_message("Usage: ?addalias <name> <alias>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    alias := strings.ToLower(strings.Join(tokens[2:], " "))
    yaml := read_yaml("yaml/friends.yaml")

    _, exists := yaml.Friends[name]

    if exists == true {

        // add phrase to yaml struct and write
        friend := yaml.Friends[name]
        friend.Aliases = append(friend.Aliases, alias)

        fmt.Println(friend.Aliases)

        yaml.Friends[name] = friend
        write_yaml("yaml/friends.yaml", yaml)
    }

    fmt.Println("ALIAS ADDED")

}
