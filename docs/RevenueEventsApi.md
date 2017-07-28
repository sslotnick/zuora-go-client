# \RevenueEventsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETRevenueEventDetails**](RevenueEventsApi.md#GETRevenueEventDetails) | **Get** /v1/revenue-events/{event-number} | Get revenue event details
[**GETRevenueEventForRevenueSchedule**](RevenueEventsApi.md#GETRevenueEventForRevenueSchedule) | **Get** /v1/revenue-events/revenue-schedules/{rs-number} | Get revenue events for a revenue schedule


# **GETRevenueEventDetails**
> GetRevenueEventDetailType GETRevenueEventDetails($eventNumber, $entityId, $entityName)

Get revenue event details

 This REST API reference describes how to get revenue event details by specifying the revenue event number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventNumber** | **string**| The number associated with the revenue event. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRevenueEventDetailType**](GETRevenueEventDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRevenueEventForRevenueSchedule**
> GetRevenueEventDetailsType GETRevenueEventForRevenueSchedule($rsNumber, $entityId, $entityName, $pageSize)

Get revenue events for a revenue schedule

 This REST API reference describes how to get all revenue events in a revenue schedule by specifying the revenue schedule number. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rsNumber** | **string**| Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 8]

### Return type

[**GetRevenueEventDetailsType**](GETRevenueEventDetailsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

