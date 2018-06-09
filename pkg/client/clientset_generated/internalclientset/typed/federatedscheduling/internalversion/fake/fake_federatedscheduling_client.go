/*
Copyright 2018 The Federation v2 Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package fake

import (
	internalversion "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/internalclientset/typed/federatedscheduling/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeFederatedscheduling struct {
	*testing.Fake
}

func (c *FakeFederatedscheduling) ReplicaSchedulingPreferences(namespace string) internalversion.ReplicaSchedulingPreferenceInterface {
	return &FakeReplicaSchedulingPreferences{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFederatedscheduling) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}