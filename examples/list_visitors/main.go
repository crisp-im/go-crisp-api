// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp/v3"
  "fmt"
)

const (
  CONFIG_TOKEN_IDENTIFIER = "7c3ef21c-1e04-41ce-8c06-5605c346f73e"
  CONFIG_TOKEN_KEY = "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a"
)

func main() {
  client := crisp.New()
  client.Authenticate(CONFIG_TOKEN_IDENTIFIER, CONFIG_TOKEN_KEY)

  // List for `website_id`
  visitors, _, err := client.Website.ListVisitors("8c842203-7ed8-4e29-a608-7cf78a7d2fcc", 1)

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Visitor list (raw): %v\n", visitors)

    for _, visitor := range *visitors {
      fmt.Printf("Got visitor `session_id`: %s\n", *visitor.SessionID)
    }
  }
}
