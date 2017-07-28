# \CustomExchangeRatesApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETCustomExchangeRates**](CustomExchangeRatesApi.md#GETCustomExchangeRates) | **Get** /v1/custom-exchange-rates/{currency} | Get custom foreign currency exchange rates


# **GETCustomExchangeRates**
> GetCustomExchangeRatesType GETCustomExchangeRates($currency, $startDate, $endDate, $entityId, $entityName)

Get custom foreign currency exchange rates

This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).   This reference describes how to query custom foreign exchange rates from Zuora. You can use this API method to query exchange rates only if you use a custom exchange rate provider and upload rates with the Import Foreign Exchange Rates mass action.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string**| The target currency. The exchange rates in the response are calculated in relation to the target currency.  The value must be a three-letter currency code, for example, USD.   | 
 **startDate** | **string**| Start date of the date range for which you want to get exchange rates.  The date must be in yyyy-mm-dd format, for example, 2016-01-15. The start date cannot be later than the end date.  | 
 **endDate** | **string**| End date of the date range for which you want to get exchange rates.  The date must be in yyyy-mm-dd format, for example, 2016-01-16. The end date can be a maximum of 90 days after the start date.  | 
 **entityId** | **string**| The Id of the entity that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 
 **entityName** | **string**| The [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name) that you want to access. Note that you must have permission to access the entity. For more information, see [REST Authentication](https://www.zuora.com/developer/api-reference/#section/Authentication/Entity-Id-and-Entity-Name). | [optional] 

### Return type

[**GetCustomExchangeRatesType**](GETCustomExchangeRatesType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

