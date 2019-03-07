# squadbot

#### Cami Carballo, Eugene (Maru) Choi, Kelly Dodson

## Dependencies

- github.com/nlopes/slack

```
go get -u github.com/nlopes/slack
```

- gopkg.in/yaml.v2

```
go get -u gopkg.in/yaml.v2
```

- github.com/perterhellberg/giphy

```
go get -u github.com/peterhellberg/giphy
```


## Usage

```
go build 
```

```
./squadbot
```

## Commands

### Add Methods:
```
*addfriend <name>
```

```
*addphrase <name> <phrase>
```

```
*addalias <name> <alias>
```

```
*addjoke <joke>
```

```
*addinsult <insult>
```

### Remove Methods:
```
*rmfriend <name>
```

```
*rmphrase <name> <phrase>
```

```
*rmalias <name> <alias>
```

```
*rmjoke <joke>
```

```
*rminsult <insult>
```

### Speak Methods:
```
*speak <name>
```

```
*joke
```

```
*aka <name> 
```

```
*insult
```

```
*show <insults | jokes | friends>
```
