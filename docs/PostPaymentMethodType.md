# PostPaymentMethodType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | **string** | ID of the customer account to update.  | [default to null]
**CardHolderInfo** | [**PostPaymentMethodTypeCardHolderInfo**](POSTPaymentMethodType_cardHolderInfo.md) |  | [optional] [default to null]
**CreditCardNumber** | **string** | Credit card number, a string of up to 16 characters. This field can only be set when creating a new payment method; it cannot be queried or updated.  | [default to null]
**CreditCardType** | **string** | Possible values are: &#x60;Visa&#x60;, &#x60;MasterCard&#x60;, &#x60;AmericanExpress&#x60;, &#x60;Discover&#x60;.  | [default to null]
**DefaultPaymentMethod** | **bool** | Specify true to make this card the default payment method; otherwise, omit this parameter to keep the current default payment method.  | [optional] [default to null]
**ExpirationMonth** | **string** | Two-digit expiration month (01-12).  | [default to null]
**ExpirationYear** | **string** | Four-digit expiration year.  | [default to null]
**SecurityCode** | **string** | The CVV or CVV2 security code for the credit card or debit card. Only required if changing expirationMonth, expirationYear, or cardHolderName. To ensure PCI compliance, this value isn&#39;t stored and can&#39;t be queried.   | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


