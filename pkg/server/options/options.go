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

package options

import (
	"os"
	"path/filepath"

	"yunion.io/x/log"
	"yunion.io/x/onecloud/pkg/cloudcommon/options"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/utils"
)

type BaremanOptions struct {
	options.BaseOptions
	options.DBOptions

	ListenInterface       string `help:"Master net interface of bareman server" default:"eth0"`
	AccessAddress         string `help:"Management IP address of baremetal server, only need to be used when multiple address bind to listen_interface"`
	ListenAddress         string `help:"PXE serve IP address to select when multiple address bind to listen_interface"`
	BaremetalsPath        string `default:"/opt/bareman/baremetals" help:"Path for baremetals configuration files"`
	TftpRoot              string `help:"TFTP root directory" default:"/opt/bareman"`
	AutoRegisterBaremetal bool   `help:"Automatically create a baremetal instance"`

	LinuxDefaultRootUser      bool   `default:"false" help:"Default account for linux system is root"`
	IpmiLanPortShared         bool   `default:"false" help:"IPMI Lan port shared or dedicated"`
	DhcpLeaseTime             int    `default:"100663296" help:"DHCP lease time in seconds"`  // 0x6000000
	DhcpRenewalTime           int    `default:"67108864" help:"DHCP renewal time in seconds"` // 0x4000000
	EnableGeneralGuestDhcp    bool   `default:"false" help:"Enable DHCP service for general guest, e.g. those on VMware ESXi or Xen"`
	ForceDhcpProbeIpmi        bool   `default:"false" help:"Force DHCP probe IPMI interface network connection"`
	TftpBlockSizeInBytes      int    `default:"1024" help:"tftp block size, default is 1024"`
	TftpMaxTimeoutRetries     int    `default:"50" help:"Maximal tftp timeout retries, default is 50"`
	LengthyWorkerCount        int    `default:"8" help:"Parallel worker count for lengthy tasks"`
	ShortWorkerCount          int    `default:"8" help:"Parallel worker count for short-lived tasks"`
	TaskWorkerCount           int    `default:"32" help:"Parallel worker count for tasks"`
	DefaultIpmiPassword       string `help:"Default IPMI passowrd"`
	DefaultStrongIpmiPassword string `help:"Default strong IPMI passowrd"`

	WindowsDefaultAdminUser bool `default:"true" help:"Default account for Windows system is Administrator"`

	CachePath     string `help:"local image cache directory"`
	EnablePxeBoot bool   `help:"Enable DHCP PXE boot" default:"true"`
	BootIsoPath   string `help:"iso boot image path"`

	StatusProbeIntervalSeconds int `help:"interval to probe baremetal status, default is 60 seconds" default:"60"`
	LogFetchIntervalSeconds    int `help:"interval to fetch baremetal log, default is 900 seconds" default:"900"`
	SendMetricsIntervalSeconds int `help:"interval to send baremetal metrics, default is 300 seconds" default:"300"`

	TftpFileMap map[string]string `help:"map of filename to real file path for tftp"`
	BootLoader  string            `help:"PXE boot loader" default:"grub"`
}

const (
	BOOT_LOADER_SYSLINUX = "syslinux"
	BOOT_LOADER_GRUB     = "grub"
)

var (
	baremanOptions *BaremanOptions
)

func GetOptions() *BaremanOptions {
	if baremanOptions == nil {
		baremanOptions = new(BaremanOptions)
	}
	return baremanOptions
}

func Init() error {
	o := GetOptions()
	options.ParseOptions(o, os.Args, "bareman.conf", "bareman")

	if len(o.CachePath) == 0 {
		o.CachePath = filepath.Join(filepath.Dir(o.BaremetalsPath), "bm_image_cache")
		log.Infof("No cachepath, use default %s", o.CachePath)
	}
	if len(o.BootIsoPath) == 0 {
		o.BootIsoPath = filepath.Join(filepath.Dir(o.BaremetalsPath), "bm_boot_iso")
		log.Infof("No BootIsoPath, use default %s", o.BootIsoPath)
		err := os.MkdirAll(o.BootIsoPath, os.FileMode(0760))
		if err != nil {
			return errors.Errorf("fail to create BootIsoPath %s", o.BootIsoPath)
		}
	}
	if !utils.IsInStringArray(o.BootLoader, []string{BOOT_LOADER_GRUB, BOOT_LOADER_SYSLINUX}) {
		return errors.Errorf("invalid boot_loader option: %q", o.BootLoader)
	}
	return nil
}
