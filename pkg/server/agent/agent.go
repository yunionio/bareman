// Copyright 2022 Yunion
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

package agent

import (
	"yunion.io/x/pkg/errors"

	"github.com/yunionio/bareman/pkg/server/manager"
	"github.com/yunionio/bareman/pkg/server/options"
	"github.com/yunionio/bareman/pkg/server/pxe"
)

var (
	bmAgent Agent
)

type Agent interface{}

type baremetalAgent struct {
	pxeServer *pxe.Server
	manager   manager.BaremetalManager
}

func NewAgent(o *options.BaremanOptions) (Agent, error) {
	manager, err := manager.NewBaremetalManager(o.BaremetalsPath)
	if err != nil {
		return nil, errors.Wrap(err, "NewBaremetalManager")
	}

	pxeServer := &pxe.Server{
		TFTPRootDir: o.TftpRoot,
	}

	return &baremetalAgent{
		pxeServer: pxeServer,
		manager:   manager,
	}, nil
}
