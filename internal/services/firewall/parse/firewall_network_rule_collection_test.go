// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = FirewallNetworkRuleCollectionId{}

func TestFirewallNetworkRuleCollectionIDFormatter(t *testing.T) {
	actual := NewFirewallNetworkRuleCollectionID("00000000-0000-0000-0000-000000000000", "mygroup1", "myfirewall", "networkRuleCollection1").ID()
	expected := "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/networkRuleCollections/networkRuleCollection1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestFirewallNetworkRuleCollectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *FirewallNetworkRuleCollectionId
	}{
		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/",
			Error: true,
		},

		{
			// missing AzureFirewallName
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/",
			Error: true,
		},

		{
			// missing value for AzureFirewallName
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/",
			Error: true,
		},

		{
			// missing NetworkRuleCollectionName
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/",
			Error: true,
		},

		{
			// missing value for NetworkRuleCollectionName
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/networkRuleCollections/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/azureFirewalls/myfirewall/networkRuleCollections/networkRuleCollection1",
			Expected: &FirewallNetworkRuleCollectionId{
				SubscriptionId:            "00000000-0000-0000-0000-000000000000",
				ResourceGroup:             "mygroup1",
				AzureFirewallName:         "myfirewall",
				NetworkRuleCollectionName: "networkRuleCollection1",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/00000000-0000-0000-0000-000000000000/RESOURCEGROUPS/MYGROUP1/PROVIDERS/MICROSOFT.NETWORK/AZUREFIREWALLS/MYFIREWALL/NETWORKRULECOLLECTIONS/NETWORKRULECOLLECTION1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := FirewallNetworkRuleCollectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.AzureFirewallName != v.Expected.AzureFirewallName {
			t.Fatalf("Expected %q but got %q for AzureFirewallName", v.Expected.AzureFirewallName, actual.AzureFirewallName)
		}
		if actual.NetworkRuleCollectionName != v.Expected.NetworkRuleCollectionName {
			t.Fatalf("Expected %q but got %q for NetworkRuleCollectionName", v.Expected.NetworkRuleCollectionName, actual.NetworkRuleCollectionName)
		}
	}
}
