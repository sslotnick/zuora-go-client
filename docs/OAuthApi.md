# \OAuthApi

All URIs are relative to *https://rest.zuora.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](OAuthApi.md#CreateToken) | **Post** /oauth/token | Generate an OAuth token


# **CreateToken**
> TokenResponse CreateToken($clientId, $clientSecret, $entityId, $grantType)

Generate an OAuth token

Generates a bearer token for a user to access the Zuora REST API.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| A unique string representing the registration information provided by the client, the recipient of the token.   | 
 **clientSecret** | **string**| The secret passphrase configured for the OAuth client. | 
 **entityId** | **string**| Entity ID associated with this bearer token | [optional] 
 **grantType** | **string**| The type of authentication being used to obtain the token. Use \&quot;client_credentials\&quot;.  | [optional] 

### Return type

[**TokenResponse**](tokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json; charset=utf-8
 - **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

