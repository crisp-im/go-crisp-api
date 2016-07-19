// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp"
  "fmt"
)

func main() {
  client := crisp.NewClient()
  client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")

  // Resolve conversation for `website_id` and `session_id`
  conversation, _, err := client.Website.GetConversation("8c842203-7ed8-4e29-a608-7cf78a7d2fcc", "session_19e5240f-0a8d-461e-a661-a3123fc6eec9")

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Conversation (raw): %s\n", conversation)
  }
}
