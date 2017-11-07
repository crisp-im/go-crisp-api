// Copyright 2017 Crisp IM SARL All rights reserved.
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
  // Notice: set event list to '[]string{}' to listen to all event namespaces
  client.Events.Listen([]string{"message:send", "message:received", "message:compose:send", "message:compose:receive"}, func(reg *crisp.EventsRegister) {
    fmt.Print("Now listening for events\n")

    reg.On("message:send/text", func(evt crisp.EventsReceiveTextMessage) {
      fmt.Printf("[message:send/text] %s\n", evt)
    })

    reg.On("message:send/file", func(evt crisp.EventsReceiveFileMessage) {
      fmt.Printf("[message:send/file] %s\n", evt)
    })

    reg.On("message:send/animation", func(evt crisp.EventsReceiveAnimationMessage) {
      fmt.Printf("[message:send/animation] %s\n", evt)
    })

    reg.On("message:send/audio", func(evt crisp.EventsReceiveAudioMessage) {
      fmt.Printf("[message:send/audio] %s\n", evt)
    })

    reg.On("message:send/picker", func(evt crisp.EventsReceivePickerMessage) {
      fmt.Printf("[message:send/picker] %s\n", evt)
    })

    reg.On("message:send/field", func(evt crisp.EventsReceiveFieldMessage) {
      fmt.Printf("[message:send/field] %s\n", evt)
    })

    reg.On("message:send/note", func(evt crisp.EventsReceiveNoteMessage) {
      fmt.Printf("[message:send/note] %s\n", evt)
    })

    reg.On("message:received/text", func(evt crisp.EventsReceiveTextMessage) {
      fmt.Printf("[message:received/text] %s\n", evt)
    })

    reg.On("message:received/file", func(evt crisp.EventsReceiveFileMessage) {
      fmt.Printf("[message:received/file] %s\n", evt)
    })

    reg.On("message:received/animation", func(evt crisp.EventsReceiveAnimationMessage) {
      fmt.Printf("[message:received/animation] %s\n", evt)
    })

    reg.On("message:received/audio", func(evt crisp.EventsReceiveAudioMessage) {
      fmt.Printf("[message:received/audio] %s\n", evt)
    })

    reg.On("message:received/picker", func(evt crisp.EventsReceivePickerMessage) {
      fmt.Printf("[message:received/picker] %s\n", evt)
    })

    reg.On("message:received/field", func(evt crisp.EventsReceiveFieldMessage) {
      fmt.Printf("[message:received/field] %s\n", evt)
    })

    reg.On("message:received/note", func(evt crisp.EventsReceiveNoteMessage) {
      fmt.Printf("[message:received/note] %s\n", evt)
    })

    reg.On("message:compose:send", func(evt crisp.EventsReceiveMessageComposeSend) {
      fmt.Printf("[message:compose:send] %s\n", evt)
    })

    reg.On("message:compose:receive", func(evt crisp.EventsReceiveMessageComposeReceive) {
      fmt.Printf("[message:compose:receive] %s\n", evt)
    })
  })

  for {
    time.Sleep(1 * time.Second)
  }
}
