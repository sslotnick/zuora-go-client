# \RevenueItemsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETRevenueItemsByChargeRevenueEventNumber**](RevenueItemsApi.md#GETRevenueItemsByChargeRevenueEventNumber) | **Get** /v1/revenue-items/revenue-events/{event-number} | Get revenue items by revenue event number
[**GETRevenueItemsByChargeRevenueSummaryNumber**](RevenueItemsApi.md#GETRevenueItemsByChargeRevenueSummaryNumber) | **Get** /v1/revenue-items/charge-revenue-summaries/{crs-number} | Get revenue items by charge revenue summary number
[**GETRevenueItemsByRevenueSchedule**](RevenueItemsApi.md#GETRevenueItemsByRevenueSchedule) | **Get** /v1/revenue-items/revenue-schedules/{rs-number} | Get revenue items by revenue schedule
[**PUTCustomFieldsonRevenueItemsByRevenueEvent**](RevenueItemsApi.md#PUTCustomFieldsonRevenueItemsByRevenueEvent) | **Put** /v1/revenue-items/revenue-events/{event-number} | Update custom fields on revenue items by revenue event number
[**PUTCustomFieldsonRevenueItemsByRevenueSchedule**](RevenueItemsApi.md#PUTCustomFieldsonRevenueItemsByRevenueSchedule) | **Put** /v1/revenue-items/revenue-schedules/{rs-number} | Update custom fields on revenue items by revenue schedule number


# **GETRevenueItemsByChargeRevenueEventNumber**
> GetRevenueItemsType GETRevenueItemsByChargeRevenueEventNumber($eventNumber, $entityId, $entityName, $pageSize)

Get revenue items by revenue event number

This REST API reference describes how to get the details of each revenue item in a revenue event by specifying the revenue event number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventNumber** | **string**| The number associated with the revenue event. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 300]

### Return type

[**GetRevenueItemsType**](GETRevenueItemsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRevenueItemsByChargeRevenueSummaryNumber**
> GetRevenueItemsType GETRevenueItemsByChargeRevenueSummaryNumber($crsNumber, $entityId, $entityName, $pageSize)

Get revenue items by charge revenue summary number

This REST API reference describes how to get the details for each revenue item in a charge revenue summary by specifying the charge revenue summary number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crsNumber** | **string**| The charge revenue summary number. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 300]

### Return type

[**GetRevenueItemsType**](GETRevenueItemsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRevenueItemsByRevenueSchedule**
> GetRsRevenueItemsType GETRevenueItemsByRevenueSchedule($rsNumber, $entityId, $entityName, $pageSize)

Get revenue items by revenue schedule

This REST API reference describes how to get the details for each revenue items in a revenue schedule by specifying the revenue schedule number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 300]

### Return type

[**GetRsRevenueItemsType**](GETRsRevenueItemsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTCustomFieldsonRevenueItemsByRevenueEvent**
> CommonResponseType PUTCustomFieldsonRevenueItemsByRevenueEvent($eventNumber, $request, $entityId, $entityName)

Update custom fields on revenue items by revenue event number

This REST API reference describes how to update custom fields on revenue items by specifying the revenue event number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventNumber** | **string**| The number associated with the revenue event. | 
 **request** | [**PutEventRiDetailType**](PutEventRiDetailType.md)|  | 
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

# **PUTCustomFieldsonRevenueItemsByRevenueSchedule**
> CommonResponseType PUTCustomFieldsonRevenueItemsByRevenueSchedule($rsNumber, $request, $entityId, $entityName)

Update custom fields on revenue items by revenue schedule number

This REST API reference describes how to update custom fields on revenue Items by specifying the revenue schedule number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;. | 
 **request** | [**PutScheduleRiDetailType**](PutScheduleRiDetailType.md)|  | 
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

