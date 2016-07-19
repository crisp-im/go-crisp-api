// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp"
  "fmt"
  "time"
)

func main() {
  client := crisp.New()
  client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")

  // Subscribe to realtime events
  client.Events.Listen(func(reg *crisp.EventsRegister, err error) {
    if err != nil {
      fmt.Print("Could not listen: %s\n", err)
    } else {
      fmt.Print("Now listening for events\n")

      reg.On("message:send", func(evt crisp.EventsReceiveMessageSend) {
        fmt.Printf("[message:send] %s\n", evt)
      })

      reg.On("message:received", func(evt crisp.EventsReceiveMessageReceived) {
        fmt.Printf("[message:received] %s\n", evt)
      })

      reg.On("message:compose:send", func(evt crisp.EventsReceiveMessageComposeSend) {
        fmt.Printf("[message:compose:send] %s\n", evt)
      })

      reg.On("message:compose:receive", func(evt crisp.EventsReceiveMessageComposeReceive) {
        fmt.Printf("[message:compose:receive] %s\n", evt)
      })
    }
  })

  for {
    time.Sleep(1 * time.Second)
  }
}
