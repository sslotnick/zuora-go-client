# \AccountingCodesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAccountingCode**](AccountingCodesApi.md#DELETEAccountingCode) | **Delete** /v1/accounting-codes/{ac-id} | Delete accounting code
[**GETAccountingCode**](AccountingCodesApi.md#GETAccountingCode) | **Get** /v1/accounting-codes/{ac-id} | Query an accounting code
[**GETAllAccountingCodes**](AccountingCodesApi.md#GETAllAccountingCodes) | **Get** /v1/accounting-codes | Get all accounting codes
[**POSTAccountingCode**](AccountingCodesApi.md#POSTAccountingCode) | **Post** /v1/accounting-codes | Create accounting code
[**PUTAccountingCode**](AccountingCodesApi.md#PUTAccountingCode) | **Put** /v1/accounting-codes/{ac-id} | Update an accounting code
[**PUTActivateAccountingCode**](AccountingCodesApi.md#PUTActivateAccountingCode) | **Put** /v1/accounting-codes/{ac-id}/activate | Activate accounting code
[**PUTDeactivateAccountingCode**](AccountingCodesApi.md#PUTDeactivateAccountingCode) | **Put** /v1/accounting-codes/{ac-id}/deactivate | Deactivate accounting code


# **DELETEAccountingCode**
> CommonResponseType DELETEAccountingCode($acId, $entityId, $entityName)

Delete accounting code

This reference describes how to delete an accounting code through the REST API. ## Prerequisites If you have Zuora Finance enabled on your tenant, then you must have the Delete Unused Accounting Code permission. ## Limitations You can only delete accounting codes that have never been associated with any transactions. An accounting code must be deactivated before you can delete it. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acId** | **string**| ID of the accounting code you want to delete. | 
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

# **GETAccountingCode**
> GetAccountingCodeItemType GETAccountingCode($acId, $entityId, $entityName)

Query an accounting code

This reference describes how to query an accounting code through the REST API.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acId** | **string**| ID of the accounting code you want to query. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetAccountingCodeItemType**](GETAccountingCodeItemType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETAllAccountingCodes**
> GetAccountingCodesType GETAllAccountingCodes($entityId, $entityName, $pageSize)

Get all accounting codes

This reference describes how to query all accounting codes in your chart of accounts through the REST API.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 300]

### Return type

[**GetAccountingCodesType**](GETAccountingCodesType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTAccountingCode**
> PostAccountingCodeResponseType POSTAccountingCode($request, $entityId, $entityName)

Create accounting code

This reference describes how to create a new accounting code through the REST API.  The accounting code will be active as soon as it has been created.  ## Prerequisites   If you have Zuora Finance enabled on your tenant, you must have the  Configure Accounting Codes permission.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostAccountingCodeType**](PostAccountingCodeType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostAccountingCodeResponseType**](POSTAccountingCodeResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTAccountingCode**
> CommonResponseType PUTAccountingCode($acId, $request, $entityId, $entityName)

Update an accounting code

This reference describes how to update an existing accounting code through the REST API. ## Prerequisites   If you have Zuora Finance enabled on your tenant, you must have the  Manage Accounting Code permission.  ## Limitations You can only update accounting codes that are not already associated with any transactions. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acId** | **string**| ID of the accounting code you want to update. | 
 **request** | [**PutAccountingCodeType**](PutAccountingCodeType.md)|  | 
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

# **PUTActivateAccountingCode**
> CommonResponseType PUTActivateAccountingCode($acId, $entityId, $entityName)

Activate accounting code

This reference describes how to activate an accounting code through the REST API.  Prerequisites ------------- If you have Zuora Finance enabled on your tenant, you must have the Manage Accounting Code permission.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acId** | **string**| ID of the accounting code you want to activate. | 
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

# **PUTDeactivateAccountingCode**
> CommonResponseType PUTDeactivateAccountingCode($acId, $entityId, $entityName)

Deactivate accounting code

This reference describes how to deactivate an accounting code through the REST API.  ## Prerequisites If you have Zuora Finance enabled on your tenant, you must have the Manage Accounting Code permission. ## Limitations You can only deactivate accounting codes that are not associated with any transactions.  You cannot disable accounting codes of type AccountsReceivable. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acId** | **string**| ID of the accounting code you want to deactivate. | 
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

