// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp/v4"
  "fmt"
)

const (
  CONFIG_TOKEN_IDENTIFIER = "43f34724-9eeb-4474-9cec-560250754dec"
  CONFIG_TOKEN_KEY = "d12e60c5d2aa264b90997a641b6474ffd6602b66d8e8abc49634c404f06fa7d0"
)

func main() {
  client := crisp.New()
  client.AuthenticateTier("plugin", CONFIG_TOKEN_IDENTIFIER, CONFIG_TOKEN_KEY)

  // Build text message
  message := crisp.ConversationTextMessageNew{Type: "text", From: "operator", Origin: "chat", Content: "Hello :)", Fingerprint: 12345}

  // Send text message for `website_id` and `session_id`
  _, _, err := client.Website.SendTextMessageInConversation("8c842203-7ed8-4e29-a608-7cf78a7d2fcc", "session_19e5240f-0a8d-461e-a661-a3123fc6eec9", message)

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Print("Message sent")
  }
}
