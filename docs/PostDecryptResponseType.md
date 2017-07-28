# PostDecryptResponseType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecryptedSignature** | **string** | The string of a list of the following items: Payment Pages 2.0 URL, tenant ID, timestamp,the Payment Page ID  The items are separated by &#39;#&#39;, e.g., \&quot;/apps/publichostedpagelite.do#12271#rvBp1AxBJwk6FrT7aqFuABIINiRbwJCc #1418848373103#2c92c0f948f899\&quot;  | [optional] [default to null]
**PublicKey** | **string** | The public key passed in as a request parameter.  | [optional] [default to null]
**Signature** | **string** | The signature passed in as a request parameter.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


