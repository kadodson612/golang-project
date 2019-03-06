package main

import (
//    "fmt"
//    "regexp"
//    "math/rand"
//    "time"
    "strings"

    "github.com/nlopes/slack"
)

func show_items(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check arguments
    if len(tokens) != 2 {
        send_message("Usage: *show < jokes | insults | friends >", ev, rtm)
    }

    item := tokens[1]

    switch item {

    case "jokes":
        //yaml := read_yaml("yaml/jokes.yaml")
    case "insults":
        //yaml := read_yaml("yaml/insults.yaml")
    case "friends":
        yaml := read_yaml("yaml/friends.yaml")

        var friends []string
        for f := range yaml.Friends {
            friends = append(friends, f)
        }
        send_message(strings.Join(friends, ", "), ev, rtm)
    }

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
