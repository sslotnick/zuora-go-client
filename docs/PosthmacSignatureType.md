# PosthmacSignatureType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | **string** | Customer account number or ID.  Specifies this field only when creating signatures for Create payment method.  | [optional] [default to null]
**Method** | **string** | Possible values are: &#39;GET&#39;, &#39;POST&#39;, &#39;PUT&#39;, &#39;DELETE&#39;, &#39;OPTIONS&#39;.  | [default to null]
**Name** | **string** | Account name.  Specifies this field only when creating signatures for Create account.  | [optional] [default to null]
**PageId** | **string** | The page id of your Payment Pages 2.0 form. Click **Show Page Id** next to the Payment Page name in the Hosted Page List to retrieve the page id.  Specifies this field only when creating signatures for RSA Signatures.  | [optional] [default to null]
**Uri** | **string** | The URI of the API object the customer will make a CORS enabled call to. e.g. \&quot;https://rest.zuora.com/v1/payment-methods/credit-cards\&quot;  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


