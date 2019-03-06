package main

import (
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

type Config struct {
    Friends map[string]Friend
    Jokes []string
    Insults []string
}

type Friend struct {
    Phrases []string
    Aliases []string
}

func read_yaml(file string) Config {

    var config Config

    source, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }

    return(config)

}

func write_yaml(file string, config Config) {

    to_write, err := yaml.Marshal(&config)
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(file, to_write, 0644)
    if err != nil {
        panic(err)
    }

}
