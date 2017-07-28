# \AccountingPeriodsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAccountingPeriod**](AccountingPeriodsApi.md#DELETEAccountingPeriod) | **Delete** /v1/accounting-periods/{ap-id} | Delete accounting period
[**GETAccountingPeriod**](AccountingPeriodsApi.md#GETAccountingPeriod) | **Get** /v1/accounting-periods/{ap-id} | Get accounting period
[**GETAllAccountingPeriods**](AccountingPeriodsApi.md#GETAllAccountingPeriods) | **Get** /v1/accounting-periods | Get all accounting periods
[**POSTAccountingPeriod**](AccountingPeriodsApi.md#POSTAccountingPeriod) | **Post** /v1/accounting-periods | Create accounting period
[**PUTCloseAccountingPeriod**](AccountingPeriodsApi.md#PUTCloseAccountingPeriod) | **Put** /v1/accounting-periods/{ap-id}/close | Close accounting period
[**PUTPendingCloseAccountingPeriod**](AccountingPeriodsApi.md#PUTPendingCloseAccountingPeriod) | **Put** /v1/accounting-periods/{ap-id}/pending-close | Set accounting period to pending close
[**PUTReopenAccountingPeriod**](AccountingPeriodsApi.md#PUTReopenAccountingPeriod) | **Put** /v1/accounting-periods/{ap-id}/reopen | Re-open accounting period
[**PUTRunTrialBalance**](AccountingPeriodsApi.md#PUTRunTrialBalance) | **Put** /v1/accounting-periods/{ap-id}/run-trial-balance | Run trial balance
[**PUTUpdateAccountingPeriod**](AccountingPeriodsApi.md#PUTUpdateAccountingPeriod) | **Put** /v1/accounting-periods/{ap-id} | Update accounting period


# **DELETEAccountingPeriod**
> CommonResponseType DELETEAccountingPeriod($apId, $entityId, $entityName)

Delete accounting period

 Deletes an accounting period.  Prerequisites -------------   * You must have Zuora Finance enabled on your tenant.   * You must have the Delete Accounting Period user permission. See [Finance Roles](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/User_Roles/f_Finance_Roles).   Limitations -----------  The accounting period to be deleted:  * Must be the most recent accounting period  * Must be an open accounting period  * Must have no revenue distributed into it  * Must not have any active journal entries  * Must not be the open-ended accounting period  * Must not be in the process of running a trial balance 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apId** | **string**| ID of the accounting period you want to delete. | 
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

# **GETAccountingPeriod**
> GetAccountingPeriodType GETAccountingPeriod($apId, $entityId, $entityName)

Get accounting period

Retrieves an accounting period. Prerequisites -------------  You must have Zuora Finance enabled on your tenant. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apId** | **string**| ID of the accounting period you want to get. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetAccountingPeriodType**](GETAccountingPeriodType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETAllAccountingPeriods**
> GetAccountingPeriodsType GETAllAccountingPeriods($entityId, $entityName, $pageSize)

Get all accounting periods

Retrieves all accounting periods on your tenant.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 300]

### Return type

[**GetAccountingPeriodsType**](GETAccountingPeriodsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTAccountingPeriod**
> PostAccountingPeriodResponseType POSTAccountingPeriod($request, $entityId, $entityName)

Create accounting period

Creates an accounting period. Prerequisites ------------- * You must have Zuora Finance enabled on your tenant. * You must have the Create Accounting Period user permission.  Limitations ----------- * When creating the first accounting period on your tenant, the start date must be equal to or earlier than the date of the earliest transaction on the tenant. * Start and end dates of accounting periods must be contiguous. For example, if one accounting period ends on January 31, the next period must start on February 1. * If you have the Revenue Recognition Package and have enabled the \"Monthly recognition over time\" revenue recognition model, the accounting period start date and end date must be on the first day and last day of the month, respectively. Note that the start and end dates do not necessarily have to be in the same month.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostAccountingPeriodType**](PostAccountingPeriodType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostAccountingPeriodResponseType**](POSTAccountingPeriodResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTCloseAccountingPeriod**
> CommonResponseType PUTCloseAccountingPeriod($apId, $entityId, $entityName)

Close accounting period

Close an accounting period by accounting period ID.  Prerequisites ------------- You must have Zuora Finance enabled on your tenant. You must have the Manage Close Process and Run Trial Balance user permissions.  Limitations ----------- * The accounting period cannot already be closed. * The accounting period cannot be in the process of running a trial balance. * All earlier accounting periods must be closed. * There must be no required action items for the accounting period. See Reconcile Transactions Before Closing an Accounting Period for more information.  Notes ----- When you close an accounting period in Zuora, a trial balance is automatically run for that period. A successful response means only that the accounting period is now closed, but does not mean that the trial balance has successfully completed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apId** | **string**| ID of the accounting period you want to close. | 
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

# **PUTPendingCloseAccountingPeriod**
> CommonResponseType PUTPendingCloseAccountingPeriod($apId, $entityId, $entityName)

Set accounting period to pending close

Sets an accounting period to pending close.   Prerequisites -------------  * You must have Zuora Finance enabled on your tenant. * You must have the Manage Close Process and Run Trial Balance user permissions.               Limitations   -----------    * The accounting period cannot be closed or pending close.    * The accounting period cannot be in the process of running a trial balance.    * All earlier accounting periods must be closed.     Notes ----- When you set an accounting period to pending close in Zuora, a trial balance is automatically run for that period. A response of `{ \"success\": true }`  means only that the accounting period status is now pending close, but does not mean that the trial balance has successfully completed. You can use the Get Accounting Period REST API call to view details about the outcome of the trial balance. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apId** | **string**| ID of the accounting period you want to set to pending close. | 
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

# **PUTReopenAccountingPeriod**
> CommonResponseType PUTReopenAccountingPeriod($apId, $entityId, $entityName)

Re-open accounting period

Re-opens an accounting period. Prerequisites ------------- * You must have Zuora Finance enabled on your tenant. * You must have the Manage Close Process and Run Trial Balance user permissions.  Limitations ----------- * The accounting period must be closed or pending close. * You can only re-open an accounting period that is immediately previous to an open period.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apId** | **string**| ID of the accounting period that you want to re-open. | 
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

# **PUTRunTrialBalance**
> CommonResponseType PUTRunTrialBalance($apId, $entityId, $entityName)

Run trial balance

Runs the trial balance for an accounting period.   Prerequisites -------------  * You must have Zuora Finance enabled on your tenant.  * You must have the Manage Close Process and Run Trial Balance user permissions. See [Finance Roles](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/User_Roles/f_Finance_Roles).             Limitations  -----------    * The accounting period must be open.    * The accounting period cannot already be in the process of running a trial balance.   Notes ----- The trial balance is run asynchronously. A response of `{ \"success\": true }` means only that the trial balance has started processing, but does not mean that the trial balance has successfully completed. You can use the [Get Accounting Period](https://knowledgecenter.zuora.com/DC_Developers/REST_API/B_REST_API_reference/Accounting_Periods/Get_Accounting_Period) REST API call to view details about the outcome of the trial balance. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apId** | **string**| ID of the accounting period for which you want to run a trial balance. | 
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

# **PUTUpdateAccountingPeriod**
> CommonResponseType PUTUpdateAccountingPeriod($apId, $request, $entityId, $entityName)

Update accounting period

 Updates an accounting period.  Prerequisites -------------  * You must have Zuora Finance enabled on your tenant.  * You must have the Create Accounting Period user permission. See [Finance Roles](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/User_Roles/f_Finance_Roles).  Limitations -----------  * You can update the start date of only the earliest accounting period on your tenant. You cannot update the start date of later periods.  * If you update the earliest accounting period, the start date must be equal to or earlier than the date of the earliest transaction on the tenant.  * Start and end dates of accounting periods must be contiguous. For example, if one accounting period ends on January 31, the next period must start on February 1.  * If you have the Revenue Recognition Package and have enabled the \"Monthly recognition over time\" revenue recognition model, the accounting period start date and end date must be on the first day and last day of the month, respectively. Note that the start and end dates do not necessarily have to be in the same month.  * You cannot update the start date or end date of an accounting period if:   * Any revenue has been distributed into the period.   * The period has any active journal entries. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apId** | **string**| ID of the accounting period you want to update. | 
 **request** | [**PutAccountingPeriodType**](PutAccountingPeriodType.md)|  | 
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

