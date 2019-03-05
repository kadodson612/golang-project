package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/peterhellberg/giphy"
)

func get_gif(tokens []string, ev *slack.MessageEvent, rtm *slack.RTM){

    text := strings.Join(tokens[1:], " ")

    g := giphy.DefaultClient
    res, err := g.Search([]string{text})

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    if len(res.Data) > 0 {
        rtm.SendMessage(rtm.NewOutgoingMessage(res.Data[0].BitlyGifURL, ev.Channel))
    }

}
