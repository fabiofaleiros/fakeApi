// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package entity

import "encoding/json"

type Endpoint map[string]interface{}

func (c *Endpoint) Unmarshal(b []byte) error {
	return json.Unmarshal(b, c)
}
