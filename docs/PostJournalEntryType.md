# PostJournalEntryType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingPeriodName** | **string** | Name of the accounting period. The open-ended accounting period is named &#x60;Open-Ended&#x60;.  | [default to null]
**Currency** | **string** | The type of currency used. Currency must be active.  | [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**JournalEntryDate** | [**time.Time**](time.Time.md) | Date of the journal entry.  | [default to null]
**JournalEntryItems** | [**[]PostJournalEntryItemType**](POSTJournalEntryItemType.md) | Key name that represents the list of journal entry items.  | [default to null]
**Notes** | **string** | The number associated with the revenue event.  Character limit: 2,000  | [optional] [default to null]
**Segments** | [**[]PostJournalEntrySegmentType**](POSTJournalEntrySegmentType.md) | List of segments that apply to the summary journal entry.  | [optional] [default to null]
**TransferredToAccounting** | **string** | Status shows whether the journal entry has been transferred to an accounting system.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


