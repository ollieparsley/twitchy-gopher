# Twitchy-Gopher
A golang client library for Twitch. Aiming to support V5 of the API and video uploading.

![Travis-CI build status](https://travis-ci.org/ollieparsley/twitchy-gopher.svg?branch=master)

![The Twitchy Gopher](https://raw.githubusercontent.com/ollieparsley/twitchy-gopher/master/twitchy-gopher.png)

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
      AccessTokenSecret: 'my-users-access-token-secret',
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
