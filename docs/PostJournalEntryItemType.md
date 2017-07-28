# PostJournalEntryItemType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingCodeName** | **string** | Name of the accounting code.  | [default to null]
**AccountingCodeType** | **string** | Accounting code type. This field is required if &#x60;accountingCodeName&#x60; is not unique.  Note that &#x60;On-Account Receivable&#x60; is only available if you enable the Advanced AR Settlement feature.   | [optional] [default to null]
**Amount** | **string** | Journal entry item amount in transaction currency.  | [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**HomeCurrencyAmount** | **string** | Journal entry item amount in home currency.  This field is required if you have set your home currency for foreign currency conversion. Otherwise, do not pass this field.  | [optional] [default to null]
**Type_** | **string** | Type of journal entry item.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


