# \UsersApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETEntitiesUserAccessible**](UsersApi.md#GETEntitiesUserAccessible) | **Get** /v1/users/{username}/accessible-entities | Multi-entity: Get entities that a user can access
[**PUTAcceptUserAccess**](UsersApi.md#PUTAcceptUserAccess) | **Put** /v1/users/{username}/accept-access | Multi-entity: Accept user access
[**PUTDenyUserAccess**](UsersApi.md#PUTDenyUserAccess) | **Put** /v1/users/{username}/deny-access | Multi-entity: Deny user access
[**PUTSendUserAccessRequests**](UsersApi.md#PUTSendUserAccessRequests) | **Put** /v1/users/{username}/request-access | Multi-entity: Send user access requests


# **GETEntitiesUserAccessible**
> GetEntitiesUserAccessibleResponseType GETEntitiesUserAccessible($username, $entityId, $entityName)

Multi-entity: Get entities that a user can access

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves detailed information about all the entities that a user has permission to access.  ## User Access Permission You can make the call as any entity user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Specify the login user name that you want to retrieve.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetEntitiesUserAccessibleResponseType**](GETEntitiesUserAccessibleResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTAcceptUserAccess**
> PutAcceptUserAccessResponseType PUTAcceptUserAccess($username, $entityId, $entityName)

Multi-entity: Accept user access

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Accepts user access to an entity.  ## User Access Permission You must make the calls as an administrator of the entity that you want to accept the user access to.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Specify the login name of the user that you want to accept the access request for.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutAcceptUserAccessResponseType**](PUTAcceptUserAccessResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTDenyUserAccess**
> PutDenyUserAccessResponseType PUTDenyUserAccess($username, $entityId, $entityName)

Multi-entity: Deny user access

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Denies a user access to an entity.   ## User Access Permission You must make the calls as an administrator of the entity that you want to deny the user access to. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Specify the login name of the user that you want to deny the access.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutDenyUserAccessResponseType**](PUTDenyUserAccessResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTSendUserAccessRequests**
> PutSendUserAccessRequestResponseType PUTSendUserAccessRequests($username, $request, $entityId, $entityName)

Multi-entity: Send user access requests

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Sends access requests to the entities that a user wants to access.  ## User Access Permission You must make the call as an administrator of the entity, in which the request user is created. Also, this administrator must have the permission to access the entities that the request user wants to access. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string**| Specify the login name of the user who wants to access other entities.  | 
 **request** | [**PutSendUserAccessRequestType**](PutSendUserAccessRequestType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutSendUserAccessRequestResponseType**](PUTSendUserAccessRequestResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

