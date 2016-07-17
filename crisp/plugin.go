// Copyright 2016 The go-crisp-api AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

type PluginService service

type Plugin struct {
  One  *string  `json:""`
}

func (u Plugin) String() string {
  return Stringify(u)
}

// GetPluginInformation
// Reference: https://docs.crisp.im/api/v1/#plugin-one-plugin-get
func (service *PluginService) GetPluginInformation(plugin_id int) (
  *User, *Response, error
) {
  u := fmt.Sprintf("plugin/%s", plugin_id)
  req, err := service.client.NewRequest("GET", u, nil)
  if err != nil {
    return nil, nil, err
  }

  plugin := new(Plugin)
  resp, err := service.client.Do(req, plugin)
  if err != nil {
    return nil, resp, err
  }

  return plugin, resp, err
}
