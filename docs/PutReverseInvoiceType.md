# PutReverseInvoiceType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyEffectiveDate** | [**time.Time**](time.Time.md) | The date when the credit memo was applied to the invoice that will be reversed, in &#x60;yyyy-mm-dd&#x60; format. The effective date must be later than or equal to the memo date.  Default value is today&#39;s date.  | [optional] [default to null]
**MemoDate** | [**time.Time**](time.Time.md) | The date when the credit memo was created, in &#x60;yyyy-mm-dd&#x60; format. The memo date must be later than or equal to the invoice date.  Default value is today&#39;s date.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


