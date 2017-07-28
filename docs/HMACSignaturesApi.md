# \HMACSignaturesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTHMACSignatures**](HMACSignaturesApi.md#POSTHMACSignatures) | **Post** /v1/hmac-signatures | Return HMAC signatures


# **POSTHMACSignatures**
> PosthmacSignatureResponseType POSTHMACSignatures($request, $entityId, $entityName)

Return HMAC signatures

This REST API reference describes how to return unique signature and token values that used to process a CORS enabled API call. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PosthmacSignatureType**](PosthmacSignatureType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PosthmacSignatureResponseType**](POSTHMACSignatureResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

