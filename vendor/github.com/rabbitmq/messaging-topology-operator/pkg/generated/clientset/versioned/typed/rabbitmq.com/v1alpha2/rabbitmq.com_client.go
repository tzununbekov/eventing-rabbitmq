/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/rabbitmq/messaging-topology-operator/api/v1alpha2"
	"github.com/rabbitmq/messaging-topology-operator/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type RabbitmqV1alpha2Interface interface {
	RESTClient() rest.Interface
	BindingsGetter
	ExchangesGetter
	PoliciesGetter
	QueuesGetter
	UsersGetter
	VhostsGetter
}

// RabbitmqV1alpha2Client is used to interact with features provided by the rabbitmq.com group.
type RabbitmqV1alpha2Client struct {
	restClient rest.Interface
}

func (c *RabbitmqV1alpha2Client) Bindings(namespace string) BindingInterface {
	return newBindings(c, namespace)
}

func (c *RabbitmqV1alpha2Client) Exchanges(namespace string) ExchangeInterface {
	return newExchanges(c, namespace)
}

func (c *RabbitmqV1alpha2Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

func (c *RabbitmqV1alpha2Client) Queues(namespace string) QueueInterface {
	return newQueues(c, namespace)
}

func (c *RabbitmqV1alpha2Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

func (c *RabbitmqV1alpha2Client) Vhosts(namespace string) VhostInterface {
	return newVhosts(c, namespace)
}

// NewForConfig creates a new RabbitmqV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*RabbitmqV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &RabbitmqV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new RabbitmqV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *RabbitmqV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new RabbitmqV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *RabbitmqV1alpha2Client {
	return &RabbitmqV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *RabbitmqV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}