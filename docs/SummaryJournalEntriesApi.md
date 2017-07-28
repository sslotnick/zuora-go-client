# \SummaryJournalEntriesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESummaryJournalEntry**](SummaryJournalEntriesApi.md#DELETESummaryJournalEntry) | **Delete** /v1/journal-entries/{je-number} | Delete summary journal entry
[**GETAllSummaryJournalEntries**](SummaryJournalEntriesApi.md#GETAllSummaryJournalEntries) | **Get** /v1/journal-entries/journal-runs/{jr-number} | Get all summary journal entries in a journal run
[**GETSummaryJournalEntry**](SummaryJournalEntriesApi.md#GETSummaryJournalEntry) | **Get** /v1/journal-entries/{je-number} | Get summary journal entry
[**POSTSummaryJournalEntry**](SummaryJournalEntriesApi.md#POSTSummaryJournalEntry) | **Post** /v1/journal-entries | Create summary journal entry
[**PUTBasicSummaryJournalEntry**](SummaryJournalEntriesApi.md#PUTBasicSummaryJournalEntry) | **Put** /v1/journal-entries/{je-number}/basic-information | Update basic information of a summary journal entry
[**PUTSummaryJournalEntry**](SummaryJournalEntriesApi.md#PUTSummaryJournalEntry) | **Put** /v1/journal-entries/{je-number}/cancel | Cancel summary journal entry


# **DELETESummaryJournalEntry**
> CommonResponseType DELETESummaryJournalEntry($jeNumber, $entityId, $entityName)

Delete summary journal entry

This reference describes how to delete a summary journal entry using the REST API.  You must have the \"Delete Cancelled Journal Entry\" user permission enabled to delete summary journal entries.  A summary journal entry must be canceled before it can be deleted. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jeNumber** | **string**| Journal entry number in the format JE-00000001. | 
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

# **GETAllSummaryJournalEntries**
> GetJournalEntriesInJournalRunType GETAllSummaryJournalEntries($jrNumber, $entityId, $entityName, $pageSize)

Get all summary journal entries in a journal run

 This REST API reference describes how to retrieve information about all summary journal entries in a journal run. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jrNumber** | **string**| Journal run number. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 8]

### Return type

[**GetJournalEntriesInJournalRunType**](GETJournalEntriesInJournalRunType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETSummaryJournalEntry**
> GetJournalEntryDetailType GETSummaryJournalEntry($jeNumber, $entityId, $entityName)

Get summary journal entry

This REST API reference describes how to get information about a summary journal entry by its journal entry number. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jeNumber** | **string**|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetJournalEntryDetailType**](GETJournalEntryDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTSummaryJournalEntry**
> PostJournalEntryResponseType POSTSummaryJournalEntry($request, $entityId, $entityName)

Create summary journal entry

This REST API reference describes how to manually create a summary journal entry. Request and response field descriptions and sample code are provided. ## Requirements 1.The sum of debits must equal the sum of credits in the summary journal entry.  2.The following applies only if you use foreign currency conversion:   * If you have configured Aggregate transactions with different currencies during a Journal Run to \"Yes\", the value of the **currency** field must be the same as your tenant's home currency. That is, you must create journal entries using your home currency.   * All journal entries in an accounting period must either all be aggregated or all be unaggregated. You cannot have a mix of aggregated and unaggregated journal entries in the same accounting period. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostJournalEntryType**](PostJournalEntryType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostJournalEntryResponseType**](POSTJournalEntryResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTBasicSummaryJournalEntry**
> CommonResponseType PUTBasicSummaryJournalEntry($jeNumber, $request, $entityId, $entityName)

Update basic information of a summary journal entry

 This REST API reference describes how to update the basic information of a summary journal entry. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jeNumber** | **string**| Journal entry number in the format JE-00000001. | 
 **request** | [**PutBasicSummaryJournalEntryType**](PutBasicSummaryJournalEntryType.md)|  | 
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

# **PUTSummaryJournalEntry**
> CommonResponseType PUTSummaryJournalEntry($jeNumber, $entityId, $entityName)

Cancel summary journal entry

 This reference describes how to cancel a summary journal entry using the REST API.  You must have the \"Cancel Journal Entry\" user permission enabled to cancel summary journal entries.  A summary journal entry cannot be canceled if its Transferred to Accounting status is \"Yes\" or \"Processing\". 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jeNumber** | **string**| Journal entry number in the format JE-00000001. | 
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

