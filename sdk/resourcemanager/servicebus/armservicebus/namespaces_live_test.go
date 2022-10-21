//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
	"github.com/stretchr/testify/suite"
)

type NamespacesTestSuite struct {
	suite.Suite

	ctx                   context.Context
	cred                  azcore.TokenCredential
	options               *arm.ClientOptions
	authorizationRuleName string
	namespaceName         string
	location              string
	resourceGroupName     string
	subscriptionId        string
}

func (testsuite *NamespacesTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/servicebus/armservicebus/testdata")
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.authorizationRuleName = testutil.GenerateAlphaNumericID(testsuite.T(), "sbauthrule", 6)
	testsuite.namespaceName = testutil.GenerateAlphaNumericID(testsuite.T(), "sbnamespace", 6)
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
}

func (testsuite *NamespacesTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestNamespacesTestSuite(t *testing.T) {
	suite.Run(t, new(NamespacesTestSuite))
}

// Microsoft.ServiceBus/namespaces
func (testsuite *NamespacesTestSuite) TestNamespace() {
	var err error
	// From step Namespace_CheckNameAvailability
	namespacesClient, err := armservicebus.NewNamespacesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = namespacesClient.CheckNameAvailability(testsuite.ctx, armservicebus.CheckNameAvailability{
		Name: to.Ptr(testsuite.namespaceName),
	}, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_Create
	namespacesClientCreateOrUpdateResponsePoller, err := namespacesClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, armservicebus.SBNamespace{
		Location: to.Ptr(testsuite.location),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		SKU: &armservicebus.SBSKU{
			Name: to.Ptr(armservicebus.SKUNamePremium),
			Tier: to.Ptr(armservicebus.SKUTierPremium),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, namespacesClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step Namespace_Update
	_, err = namespacesClient.Update(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, armservicebus.SBNamespaceUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags: map[string]*string{
			"tag3": to.Ptr("value3"),
			"tag4": to.Ptr("value4"),
		},
	}, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_Get
	_, err = namespacesClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_List
	namespacesClientNewListPager := namespacesClient.NewListPager(nil)
	for namespacesClientNewListPager.More() {
		_, err := namespacesClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Namespace_ListByResourceGroup
	namespacesClientNewListByResourceGroupPager := namespacesClient.NewListByResourceGroupPager(testsuite.resourceGroupName, nil)
	for namespacesClientNewListByResourceGroupPager.More() {
		_, err := namespacesClientNewListByResourceGroupPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Operation_List
	operationsClient, err := armservicebus.NewOperationsClient(testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	operationsClientNewListPager := operationsClient.NewListPager(nil)
	for operationsClientNewListPager.More() {
		_, err := operationsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Namespace_CreateAuthorizationRule
	_, err = namespacesClient.CreateOrUpdateAuthorizationRule(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.authorizationRuleName, armservicebus.SBAuthorizationRule{
		Properties: &armservicebus.SBAuthorizationRuleProperties{
			Rights: []*armservicebus.AccessRights{
				to.Ptr(armservicebus.AccessRightsListen),
				to.Ptr(armservicebus.AccessRightsSend)},
		},
	}, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_GetAuthorizationRule
	_, err = namespacesClient.GetAuthorizationRule(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.authorizationRuleName, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_ListAuthorizationRules
	namespacesClientNewListAuthorizationRulesPager := namespacesClient.NewListAuthorizationRulesPager(testsuite.resourceGroupName, testsuite.namespaceName, nil)
	for namespacesClientNewListAuthorizationRulesPager.More() {
		_, err := namespacesClientNewListAuthorizationRulesPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Namespace_RegenerateKeys
	_, err = namespacesClient.RegenerateKeys(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.authorizationRuleName, armservicebus.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armservicebus.KeyTypePrimaryKey),
	}, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_ListKeys
	_, err = namespacesClient.ListKeys(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.authorizationRuleName, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_GetNetworkRuleSet
	_, err = namespacesClient.GetNetworkRuleSet(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_ListNetworkRuleSets
	namespacesClientNewListNetworkRuleSetsPager := namespacesClient.NewListNetworkRuleSetsPager(testsuite.resourceGroupName, testsuite.namespaceName, nil)
	for namespacesClientNewListNetworkRuleSetsPager.More() {
		_, err := namespacesClientNewListNetworkRuleSetsPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Namespace_DeleteAuthorizationRule
	_, err = namespacesClient.DeleteAuthorizationRule(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, testsuite.authorizationRuleName, nil)
	testsuite.Require().NoError(err)

	// From step Namespace_Delete
	namespacesClientDeleteResponsePoller, err := namespacesClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, testsuite.namespaceName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, namespacesClientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}