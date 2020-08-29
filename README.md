# Twitchy-Gopher
A golang client library for Twitch API v5. We will add support for Helix API as soon as it has parity with V5. We currently support golang versions >= 1.11.

[![Build Status](https://travis-ci.org/ollieparsley/twitchy-gopher.svg?branch=master)](https://travis-ci.org/ollieparsley/twitchy-gopher) [![Coverage Status](https://coveralls.io/repos/github/ollieparsley/twitchy-gopher/badge.svg)](https://coveralls.io/github/ollieparsley/twitchy-gopher)

![The Twitchy Gopher](https://raw.githubusercontent.com/ollieparsley/twitchy-gopher/master/twitchy-gopher.png)

# Progress

This is a work in progress and is not complete, we have several endpoints to add. We will continue with API v5 until the Helix API has parity, currently this is a long way off (according to Twitch dev documentation).

# Installing

```
go get -u github.com/ollieparsley/twitchy-gopher
```

# Usage
This example is just to get a session and then get the basic authentication details. Once you create client/session you can make requests to the Twitch API.

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
