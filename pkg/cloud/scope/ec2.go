/*
Copyright 2018 The Kubernetes Authors.

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

package scope

import (
	"github.com/go-logr/logr"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	expinfrav1 "sigs.k8s.io/cluster-api-provider-aws/exp/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud"
	"sigs.k8s.io/cluster-api/util/conditions"
)

// Scope is the interface for the scoep to be used with the ec2 service
type EC2Scope interface {
	cloud.ClusterScoper

	// VPC returns the cluster VPC.
	VPC() *infrav1.VPCSpec

	// Subnets returns the cluster subnets.
	Subnets() infrav1.Subnets

	// Network returns the cluster network object.
	Network() *infrav1.Network

	// SecurityGroups returns the cluster security groups as a map, it creates the map if empty.
	SecurityGroups() map[infrav1.SecurityGroupRole]infrav1.SecurityGroup

	// Bastion returns the bastion details for the cluster.
	Bastion() *infrav1.Bastion

	// SetBastionInstance sets the bastion instance in the status of the cluster.
	SetBastionInstance(instance *infrav1.Instance)

	// SSHKeyName returns the SSH key name to use for instances.
	SSHKeyName() *string

	// ImageLookupFormat returns the format string to use when looking up AMIs
	ImageLookupFormat() string

	// ImageLookupOrg returns the organization name to use when looking up AMIs
	ImageLookupOrg() string

	// ImageLookupBaseOS returns the base operating system name to use when looking up AMIs
	ImageLookupBaseOS() string
}

// LaunchTemplateOwner encapsulates a scope referencing a LaunchTemplate
type LaunchTemplateOwner interface {
	logr.Logger
	Name() string
	AdditionalTags() infrav1.Tags
	LaunchTemplateSpec() *expinfrav1.AWSLaunchTemplate
	StatusLaunchTemplateID() string
	CoreSecurityGroups(s EC2Scope) ([]string, error)
	OwnerObject() conditions.Setter
}
