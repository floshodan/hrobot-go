# hrobot - A go library for the Hetzner Robot API 

The package hrobot is a libary for the Hetzner Robot API. 
The public API documentation is available [robot.your-server.de ](https://robot.your-server.de/doc/webservice/en.html#preface)

> Please note this is not an official Hetzner product, use at own risk!  

## Getting started

``` go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/floshodan/hrobot-go/hrobot"
)

func main() {

	//export HETZNER_TOKEN=USERNAME:PASSWORD
	client := hrobot.NewClient(hrobot.WithToken(os.Getenv("HETZNER_TOKEN")))
    
    // get list of server
    server, _, err := client.Server.List(context.Background())
    if err != nil {
        log.Fatalf("error retrieving server list: %s\n", err)
    }
	fmt.Printf("%+v\n", server[0])

    // retriev information for specific server id
	server2, _, err := client.Server.GetServerById(context.Background(), "1337")
    if err != nil {
        log.Fatalf("error retrieving server: %s\n", err)
    }
	fmt.Printf("%+v\n", server2)
}

```

## Authentication

To authenticate the client we can use two different inbuild methods: 

1) By "Token" eg as an environment variable: 

``` go
	client := hrobot.NewClient(hrobot.WithToken(os.Getenv("HETZNER_TOKEN")))
```
The "Token" has the following structure: "username:password". 

To use the Token as an enviroment variable as in the example above you can export a variable: `export HETZNER_TOKEN="username:password"` in your terminal. To make it persitent in your system you can put the export command in your `~/.profile` file.


2) Use the WithBasicAuth method which takes in (username, password) as string parameters. 
``` go 
	client := hrobot.WithBasicAuth(username, password)
```


