/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package v1

import (
	"github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-custom-scheduler/config"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-netservice/pkg/netservice"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-netservice/pkg/netservice/types"

	"fmt"
	"strings"
)

//BcsConfig config item for ipscheduler
type BcsConfig struct {
	ZkHost   string         `json:"zkHost"`
	TLS      *types.SSLInfo `json:"tls,omitempty"`
	Interval int            `json:"interval,omitempty"`
}

func createNetSvcClient(conf *config.CustomSchedulerConfig) (netservice.Client, error) {
	bcsConf := newBcsConf(conf)

	var client netservice.Client
	var clientErr error
	if bcsConf.TLS == nil {
		client, clientErr = netservice.NewClient()
	} else {
		client, clientErr = netservice.NewTLSClient(bcsConf.TLS.CACert, bcsConf.TLS.Key, bcsConf.TLS.PubKey, bcsConf.TLS.Passwd)
	}
	if clientErr != nil {
		return nil, clientErr
	}
	//client get bcs-netservice info
	hosts := strings.Split(bcsConf.ZkHost, ";")
	if err := client.GetNetService(hosts); err != nil {
		return nil, fmt.Errorf("get netservice failed, %s", err.Error())
	}
	return client, nil
}

func newBcsConf(conf *config.CustomSchedulerConfig) BcsConfig {
	bcsConf := BcsConfig{
		ZkHost: conf.ZkHosts,
		TLS: &types.SSLInfo{
			CACert: conf.ClientCert.CAFile,
			Key:    conf.ClientCert.KeyFile,
			PubKey: conf.ClientCert.CertFile,
			Passwd: conf.ClientCert.CertPasswd,
		},
	}

	return bcsConf
}
