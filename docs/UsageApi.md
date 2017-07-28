# \UsageApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETUsage**](UsageApi.md#GETUsage) | **Get** /v1/usage/accounts/{account-key} | Get usage
[**ObjectDELETEUsage**](UsageApi.md#ObjectDELETEUsage) | **Delete** /v1/object/usage/{id} | CRUD: Delete Usage
[**ObjectGETUsage**](UsageApi.md#ObjectGETUsage) | **Get** /v1/object/usage/{id} | CRUD: Retrieve Usage
[**ObjectPOSTUsage**](UsageApi.md#ObjectPOSTUsage) | **Post** /v1/object/usage | CRUD: Create Usage
[**ObjectPUTUsage**](UsageApi.md#ObjectPUTUsage) | **Put** /v1/object/usage/{id} | CRUD: Update Usage
[**POSTUsage**](UsageApi.md#POSTUsage) | **Post** /v1/usage | Post usage


# **GETUsage**
> GetUsageWrapper GETUsage($accountKey, $entityId, $entityName, $pageSize)

Get usage

This REST API reference describes how to retrieve usage details for an account. Usage data is returned in reverse chronological order. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKey** | **string**| Account number or account ID. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetUsageWrapper**](GETUsageWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectDELETEUsage**
> ProxyDeleteResponse ObjectDELETEUsage($id, $entityId, $entityName)

CRUD: Delete Usage




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

# **ObjectGETUsage**
> ProxyGetUsage ObjectGETUsage($id, $entityId, $entityName, $fields)

CRUD: Retrieve Usage




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **fields** | **string**| Object fields to return | [optional] 

### Return type

[**ProxyGetUsage**](ProxyGetUsage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectPOSTUsage**
> ProxyCreateOrModifyResponse ObjectPOSTUsage($createRequest, $entityId, $entityName)

CRUD: Create Usage




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**ProxyCreateUsage**](ProxyCreateUsage.md)|  | 
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

# **ObjectPUTUsage**
> ProxyCreateOrModifyResponse ObjectPUTUsage($id, $modifyRequest, $entityId, $entityName)

CRUD: Update Usage




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **modifyRequest** | [**ProxyModifyUsage**](ProxyModifyUsage.md)|  | 
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

# **POSTUsage**
> PostUsageResponseType POSTUsage($contentType, $entityId, $entityName)

Post usage

 This REST API reference describes how to post or import usage data for one or more accounts in the CSV  format. There are no path or query parameters. The data is uploade using the HTTP multipart/form-data POST method and applied to the user's tenant.   ## How this REST API Call Works The content of the upload file must follow the format used by the UI import tool. It must be a comma-separated (CSV) file with a corresponding .csv extension. The file size must not exceed 4MB. Click [here](https://knowledgecenter.zuora.com/@api/deki/files/4105/UsageFileFormat.csv) to download the usage file template.  At the completion of the upload, before actually processing the file contents, theAPI returns a response containing the byte count of the received file and a URL for checking the status of the import process.  Of the five possible results displayed at that URL Pending, Processing, Completed, Canceled, and Failed) only a Completed status indicates that the import was successful.  The operation is atomic; if any record fails, the file is rejected.  In that case, the entire import is rolled back and all stored data is returned to its original state.  To view the actual import status, enter the resulting status URL from the checkImportStatus response using a tool such as POSTMAN.This additional step provides more information about why the import may have failed.  To manage the information after a successful upload, use the web-based UI.  ## Upload File Format The upload file uses the following headings:  | Heading         | Description   | Required | |-----------------|--------|----------| | ACCOUNT_ID      | Enter the account number, e.g., the default account number, such as A00000001, or your custom account number.,Although this field is labeled as Account_Id, it is not the actual Account ID nor Account Name.  | Yes      | | UOM             | Enter the unit of measure. This must match the UOM for the usage. | Yes      | | QTY             | Enter the quantity.  | Yes      | | STARTDATE       | Enter the start date of the usage.,This date determines the invoice item service period the associated usage is billed to. Date format is based on locale of the current user. Default date format: `MM/DD/YYYY` | Yes      | | ENDDATE         | Enter the end date of the usage.,This is not used in calculations for usage billing and is optional. Date format is based on locale of the current user. Default date format: `MM/DD/YYYY`    | Yes      | | SUBSCRIPTION_ID | Enter the subscription number or subscription name. If you created the subscription in the Zuora application, Zuora created a number automatically in a format similar to A-S00000001. If you do not provide a value for this field, the associated usage will be added to all subscriptions for the specified Account that use this Unit Of Measure. If your Accounts can have multiple subscriptions and you do not want double or triple counting of usage, you must specify the Subscription or Charge ID in each usage record.  | Yes      | | CHARGE_ID       | Enter the charge number (not the charge name). You can see the charge ID, e.g., C-00000001, when you add your rate plan to your subscription and view your individual charges. If your Accounts can have multiple subscriptions and you do not want double or triple counting of usage, you must specify the specific Subscription or Charge ID in each usage record. This field is related to the Charge Number on the subscription rate plan.                       | Yes      | | DESCRIPTION     | Enter a description for the charge. | No       | 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**| Must be set to \&quot;multipart/form-data\&quot;.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostUsageResponseType**](POSTUsageResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

