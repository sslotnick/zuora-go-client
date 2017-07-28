# \InvoiceItemAdjustmentsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObjectDELETEInvoiceItemAdjustment**](InvoiceItemAdjustmentsApi.md#ObjectDELETEInvoiceItemAdjustment) | **Delete** /v1/object/invoice-item-adjustment/{id} | CRUD: Delete InvoiceItemAdjustment
[**ObjectGETInvoiceItemAdjustment**](InvoiceItemAdjustmentsApi.md#ObjectGETInvoiceItemAdjustment) | **Get** /v1/object/invoice-item-adjustment/{id} | CRUD: Retrieve InvoiceItemAdjustment


# **ObjectDELETEInvoiceItemAdjustment**
> ProxyDeleteResponse ObjectDELETEInvoiceItemAdjustment($id, $entityId, $entityName)

CRUD: Delete InvoiceItemAdjustment




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

# **ObjectGETInvoiceItemAdjustment**
> ProxyGetInvoiceItemAdjustment ObjectGETInvoiceItemAdjustment($id, $entityId, $entityName, $fields)

CRUD: Retrieve InvoiceItemAdjustment




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **fields** | **string**| Object fields to return | [optional] 

### Return type

[**ProxyGetInvoiceItemAdjustment**](ProxyGetInvoiceItemAdjustment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

