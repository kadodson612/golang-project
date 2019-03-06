package main

import (
    "strings"

    "github.com/nlopes/slack"
)

func show(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check arguments
    if len(tokens) != 2 {
        send_message("Usage: *show < jokes | insults | friends >", ev, rtm)
    }

    item := tokens[1]

    yaml := read_yaml(YAML_FILE)

    switch item {

    case "jokes":
        send_message(strings.Join(yaml.Jokes, "\n"), ev, rtm)
    case "insults":
        send_message(strings.Join(yaml.Insults, "\n"), ev, rtm)
    case "friends":
        var friends []string
        for f := range yaml.Friends {
            friends = append(friends, f)
        }
        send_message(strings.Join(friends, "\n"), ev, rtm)
    }

}


func joke(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

}

func insult(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

}

func speak(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

}

func aka(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

}

func check_message(text string) string {

//    config := read_yaml("phrases.yaml")
//
//    for k, v := range config.Phrases {
//        //fmt.Printf("key[%s] value[%s]\n", k, v)
//	    s := fmt.Sprintf(".*%s.*", k)
//      matched,err := regexp.MatchString(s, text)
//
//	    if(matched == true) {
//            rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
//            return(v[rand.Intn(len(v))])
//	    }
//
//    }
//
    return("")
}
