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

package hmacauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"

	"mosn.io/htnn/types/plugins/hmacauth"
)

func TestConsumerConfig(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   string
	}{
		{
			name:  "name in signedHeaders",
			input: `{"accessKey":"a", "secretKey":"s", "signedHeaders":[""]}`,
			err:   "value length must be at least 1 runes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &hmacauth.ConsumerConfig{}
			err := protojson.Unmarshal([]byte(tt.input), conf)
			if err == nil {
				err = conf.Validate()
			}
			if tt.err == "" {
				assert.Nil(t, err)
			} else {
				assert.ErrorContains(t, err, tt.err)
			}
		})
	}
}
