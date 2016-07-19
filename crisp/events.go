// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "github.com/graarh/golang-socketio"
  "github.com/graarh/golang-socketio/transport"
  "time"
  "fmt"
)


// EventsService service
type EventsService service

// EventsRegister register
type EventsRegister struct {
  // TODO
}

// EventsPayloadAuthentication maps authenticate payload
type EventsPayloadAuthentication struct {
  Username  string  `json:"username"`
  Password  string  `json:"password"`
}

// EventsMessage message
type EventsMessage struct {
  ID      int    `json:"id"`
  Channel string `json:"channel"`
  Text    string `json:"text"`
}


// Listen starts listening for incoming realtime events.
func (service *EventsService) Listen(handleConnected func(*gosocketio.Client, error)) {
  // TODO: dynamic socket configuration!
  so, err := gosocketio.Dial(
    gosocketio.GetUrl("relay-app.crisp.im.dev", 3050, false),
    //gosocketio.GetUrl("relay-app.crisp.im", 443, true),
    transport.GetDefaultWebsocketTransport(),
  )

  if err == nil {
    so.On("authenticated", func(chnl *gosocketio.Channel, msg EventsMessage) {
      fmt.Printf("authenticated => %v", msg)

      handleConnected(so, err)
    })

    time.Sleep(4 * time.Second)

    // Authenticate to socket
    so.Channel.Emit("authentication", EventsPayloadAuthentication{"lol", "hey"})

    time.Sleep(60 * time.Second)
  } else {
    // TODO: callback handle on error
    fmt.Printf("TODO Error: %v", err)
  }
}
