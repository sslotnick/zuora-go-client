# \AttachmentsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAttachments**](AttachmentsApi.md#DELETEAttachments) | **Delete** /v1/attachments/{attachment-id} | Delete attachments
[**GETAttachments**](AttachmentsApi.md#GETAttachments) | **Get** /v1/attachments/{attachment-id} | View attachments
[**GETAttachmentsList**](AttachmentsApi.md#GETAttachmentsList) | **Get** /v1/attachments/{object-type}/{object-key} | View attachments list
[**POSTAttachments**](AttachmentsApi.md#POSTAttachments) | **Post** /v1/attachments | Add attachments
[**PUTAttachments**](AttachmentsApi.md#PUTAttachments) | **Put** /v1/attachments/{attachment-id} | Edit attachments


# **DELETEAttachments**
> CommonResponseType DELETEAttachments($attachmentId, $entityId, $entityName)

Delete attachments

Use the Delete Attachment REST request to delete an attachment from a Zuora object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentId** | **string**| Id of the attachment to be deleted. | 
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

# **GETAttachments**
> GetAttachmentResponseType GETAttachments($attachmentId, $entityId, $entityName)

View attachments

Use the View Attachment REST request to retrieve information about an attachment document.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentId** | **string**| Id of the attachment you want to view. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetAttachmentResponseType**](GETAttachmentResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETAttachmentsList**
> GetAttachmentsResponseType GETAttachmentsList($objectType, $objectKey, $entityId, $entityName, $pageSize)

View attachments list

Use the View Attachment REST request to get a list of attachments on an account, an invoice, or a subscription.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **string**|  Specify one of the following: * Account  * Subscription  * Invoice  | 
 **objectKey** | **string**| Id of the attachment. This id is included in the response from the Add Attachment request. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetAttachmentsResponseType**](GETAttachmentsResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTAttachments**
> PostAttachmentResponseType POSTAttachments($associatedObjectType, $associatedObjectKey, $request, $entityId, $entityName, $description)

Add attachments

Use the Add Attachment REST request with a multipart/form-data to attach a document file to an Account, a Subscription, or an Invoice.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **associatedObjectType** | **string**| Specify one of the following values: Account, Subscription, or Invoice.  | 
 **associatedObjectKey** | **string**| For the Subscription type, specify the Subscription Number. An attachment is tied tied to the Subscription Number and thus viewable with every subscription version.  For Account and Invoice, specify the id.  | 
 **request** | [**PostAttachmentType**](PostAttachmentType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **description** | **string**| Description of the attachment document.  | [optional] 

### Return type

[**PostAttachmentResponseType**](POSTAttachmentResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTAttachments**
> CommonResponseType PUTAttachments($attachmentId, $entityId, $entityName, $request)

Edit attachments

Use the Edit Attachment REST request to make changes to the descriptive fields of an attachment, such as the description and the file name. You cannot change the actual content of the attached file in Zuora. If you need to change the actual content, you need to delete the attachment and add the updated file as a new attachment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentId** | **string**| Id of the attachment to be updated. | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **request** | [**PutAttachmentType**](PutAttachmentType.md)|  | [optional] 

### Return type

[**CommonResponseType**](CommonResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

