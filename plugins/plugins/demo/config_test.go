// Copyright The HTNN Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestPlugin(t *testing.T) {
	data := []byte(`{"hostName":"Jack"}`)
	c := &config{}
	protojson.Unmarshal(data, c)
	err := c.Validate()
	assert.Nil(t, err)
	assert.Equal(t, "Jack", c.HostName)
	err = c.Init(nil)
	assert.Nil(t, err)
}

func TestBadConfig(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   string
	}{
		{
			name:  "no hostName",
			input: `{}`,
			err:   "invalid Config.HostName: value length must be at least 1 runes",
		},
		{
			name:  "empty hostName",
			input: `{"hostName":""}`,
			err:   "invalid Config.HostName: value length must be at least 1 runes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &config{}
			err := protojson.Unmarshal([]byte(tt.input), conf)
			if err == nil {
				err = conf.Validate()
			}
			assert.NotNil(t, err)
			assert.ErrorContains(t, err, tt.err)
		})
	}
}
