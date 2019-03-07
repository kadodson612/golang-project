package main

import (
    "strings"

    "math/rand"

    "time"

    "strconv"

    "github.com/nlopes/slack"
)

func show(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {

    // check arguments
    if len(tokens) != 2 {
        send_message("Usage: *show < jokes | insults | friends >", ev, rtm)
    }

    item := tokens[1]

    yaml := read_yaml("YAML_FILE")

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
    jokeNum := 0
    if len(tokens) == 2 {
        if tokens[1] == "-a" {
            jokeNum = -1
        } else {
            num, err := strconv.Atoi(tokens[1])
            if err != nil {
                panic(err)
            } else {
                jokeNum = num
            }
        }
    }

    yaml := read_yaml("YAML_FILE")

    var jokes []string
    for _, j := range yaml.Jokes {
        jokes = append(jokes, j)
    }

    if jokeNum == 0 {
        rand.Seed(time.Now().Unix())
        joke := jokes[rand.Intn(len(jokes))]
        send_message(joke, ev, rtm)

    } else if jokeNum < 0 {
        send_message(strings.Join(jokes, "\n"), ev, rtm)

    } else if jokeNum > 0 {
        if jokeNum < len(jokes) {
            joke := jokes[jokeNum]
            send_message(joke, ev, rtm)
        } else {
            send_message("That's not a valid joke number!", ev, rtm)
        }
        
    }
}

func insult(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {
    insultNum := 0
    if len(tokens) == 2 {
        if tokens[1] == "-a" {
            insultNum = -1
        } else {
            num, err := strconv.Atoi(tokens[1])
            if err != nil {
                panic(err)
            } else {
                insultNum = num
            }
        }
    }

    yaml := read_yaml("YAML_FILE")

    var insults []string
    for _, j := range yaml.Insults {
        insults = append(insults, j)
    }

    if insultNum == 0 {
        rand.Seed(time.Now().Unix())
        insult := insults[rand.Intn(len(insults))]
        send_message(insult, ev, rtm)

    } else if insultNum < 0 {
        send_message(strings.Join(insults, "\n"), ev, rtm)

    } else if insultNum > 0 {
        if insultNum < len(insults) {
            insult := insults[insultNum]
            send_message(insult, ev, rtm)
        } else {
            send_message("That's not a valid insult number!", ev, rtm)
        }
        
    }
}

func speak(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {
    phraseNum := 0
    if len(tokens) < 2{
        send_message("Usage: *speak < name > < number(opt) >", ev, rtm)
    } else if len(tokens) == 3 {
        if tokens[2] == "-a" {
            phraseNum = -1
        } else {
            num, err := strconv.Atoi(tokens[2])
            if err != nil {
                panic(err)
            }
            phraseNum = num
        }
    }

    friend := tokens[1]

    yaml := read_yaml(YAML_FILE)

    var phrases []string

    for f,v := range yaml.Friends {
        if f == friend{
            for _, p := range v.Phrases {
                phrases = append(phrases, p)
            }
        }    
    }

    var phrase string
    
    if phraseNum == 0 {
        rand.Seed(time.Now().Unix())
        phrase = phrases[rand.Intn(len(phrases))]
        send_message(phrase, ev, rtm)
    } else if phraseNum > 0 {
        if phraseNum < len(phrases){
            phrase = phrases[phraseNum - 1]
            send_message(phrase, ev, rtm)
        } else {
            send_message("That's not a valid phrase number!", ev, rtm)
        }
        
    } else if phraseNum < 0 {
        send_message(strings.Join(phrases, "\n"), ev, rtm)
    }
}

func aka(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM) {
    var showAll bool
    if len(tokens) < 2{
        send_message("Usage: *aka < name > < -a >", ev, rtm)
    } else if len(tokens) == 3 {
        showAll = true
    }

    friend := tokens[1]

    yaml := read_yaml("YAML_FILE")

    var aliases []string

    for f,v := range yaml.Friends {
        if f == friend{
            for _, a := range v.Aliases {
                aliases = append(aliases, a)
            }
        }    
    }

    if showAll == true {
        send_message(strings.Join(aliases, "\n"), ev, rtm)

    } else {
        rand.Seed(time.Now().Unix())
        alias := aliases[rand.Intn(len(aliases))]
        send_message(alias, ev, rtm)
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
