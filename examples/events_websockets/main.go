// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp/v3"
  "fmt"
  "time"
)

const (
  CONFIG_TOKEN_IDENTIFIER = "43f34724-9eeb-4474-9cec-560250754dec"
  CONFIG_TOKEN_KEY = "d12e60c5d2aa264b90997a641b6474ffd6602b66d8e8abc49634c404f06fa7d0"
)

func main() {
  client := crisp.New()
  client.AuthenticateTier("plugin", CONFIG_TOKEN_IDENTIFIER, CONFIG_TOKEN_KEY)

  // Subscribe to realtime events (RTM API over WebSockets)
  client.Events.Listen(
    crisp.EventsModeWebSockets,

    []string{
      "message:send",
      "message:received",
      "message:compose:send",
      "message:compose:receive",
    },

    func(reg *crisp.EventsRegister) {
      fmt.Print("WebSocket channel is connected: now listening for events\n")

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

      reg.On("message:send/carousel", func(evt crisp.EventsReceiveCarouselMessage) {
        fmt.Printf("[message:send/carousel] %s\n", evt)
      })

      reg.On("message:send/note", func(evt crisp.EventsReceiveNoteMessage) {
        fmt.Printf("[message:send/note] %s\n", evt)
      })

      reg.On("message:send/event", func(evt crisp.EventsReceiveEventMessage) {
        fmt.Printf("[message:send/event] %s\n", evt)
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

      reg.On("message:received/carousel", func(evt crisp.EventsReceiveCarouselMessage) {
        fmt.Printf("[message:received/carousel] %s\n", evt)
      })

      reg.On("message:received/note", func(evt crisp.EventsReceiveNoteMessage) {
        fmt.Printf("[message:received/note] %s\n", evt)
      })

      reg.On("message:received/event", func(evt crisp.EventsReceiveEventMessage) {
        fmt.Printf("[message:received/event] %s\n", evt)
      })

      reg.On("message:compose:send", func(evt crisp.EventsReceiveMessageComposeSend) {
        fmt.Printf("[message:compose:send] %s\n", evt)
      })

      reg.On("message:compose:receive", func(evt crisp.EventsReceiveMessageComposeReceive) {
        fmt.Printf("[message:compose:receive] %s\n", evt)
      })
    },

    func() {
      fmt.Print("WebSocket channel is disconnected: will try to reconnect\n")
    },

    func(err error) {
      fmt.Printf("WebSocket channel error: %+v\n", err)
    },
  )

  for {
    time.Sleep(1 * time.Second)
  }
}
