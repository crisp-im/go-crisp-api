# go-crisp-api

[![Build Status](https://travis-ci.org/crisp-im/go-crisp-api.svg?branch=master)](https://travis-ci.org/crisp-im/go-crisp-api)

The Crisp API Golang wrapper. Authenticate, send messages, fetch conversations, access your agent accounts from your Go code.

## API Reference

The full API reference is available at: [docs.crisp.im/api/](https://docs.crisp.im/api/)

## Usage

```go
import "github.com/crisp-im/go-crisp-api/crisp"
```

Construct a new Crisp client, then use the various services on the client to
access different parts of the Crisp API. For example:

```go
client := crisp.NewClient(nil)

// Get plugin information
plugin, _, err := client.Plugin.GetPluginInformation("185fe7ee-7cc6-4b8b-884d-fda9df632c13", nil)
```
