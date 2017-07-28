# GetJournalEntryDetailTypeWithoutSuccess

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingPeriodName** | **string** | Name of the accounting period that the journal entry belongs to.  | [optional] [default to null]
**AggregateCurrency** | **bool** | Returns true if the journal entry is aggregating currencies. That is, if the journal entry was created when the &#x60;Aggregate transactions with different currencies during a JournalRun&#x60; setting was configured to \&quot;Yes\&quot;. Otherwise, returns &#x60;false&#x60;.  | [optional] [default to null]
**Currency** | **string** | Currency used.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**HomeCurrency** | **string** | Home currency used.  | [optional] [default to null]
**JournalEntryDate** | [**time.Time**](time.Time.md) | Date of the journal entry.  | [optional] [default to null]
**JournalEntryItems** | [**[]GetJournalEntryItemType**](GETJournalEntryItemType.md) | Key name that represents the list of journal entry items.  | [optional] [default to null]
**Notes** | **string** | Additional information about this record. Character limit: 2,000  | [optional] [default to null]
**Number** | **string** | Journal entry number in the format JE-00000001.  | [optional] [default to null]
**Segments** | [**[]GetJournalEntrySegmentType**](GETJournalEntrySegmentType.md) | List of segments that apply to the summary journal entry.  | [optional] [default to null]
**Status** | **string** | Status of journal entry.  | [optional] [default to null]
**TimePeriodEnd** | [**time.Time**](time.Time.md) | End date of time period included in the journal entry.  | [optional] [default to null]
**TimePeriodStart** | [**time.Time**](time.Time.md) | Start date of time period included in the journal entry.  | [optional] [default to null]
**TransactionType** | **string** | Transaction type of the transactions included in the summary journal entry.  | [optional] [default to null]
**TransferDateTime** | [**time.Time**](time.Time.md) | Date and time that transferredToAccounting was changed to &#x60;Yes&#x60;. This field is returned only when transferredToAccounting is &#x60;Yes&#x60;. Otherwise, this field is &#x60;null&#x60;.  | [optional] [default to null]
**TransferredBy** | **string** | User ID of the person who changed transferredToAccounting to &#x60;Yes&#x60;. This field is returned only when transferredToAccounting is &#x60;Yes&#x60;. Otherwise, this field is &#x60;null&#x60;.  | [optional] [default to null]
**TransferredToAccounting** | **string** | Status shows whether the journal entry has been transferred to an accounting system.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


