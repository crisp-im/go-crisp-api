// Copyright 2022 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp/v3"
  "fmt"
  "io/ioutil"
  "net/http"
)

const (
  HOOKS_LISTENER = "localhost:8080"
  HOOKS_PATH_EVENTS = "/hook/events"
)

const (
  CONFIG_TOKEN_IDENTIFIER = "43f34724-9eeb-4474-9cec-560250754dec"
  CONFIG_TOKEN_KEY = "d12e60c5d2aa264b90997a641b6474ffd6602b66d8e8abc49634c404f06fa7d0"
  CONFIG_TOKEN_SIGNING_SECRET = "914a657d8cc2efbc20593f1806f9bde2"
)

func main() {
  client := crisp.New()
  client.AuthenticateTier("plugin", CONFIG_TOKEN_IDENTIFIER, CONFIG_TOKEN_KEY)

  // Subscribe to realtime events (RTM events from Web Hooks)
  client.Events.Listen(
    crisp.EventsModeWebHooks,

    []string{
      "message:send",
      "message:received",
    },

    func(reg *crisp.EventsRegister) {
      fmt.Print("Web Hooks channel bound: now listening for events\n")

      reg.On("message:send/text", func(evt crisp.EventsReceiveTextMessage) {
        fmt.Printf("[message:send/text] %s\n", evt)
      })

      reg.On("message:received/text", func(evt crisp.EventsReceiveTextMessage) {
        fmt.Printf("[message:received/text] %s\n", evt)
      })
    },

    func() {
      fmt.Print("Web Hooks channel unbound: not listening anymore\n")
    },

    func(err error) {
      fmt.Printf("Web Hooks channel error: %+v\n", err)
    },
  )

  // Start Web Hooks HTTP server
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }

    fmt.Fprintf(w, "Web Hooks endpoint is running at: http://%s%s", HOOKS_LISTENER, HOOKS_PATH_EVENTS)
  })

  http.HandleFunc(HOOKS_PATH_EVENTS, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }
    if r.Body == nil {
      w.WriteHeader(http.StatusBadRequest)
      return
    }

    // Read Web Hook data
    hookTimestamp := r.Header.Get("X-Crisp-Request-Timestamp")
    hookSignature := r.Header.Get("X-Crisp-Signature")
    hookPayload, _ :=  ioutil.ReadAll(r.Body)

    // Verify Web Hook payload
    if client.Events.VerifyHook(CONFIG_TOKEN_SIGNING_SECRET, &hookPayload, hookTimestamp, hookSignature) != true {
      fmt.Printf("Web Hooks request could not be verified with signature: %s\n", hookSignature)

      w.WriteHeader(http.StatusForbidden)
      return
    }

    // Receive Web Hook payload?
    // Notice: receiver goes asynchronous internally, therefore there is no need to spawn a goroutine from there.
    routed, err := client.Events.ReceiveHook(&hookPayload)

    if err != nil {
      // Failed processing Web Hook payload
      fmt.Printf("Web Hooks payload processing error: %+v\n", err)

      w.WriteHeader(http.StatusBadRequest)
      return
    }

    fmt.Printf("Web Hooks payload processed (routed=%t):\n \\-> %+v\n", routed, string(hookPayload))

    // Web Hook payload received and processed
    w.WriteHeader(http.StatusOK)
  })

  fmt.Printf("Web Hooks HTTP endpoint is: http://%s%s\n", HOOKS_LISTENER, HOOKS_PATH_EVENTS)

  http.ListenAndServe(HOOKS_LISTENER, nil)
}
