# \BillingPreviewRunApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETBillingPreviewRun**](BillingPreviewRunApi.md#GETBillingPreviewRun) | **Get** /v1/billing-preview-runs/{billingPreviewRunId} | Get Billing Preview Run
[**POSTBillingPreviewRun**](BillingPreviewRunApi.md#POSTBillingPreviewRun) | **Post** /v1/billing-preview-runs | Create Billing Preview Run


# **GETBillingPreviewRun**
> GetBillingPreviewRunResponse GETBillingPreviewRun($billingPreviewRunId, $entityId, $entityName)

Get Billing Preview Run

**Note:** This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).    Retrieves a preview of future invoice items for multiple customer accounts through a billing preview run. If you have the Advanced AR Settlement feature enabled,  you can also retrieve a preview of future credit memo items. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   A billing preview run asynchronously generates a downloadable CSV file containing a preview of invoice item data and credit memo item data for a batch of customer accounts. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingPreviewRunId** | **string**| Id of the billing preview run.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetBillingPreviewRunResponse**](GetBillingPreviewRunResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTBillingPreviewRun**
> InlineResponse200 POSTBillingPreviewRun($request, $entityId, $entityName)

Create Billing Preview Run

**Note:** This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Creates billing preview runs for multiple customer accounts.  You can run up to 10 billing previews in batches concurrently. A single batch of customer accounts can only have one billing preview run at a time. So you can have up to 10 batches running at the same time. If you create a billing preview run for all customer batches, you cannot create another billing preview run until this preview run is completed. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostBillingPreviewRunParam**](PostBillingPreviewRunParam.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

