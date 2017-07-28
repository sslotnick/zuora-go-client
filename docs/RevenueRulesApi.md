# \RevenueRulesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETRevenueRecRulebyProductRatePlanCharge**](RevenueRulesApi.md#GETRevenueRecRulebyProductRatePlanCharge) | **Get** /v1/revenue-recognition-rules/product-charges/{charge-key} | Get revenue recognition rule by product rate plan charge
[**GETRevenueRecRules**](RevenueRulesApi.md#GETRevenueRecRules) | **Get** /v1/revenue-recognition-rules/subscription-charges/{charge-key} | Get revenue recognition rule by subscription charge


# **GETRevenueRecRulebyProductRatePlanCharge**
> GetRevenueRecognitionRuleAssociationType GETRevenueRecRulebyProductRatePlanCharge($chargeKey)

Get revenue recognition rule by product rate plan charge

**Note:** This feature is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   Retrieves the revenue recognition rule associated with a production rate plan charge by specifying the charge ID. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeKey** | **string**| The unique ID of a product rate plan charge. For example, 8a8082e65ba86084015bb323d3c61d82.  | 

### Return type

[**GetRevenueRecognitionRuleAssociationType**](GETRevenueRecognitionRuleAssociationType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GETRevenueRecRules**
> GetRevenueRecognitionRuleAssociationType GETRevenueRecRules($chargeKey, $entityId, $entityName)

Get revenue recognition rule by subscription charge

Retrieves the revenue recognition rule associated with a subscription charge by specifying the charge ID. Request and response field descriptions and sample code are provided. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeKey** | **string**| The unique ID of the subscription rate plan charge. For example, 402892793e173340013e173b81000012.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetRevenueRecognitionRuleAssociationType**](GETRevenueRecognitionRuleAssociationType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

