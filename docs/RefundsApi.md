# \RefundsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETERefund**](RefundsApi.md#DELETERefund) | **Delete** /v1/refunds/{refundId} | Delete refund
[**GETRefund**](RefundsApi.md#GETRefund) | **Get** /v1/refunds/{refundId} | Get refund
[**GETRefundItemPart**](RefundsApi.md#GETRefundItemPart) | **Get** /v1/refunds/{refundId}/parts/{refundpartid}/itemparts/{itempartid} | Get refund part item
[**GETRefundItemParts**](RefundsApi.md#GETRefundItemParts) | **Get** /v1/refunds/{refundId}/parts/{refundpartid}/itemparts | Get refund part items
[**GETRefundPart**](RefundsApi.md#GETRefundPart) | **Get** /v1/refunds/{refundId}/parts/{refundpartid} | Get refund part
[**GETRefundParts**](RefundsApi.md#GETRefundParts) | **Get** /v1/refunds/{refundId}/parts | Get refund parts
[**GETRefunds**](RefundsApi.md#GETRefunds) | **Get** /v1/refunds | Get all refunds
[**ObjectDELETERefund**](RefundsApi.md#ObjectDELETERefund) | **Delete** /v1/object/refund/{id} | CRUD: Delete Refund
[**ObjectGETRefund**](RefundsApi.md#ObjectGETRefund) | **Get** /v1/object/refund/{id} | CRUD: Retrieve Refund
[**ObjectPOSTRefund**](RefundsApi.md#ObjectPOSTRefund) | **Post** /v1/object/refund | CRUD: Create Refund
[**ObjectPUTRefund**](RefundsApi.md#ObjectPUTRefund) | **Put** /v1/object/refund/{id} | CRUD: Update Refund
[**POSTCancelRefund**](RefundsApi.md#POSTCancelRefund) | **Post** /v1/refunds/{refundId}/cancel | Cancel refund
[**PUTUpdateRefund**](RefundsApi.md#PUTUpdateRefund) | **Put** /v1/refunds/{refundId} | Update refund


# **DELETERefund**
> CommonResponseType DELETERefund($refundId, $entityId, $entityName)

Delete refund

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Deletes a refund. You can delete a refund with the Canceled or Error status.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.  | 
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

# **GETRefund**
> GetRefundType GETRefund($refundId, $entityId, $entityName)

Get refund

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about a specific refund. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRefundType**](GETRefundType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRefundItemPart**
> GetRefundItemPartType GETRefundItemPart($itempartid, $refundpartid, $refundId, $entityId, $entityName)

Get refund part item

**Note:** The Invoice Item Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at  [Zuora Global Support](http://support.zuora.com/).    Retrieves the information about a specific refund part item. A refund part item is a single line item in a refund part. A refund part can consist of several different types of items. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itempartid** | **string**| The unique ID of a specific refund part item. You can get the refund part item ID from the response of [Get refund part items](https://www.zuora.com/developer/api-reference/#operation/GET_RefundItemParts).  | 
 **refundpartid** | **string**| The unique ID of a specific refund part. You can get the refund part ID from the response of [Get refund parts](https://www.zuora.com/developer/api-reference/#operation/GET_RefundParts).  | 
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRefundItemPartType**](GETRefundItemPartType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRefundItemParts**
> GetRefundItemPartCollectionType GETRefundItemParts($refundpartid, $refundId, $entityId, $entityName, $pageSize)

Get refund part items

**Note:** The Invoice Item Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at  [Zuora Global Support](http://support.zuora.com/).  Retrieves the information about all items of a refund part. A refund part item is a single line item in a refund part. A refund part can consist of several different types of items. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundpartid** | **string**| The unique ID of a specific refund part. You can get the refund part ID from the response of [Get refund parts](https://www.zuora.com/developer/api-reference/#operation/GET_RefundParts).  | 
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetRefundItemPartCollectionType**](GETRefundItemPartCollectionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRefundPart**
> RefundPartResponseType GETRefundPart($refundpartid, $refundId, $entityId, $entityName)

Get refund part

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about a specific refund part. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundpartid** | **string**| The unique ID of a specific refund part. You can get the refund part ID from the response of [Get refund parts](https://www.zuora.com/developer/api-reference/#operation/GET_RefundParts).  | 
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**RefundPartResponseType**](RefundPartResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRefundParts**
> GetRefundPartCollectionType GETRefundParts($refundId, $entityId, $entityName)

Get refund parts

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about all parts of a refund. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRefundPartCollectionType**](GETRefundPartCollectionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRefunds**
> GetRefundCollectionType GETRefunds($entityId, $entityName, $pageSize, $fieldsFilterable, $sort)

Get all refunds

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about all refunds. Two types of refunds are available, electronic refunds and external refunds. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]
 **fieldsFilterable** | **string**| This parameter restricts the data returned in the response. You can use this parameter to supply a dimension you want to filter on.   A single filter uses the following form:   *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*              Filters can be combined by using &#x60;&amp;&#x60;. For example: *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value* &#x60;&amp;&#x60; *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*  *fieldsFilterable* indicates the name of a supported field that you can use to filter the response data. The supported filterable fields of this operation are as below:    - number   - status   - type   - accountID   - amount   - refundDate   - paymentId   - createdDate   - createdById   - updatedDate   - updatedById   - methodType   *field_value* indicates a valid value of the filterable field. If the type of the field value is string, you can set the field to &#x60;null&#x60; value for filtering. Then, you can get the response data with this field value being &#39;null&#39;. For more information about these fields, see the field descriptions in the **Responses** section.     Examples:  - /refunds?status&#x3D;Processed  - /refunds?amount&#x3D;4&amp;status&#x3D;Processed  - /refunds?status&#x3D;Processed&amp;type&#x3D;External&amp;sort&#x3D;+number  | [optional] 
 **sort** | **string**| This parameter restricts the order of the data returned in the response. You can use this parameter to supply a dimension you want to sort on.  A sortable field uses the following form:   *operator* *field_name*  You can use at most two sortable fields in one URL path. Use a comma to separate sortable fields. For example:  *operator* *field_name*, *operator* *field_name*    *operator* is used to mark the order of sequencing. The operator is optional. If you only specify the sortable field without any operator, the response data is sorted in descending order by this field.    - The &#x60;-&#x60; operator indicates an ascending order.   - The &#x60;+&#x60; operator indicates a descending order.  By default, the response data is displayed in descending order by refund number.  *field_name* indicates the name of a sortable field. The supported sortable fields of this operation are as below:    - number   - accountID   - amount   - refundDate   - paymentId   - createdDate   - createdById   - updatedDate   - updatedById    Examples:  - /refunds?sort&#x3D;+number  - /refunds?status&#x3D;Processed&amp;sort&#x3D;-number,+amount  | [optional] 

### Return type

[**GetRefundCollectionType**](GETRefundCollectionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectDELETERefund**
> ProxyDeleteResponse ObjectDELETERefund($id, $entityId, $entityName)

CRUD: Delete Refund




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**ProxyDeleteResponse**](ProxyDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectGETRefund**
> ProxyGetRefund ObjectGETRefund($id, $entityId, $entityName, $fields)

CRUD: Retrieve Refund




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **fields** | **string**| Object fields to return | [optional] 

### Return type

[**ProxyGetRefund**](ProxyGetRefund.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectPOSTRefund**
> ProxyCreateOrModifyResponse ObjectPOSTRefund($createRequest, $entityId, $entityName)

CRUD: Create Refund




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**ProxyCreateRefund**](ProxyCreateRefund.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**ProxyCreateOrModifyResponse**](ProxyCreateOrModifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectPUTRefund**
> ProxyCreateOrModifyResponse ObjectPUTRefund($id, $modifyRequest, $entityId, $entityName)

CRUD: Update Refund




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **modifyRequest** | [**ProxyModifyRefund**](ProxyModifyRefund.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**ProxyCreateOrModifyResponse**](ProxyCreateOrModifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTCancelRefund**
> GetRefundType POSTCancelRefund($refundId, $entityId, $entityName)

Cancel refund

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Cancels a refund. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.        | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRefundType**](GETRefundType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTUpdateRefund**
> GetRefundType PUTUpdateRefund($body, $refundId, $entityId, $entityName)

Update refund

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Updates the basic and finance information about a refund. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PutRefundType**](PutRefundType.md)|  | 
 **refundId** | **string**| The unique ID of a refund. For example, 4028905f5a87c0ff015a889e590e00c9.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRefundType**](GETRefundType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

