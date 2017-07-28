# \OperationsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**POSTBillingPreview**](OperationsApi.md#POSTBillingPreview) | **Post** /v1/operations/billing-preview | Create billing preview
[**POSTTransactionInvoicePayment**](OperationsApi.md#POSTTransactionInvoicePayment) | **Post** /v1/operations/invoice-collect | Invoice and collect


# **POSTBillingPreview**
> BillingPreviewResult POSTBillingPreview($request, $entityId, $entityName)

Create billing preview

**Note:** The Billing Preview feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).    Generates a preview of future invoice items for one customer account. Use the BillingPreview call to calculate how much a single customer will be invoiced from the most recent invoice to a specific end of term date in the future.  Additionally, you can use the BillingPreview service to access real-time data on an individual customer's usage consumption.   The BillingPreview call does not calculate taxes for charges in the subscription.  If you have the Advanced AR Settlement feature enabled, you can also generate a preview of future credit memo items for one customer account. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostBillingPreviewParam**](PostBillingPreviewParam.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**BillingPreviewResult**](BillingPreviewResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTTransactionInvoicePayment**
> PostInvoiceCollectResponseType POSTTransactionInvoicePayment($request, $entityId, $entityName)

Invoice and collect

Generates invoices and collects payments for a specified account.  This method can generate invoices and collect payments on the invoices generated, or else simply collect payment on a specified existing invoice. The customer's default payment method is used, and the full amount due is collected. The operation depends on the parameters you specify  - To generate one or more new invoices for that customer and collect payment on the generated invoice(s), leave the **invoiceId** field empty.   - To collect payment on an existing invoice, specify the invoice ID.    The operation is atomic; if any part is unsuccessful, the entire operation is rolled back.   ## Notes  Timeouts may occur when using this method on an account that has an extremely high number of subscriptions. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostInvoiceCollectType**](PostInvoiceCollectType.md)| Customer account ID or account number. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostInvoiceCollectResponseType**](POSTInvoiceCollectResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

