# \SettingsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETRevenueAutomationStartDate**](SettingsApi.md#GETRevenueAutomationStartDate) | **Get** /v1/settings/finance/revenue-automation-start-date | Get the revenue automation start date


# **GETRevenueAutomationStartDate**
> GetRevenueStartDateSettingType GETRevenueAutomationStartDate($entityId, $entityName)

Get the revenue automation start date

This REST API reference describes how to get the revenue automation start date. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRevenueStartDateSettingType**](GETRevenueStartDateSettingType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

