# \ChargeRevenueSummariesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETCRSByCRSNumber**](ChargeRevenueSummariesApi.md#GETCRSByCRSNumber) | **Get** /v1/charge-revenue-summaries/{crs-number} | Get charge summary details by CRS number
[**GETCRSByChargeID**](ChargeRevenueSummariesApi.md#GETCRSByChargeID) | **Get** /v1/charge-revenue-summaries/subscription-charges/{charge-key} | Get charge summary details by charge ID


# **GETCRSByCRSNumber**
> GetChargeRsDetailType GETCRSByCRSNumber($crsNumber, $entityId, $entityName)

Get charge summary details by CRS number

This REST API reference describes how to retrieve the details of a charge revenue summary by specifying the charge revenue summary number. The response includes all revenue items associated with the charge revenue summary. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crsNumber** | **string**| The charge revenue summary number.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetChargeRsDetailType**](GETChargeRSDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETCRSByChargeID**
> GetChargeRsDetailType GETCRSByChargeID($chargeKey, $entityId, $entityName)

Get charge summary details by charge ID

This REST API reference describes how to retrieve the details of a charge revenue summary by specifying the subscription charge ID. This response retrieves all revenue items associated with a charge revenue summary. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeKey** | **string**| ID of the subscription rate plan charge; for example, 402892793e173340013e173b81000012.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetChargeRsDetailType**](GETChargeRSDetailType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

