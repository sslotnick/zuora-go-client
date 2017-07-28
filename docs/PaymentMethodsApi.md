# \PaymentMethodsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPaymentMethods**](PaymentMethodsApi.md#DELETEPaymentMethods) | **Delete** /v1/payment-methods/{payment-method-id} | Delete payment method
[**GETPaymentMethods**](PaymentMethodsApi.md#GETPaymentMethods) | **Get** /v1/payment-methods/credit-cards/accounts/{account-key} | Get payment methods
[**ObjectDELETEPaymentMethod**](PaymentMethodsApi.md#ObjectDELETEPaymentMethod) | **Delete** /v1/object/payment-method/{id} | CRUD: Delete PaymentMethod
[**ObjectGETPaymentMethod**](PaymentMethodsApi.md#ObjectGETPaymentMethod) | **Get** /v1/object/payment-method/{id} | CRUD: Retrieve PaymentMethod
[**ObjectPOSTPaymentMethod**](PaymentMethodsApi.md#ObjectPOSTPaymentMethod) | **Post** /v1/object/payment-method | CRUD: Create PaymentMethod
[**ObjectPUTPaymentMethod**](PaymentMethodsApi.md#ObjectPUTPaymentMethod) | **Put** /v1/object/payment-method/{id} | CRUD: Update PaymentMethod
[**POSTPaymentMethods**](PaymentMethodsApi.md#POSTPaymentMethods) | **Post** /v1/payment-methods/credit-cards | Create payment method
[**POSTPaymentMethodsDecryption**](PaymentMethodsApi.md#POSTPaymentMethodsDecryption) | **Post** /v1/payment-methods/decryption | Create payment method decryption
[**PUTPaymentMethods**](PaymentMethodsApi.md#PUTPaymentMethods) | **Put** /v1/payment-methods/credit-cards/{payment-method-id} | Update payment method


# **DELETEPaymentMethods**
> CommonResponseType DELETEPaymentMethods($paymentMethodId, $entityId, $entityName)

Delete payment method

Deletes a credit card payment method from the specified customer account.  If the specified payment method is the account's default payment method, the request will fail.  In that case, you must first designate a different payment method for that customer to be the default. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodId** | **string**| Unique identifier of a payment method. (Since this ID is unique, and linked to a customer account in the system, no customer identifier is needed.) | 
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

# **GETPaymentMethods**
> GetPaymentMethodsType GETPaymentMethods($accountKey, $entityId, $entityName, $pageSize)

Get payment methods

This REST API reference describes how to retrieve all credit card information for the specified customer account.   ## Notes The response includes details credit or debit cards for the specified customer account. Card numbers are masked, e.g., \"************1234\". Cards are returned in reverse chronological order of last update.  You can send requests for bank transfer payment methods types. The response will not include bank transfer details.  The response only includes payment details on payment methods that are credit or debit cards. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKey** | **string**| Account number or account ID. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetPaymentMethodsType**](GETPaymentMethodsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectDELETEPaymentMethod**
> ProxyDeleteResponse ObjectDELETEPaymentMethod($id, $entityId, $entityName)

CRUD: Delete PaymentMethod




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

# **ObjectGETPaymentMethod**
> ProxyGetPaymentMethod ObjectGETPaymentMethod($id, $entityId, $entityName, $fields)

CRUD: Retrieve PaymentMethod




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **fields** | **string**| Object fields to return | [optional] 

### Return type

[**ProxyGetPaymentMethod**](ProxyGetPaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectPOSTPaymentMethod**
> ProxyCreateOrModifyResponse ObjectPOSTPaymentMethod($createRequest, $entityId, $entityName)

CRUD: Create PaymentMethod




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**ProxyCreatePaymentMethod**](ProxyCreatePaymentMethod.md)|  | 
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

# **ObjectPUTPaymentMethod**
> ProxyCreateOrModifyResponse ObjectPUTPaymentMethod($id, $modifyRequest, $entityId, $entityName)

CRUD: Update PaymentMethod




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **modifyRequest** | [**ProxyModifyPaymentMethod**](ProxyModifyPaymentMethod.md)|  | 
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

# **POSTPaymentMethods**
> PostPaymentMethodResponseType POSTPaymentMethods($request, $entityId, $entityName)

Create payment method

This REST API reference describes how to create a new credit card payment method for a customer account.  ## Notes  This API call is CORS Enabled. Use client-side JavaScript to invoke the call.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostPaymentMethodType**](PostPaymentMethodType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostPaymentMethodResponseType**](POSTPaymentMethodResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTPaymentMethodsDecryption**
> PostPaymentMethodResponseDecryption POSTPaymentMethodsDecryption($request, $entityId, $entityName)

Create payment method decryption

The decryption API endpoint can conditionally perform 3 tasks in one atomic call:   * Decrypt Apple Pay Payment token   * Create Credit Card Payment Method in Zuora with decrypted Apple Pay information   * Process Payment on a specified Invoice (optional) 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostPaymentMethodDecryption**](PostPaymentMethodDecryption.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostPaymentMethodResponseDecryption**](POSTPaymentMethodResponseDecryption.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTPaymentMethods**
> PutPaymentMethodResponseType PUTPaymentMethods($paymentMethodId, $request, $entityId, $entityName)

Update payment method

Updates an existing credit card payment method for the specified customer account. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodId** | **string**| Unique ID of the payment method to update. | 
 **request** | [**PutPaymentMethodType**](PutPaymentMethodType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutPaymentMethodResponseType**](PUTPaymentMethodResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

