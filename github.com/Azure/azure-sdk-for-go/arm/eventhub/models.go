package eventhub

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen specifies the listen state for access rights.
	Listen AccessRights = "Listen"
	// Manage specifies the manage state for access rights.
	Manage AccessRights = "Manage"
	// Send specifies the send state for access rights.
	Send AccessRights = "Send"
)

// EntityStatus enumerates the values for entity status.
type EntityStatus string

const (
	// Active specifies the active state for entity status.
	Active EntityStatus = "Active"
	// Creating specifies the creating state for entity status.
	Creating EntityStatus = "Creating"
	// Deleting specifies the deleting state for entity status.
	Deleting EntityStatus = "Deleting"
	// Disabled specifies the disabled state for entity status.
	Disabled EntityStatus = "Disabled"
	// ReceiveDisabled specifies the receive disabled state for entity status.
	ReceiveDisabled EntityStatus = "ReceiveDisabled"
	// Renaming specifies the renaming state for entity status.
	Renaming EntityStatus = "Renaming"
	// Restoring specifies the restoring state for entity status.
	Restoring EntityStatus = "Restoring"
	// SendDisabled specifies the send disabled state for entity status.
	SendDisabled EntityStatus = "SendDisabled"
	// Unknown specifies the unknown state for entity status.
	Unknown EntityStatus = "Unknown"
)

// Policykey enumerates the values for policykey.
type Policykey string

const (
	// PrimaryKey specifies the primary key state for policykey.
	PrimaryKey Policykey = "PrimaryKey"
	// SecondaryKey specifies the secondary key state for policykey.
	SecondaryKey Policykey = "SecondaryKey"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic specifies the basic state for sku name.
	Basic SkuName = "Basic"
	// Standard specifies the standard state for sku name.
	Standard SkuName = "Standard"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic specifies the sku tier basic state for sku tier.
	SkuTierBasic SkuTier = "Basic"
	// SkuTierPremium specifies the sku tier premium state for sku tier.
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard specifies the sku tier standard state for sku tier.
	SkuTierStandard SkuTier = "Standard"
)

// UnavailableReason enumerates the values for unavailable reason.
type UnavailableReason string

const (
	// InvalidName specifies the invalid name state for unavailable reason.
	InvalidName UnavailableReason = "InvalidName"
	// NameInLockdown specifies the name in lockdown state for unavailable
	// reason.
	NameInLockdown UnavailableReason = "NameInLockdown"
	// NameInUse specifies the name in use state for unavailable reason.
	NameInUse UnavailableReason = "NameInUse"
	// None specifies the none state for unavailable reason.
	None UnavailableReason = "None"
	// SubscriptionIsDisabled specifies the subscription is disabled state for
	// unavailable reason.
	SubscriptionIsDisabled UnavailableReason = "SubscriptionIsDisabled"
	// TooManyNamespaceInCurrentSubscription specifies the too many namespace
	// in current subscription state for unavailable reason.
	TooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// CheckNameAvailability is description of a Check Name availability request
// properties.
type CheckNameAvailability struct {
	Name *string `json:"name,omitempty"`
}

// CheckNameAvailabilityResult is description of a Check Name availability
// request properties.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool             `json:"nameAvailable,omitempty"`
	Reason            UnavailableReason `json:"reason,omitempty"`
	Message           *string           `json:"message,omitempty"`
}

// ConsumerGroup is description of the consumer group resource.
type ConsumerGroup struct {
	autorest.Response        `json:"-"`
	ID                       *string `json:"id,omitempty"`
	Name                     *string `json:"name,omitempty"`
	Type                     *string `json:"type,omitempty"`
	*ConsumerGroupProperties `json:"properties,omitempty"`
}

// ConsumerGroupListResult is the response to the List Consumer Group
// operation.
type ConsumerGroupListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ConsumerGroup `json:"value,omitempty"`
	NextLink          *string          `json:"nextLink,omitempty"`
}

// ConsumerGroupListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ConsumerGroupListResult) ConsumerGroupListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ConsumerGroupProperties is description of a Consumer Group properties.
type ConsumerGroupProperties struct {
	CreatedAt    *date.Time `json:"createdAt,omitempty"`
	EventHubPath *string    `json:"eventHubPath,omitempty"`
	UpdatedAt    *date.Time `json:"updatedAt,omitempty"`
	UserMetadata *string    `json:"userMetadata,omitempty"`
}

