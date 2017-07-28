# \MassUpdaterApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETMassUpdater**](MassUpdaterApi.md#GETMassUpdater) | **Get** /v1/bulk/{bulk-key} | Get mass action result
[**POSTMassUpdater**](MassUpdaterApi.md#POSTMassUpdater) | **Post** /v1/bulk | Perform mass action
[**PUTMassUpdater**](MassUpdaterApi.md#PUTMassUpdater) | **Put** /v1/bulk/{bulk-key}/stop | Stop mass action


# **GETMassUpdater**
> GetMassUpdateType GETMassUpdater($bulkKey, $entityId, $entityName)

Get mass action result

This reference describes how to get information about the result of a mass action through the REST API.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkKey** | **string**| String of 32 characters that identifies a mass action. You get the bulk-key after performing a mass action through the REST API.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetMassUpdateType**](GETMassUpdateType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTMassUpdater**
> PostMassUpdateResponseType POSTMassUpdater($contentType, $request, $entityId, $entityName)

Perform mass action

This reference describes how to perform a mass action through the REST API.   Using this API method, you send a multipart/form-data request containing a `.csv` file with data about the mass action you want to perform. Zuora returns a key and then asynchronously processes the mass action. You can use the key to get details about the result of the mass action. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**| Must be multipart/form-data.  | 
 **request** | [**PostMassUpdateType**](PostMassUpdateType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostMassUpdateResponseType**](POSTMassUpdateResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTMassUpdater**
> CommonResponseType PUTMassUpdater($bulkKey, $entityId, $entityName)

Stop mass action

This reference describes how to stop a mass action through the REST API. You can stop a mass action when its status is Pending or Processing. After you have stopped a mass action, you can get the mass action result to see details of the mass action.  - If you stop a mass action when its status is Pending, no response file is generated because no records have been processed.  - If you stop a mass action when its status is Processing, a response file is generated. You can check the response file to see which records have been processed and which have not. In the response file, the **Success** column has the value `Y` (successful) or `N` (failed) for processed records, and a blank value for unprocessed records.  Records that have already been processed when a mass action is stopped are not rolled back. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkKey** | **string**| String of 32 characters that identifies a mass action. You get the bulk-key after performing a mass action through the REST API.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**CommonResponseType**](CommonResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

