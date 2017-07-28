# PostAccountTypeCreditCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardHolderInfo** | [**PostAccountTypeCreditCardCardHolderInfo**](POSTAccountType_creditCard_cardHolderInfo.md) |  | [optional] [default to null]
**CardNumber** | **string** | Card number, up to 16 characters. Once created, this field can&#39;t be updated or queried, and is only available in masked format (e.g., XXXX-XXXX-XXXX-1234).  | [default to null]
**CardType** | **string** | Possible values are: &#x60;Visa&#x60;, &#x60;MasterCard&#x60;, &#x60;AmericanExpress&#x60;, &#x60;Discover&#x60;.  | [default to null]
**ExpirationMonth** | **string** | Two-digit expiration month (01-12).  | [default to null]
**ExpirationYear** | **string** | Four-digit expiration year.  | [default to null]
**SecurityCode** | **string** | The CVV or CVV2 security code of the card. To ensure PCI compliance, this value isn&#39;t stored and can&#39;t be queried.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


