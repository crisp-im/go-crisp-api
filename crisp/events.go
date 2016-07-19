// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "github.com/graarh/golang-socketio"
  "github.com/graarh/golang-socketio/transport"
  "errors"
  "net"
  "net/url"
  "strconv"
)


// EventsService service
type EventsService service

// EventsNamed generic data mapper
type EventsNamed interface {
  Name()  string
}

// EventsRegister register
type EventsRegister struct {
  Handlers  map[string]func(EventsNamed)
}

// EventsSendAuthentication maps authentication
type EventsSendAuthentication struct {
  Username  string  `json:"username"`
  Password  string  `json:"password"`
}

// EventsReceiveMessageSend maps message:send
type EventsReceiveMessageSend struct {
  WebsiteID    *string                          `json:"website_id"`
  SessionID    *string                          `json:"session_id"`
  From         *string                          `json:"from"`
  Type         *string                          `json:"type"`
  Origin       *string                          `json:"origin"`
  Content      *string                          `json:"content"`
  Stamped      *bool                            `json:"stamped"`
  Timestamp    *uint                            `json:"timestamp"`
  Fingerprint  *int                             `json:"fingerprint"`
  User         *EventsReceiveMessageCommonUser  `json:"user"`
}

// EventsReceiveMessageCommonUser maps message:*/user
type EventsReceiveMessageCommonUser struct {
  UserID    *string  `json:"user_id"`
  Nickname  *string  `json:"nickname"`
}


// String returns the string representation of EventsReceiveMessageSend
func (evt EventsReceiveMessageSend) String() string {
  return Stringify(evt)
}

// Name returns the event name of EventsReceiveMessageSend
func (evt EventsReceiveMessageSend) Name() string {
  return "message:send"
}


// Listen starts listening for incoming realtime events.
func (service *EventsService) Listen(handleDone func(*EventsRegister, error)) {
  var secure bool

  u, _ := url.Parse(service.client.config.RealtimeEndpointURL)
  host, portString, _ := net.SplitHostPort(u.Host)

  port, _ := strconv.ParseInt(portString, 10, 64)

  if u.Scheme == "https" {
    secure = true
  } else {
    secure = false
  }

  so, err := gosocketio.Dial(
    gosocketio.GetUrl(host, int(port), secure),
    transport.GetDefaultWebsocketTransport(),
  )

  if err == nil {
    so.On("authenticated", func(chnl *gosocketio.Channel, authenticated bool) {
      if authenticated == true {
        reg := EventsRegister{Handlers: make(map[string]func(EventsNamed))}

        reg.BindEvents(so)

        handleDone(&reg, nil)
      } else {
        handleDone(nil, errors.New("Authentication failed"))
      }
    })

    // Authenticate to socket
    if service.client.basicAuth.Available == true {
      so.Channel.Emit("authentication", EventsSendAuthentication{service.client.basicAuth.Username, service.client.basicAuth.Password})
    } else {
      handleDone(nil, errors.New("Not authenticated"))
    }
  } else {
    handleDone(nil, err)
  }
}


// On registers an event handler on event name
func (register *EventsRegister) On(eventName string, handleEvent func(EventsNamed)) {
  register.Handlers[eventName] = handleEvent
}

// BindEvents listens for recognized incoming events
func (register *EventsRegister) BindEvents(so *gosocketio.Client) {
  so.On("message:send", func(chnl *gosocketio.Channel, message EventsReceiveMessageSend) {
    if handler, ok := register.Handlers["message:send"]; ok {
      handler(message)
    }
  })
}
