/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=rancheros.cattle.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineInventoryList is a list of MachineInventory resources
type MachineInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MachineInventory `json:"items"`
}

func NewMachineInventory(namespace, name string, obj MachineInventory) *MachineInventory {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("MachineInventory").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineRegistrationList is a list of MachineRegistration resources
type MachineRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MachineRegistration `json:"items"`
}

func NewMachineRegistration(namespace, name string, obj MachineRegistration) *MachineRegistration {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("MachineRegistration").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedOSImageList is a list of ManagedOSImage resources
type ManagedOSImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ManagedOSImage `json:"items"`
}

func NewManagedOSImage(namespace, name string, obj ManagedOSImage) *ManagedOSImage {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("ManagedOSImage").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}
