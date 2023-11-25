// Copyright 2023 zyy17 <zyylsxm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package factory

import (
	"fmt"

	. "github.com/zyy17/pubsub/pkg/provider"
	"github.com/zyy17/pubsub/pkg/provider/etcd"
)

// NewProvider is a factory method to create a new provider with the given options.
func NewProvider(opts *Options) (Provider, error) {
	switch opts.StoreType {
	case StoreTypeEtcd:
		return etcd.NewEtcdProvider(opts)
	default:
		return nil, fmt.Errorf("unknown store type: %s", opts.StoreType)
	}
}
