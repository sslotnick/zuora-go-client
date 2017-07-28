# \DebitMemosApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEDebitMemo**](DebitMemosApi.md#DELETEDebitMemo) | **Delete** /v1/debitmemos/{debitMemoId} | Delete debit memo
[**GETDebitMemo**](DebitMemosApi.md#GETDebitMemo) | **Get** /v1/debitmemos/{debitMemoId} | Get debit memo
[**GETDebitMemoItem**](DebitMemosApi.md#GETDebitMemoItem) | **Get** /v1/debitmemos/{debitMemoId}/items/{dmitemid} | Get debit memo item
[**GETDebitMemoItems**](DebitMemosApi.md#GETDebitMemoItems) | **Get** /v1/debitmemos/{debitMemoId}/items | Get debit memo items
[**GETDebitMemos**](DebitMemosApi.md#GETDebitMemos) | **Get** /v1/debitmemos | Get debit memos
[**POSTCancelDebitMemo**](DebitMemosApi.md#POSTCancelDebitMemo) | **Post** /v1/debitmemos/{debitMemoId}/cancel | Cancel debit memo
[**POSTDebitMemoFromInvoice**](DebitMemosApi.md#POSTDebitMemoFromInvoice) | **Post** /v1/debitmemos/invoice/{invoiceId} | Create debit memo from invoice
[**POSTDebitMemoFromPrpc**](DebitMemosApi.md#POSTDebitMemoFromPrpc) | **Post** /v1/debitmemos | Create debit memo from charge
[**POSTDebitMemoPDF**](DebitMemosApi.md#POSTDebitMemoPDF) | **Post** /v1/debitmemos/{debitMemoId}/pdfs | Create debit memo PDF
[**POSTEmailDebitMemo**](DebitMemosApi.md#POSTEmailDebitMemo) | **Post** /v1/debitmemos/{debitMemoId}/emails | Email debit memo
[**POSTPostDebitMemo**](DebitMemosApi.md#POSTPostDebitMemo) | **Post** /v1/debitmemos/{debitMemoId}/post | Post debit memo
[**PUTDebitMemo**](DebitMemosApi.md#PUTDebitMemo) | **Put** /v1/debitmemos/{debitMemoId} | Update debit memo


# **DELETEDebitMemo**
> CommonResponseType DELETEDebitMemo($debitMemoId, $entityId, $entityName)

Delete debit memo

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Deletes a debit memo. Only debit memos with the Cancelled status can be deleted.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debitMemoId** | **string**| The unique ID of a debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
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

# **GETDebitMemo**
> GetDebitMemoType GETDebitMemo($debitMemoId, $entityId, $entityName)

Get debit memo

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about a specific debit memo. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debitMemoId** | **string**| The unique ID of a debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetDebitMemoType**](GETDebitMemoType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETDebitMemoItem**
> GetDebitMemoItemType GETDebitMemoItem($dmitemid, $debitMemoId, $entityId, $entityName)

Get debit memo item

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Retrieves the information about a specific item of a debit memo. A debit memo item is a single line item in a debit memo. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dmitemid** | **string**| The unique ID of a debit memo item. You can get the debit memo item ID from the response of [Get debit memo items](https://www.zuora.com/developer/api-reference/#operation/GET_DebitMemoItems).  | 
 **debitMemoId** | **string**| The unique ID of a debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetDebitMemoItemType**](GETDebitMemoItemType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETDebitMemoItems**
> GetDebitMemoItemCollectionType GETDebitMemoItems($debitMemoId, $entityId, $entityName, $pageSize, $fieldsFilterable, $sort)

Get debit memo items

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about all items of a debit memo. A debit memo item is a single line item in a debit memo.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debitMemoId** | **string**| The unique ID of a debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]
 **fieldsFilterable** | **string**| This parameter restricts the data returned in the response. You can use this parameter to supply a dimension you want to filter on.   A single filter uses the following form:   *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*              Filters can be combined by using &#x60;&amp;&#x60;. For example: *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value* &#x60;&amp;&#x60; *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*  *fieldsFilterable* indicates the name of a supported field that you can use to filter the response data. The supported filterable fields of this operation are as below:    - id   - amount   - beAppliedAmount   - sku   - skuName   - serviceStartDate   - serviceEndDate   - sourceItemId   - createdDate   - createdById   - updatedDate   - updatedById   - subscriptionId   *field_value* indicates a valid value of the filterable field. If the type of the field value is string, you can set the field to &#x60;null&#x60; value for filtering. Then, you can get the response data with this field value being &#39;null&#39;. For more information about these fields, see the field descriptions in the **Responses** section.     Examples:  - /debitmemos/402890245c7ca371015c7cb40b28001f/items?amount&#x3D;100  - /debitmemos/402890245c7ca371015c7cb40b28001f/items?amount&#x3D;100&amp;sort&#x3D;createdDate  | [optional] 
 **sort** | **string**| This parameter restricts the order of the data returned in the response. You can use this parameter to supply a dimension you want to sort on.  A sortable field uses the following form:   *operator* *field_name*  You can use at most two sortable fields in one URL path. Use a comma to separate sortable fields. For example:  *operator* *field_name*, *operator* *field_name*    *operator* is used to mark the order of sequencing. The operator is optional. If you only specify the sortable field without any operator, the response data is sorted in descending order by this field.    - The &#x60;-&#x60; operator indicates an ascending order.   - The &#x60;+&#x60; operator indicates a descending order.  By default, the response data is displayed in descending order by updated date.  *field_name* indicates the name of a sortable field. The supported sortable fields of this operation are as below:    - id   - amount   - beAppliedAmount   - sku   - skuName   - serviceStartDate   - serviceEndDate   - sourceItemId   - createdDate   - createdById   - updatedDate   - updatedById   - subscriptionId    Examples:  - /debitmemos/402890245c7ca371015c7cb40b28001f/items?sort&#x3D;createdDate  - /debitmemos/402890245c7ca371015c7cb40b28001f/items?amount&#x3D;100&amp;sort&#x3D;createdDate  | [optional] 

### Return type

[**GetDebitMemoItemCollectionType**](GETDebitMemoItemCollectionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETDebitMemos**
> GetDebitMemoCollectionType GETDebitMemos($entityId, $entityName, $pageSize, $fieldsFilterable, $sort)

Get debit memos

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Retrieves the information about all debit memos associated with all customer accounts. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]
 **fieldsFilterable** | **string**| This parameter restricts the data returned in the response. You can use this parameter to supply a dimension you want to filter on.   A single filter uses the following form:   *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*              Filters can be combined by using &#x60;&amp;&#x60;. For example: *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value* &#x60;&amp;&#x60; *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*  *fieldsFilterable* indicates the name of a supported field that you can use to filter the response data. The supported filterable fields of this operation are as below:    - number   - accountID   - debitMemoDate   - targetDate   - dueDate   - currency   - status   - amount   - taxAmount   - totalTaxExemptAmount   - balance   - beAppliedAmount   - referredInvoiceId   - createdDate   - createdById   - updatedDate   - updatedById   *field_value* indicates a valid value of the filterable field. If the type of the field value is string, you can set the field to &#x60;null&#x60; value for filtering. Then, you can get the response data with this field value being &#39;null&#39;. For more information about these fields, see the field descriptions in the **Responses** section.     Examples:  - /debitmemos?status&#x3D;Processed  - /debitmemos?referredInvoiceId&#x3D;null&amp;status&#x3D;Draft  - /debitmemos?status&#x3D;Processed&amp;type&#x3D;External&amp;sort&#x3D;+number  | [optional] 
 **sort** | **string**| This parameter restricts the order of the data returned in the response. You can use this parameter to supply a dimension you want to sort on.  A sortable field uses the following form:   *operator* *field_name*  You can use at most two sortable fields in one URL path. Use a comma to separate sortable fields. For example:  *operator* *field_name*, *operator* *field_name*    *operator* is used to mark the order of sequencing. The operator is optional. If you only specify the sortable field without any operator, the response data is sorted in descending order by this field.    - The &#x60;-&#x60; operator indicates an ascending order.   - The &#x60;+&#x60; operator indicates a descending order.  By default, the response data is displayed in descending order by debit memo number.  *field_name* indicates the name of a sortable field. The supported sortable fields of this operation are as below:    - number   - accountID   - debitMemoDate   - targetDate   - dueDate   - amount   - taxAmount   - totalTaxExemptAmount   - balance   - beAppliedAmount   - referredInvoiceId   - createdDate   - createdById   - updatedDate   - updatedById    Examples:  - /debitmemos?sort&#x3D;+number  - /debitmemos?status&#x3D;Processed&amp;sort&#x3D;-number,+amount  | [optional] 

### Return type

[**GetDebitMemoCollectionType**](GETDebitMemoCollectionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTCancelDebitMemo**
> GetDebitMemoType POSTCancelDebitMemo($debitMemoId, $entityId, $entityName)

Cancel debit memo

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Cancels a debit memo. Only debit memos with the Draft status can be cancelled. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debitMemoId** | **string**| The unique ID of a debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetDebitMemoType**](GETDebitMemoType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTDebitMemoFromInvoice**
> GetDebitMemoType POSTDebitMemoFromInvoice($invoiceId, $body, $entityId, $entityName)

Create debit memo from invoice

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Creates an ad-hoc debit memo from an invoice. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceId** | **string**| | The ID of an invoice that you want to create a debit memo from.  | 
 **body** | [**DebitMemoFromInvoiceType**](DebitMemoFromInvoiceType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetDebitMemoType**](GETDebitMemoType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTDebitMemoFromPrpc**
> GetDebitMemoType POSTDebitMemoFromPrpc($body, $entityId, $entityName)

Create debit memo from charge

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Creates an ad-hoc debit memo from a product rate plan charge. Zuora only supports one-time positive charges for debit memos. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DebitMemoFromChargeType**](DebitMemoFromChargeType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetDebitMemoType**](GETDebitMemoType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTDebitMemoPDF**
> PostMemoPdfResponse POSTDebitMemoPDF($debitMemoId, $entityId, $entityName)

Create debit memo PDF

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Creates PDF for a specified debit memo. To access the generated PDF, you can download it by clicking **View PDF** on the detailed debit memo page through the Zuora UI. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debitMemoId** | **string**| The unique ID of the debit memo that you want to create PDF for. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**PostMemoPdfResponse**](POSTMemoPdfResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTEmailDebitMemo**
> CommonResponseType POSTEmailDebitMemo($request, $debitMemoId, $entityId, $entityName)

Email debit memo

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Sends a posted debit memo to the specified email addresses manually.    ## Notes   - You must activate the **Email Debit Memo | Manually email Debit Memo** notification before emailing debit memos. To include the debit memo PDF in the email, select the **Include Debit Memo PDF** check box in the **Edit notification** dialog from the Zuora UI. See [Create and Edit Notifications](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/Notifications/C_Create_Notifications#section_2) for more information.     - Zuora sends the email messages based on the email template you set. You can set the email template to use in the **Delivery Options** panel of the **Edit notification** dialog from the Zuora UI. By default, the **Manual Email for Debit Memo Default Template** template is used. See [Create and Edit Email Templates](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/Notifications/Create_Email_Templates) for more information.     - The debit memos are sent only to the work email addresses or personal email addresses of the Bill To contact if the following conditions are all met:      * The `useEmailTemplateSetting` field is set to `false`.     * The email addresses are not specified in the `emailAddresses` field. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**PostDebitMemoEmailType**](PostDebitMemoEmailType.md)|  | 
 **debitMemoId** | **string**| The ID of a posted debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
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

# **POSTPostDebitMemo**
> GetDebitMemoType POSTPostDebitMemo($debitMemoId, $entityId, $entityName)

Post debit memo

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Posts a debit memo to activate it. You can post debit memos only if you have the [Billing permissions](https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/User_Roles/d_Billing_Roles#Billing_Permissions). 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debitMemoId** | **string**| The unique ID of a debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetDebitMemoType**](GETDebitMemoType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTDebitMemo**
> GetDebitMemoType PUTDebitMemo($body, $debitMemoId, $entityId, $entityName)

Update debit memo

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Updates the basic and finance information about a debit memo. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PutDebitMemoType**](PutDebitMemoType.md)|  | 
 **debitMemoId** | **string**| The unique ID of a debit memo. For example, 8a8082e65b27f6c3015ba419f3c2644e.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetDebitMemoType**](GETDebitMemoType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

