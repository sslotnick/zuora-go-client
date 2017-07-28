# \TransactionsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETTransactionInvoice**](TransactionsApi.md#GETTransactionInvoice) | **Get** /v1/transactions/invoices/accounts/{account-key} | Get invoices
[**GETTransactionPayment**](TransactionsApi.md#GETTransactionPayment) | **Get** /v1/transactions/payments/accounts/{account-key} | Get payments


# **GETTransactionInvoice**
> GetInvoiceFileWrapper GETTransactionInvoice($accountKey, $entityId, $entityName, $pageSize)

Get invoices

Retrieves invoices for a specified account.  Invoices are returned in reverse chronological order by **updatedDate**. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKey** | **string**| Account number or account ID.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetInvoiceFileWrapper**](GETInvoiceFileWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETTransactionPayment**
> GetPaymentsType GETTransactionPayment($accountKey, $entityId, $entityName, $pageSize)

Get payments

Retrieves payments for a specified account. Payments are returned in reverse chronological order by **updatedDate**. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKey** | **string**| Account number or account ID. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetPaymentsType**](GETPaymentsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

