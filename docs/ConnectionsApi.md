# \ConnectionsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTConnections**](ConnectionsApi.md#POSTConnections) | **Post** /v1/connections | Establish connection to Zuora REST API service


# **POSTConnections**
> CommonResponseType POSTConnections($apiAccessKeyId, $apiSecretAccessKey, $contentType, $entityId, $entityName)

Establish connection to Zuora REST API service

Establishes a connection to the Zuora REST API service based on a valid user credentials.  This call authenticates the user and returns an API session cookie that's used to authorize subsequent calls to the REST API. A call to `connections` is a required first step before using the Zuora REST API to access data.  The credentials must belong to a user account that has permission to access the API service.  As noted elsewhere, it's strongly recommended that an account used for Zuora API activity is never used to log into the Zuora UI.  Once an account is used to log into the UI, it may be subject to periodic forced password changes, which may eventually lead to authentication failures when using the API. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAccessKeyId** | **string**| Account username  | 
 **apiSecretAccessKey** | **string**| Account password  | 
 **contentType** | **string**| Must be set to \&quot;application/json\&quot;  | 
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

