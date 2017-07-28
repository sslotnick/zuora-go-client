# PutPaymentMethodType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | First address line, 255 characters or less.  | [optional] [default to null]
**AddressLine2** | **string** | Second address line, 255 characters or less.  | [optional] [default to null]
**CardHolderName** | **string** | The full name as it appears on the card, e.g., \&quot;John J Smith\&quot;, 50 characters or less.  | [optional] [default to null]
**City** | **string** | City, 40 characters or less.  | [optional] [default to null]
**Country** | **string** | Country; must be a valid country name or abbreviation.  | [optional] [default to null]
**DefaultPaymentMethod** | **bool** | Specify \&quot;true\&quot; to make this card the default payment method; otherwise, omit this parameter to keep the current default payment method.  | [optional] [default to null]
**Email** | **string** | Card holder&#39;s email address, 80 characters or less.  | [optional] [default to null]
**ExpirationMonth** | **string** | Two-digit expiration month (01-12).  | [optional] [default to null]
**ExpirationYear** | **string** | Four-digit expiration year.  | [optional] [default to null]
**Phone** | **string** | Phone number, 40 characters or less.  | [optional] [default to null]
**SecurityCode** | **string** | The CVV or CVV2 security code for the credit card or debit card. Only required if changing expirationMonth, expirationYear, or cardHolderName. To ensure PCI compliance, this value isn&#39;t stored and can&#39;t be queried.   | [optional] [default to null]
**State** | **string** | State; must be a valid state name or 2-character abbreviation.  | [optional] [default to null]
**ZipCode** | **string** | Zip code, 20 characters or less.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


