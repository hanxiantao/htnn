// Copyright The HTNN Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugins

import (
	_ "mosn.io/htnn/controller/plugins/bandwidthlimit"
	_ "mosn.io/htnn/controller/plugins/buffer"
	_ "mosn.io/htnn/controller/plugins/cors"
	_ "mosn.io/htnn/controller/plugins/extproc"
	_ "mosn.io/htnn/controller/plugins/fault"
	_ "mosn.io/htnn/controller/plugins/listenerpatch"
	_ "mosn.io/htnn/controller/plugins/localratelimit"
	_ "mosn.io/htnn/controller/plugins/lua"
	_ "mosn.io/htnn/controller/plugins/networkrbac"
	_ "mosn.io/htnn/controller/plugins/routepatch"
	_ "mosn.io/htnn/controller/plugins/tlsinspector"
	_ "mosn.io/htnn/types/plugins"
)
