# \EntitiesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEEntities**](EntitiesApi.md#DELETEEntities) | **Delete** /v1/entities/{id} | Multi-entity: Delete entity
[**GETEntities**](EntitiesApi.md#GETEntities) | **Get** /v1/entities | Multi-entity: Get entities
[**GETEntityById**](EntitiesApi.md#GETEntityById) | **Get** /v1/entities/{id} | Multi-entity: Get entity by Id
[**POSTEntities**](EntitiesApi.md#POSTEntities) | **Post** /v1/entities | Multi-entity: Create entity
[**PUTEntities**](EntitiesApi.md#PUTEntities) | **Put** /v1/entities/{id} | Multi-entity: Update entity
[**PUTProvisionEntity**](EntitiesApi.md#PUTProvisionEntity) | **Put** /v1/entities/{id}/provision | Multi-entity: Provision entity


# **DELETEEntities**
> DeletEntityResponseType DELETEEntities($id, $entityId, $entityName)

Multi-entity: Delete entity

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Removes an entity and its sub-entities from a multi-entity hierarchy. You can only remove unprovisioned entities. An error occurred when you remove a provisioned entity.  ## User Access Permission You must make the call as a global entity administrator.    


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Specify the Id of the entity that you want to delete. You can get the entity Id from the GET Entities call. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**DeletEntityResponseType**](DELETEntityResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETEntities**
> GetEntitiesResponseType GETEntities($entityId, $entityName, $provisioned)

Multi-entity: Get entities

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).    Retrieves detailed information of certain entities in a multi-entity hierarchy.   You can retrieve:   - Provisioned entities     - Unprovisioned entities     - Both provisioned and unprovisioned entities  ## User Access Permission  You can make the call as any entity user. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **provisioned** | **string**| Specify whether to retrieve provisioned or unprovisioned entities:  - &#x60;true&#x60;: Provisioned entities  - &#x60;false&#x60;: Unprovisioned entities   If you do not specify this field in the request, both the provisioned and unprovisioned entities are returned.  | [optional] 

### Return type

[**GetEntitiesResponseType**](GETEntitiesResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETEntityById**
> GetEntitiesResponseTypeWithId GETEntityById($id, $entityId, $entityName)

Multi-entity: Get entity by Id

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves detailed information about a specified entity.  ## User Access Permission You can make the call as any entity user.      


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Specify the Id of the entity that you want to retrieve. You can get the entity Id from the GET Entities call. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetEntitiesResponseTypeWithId**](GETEntitiesResponseTypeWithId.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTEntities**
> CreateEntityResponseType POSTEntities($request, $entityId, $entityName)

Multi-entity: Create entity

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Creates an entity in a multi-entity hierarchy.  ## User Access Permission You must make the call as a global entity administrator.  ## Notes * We recommend that you assign only one administrator to manage the entity hierarchy, because an administrator of the global entity by default can only access to the entities that are created by themselves. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateEntityType**](CreateEntityType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**CreateEntityResponseType**](CreateEntityResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTEntities**
> UpdateEntityResponseType PUTEntities($id, $request, $entityId, $entityName)

Multi-entity: Update entity

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Edits the following information about an unprovisioned entity:   - Name    - Display name    - Locale    - Timezone  ## User Access Permission You must make the call as a global entity administrator.  ## Notes * You are not allowed to edit the locale and time zone of the provisioned entities through the REST API. * You are not allowed to edit the display name of the global entity. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The Id of the entity that you want to edit. You can get the entity Id from the GET Entities call. | 
 **request** | [**UpdateEntityType**](UpdateEntityType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**UpdateEntityResponseType**](UpdateEntityResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTProvisionEntity**
> ProvisionEntityResponseType PUTProvisionEntity($id, $entityId, $entityName)

Multi-entity: Provision entity

**Note:** The Multi-entity feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Provisions an entity. You can only provision an entity if its parent entity is provisioned.  ## User Access Permission You must make the call as a global entity administrator.   ## Notes * Zuora does not allow you to remove a provisioned entity from the multi-entity hierarchy. So before you provision an entity, make sure that you put the entity in the correct place in the multi-entity hierarchy.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Specify the Id of the entity that you want to provision.  You can get the entity Id from the GET Entities call. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**ProvisionEntityResponseType**](ProvisionEntityResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

