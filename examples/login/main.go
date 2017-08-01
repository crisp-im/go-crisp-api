// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/crisp-im/go-crisp-api/crisp"
  "fmt"
)

func main() {
  client := crisp.New()

  // Login user `john@acme-inc.com`
  session, _, err := client.User.CreateNewSession("john@acme-inc.com", "SecurePassword")

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Session response (raw): %s\n", session)

    fmt.Printf("Got identifier: %s\n", *session.Identifier)
    fmt.Printf("Got key: %s\n", *session.Key)
  }
}
