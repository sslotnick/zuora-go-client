# \EntityConnectionsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETEntityConnections**](EntityConnectionsApi.md#GETEntityConnections) | **Get** /v1/entity-connections | Multi-entity: Get connections
[**POSTEntityConnections**](EntityConnectionsApi.md#POSTEntityConnections) | **Post** /v1/entity-connections | Multi-entity: Initiate connection
[**PUTEntityConnectionsAccept**](EntityConnectionsApi.md#PUTEntityConnectionsAccept) | **Put** /v1/entity-connections/{connection-id}/accept | Multi-entity: Accept connection
[**PUTEntityConnectionsDeny**](EntityConnectionsApi.md#PUTEntityConnectionsDeny) | **Put** /v1/entity-connections/{connection-id}/deny | Multi-entity: Deny connection
[**PUTEntityConnectionsDisconnect**](EntityConnectionsApi.md#PUTEntityConnectionsDisconnect) | **Put** /v1/entity-connections/{connection-id}/disconnect | Multi-entity: Disconnect connection


# **GETEntityConnections**
> GetEntityConnectionsResponseType GETEntityConnections($entityId, $entityName, $pageSize, $type_)

Multi-entity: Get connections

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Retrieves information about certain connections for a specified entity. You can specify the entity to retrieve in the `entityId` or `entityName` request header.  You can retrieve:  - Inbound connections  - Outbound connections  - Both inbound and outbound connections  ## User Access Permission You can make the call as any entity user.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 30]
 **type_** | **string**| Specifies whether to retrieve inbound or outbound connections for an entity.  Possible values:  - &#x60;inbound&#x60;: All the incoming connections to the entity.  - &#x60;outbound&#x60;: All the outgoing connections from the entity.  If you do not specify this field in the request, both the inbound and outbound connections are returned.  | [optional] 

### Return type

[**GetEntityConnectionsResponseType**](GETEntityConnectionsResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTEntityConnections**
> PostEntityConnectionsResponseType POSTEntityConnections($entityId, $entityName, $request)

Multi-entity: Initiate connection

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Initiates a connection request from a source entity to a target entity.  ## User Access Permission You must make the call as a source entity administrator. Also, this administrator must have permission to access to the target entity. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **request** | [**PostEntityConnectionsType**](PostEntityConnectionsType.md)|  | [optional] 

### Return type

[**PostEntityConnectionsResponseType**](POSTEntityConnectionsResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTEntityConnectionsAccept**
> PutEntityConnectionsAcceptResponseType PUTEntityConnectionsAccept($connectionId, $entityId, $entityName)

Multi-entity: Accept connection

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Accepts a connection request.  ## User Access Permission You must make the call as an entity administrator to accept a connection request. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionId** | **string**| The ID of the connection that you want to accept.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutEntityConnectionsAcceptResponseType**](PUTEntityConnectionsAcceptResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTEntityConnectionsDeny**
> PutEntityConnectionsDenyResponseType PUTEntityConnectionsDeny($connectionId, $entityId, $entityName)

Multi-entity: Deny connection

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Denies a connection request.  ## User Access Permission You must make the call as an entity administrator to deny a connection request. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionId** | **string**| The ID of the connection that you want to deny.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutEntityConnectionsDenyResponseType**](PUTEntityConnectionsDenyResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTEntityConnectionsDisconnect**
> PutEntityConnectionsDisconnectResponseType PUTEntityConnectionsDisconnect($connectionId, $entityId, $entityName)

Multi-entity: Disconnect connection

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Disconnects an established connection request.  ## User Access Permission You must make the call as an administrator of the target entity or source entity. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionId** | **string**| The ID of the connection that you want to disconnect.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutEntityConnectionsDisconnectResponseType**](PUTEntityConnectionsDisconnectResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

