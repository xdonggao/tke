/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1 "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
)

type FakePlatformV1 struct {
	*testing.Fake
}

func (c *FakePlatformV1) CSIOperators() v1.CSIOperatorInterface {
	return &FakeCSIOperators{c}
}

func (c *FakePlatformV1) Clusters() v1.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakePlatformV1) ClusterAddons() v1.ClusterAddonInterface {
	return &FakeClusterAddons{c}
}

func (c *FakePlatformV1) ClusterAddonTypes() v1.ClusterAddonTypeInterface {
	return &FakeClusterAddonTypes{c}
}

func (c *FakePlatformV1) ClusterAuthentications(namespace string) v1.ClusterAuthenticationInterface {
	return &FakeClusterAuthentications{c, namespace}
}

func (c *FakePlatformV1) ClusterCredentials() v1.ClusterCredentialInterface {
	return &FakeClusterCredentials{c}
}

func (c *FakePlatformV1) ConfigMaps() v1.ConfigMapInterface {
	return &FakeConfigMaps{c}
}

func (c *FakePlatformV1) CronHPAs() v1.CronHPAInterface {
	return &FakeCronHPAs{c}
}

func (c *FakePlatformV1) Helms() v1.HelmInterface {
	return &FakeHelms{c}
}

func (c *FakePlatformV1) IPAMs() v1.IPAMInterface {
	return &FakeIPAMs{c}
}

func (c *FakePlatformV1) LBCFs() v1.LBCFInterface {
	return &FakeLBCFs{c}
}

func (c *FakePlatformV1) LogCollectors() v1.LogCollectorInterface {
	return &FakeLogCollectors{c}
}

func (c *FakePlatformV1) Machines() v1.MachineInterface {
	return &FakeMachines{c}
}

func (c *FakePlatformV1) PersistentEvents() v1.PersistentEventInterface {
	return &FakePersistentEvents{c}
}

func (c *FakePlatformV1) Prometheuses() v1.PrometheusInterface {
	return &FakePrometheuses{c}
}

func (c *FakePlatformV1) Registries() v1.RegistryInterface {
	return &FakeRegistries{c}
}

func (c *FakePlatformV1) TappControllers() v1.TappControllerInterface {
	return &FakeTappControllers{c}
}

func (c *FakePlatformV1) VolumeDecorators() v1.VolumeDecoratorInterface {
	return &FakeVolumeDecorators{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePlatformV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
