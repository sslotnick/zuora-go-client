# DebitMemoFromChargeType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account associated with the debit memo.  | [default to null]
**Charges** | [**[]DebitMemoFromChargeDetailType**](DebitMemoFromChargeDetailType.md) | Container for product rate plan charges.  | [optional] [default to null]
**Comment** | **string** | Comments about the debit memo.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The date when the debit memo takes effect.  | [optional] [default to null]
**ReasonCode** | **string** | A code identifying the reason for the transaction. The value must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


