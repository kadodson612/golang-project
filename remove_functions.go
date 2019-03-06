package main

import (
    "strings"

    "github.com/nlopes/slack"
)

// REMOVE FRIEND
func remove_friend(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) != 2 {
        send_message("Usage: *addfriend <name>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
    yaml := read_yaml(YAML_FILE)

    // check if friend exists
    _, exists := yaml.Friends[name]

    if exists != true {

        //TODO: remove friend and write

    }

}

// REMOVE PHRASE
func remove_phrase(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 3 {
        send_message("Usage: *addphrase <name> <phrase>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
//    phrase := strings.ToLower(strings.Join(tokens[2:], " "))
    yaml := read_yaml(YAML_FILE)

    _, exists := yaml.Friends[name]

    if exists == true {

        //TODO: remove phrase and write

    }

}

// REMOVE JOKE
func remove_joke(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 2 {
        send_message("Usage: *addjoke <joke>", ev, rtm)
    }

//    joke := strings.ToLower(strings.Join(tokens[1:], " "))
//    yaml := read_yaml(YAML_FILE)

    //TODO: remove joke and write

}

// REMOVE INSULT
func remove_insult(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 2 {
        send_message("Usage: *addinsult <insult>", ev, rtm)
    }

//    insult := strings.ToLower(strings.Join(tokens[1:], " "))
//    yaml := read_yaml(YAML_FILE)

    //TODO: remove insult and write

}

// REMOVE ALIAS
func remove_alias(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check for arguments
    if len(tokens) < 3 {
        send_message("Usage: *addalias <name> <alias>", ev, rtm)
    }

    name := strings.ToLower(tokens[1])
//    alias := strings.ToLower(strings.Join(tokens[2:], " "))
    yaml := read_yaml(YAML_FILE)

    _, exists := yaml.Friends[name]

    if exists == true {

        //TODO: remove alias and write

    }

}
