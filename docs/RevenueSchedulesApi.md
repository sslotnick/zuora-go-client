# \RevenueSchedulesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETERS**](RevenueSchedulesApi.md#DELETERS) | **Delete** /v1/revenue-schedules/{rs-number} | Delete revenue schedule
[**GETRS**](RevenueSchedulesApi.md#GETRS) | **Get** /v1/revenue-schedules/{rs-number} | Get revenue schedule details
[**GETRSbyCreditMemoItem**](RevenueSchedulesApi.md#GETRSbyCreditMemoItem) | **Get** /v1/revenue-schedules/credit-memo-items/{cmi-id} | Get revenue schedule by credit memo item ID 
[**GETRSbyDebitMemoItem**](RevenueSchedulesApi.md#GETRSbyDebitMemoItem) | **Get** /v1/revenue-schedules/debit-memo-items/{dmi-id} | Get revenue schedule by debit memo item ID 
[**GETRSbyInvoiceItem**](RevenueSchedulesApi.md#GETRSbyInvoiceItem) | **Get** /v1/revenue-schedules/invoice-items/{invoice-item-id} | Get revenue schedule by invoice item ID
[**GETRSbyInvoiceItemAdjustment**](RevenueSchedulesApi.md#GETRSbyInvoiceItemAdjustment) | **Get** /v1/revenue-schedules/invoice-item-adjustments/{invoice-item-adj-id}/ | Get revenue schedule by invoice item adjustment
[**GETRSbyProductChargeAndBillingAccount**](RevenueSchedulesApi.md#GETRSbyProductChargeAndBillingAccount) | **Get** /v1/revenue-schedules/product-charges/{charge-key}/{account-key} | Get all revenue schedules of product charge by charge ID and billing account ID 
[**GETRSforSubscCharge**](RevenueSchedulesApi.md#GETRSforSubscCharge) | **Get** /v1/revenue-schedules/subscription-charges/{charge-key} | Get revenue schedule by subscription charge
[**POSTRSforCrditMemoItemManualDistribution**](RevenueSchedulesApi.md#POSTRSforCrditMemoItemManualDistribution) | **Post** /v1/revenue-schedules/credit-memo-items/{cmi-id} | Create revenue schedule for credit memo item (manual distribution) 
[**POSTRSforCreditMemoItemDistributeByDateRange**](RevenueSchedulesApi.md#POSTRSforCreditMemoItemDistributeByDateRange) | **Post** /v1/revenue-schedules/credit-memo-items/{cmi-id}/distribute-revenue-with-date-range | Create revenue schedule for credit memo item (distribute by date range) 
[**POSTRSforDebitMemoItemDistributeByDateRange**](RevenueSchedulesApi.md#POSTRSforDebitMemoItemDistributeByDateRange) | **Post** /v1/revenue-schedules/debit-memo-items/{dmi-id}/distribute-revenue-with-date-range | Create revenue schedule for debit memo item (distribute by date range) 
[**POSTRSforDebitMemoItemManualDistribution**](RevenueSchedulesApi.md#POSTRSforDebitMemoItemManualDistribution) | **Post** /v1/revenue-schedules/debit-memo-items/{dmi-id} | Create revenue schedule for debit memo item (manual distribution) 
[**POSTRSforInvoiceItemAdjustmentDistributeByDateRange**](RevenueSchedulesApi.md#POSTRSforInvoiceItemAdjustmentDistributeByDateRange) | **Post** /v1/revenue-schedules/invoice-item-adjustments/{invoice-item-adj-key}/distribute-revenue-with-date-range | Create revenue schedule for Invoice Item Adjustment (distribute by date range)
[**POSTRSforInvoiceItemAdjustmentManualDistribution**](RevenueSchedulesApi.md#POSTRSforInvoiceItemAdjustmentManualDistribution) | **Post** /v1/revenue-schedules/invoice-item-adjustments/{invoice-item-adj-key} | Create revenue schedule for Invoice Item Adjustment (manual distribution)
[**POSTRSforInvoiceItemDistributeByDateRange**](RevenueSchedulesApi.md#POSTRSforInvoiceItemDistributeByDateRange) | **Post** /v1/revenue-schedules/invoice-items/{invoice-item-id}/distribute-revenue-with-date-range | Create revenue schedule for Invoice Item (distribute by date range)
[**POSTRSforInvoiceItemManualDistribution**](RevenueSchedulesApi.md#POSTRSforInvoiceItemManualDistribution) | **Post** /v1/revenue-schedules/invoice-items/{invoice-item-id} | Create revenue schedule for Invoice Item (manual distribution)
[**POSTRSforSubsCharge**](RevenueSchedulesApi.md#POSTRSforSubsCharge) | **Post** /v1/revenue-schedules/subscription-charges/{charge-key} | Create revenue schedule on subscription charge
[**PUTRSBasicInfo**](RevenueSchedulesApi.md#PUTRSBasicInfo) | **Put** /v1/revenue-schedules/{rs-number}/basic-information | Update revenue schedule basic information
[**PUTRevenueAcrossAP**](RevenueSchedulesApi.md#PUTRevenueAcrossAP) | **Put** /v1/revenue-schedules/{rs-number}/distribute-revenue-across-accounting-periods | Distribute revenue across accounting periods
[**PUTRevenueByRecognitionStartandEndDates**](RevenueSchedulesApi.md#PUTRevenueByRecognitionStartandEndDates) | **Put** /v1/revenue-schedules/{rs-number}/distribute-revenue-with-date-range | Distribute revenue by recognition start and end dates
[**PUTRevenueSpecificDate**](RevenueSchedulesApi.md#PUTRevenueSpecificDate) | **Put** /v1/revenue-schedules/{rs-number}/distribute-revenue-on-specific-date | Distribute revenue on specific date


# **DELETERS**
> CommonResponseType DELETERS($rsNumber, $entityId, $entityName)

Delete revenue schedule

Deletes a revenue schedule by specifying its revenue schedule number ## Prerequisites You must have the Delete Custom Revenue Schedule permissions in Zuora Finance. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**|  Revenue schedule number of the revenue schedule you want to delete, for example, RS-00000256. To be deleted, the revenue schedule: * Must be using a custom unlimited recognition rule. * Cannot have any revenue in a closed accounting period. * Cannot be included in a summary journal entry. * Cannot have a revenue schedule date in a closed accounting period.  | 
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

# **GETRS**
> GetrsDetailType GETRS($rsNumber, $entityId, $entityName)

Get revenue schedule details

Retrieves the details of a revenue schedule by specifying the revenue schedule number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetrsDetailType**](GETRSDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRSbyCreditMemoItem**
> GetrsDetailType GETRSbyCreditMemoItem($cmiId)

Get revenue schedule by credit memo item ID 

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Retrieves the details about a revenue schedule by specifying a valid credit memo item ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmiId** | **string**| The unique ID of a credit memo item. You can get the credit memo item ID from the response of [Get credit memo items](https://www.zuora.com/developer/api-reference/#operation/GET_CreditMemoItems).  | 

### Return type

[**GetrsDetailType**](GETRSDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRSbyDebitMemoItem**
> GetrsDetailType GETRSbyDebitMemoItem($dmiId)

Get revenue schedule by debit memo item ID 

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Retrieves the details about a revenue schedule by specifying a valid debit memo item ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dmiId** | **string**| The unique ID of a debit memo item. You can get the debit memo item ID from the response of [Get debit memo items](https://www.zuora.com/developer/api-reference/#operation/GET_DebitMemoItems).  | 

### Return type

[**GetrsDetailType**](GETRSDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRSbyInvoiceItem**
> GetrsDetailType GETRSbyInvoiceItem($invoiceItemId, $entityId, $entityName)

Get revenue schedule by invoice item ID

Retrieves the details of a revenue schedule by specifying the invoice item ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceItemId** | **string**| A valid Invoice Item ID. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetrsDetailType**](GETRSDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRSbyInvoiceItemAdjustment**
> GetrsDetailType GETRSbyInvoiceItemAdjustment($invoiceItemAdjId, $entityId, $entityName)

Get revenue schedule by invoice item adjustment

Retrieves the details of a revenue schedule by specifying a valid invoice item adjustment identifier. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceItemAdjId** | **string**| ID or number of the Invoice Item Adjustment, for example, e20b07fd416dcfcf0141c81164fd0a72. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetrsDetailType**](GETRSDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRSbyProductChargeAndBillingAccount**
> GetrsDetailsByProductChargeType GETRSbyProductChargeAndBillingAccount($accountKey, $chargeKey, $pageSize)

Get all revenue schedules of product charge by charge ID and billing account ID 

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Retrieves the details about all revenue schedules of a product rate plan charge by specifying the charge ID and billing account ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKey** | **string**| The account number or account ID.  | 
 **chargeKey** | **string**| The unique ID of a product rate plan charge. For example, 8a8082e65ba86084015bb323d3c61d82.  | 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 8]

### Return type

[**GetrsDetailsByProductChargeType**](GETRSDetailsByProductChargeType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRSforSubscCharge**
> GetrsDetailsByChargeType GETRSforSubscCharge($chargeKey, $entityId, $entityName, $pageSize)

Get revenue schedule by subscription charge

Retrieves the revenue schedule details by specifying subscription charge ID. Request and response field descriptions and sample code are provided 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeKey** | **string**| ID of the subscription rate plan charge; for example, 402892793e173340013e173b81000012. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 8]

### Return type

[**GetrsDetailsByChargeType**](GETRSDetailsByChargeType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforCrditMemoItemManualDistribution**
> PostRevenueScheduleByTransactionResponseType POSTRSforCrditMemoItemManualDistribution($cmiId, $body)

Create revenue schedule for credit memo item (manual distribution) 

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Creates a revenue schedule for a credit memo item, and manually distribute the revenue. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmiId** | **string**| The unique ID of a credit memo item. You can get the credit memo item ID from the response of [Get credit memo items](https://www.zuora.com/developer/api-reference/#operation/GET_CreditMemoItems).  | 
 **body** | [**PostRevenueScheduleByTransactionType**](PostRevenueScheduleByTransactionType.md)|  | 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforCreditMemoItemDistributeByDateRange**
> PostRevenueScheduleByTransactionResponseType POSTRSforCreditMemoItemDistributeByDateRange($cmiId, $body)

Create revenue schedule for credit memo item (distribute by date range) 

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Creates a revenue schedule for a credit memo item, and automatically distribute the revenue by specifying the recognition start and end dates. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmiId** | **string**| The unique ID of a credit memo item. You can get the credit memo item ID from the response of [Get credit memo items](https://www.zuora.com/developer/api-reference/#operation/GET_CreditMemoItems).  | 
 **body** | [**PostRevenueScheduleByTransactionRatablyType**](PostRevenueScheduleByTransactionRatablyType.md)|  | 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforDebitMemoItemDistributeByDateRange**
> PostRevenueScheduleByTransactionResponseType POSTRSforDebitMemoItemDistributeByDateRange($dmiId, $body)

Create revenue schedule for debit memo item (distribute by date range) 

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Creates a revenue schedule for a debit memo item, and automatically distribute the revenue by specifying the recognition start and end dates. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dmiId** | **string**| The unique ID of a debit memo item. You can get the debit memo item ID from the response of [Get debit memo items](https://www.zuora.com/developer/api-reference/#operation/GET_DebitMemoItems).  | 
 **body** | [**PostRevenueScheduleByTransactionRatablyType**](PostRevenueScheduleByTransactionRatablyType.md)|  | 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforDebitMemoItemManualDistribution**
> PostRevenueScheduleByTransactionResponseType POSTRSforDebitMemoItemManualDistribution($dmiId, $body)

Create revenue schedule for debit memo item (manual distribution) 

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Creates a revenue schedule for a debit memo item, and manually distribute the revenue. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dmiId** | **string**| The unique ID of a debit memo item. You can get the debit memo item ID from the response of [Get debit memo items](https://www.zuora.com/developer/api-reference/#operation/GET_DebitMemoItems).  | 
 **body** | [**PostRevenueScheduleByTransactionType**](PostRevenueScheduleByTransactionType.md)|  | 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforInvoiceItemAdjustmentDistributeByDateRange**
> PostRevenueScheduleByTransactionResponseType POSTRSforInvoiceItemAdjustmentDistributeByDateRange($invoiceItemAdjKey, $request, $entityId, $entityName)

Create revenue schedule for Invoice Item Adjustment (distribute by date range)

Creates a revenue schedule for an Invoice Item Adjustment and distribute the revenue by specifying the recognition start and end dates. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceItemAdjKey** | **string**| ID or number of the Invoice Item Adjustment, for example, e20b07fd416dcfcf0141c81164fd0a72. If the specified Invoice Item Adjustment is already associated with a revenue schedule, the call will fail.  | 
 **request** | [**PostRevenueScheduleByDateRangeType**](PostRevenueScheduleByDateRangeType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforInvoiceItemAdjustmentManualDistribution**
> PostRevenueScheduleByTransactionResponseType POSTRSforInvoiceItemAdjustmentManualDistribution($invoiceItemAdjKey, $request, $entityId, $entityName)

Create revenue schedule for Invoice Item Adjustment (manual distribution)

Creates a revenue schedule for an Invoice Item Adjustment and manually distribute the revenue. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceItemAdjKey** | **string**| ID or number of the Invoice Item Adjustment, for example, e20b07fd416dcfcf0141c81164fd0a72. If the specified Invoice Item Adjustment is already associated with a revenue schedule, the call will fail.  | 
 **request** | [**PostRevenueScheduleByTransactionType**](PostRevenueScheduleByTransactionType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforInvoiceItemDistributeByDateRange**
> PostRevenueScheduleByTransactionResponseType POSTRSforInvoiceItemDistributeByDateRange($invoiceItemId, $request, $entityId, $entityName)

Create revenue schedule for Invoice Item (distribute by date range)

Creates a revenue schedule for an Invoice Item and distribute the revenue by specifying the recognition start and end dates. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceItemId** | **string**| ID of the Invoice Item, for example, e20b07fd416dcfcf0141c81164fd0a75. If the specified Invoice Item is already associated with a revenue schedule, the call will fail.  | 
 **request** | [**PostRevenueScheduleByDateRangeType**](PostRevenueScheduleByDateRangeType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforInvoiceItemManualDistribution**
> PostRevenueScheduleByTransactionResponseType POSTRSforInvoiceItemManualDistribution($invoiceItemId, $request, $entityId, $entityName)

Create revenue schedule for Invoice Item (manual distribution)

Creates a revenue schedule for an Invoice Item and manually distribute the revenue. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceItemId** | **string**| ID of the Invoice Item, for example, e20b07fd416dcfcf0141c81164fd0a75. If the specified Invoice Item is already associated with a revenue schedule, the call will fail.  | 
 **request** | [**PostRevenueScheduleByTransactionType**](PostRevenueScheduleByTransactionType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostRevenueScheduleByTransactionResponseType**](POSTRevenueScheduleByTransactionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRSforSubsCharge**
> PostRevenueScheduleByChargeResponseType POSTRSforSubsCharge($chargeKey, $request, $entityId, $entityName)

Create revenue schedule on subscription charge

Creates a revenue schedule by specifying the subscription charge. This method is for custom unlimited revenue recognition only. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeKey** | **string**| ID of the subscription rate plan charge; for example, 402892793e173340013e173b81000012. | 
 **request** | [**PostRevenueScheduleByChargeType**](PostRevenueScheduleByChargeType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostRevenueScheduleByChargeResponseType**](POSTRevenueScheduleByChargeResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTRSBasicInfo**
> CommonResponseType PUTRSBasicInfo($rsNumber, $request, $entityId, $entityName)

Update revenue schedule basic information

Retrieves basic information of a revenue schedule by specifying the revenue schedule number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;.  | 
 **request** | [**PutrsBasicInfoType**](PutrsBasicInfoType.md)|  | 
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

# **PUTRevenueAcrossAP**
> PutRevenueScheduleResponseType PUTRevenueAcrossAP($rsNumber, $request, $entityId, $entityName)

Distribute revenue across accounting periods

Distributes revenue by specifying the revenue schedule number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;.  | 
 **request** | [**PutAllocateManuallyType**](PutAllocateManuallyType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutRevenueScheduleResponseType**](PUTRevenueScheduleResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTRevenueByRecognitionStartandEndDates**
> PutRevenueScheduleResponseType PUTRevenueByRecognitionStartandEndDates($rsNumber, $request, $entityId, $entityName)

Distribute revenue by recognition start and end dates

Distributes revenue by specifying the recognition start and end dates. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. Specify the revenue schedule whose revenue you want to distribute.    The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;.  | 
 **request** | [**PutrsTermType**](PutrsTermType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutRevenueScheduleResponseType**](PUTRevenueScheduleResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTRevenueSpecificDate**
> PutRevenueScheduleResponseType PUTRevenueSpecificDate($rsNumber, $request, $entityId, $entityName)

Distribute revenue on specific date

Distributes revenue on a specific recognition date. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;.  | 
 **request** | [**PutSpecificDateAllocationType**](PutSpecificDateAllocationType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PutRevenueScheduleResponseType**](PUTRevenueScheduleResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

