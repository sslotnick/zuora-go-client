# GetJournalRunType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateCurrency** | **bool** |  | [optional] [default to null]
**ExecutedOn** | [**time.Time**](time.Time.md) | Date and time the journal run was executed.  | [optional] [default to null]
**JournalEntryDate** | [**time.Time**](time.Time.md) | Date of the journal entry.  | [optional] [default to null]
**Number** | **string** | Journal run number.  | [optional] [default to null]
**SegmentationRuleName** | **string** | Name of GL segmentation rule used in the journal run.  | [optional] [default to null]
**Status** | **string** | Status of the journal run.   | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**TargetEndDate** | [**time.Time**](time.Time.md) | The target end date of the journal run.  | [optional] [default to null]
**TargetStartDate** | [**time.Time**](time.Time.md) | The target start date of the journal run.  | [optional] [default to null]
**TotalJournalEntryCount** | **int64** | Total number of journal entries in the journal run.  | [optional] [default to null]
**TransactionTypes** | [**[]GetJournalRunTransactionType**](GETJournalRunTransactionType.md) | Transaction types included in the journal run.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


