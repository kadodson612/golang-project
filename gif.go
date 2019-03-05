package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/peterhellberg/giphy"

)

func get_gif(tokens []string) {

    text := strings.Join(tokens[1:], " ")

    g := giphy.DefaultClient
    res, err := g.Search([]string{text})

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}
