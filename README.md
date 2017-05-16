# Twitchy-Gopher
A golang client library for Twitch. Aiming to support V5 of the API and video uploading.

[![Build Status](https://travis-ci.org/ollieparsley/twitchy-gopher.svg?branch=master)](https://travis-ci.org/ollieparsley/twitchy-gopher) [![Coverage Status](https://coveralls.io/repos/github/ollieparsley/twitchy-gopher/badge.svg)](https://coveralls.io/github/ollieparsley/twitchy-gopher)

![The Twitchy Gopher](https://raw.githubusercontent.com/ollieparsley/twitchy-gopher/master/twitchy-gopher.png)

# Progress

This is a work in progress and is not complete. Version 5 of the Twitch API is also not fully released.

# Installing

```
go get -u github.com/ollieparsley/twitchy-gopher
```

# Usage
This has only been tested with Go 1.5-1.7. This example is just to get a session. From a session you can make requests to the Twitch API.

```
package main

import (
    "fmt"

    "github.com/ollieparsley/twitchy-gopher/twitch"
)

func main() {
	// Create new client
    t := twitch.NewClient(&twitch.OAuthConfig{
      ClientID: 'my-client-id',
      ClientSecret: 'my-client-secret',
      AccessToken: 'my-users-access-token',
    }, &http.Client{})

    // Make requests like this
    output, errorOutput := client.GetRoot()
    if errorOutput != nil {
        panic(errorOutput)
    }
}
```

# License
This SDK is distributed under the MIT License. See LICENSE for more information.

# Logo
The original gopher created by Takuya Ueda (https://twitter.com/tenntenn). Licensed under the Creative Commons 3.0 Attributions license.
