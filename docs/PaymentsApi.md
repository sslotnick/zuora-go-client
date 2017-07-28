# \PaymentsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPayment**](PaymentsApi.md#DELETEPayment) | **Delete** /v1/payments/{paymentId} | Delete payment
[**GETPayemntItemParts**](PaymentsApi.md#GETPayemntItemParts) | **Get** /v1/payments/{paymentId}/parts/{partid}/itemparts | Get payment part items
[**GETPayment**](PaymentsApi.md#GETPayment) | **Get** /v1/payments/{paymentId} | Get payment
[**GETPaymentItemPart**](PaymentsApi.md#GETPaymentItemPart) | **Get** /v1/payments/{paymentId}/parts/{partid}/itemparts/{itempartid} | Get payment part item
[**GETPaymentPart**](PaymentsApi.md#GETPaymentPart) | **Get** /v1/payments/{paymentId}/parts/{partid} | Get payment part
[**GETPaymentParts**](PaymentsApi.md#GETPaymentParts) | **Get** /v1/payments/{paymentId}/parts | Get payment parts
[**GETRetrieveAllPayments**](PaymentsApi.md#GETRetrieveAllPayments) | **Get** /v1/payments | Get all payments
[**ObjectDELETEPayment**](PaymentsApi.md#ObjectDELETEPayment) | **Delete** /v1/object/payment/{id} | CRUD: Delete Payment
[**ObjectGETPayment**](PaymentsApi.md#ObjectGETPayment) | **Get** /v1/object/payment/{id} | CRUD: Retrieve Payment
[**ObjectPOSTPayment**](PaymentsApi.md#ObjectPOSTPayment) | **Post** /v1/object/payment | CRUD: Create Payment
[**ObjectPUTPayment**](PaymentsApi.md#ObjectPUTPayment) | **Put** /v1/object/payment/{id} | CRUD: Update Payment
[**POSTApplyPayment**](PaymentsApi.md#POSTApplyPayment) | **Post** /v1/payments/{paymentId}/apply | Apply payment
[**POSTCancelPayment**](PaymentsApi.md#POSTCancelPayment) | **Post** /v1/payments/{paymentId}/cancel | Cancel payment
[**POSTCreatePayment**](PaymentsApi.md#POSTCreatePayment) | **Post** /v1/payments | Create payment
[**POSTRefundPayment**](PaymentsApi.md#POSTRefundPayment) | **Post** /v1/payments/{paymentId}/refund | Refund payment
[**POSTTransferPayemnt**](PaymentsApi.md#POSTTransferPayemnt) | **Post** /v1/payments/{paymentId}/accounts | Transfer payment
[**POSTUNAPPLYPAYMENT**](PaymentsApi.md#POSTUNAPPLYPAYMENT) | **Post** /v1/payments/{paymentId}/unapply | Unapply payment
[**PUTUpdatePayemnt**](PaymentsApi.md#PUTUpdatePayemnt) | **Put** /v1/payments/{paymentId} | Update payment


# **DELETEPayment**
> CommonResponseType DELETEPayment($paymentId, $entityId, $entityName)

Delete payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Deletes a payment. Only payments with the Cancelled status can be deleted.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentId** | **string**| The unique ID of an unapplied payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
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

# **GETPayemntItemParts**
> GetPaymentItemPartCollectionType GETPayemntItemParts($partid, $paymentId, $entityId, $entityName, $pageSize)

Get payment part items

**Note:** The Invoice Item Settlement feature is in **Limited Availability**, and it must be used together with other Advanced AR Settlement features (Unapplied Payments, and Credit and Debit memos). If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about all items of a payment part. A payment part item is a single line item in a payment part. A payment part can consist of several different types of items. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partid** | **string**| The unique ID of a specific payment part. You can get the payment part ID from the response of [Get payment parts](https://www.zuora.com/developer/api-reference/#operation/GET_PaymentParts).  | 
 **paymentId** | **string**| The unique ID of a payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetPaymentItemPartCollectionType**](GETPaymentItemPartCollectionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETPayment**
> GetarPaymentType GETPayment($paymentId, $entityId, $entityName)

Get payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about one specific payment. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentId** | **string**| The unique ID of a payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetarPaymentType**](GETARPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETPaymentItemPart**
> GetPaymentItemPartType GETPaymentItemPart($partid, $itempartid, $paymentId, $entityId, $entityName)

Get payment part item

**Note:** The Invoice Item Settlement feature is in **Limited Availability**, and it must be used together with other Advanced AR Settlement features (Unapplied Payments, and Credit and Debit memos). If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  Retrieves the information about a specific payment part item. A payment part item is a single line item in a payment part. A payment part can consist of several different types of items. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partid** | **string**| The unique ID of a specific payment part. You can get the payment part ID from the response of [Get payment parts](https://www.zuora.com/developer/api-reference/#operation/GET_PaymentParts).  | 
 **itempartid** | **string**| The unique ID of a specific payment part item. You can get the payment part item ID from the response of [Get payment part items](https://www.zuora.com/developer/api-reference/#operation/GET_PayemntItemParts).  | 
 **paymentId** | **string**| The unique ID of a payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetPaymentItemPartType**](GETPaymentItemPartType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETPaymentPart**
> GetPaymentPartType GETPaymentPart($partid, $paymentId, $entityId, $entityName)

Get payment part

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about a specific payment part. A payment can consist of an unapplied part, and several parts applied to invoices and debit memos. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partid** | **string**| The unique ID of a specific payment part. You can get the payment part ID from the response of [Get payment parts](https://www.zuora.com/developer/api-reference/#operation/GET_PaymentParts).  | 
 **paymentId** | **string**| The unique ID of a payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetPaymentPartType**](GETPaymentPartType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETPaymentParts**
> GetPaymentPartsCollectionType GETPaymentParts($paymentId, $entityId, $entityName, $pageSize)

Get payment parts

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about all parts of a payment. A payment can consist of an unapplied part, and several parts applied to invoices and debit memos. You can use this operation to get all the applied and unapplied portions of a payment. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentId** | **string**| The unique ID of a payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]

### Return type

[**GetPaymentPartsCollectionType**](GETPaymentPartsCollectionType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRetrieveAllPayments**
> PaymentCollectionResponseType GETRetrieveAllPayments($entityId, $entityName, $pageSize, $fieldsFilterable, $sort)

Get all payments

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the information about all payments from all your customer accounts. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **pageSize** | **int32**| Number of rows returned per page.  | [optional] [default to 20]
 **fieldsFilterable** | **string**| This parameter restricts the data returned in the response. You can use this parameter to supply a dimension you want to filter on.   A single filter uses the following form:   *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*              Filters can be combined by using &#x60;&amp;&#x60;. For example: *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value* &amp; *fieldsFilterable* &#x60;&#x3D;&#x60; *field_value*  *fieldsFilterable* indicates the name of a supported field that you can use to filter the response data. The supported filterable fields of this operation are as below:    - number   - status   - type   - accountID   - amount   - appliedAmount   - unappliedAmount   - refundAmount   - creditBalanceAmount   - currency   - effectiveDate   - createdDate   - createdById   - updatedDate   - updatedById   *field_value* indicates a valid value of the filterable field. If the type of the field value is string, you can set the field to &#x60;null&#x60; value for filtering. Then, you can get the response data with this field value being &#39;null&#39;. For more information about these fields, see the field descriptions in the **Responses** section.     Examples:  - /payments?status&#x3D;Processed  - /payments?currency&#x3D;USD&amp;status&#x3D;Processed  - /payments?status&#x3D;Processed&amp;type&#x3D;External&amp;sort&#x3D;+number  | [optional] 
 **sort** | **string**| This parameter restricts the order of the data returned in the response. You can use this parameter to supply a dimension you want to sort on.  A sortable field uses the following form:   *operator* *field_name*  You can use at most two sortable fields in one URL path. Use a comma to separate sortable fields. For example:  *operator* *field_name*, *operator* *field_name*    *operator* is used to mark the order of sequencing. The operator is optional. If you only specify the sortable field without any operator, the response data is sorted in descending order by this field.    - The &#x60;-&#x60; operator indicates an ascending order.   - The &#x60;+&#x60; operator indicates a descending order.  By default, the response data is displayed in descending order by payment number.  *field_name* indicates the name of a sortable field. The supported sortable fields of this operation are as below:    - number   - accountID   - amount   - appliedAmount   - unappliedAmount   - refundAmount   - creditBalanceAmount   - effectiveDate   - createdDate   - createdById   - updatedDate   - updatedById    Examples:  - /payments?sort&#x3D;+number  - /payments?status&#x3D;Processed&amp;sort&#x3D;-number,+amount  | [optional] 

### Return type

[**PaymentCollectionResponseType**](PaymentCollectionResponseType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectDELETEPayment**
> ProxyDeleteResponse ObjectDELETEPayment($id, $entityId, $entityName)

CRUD: Delete Payment




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

# **ObjectGETPayment**
> ProxyGetPayment ObjectGETPayment($id, $entityId, $entityName, $fields)

CRUD: Retrieve Payment




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **fields** | **string**| Object fields to return | [optional] 

### Return type

[**ProxyGetPayment**](ProxyGetPayment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObjectPOSTPayment**
> ProxyCreateOrModifyResponse ObjectPOSTPayment($createRequest, $entityId, $entityName)

CRUD: Create Payment




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**ProxyCreatePayment**](ProxyCreatePayment.md)|  | 
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

# **ObjectPUTPayment**
> ProxyCreateOrModifyResponse ObjectPUTPayment($id, $modifyRequest, $entityId, $entityName)

CRUD: Update Payment




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **modifyRequest** | [**ProxyModifyPayment**](ProxyModifyPayment.md)|  | 
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

# **POSTApplyPayment**
> GetarPaymentType POSTApplyPayment($body, $paymentId, $entityId, $entityName)

Apply payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Applies an unapplied payment to invoices and debit memos.  For more information, see [Apply Unapplied Payments to Invoices and Debit Memos](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/A_Unapplied_Payments/Management_of_Unapplied_Payments/Apply_Unapplied_Payments_to_Invoices_and_Debit_Memos). 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplyPaymentType**](ApplyPaymentType.md)|  | 
 **paymentId** | **string**| The unique ID of an unapplied payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetarPaymentType**](GETARPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTCancelPayment**
> GetarPaymentType POSTCancelPayment($paymentId, $entityId, $entityName)

Cancel payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Cancels a payment. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentId** | **string**| The unique ID of an unapplied payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetarPaymentType**](GETARPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTCreatePayment**
> GetarPaymentType POSTCreatePayment($body, $entityId, $entityName)

Create payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Creates a payment for the following scenarios:  - A full payment on an invoice or debit memo - A partial payment - A payment for several invoices and debit memos - An unapplied payment   If you do not know to which customer account the payment belongs, you can create a payment without specifying a customer account.  For more information, see [Create Payments](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/A_Unapplied_Payments/Management_of_Unapplied_Payments/AA_Create_Payments) and [Create Payments Without Specifying Customer Accounts](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/A_Unapplied_Payments/Management_of_Unapplied_Payments/AA_Create_Payments_Without_Specifying_Customer_Accounts).      


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreatePaymentType**](CreatePaymentType.md)|  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetarPaymentType**](GETARPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTRefundPayment**
> GetRefundPaymentType POSTRefundPayment($body, $paymentId, $entityId, $entityName)

Refund payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Refunds a full or partial unapplied payment to your customers. To refund applied payments, you must unapply the applied payments from the invoices or debit memos, and then refund the unapplied payments to customers.  For more information, see [Refund Payments](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/A_Unapplied_Payments/Management_of_Unapplied_Payments/Z_Refund_Payments). 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PostRefundType**](PostRefundType.md)|  | 
 **paymentId** | **string**| The unique ID of an unapplied payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRefundPaymentType**](GETRefundPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTTransferPayemnt**
> GetarPaymentType POSTTransferPayemnt($body, $paymentId, $entityId, $entityName)

Transfer payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Transfers an unapplied payment.  For more information, see [Transfer Unapplied Payments](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/A_Unapplied_Payments/Management_of_Unapplied_Payments/Transfer_Unapplied_Payments). 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TransferPaymentType**](TransferPaymentType.md)|  | 
 **paymentId** | **string**| The unique ID of an unapplied payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetarPaymentType**](GETARPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **POSTUNAPPLYPAYMENT**
> GetarPaymentType POSTUNAPPLYPAYMENT($body, $paymentId, $entityId, $entityName)

Unapply payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Unapplies an applied payment from invoices and debit memos.  For more information, see [Unapply Payments from Invoices and Debit Memos](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/A_Unapplied_Payments/Management_of_Unapplied_Payments/Unapply_Payments_from_Invoices_and_Debit_Memos). 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UnapplyPaymentType**](UnapplyPaymentType.md)|  | 
 **paymentId** | **string**| The unique ID of an applied payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetarPaymentType**](GETARPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PUTUpdatePayemnt**
> GetarPaymentType PUTUpdatePayemnt($body, $paymentId, $entityId, $entityName)

Update payment

**Note:** The Advanced AR Settlement feature is in **Limited Availability**. This feature includes Unapplied Payments, Credit and Debit Memo, and Invoice Item Settlement. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Updates a payment. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdatePaymentType**](UpdatePaymentType.md)|  | 
 **paymentId** | **string**| The unique ID of an unapplied payment. For example, 8a8082e65b27f6c3015b89e4344c16b1.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetarPaymentType**](GETARPaymentType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

