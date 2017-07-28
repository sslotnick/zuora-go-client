# PutBasicSummaryJournalEntryType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**JournalEntryItems** | [**[]PutJournalEntryItemType**](PUTJournalEntryItemType.md) | Key name that represents the list of journal entry items.  | [optional] [default to null]
**Notes** | **string** | Additional information about this record.  ***Character limit:*** 2,000  | [optional] [default to null]
**TransferredToAccounting** | **string** | Status shows whether the journal entry has been transferred to an accounting system.   This field cannot be changed after the summary journal entry has been canceled.  **Note:** The Zuora Finance ***Override Transferred to Accounting*** permission is required to change &#x60;transferredToAccounting&#x60; from &#x60;Yes&#x60; to any other value.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


