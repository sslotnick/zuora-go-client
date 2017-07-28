# PostGenerateBillingDocumentType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPost** | **bool** | Determines whether to auto post the billing documents once the draft billing documents are generated.   If an error occurred during posting billing documents, the draft billing documents are not generated too.  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The date on which to generate the billing documents, in &#x60;yyyy-mm-dd&#x60; format.  | [optional] [default to null]
**SubscriptionIds** | **[]string** | The IDs of the subscriptions that you want to create the billing documents for.   | [optional] [default to null]
**TargetDate** | [**time.Time**](time.Time.md) | The date used to determine which charges are to be billed, in &#x60;yyyy-mm-dd&#x60; format.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


