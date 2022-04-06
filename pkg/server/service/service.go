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

package service

import (
	"log"

	app_common "yunion.io/x/onecloud/pkg/cloudcommon/app"
	"yunion.io/x/onecloud/pkg/cloudcommon/service"

	o "github.com/yunionio/bareman/pkg/server/options"
)

type BaremanService interface {
	Start()
}

type baremanService struct {
	service.SServiceBase
}

func New() BaremanService {
	return new(baremanService)
}

func (s *baremanService) Start() {
	if err := o.Init(); err != nil {
		log.Fatalf("Init options error: %v", err)
	}

	app := app_common.InitApp(&o.GetOptions().BaseOptions, true)

}
