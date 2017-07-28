# \PaymentMethodSnapshotsApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObjectGETPaymentMethodSnapshot**](PaymentMethodSnapshotsApi.md#ObjectGETPaymentMethodSnapshot) | **Get** /v1/object/payment-method-snapshot/{id} | CRUD: Retrieve PaymentMethodSnapshot


# **ObjectGETPaymentMethodSnapshot**
> ProxyGetPaymentMethodSnapshot ObjectGETPaymentMethodSnapshot($id, $entityId, $entityName, $fields)

CRUD: Retrieve PaymentMethodSnapshot

This REST API reference describes how to retrieve a Payment Method Snapshot.  A Payment Method Snapshot is a copy of the particular Payment Method used in a transaction. If the Payment Method is deleted, the Payment Method Snapshot continues to retain the data used in each of the past transactions.  ## Notes The following Payment Method fields are not available in Payment Method Snapshots:  * `Active` * `AchAddress1` * `AchAddress2` * `CreatedById` * `CreatedDate` * `UpdatedById` * `UpdatedDate`  The Payment Method Snapshot field `PaymentMethodId` is not available in Payment Methods. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Object id | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **fields** | **string**| Object fields to return | [optional] 

### Return type

[**ProxyGetPaymentMethodSnapshot**](ProxyGetPaymentMethodSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

