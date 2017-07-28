# \BillingDocumentsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTGenerateBillingDocuments**](BillingDocumentsApi.md#POSTGenerateBillingDocuments) | **Post** /v1/accounts/{id}/billing-documents/generate | Generate billing documents by account


# **POSTGenerateBillingDocuments**
> GenerateBillingDocumentResponseType POSTGenerateBillingDocuments($body, $id, $entityId, $entityName)

Generate billing documents by account

Generates draft or posted billing documents for a specified account. You can also generate billing documents for specified subscriptions of a specified account. The billing documents contain invoices and credit memos. To generate credit memos, you must have the Advanced AR Settlement feature enabled. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PostGenerateBillingDocumentType**](PostGenerateBillingDocumentType.md)|  | 
 **id** | **string**| The ID of the customer account that billing documents are generated for. For example, 8a8082e65b27f6c3015ba3e326b26419.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GenerateBillingDocumentResponseType**](GenerateBillingDocumentResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

