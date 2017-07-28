# \GetFilesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETFiles**](GetFilesApi.md#GETFiles) | **Get** /v1/files/{file-id} | Get files


# **GETFiles**
> GETFiles($fileId, $entityId, $entityName)

Get files

Retrieve a file by specifying its file ID. You can retrieve accounting period reports, such as the Revenue Detail report, as well as other files such as invoice PDFs.  ## Example  ```curl curl -i -k -H \"apiAccessKeyId:$USER_NAME\" -H \"apiSecretAccessKey:$PASSWORD\" -H \"Accept:application/pdf\" -X GET https://rest.zuora.com/v1/files/2c92c08c55534cf00155581fb474314d -o /Users/jsmith/Downloads/invoiceFile1.pdf ``` The invoice PDF is downloaded to /Users/jsmith/Downloads and named invoiceFile1.pdf.  ## File Size Limitation The maximum export file size is 2047MB. If you have large data requests that go over this limit, you will get the following 403 HTTP response code from Zuora.  `<security:max-object-size>2047MB</security:max-object-size>`  Submit a request at [Zuora Global Support](https://zuora.zendesk.com/agent/) if you require additional assistance.  We can work with you to determine if large file optimization is an option for you. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileId** | **string**| The Zuora ID of the file you want to retrieve.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

