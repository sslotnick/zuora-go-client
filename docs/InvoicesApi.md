# \InvoicesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObjectDELETEInvoice**](InvoicesApi.md#ObjectDELETEInvoice) | **Delete** /v1/object/invoice/{id} | CRUD: Delete Invoice
[**ObjectGETInvoice**](InvoicesApi.md#ObjectGETInvoice) | **Get** /v1/object/invoice/{id} | CRUD: Retrieve Invoice
[**ObjectPOSTInvoice**](InvoicesApi.md#ObjectPOSTInvoice) | **Post** /v1/object/invoice | CRUD: Create Invoice
[**ObjectPUTInvoice**](InvoicesApi.md#ObjectPUTInvoice) | **Put** /v1/object/invoice/{id} | CRUD: Update Invoice
[**POSTEmailInvoice**](InvoicesApi.md#POSTEmailInvoice) | **Post** /v1/invoices/{invoiceId}/emails | Email invoice
[**PUTReverseInvoice**](InvoicesApi.md#PUTReverseInvoice) | **Put** /v1/invoices/{invoiceId}/reverse | Reverse invoice


# **ObjectDELETEInvoice**
> ProxyDeleteResponse ObjectDELETEInvoice($id, $entityId, $entityName)

CRUD: Delete Invoice




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

# **ObjectGETInvoice**
> ProxyGetInvoice ObjectGETInvoice($id, $entityId, $entityName, $fields)

CRUD: Retrieve Invoice




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **fields** | **string**| Object fields to return | [optional] 

### Return type

[**ProxyGetInvoice**](ProxyGetInvoice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectPOSTInvoice**
> ProxyCreateOrModifyResponse ObjectPOSTInvoice($createRequest, $entityId, $entityName)

CRUD: Create Invoice




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**ProxyCreateInvoice**](ProxyCreateInvoice.md)|  | 
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

# **ObjectPUTInvoice**
> ProxyCreateOrModifyResponse ObjectPUTInvoice($id, $modifyRequest, $entityId, $entityName)

CRUD: Update Invoice




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **modifyRequest** | [**ProxyModifyInvoice**](ProxyModifyInvoice.md)|  | 
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

# **POSTEmailInvoice**
> CommonResponseType POSTEmailInvoice($request, $invoiceId, $entityId, $entityName)

Email invoice

Sends a posted invoice to the specified email addresses manually.    ## Notes   - You must activate the **Manual Email For Invoice | Manual Email For Invoice** notification before emailing invoices. To include the invoice PDF in the email, select the **Include Invoice PDF** check box in the **Edit notification** dialog from the Zuora UI. See [Create and Edit Notifications](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/Notifications/C_Create_Notifications#section_2) for more information.     - Zuora sends the email messages based on the email template you set. You can set the email template to use in the **Delivery Options** panel of the **Edit notification** dialog from the Zuora UI. By default, the **Invoice Posted Default Email Template** template is used. See [Create and Edit Email Templates](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/Notifications/Create_Email_Templates) for more information.     - The invoices are sent only to the work email addresses or personal email addresses of the Bill To contact if the following conditions are all met:      * The `useEmailTemplateSetting` field is set to `false`.     * The email addresses are not specified in the `emailAddresses` field. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostInvoiceEmailRequestType**](PostInvoiceEmailRequestType.md)|  | 
 **invoiceId** | **string**| The ID of the invoice. For example, 2c92c8955bd63cc1015bd7c151af02ab.  | 
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

# **PUTReverseInvoice**
> PutReverseInvoiceResponseType PUTReverseInvoice($invoiceId, $request, $entityId, $entityName)

Reverse invoice

Reverses posted invoices.    **Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  **Restrictions**  You are not allowed to reverse an invoice if one of the following restrictions is met:  * Payments and credit memos are applied to the invoice. * The invoice is split. * The invoice is not in Posted status. * The total amount of the invoice is less than 0 (a negative invoice). * Using Tax Connector for Extension Platform to calculate taxes.  See [Reverse Posted Invoices](https://knowledgecenter.zuora.com/CB_Billing/IA_Invoices/Reverse_Posted_Invoices) for more information. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **string**| The ID of the invoice. For example, 2c92c8955bd63cc1015bd7c151af02ab.  | 
 **request** | [**PutReverseInvoiceType**](PutReverseInvoiceType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutReverseInvoiceResponseType**](PutReverseInvoiceResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

