# \NotificationHistoryApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETCalloutHistory**](NotificationHistoryApi.md#GETCalloutHistory) | **Get** /v1/notification-history/callout | Get callout notification histories
[**GETEmailHistory**](NotificationHistoryApi.md#GETEmailHistory) | **Get** /v1/notification-history/email | Get email notification histories


# **GETCalloutHistory**
> GetCalloutHistoryVOsType GETCalloutHistory($entityId, $entityName, $pageSize, $endTime, $startTime, $objectId, $failedOnly, $eventCategory, $includeResponseContent)

Get callout notification histories

This REST API reference describes how to get a notification history for callouts. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]
 **endTime** | **time.Time**| The final date and time of records to be returned. Defaults to now. Use format yyyy-MM-ddTHH:mm:ss. | [optional] 
 **startTime** | **time.Time**| The initial date and time of records to be returned. Defaults to (end time - 1 day). Use format yyyy-MM-ddTHH:mm:ss. | [optional] 
 **objectId** | **string**| The ID of an object that triggered a callout notification. | [optional] 
 **failedOnly** | **bool**| If &#x60;true&#x60;, only return failed records. If &#x60;false&#x60;, return all records in the given date range. The default value is &#x60;true&#x60;. | [optional] 
 **eventCategory** | **string**| Category of records to be returned by event category. | [optional] 
 **includeResponseContent** | **bool**|  | [optional] 

### Return type

[**GetCalloutHistoryVOsType**](GETCalloutHistoryVOsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETEmailHistory**
> GetEmailHistoryVOsType GETEmailHistory($entityId, $entityName, $pageSize, $endTime, $startTime, $objectId, $failedOnly, $eventCategory)

Get email notification histories

This REST API reference describes how to get a notification history for notification emails.   ## Notes Request parameters and their values may be appended with a \"?\" following the HTTPS GET request.  Additional request parameter are separated by \"&\".   For example:  `GET https://rest.zuora.com/v1/notification-history/email?startTime=2015-01-12T00:00:00&endTime=2015-01-15T00:00:00&failedOnly=false&eventCategory=1000&pageSize=1` 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]
 **endTime** | **time.Time**| The end date and time of records to be returned. Defaults to now. Use format yyyy-MM-ddTHH:mm:ss. The maximum date range (endTime - startTime) is three days. | [optional] 
 **startTime** | **time.Time**| The initial date and time of records to be returned. Defaults to (end time - 1 day). Use format yyyy-MM-ddTHH:mm:ss. The maximum date range (endTime - startTime) is three days. | [optional] 
 **objectId** | **string**| The Id of an object that triggered an email notification. | [optional] 
 **failedOnly** | **bool**| If &#x60;true&#x60;, only returns failed records. When &#x60;false&#x60;, returns all records in the given date range. Defaults to &#x60;true&#x60; when not specified. | [optional] 
 **eventCategory** | **string**| Category of records to be returned by event category. | [optional] 

### Return type

[**GetEmailHistoryVOsType**](GETEmailHistoryVOsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

