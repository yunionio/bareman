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

package provider

import "net"

type Network interface {
	GetType() string
	GetPXEConfig(relayAddr net.IP) (*NetworkConfig, error)
}

type NetworkConfig struct {
	DHCP    string `json:"dhcp"`
	Gateway string `json:"gateway"`
	IPStart string `json:"ip_start"`
	IPEnd   string `json:"ip_end"`
	IPMask  int    `json:"ip_mask"`
	Name    string `json:"name"`
	VlanId  int    `json:"vlan_id"`
	WireId  string `json:"wire_id"`
}