// ErrorResponse is error reponse indicates EventHub service is not able to
// process the incoming request. The reason is provided in the error message.
type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// EventHub is description of the Event Hub resource.
type EventHub struct {
	autorest.Response `json:"-"`
	ID                *string `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *string `json:"type,omitempty"`
	*Properties       `json:"properties,omitempty"`
}

// ListResult is the response of the List Event Hubs operation.
type ListResult struct {
	autorest.Response `json:"-"`
	Value             *[]EventHub `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// ListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListResult) ListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Namespace is description of a namespace resource.
type Namespace struct {
	autorest.Response    `json:"-"`
	ID                   *string             `json:"id,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	Type                 *string             `json:"type,omitempty"`
	Location             *string             `json:"location,omitempty"`
	Tags                 *map[string]*string `json:"tags,omitempty"`
	Sku                  *Sku                `json:"sku,omitempty"`
	*NamespaceProperties `json:"properties,omitempty"`
}

// NamespaceListResult is the response of the List Namespace operation.
type NamespaceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Namespace `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// NamespaceListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client NamespaceListResult) NamespaceListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// NamespaceProperties is properties of the namespace.
type NamespaceProperties struct {
	ProvisioningState  *string    `json:"provisioningState,omitempty"`
	CreatedAt          *date.Time `json:"createdAt,omitempty"`
	UpdatedAt          *date.Time `json:"updatedAt,omitempty"`
	ServiceBusEndpoint *string    `json:"serviceBusEndpoint,omitempty"`
}

// NamespaceUpdateParameter is parameters supplied to the Patch Namespace
// operation.
type NamespaceUpdateParameter struct {
	Tags *map[string]*string `json:"tags,omitempty"`
	Sku  *Sku                `json:"sku,omitempty"`
}

// Operation is a EventHub REST API operation
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider  *string `json:"provider,omitempty"`
	Resource  *string `json:"resource,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult is result of the request to list EventHub operations. It
// contains a list of operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Properties is description of a EventHub Properties.
type Properties struct {
	CreatedAt              *date.Time   `json:"createdAt,omitempty"`
	MessageRetentionInDays *int64       `json:"messageRetentionInDays,omitempty"`
	PartitionCount         *int64       `json:"partitionCount,omitempty"`
	PartitionIds           *[]string    `json:"partitionIds,omitempty"`
	Status                 EntityStatus `json:"status,omitempty"`
	UpdatedAt              *date.Time   `json:"updatedAt,omitempty"`
}

// RegenerateKeysParameters is parameters supplied to the Regenerate
// Authorization Rule operation.
type RegenerateKeysParameters struct {
	Policykey Policykey `json:"policykey,omitempty"`
}

// Resource is the Resource definition
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ResourceListKeys is namespace/EventHub Connection String
type ResourceListKeys struct {
	autorest.Response         `json:"-"`
	PrimaryConnectionString   *string `json:"primaryConnectionString,omitempty"`
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`
	PrimaryKey                *string `json:"primaryKey,omitempty"`
	SecondaryKey              *string `json:"secondaryKey,omitempty"`
	KeyName                   *string `json:"keyName,omitempty"`
}

// SharedAccessAuthorizationRule is description of a namespace authorization
// rule.
type SharedAccessAuthorizationRule struct {
	autorest.Response                        `json:"-"`
	ID                                       *string `json:"id,omitempty"`
	Name                                     *string `json:"name,omitempty"`
	Type                                     *string `json:"type,omitempty"`
	*SharedAccessAuthorizationRuleProperties `json:"properties,omitempty"`
}

// SharedAccessAuthorizationRuleListResult is the response from the List
// Namespace operation.
type SharedAccessAuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SharedAccessAuthorizationRule `json:"value,omitempty"`
	NextLink          *string                          `json:"nextLink,omitempty"`
}

// SharedAccessAuthorizationRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SharedAccessAuthorizationRuleListResult) SharedAccessAuthorizationRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SharedAccessAuthorizationRuleProperties is sharedAccessAuthorizationRule
// properties.
type SharedAccessAuthorizationRuleProperties struct {
	Rights *[]AccessRights `json:"rights,omitempty"`
}

// Sku is sKU of the namespace.
type Sku struct {
	Name     SkuName `json:"name,omitempty"`
	Tier     SkuTier `json:"tier,omitempty"`
	Capacity *int32  `json:"capacity,omitempty"`
}

// TrackedResource is definition of Resource
type TrackedResource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}