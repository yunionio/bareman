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

package manager

import (
	"io/ioutil"
	"os"

	"yunion.io/x/onecloud/pkg/util/procutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/regutils"
	"yunion.io/x/pkg/util/workqueue"
)

type BaremetalManager interface{}

type baremetalManager struct {
	// agent      interface{}
	homePath   string
	baremetals *baremetalMap
}

func NewBaremetalManager(homePath string) (BaremetalManager, error) {
	if err := os.MkdirAll(homePath, 0755); err != nil {
		return nil, errors.Wrapf(err, "mkdir %s", homePath)
	}
	return &baremetalManager{
		homePath:   homePath,
		baremetals: newBaremetalMap(),
	}, nil
}

func (m *baremetalManager) killAllIPMITool() {
	procutils.NewCommand("killall", "-9", "ipmitool").Run()
}

func (m *baremetalManager) loadConfigs() ([]os.FileInfo, error) {
	m.killAllIPMITool()
	files, err := ioutil.ReadDir(m.homePath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (m *baremetalManager) initBaremetals(files []os.FileInfo) error {
	bmIds := make([]string, 0)
	for _, file := range files {
		if file.IsDir() && regutils.MatchUUID(file.Name()) {
			bmIds = append(bmIds, file.Name())
		}
	}

	errsChannel := make(chan error, len(bmIds))
	initBaremetal := func(i int) {
		bmId := bmIds[i]
		err := m.initBaremetal(bmId, true)
		if err != nil {
			errsChannel <- err
			return
		}
	}
	workqueue.Parallelize(4, len(bmIds), initBaremetal)
	errs := make([]error, 0)
	if len(errsChannel) > 0 {
		length := len(errsChannel)
		for ; length > 0; length-- {
			errs = append(errs, <-errsChannel)
		}
	}
	return errors.NewAggregate(errs)
}

func (m *baremetalManager) initBaremetal(bmId string, update bool) error {
	return nil
}
