# PostSubscriptionPreviewTypePreviewAccountInfoBillToContact

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | The city of the bill-to address. The value should be 40 characters or less.  | [optional] [default to null]
**Country** | **string** | The country of the bill-to address. The value must be a valid country name or abbreviation.  **Note:** You must specify this field if you are using Zuora Tax for this account.  | [optional] [default to null]
**County** | **string** | The county of the bill-to address. The value should be 32 characters or less.  | [optional] [default to null]
**State** | **string** | The state of the bill-to address. The value must be a valid state or province name or 2-character abbreviation.  **Note:** You must specify this field if you are using Zuora Tax for this account and the country is &#x60;USA&#x60; or &#x60;Canada&#x60;.  | [optional] [default to null]
**TaxRegion** | **string** | If using Zuora Tax, a region string as optionally defined in your tax rules.  | [optional] [default to null]
**ZipCode** | **string** | The zip code of the bill-to address. The value should be 20 characters or less.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


