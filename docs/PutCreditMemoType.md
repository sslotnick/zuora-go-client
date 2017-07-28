# PutCreditMemoType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApplyUponPosting** | **bool** | Whether the credit memo automatically applies to the invoice upon posting.  | [optional] [default to null]
**Comment** | **string** | Comments about the credit memo.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The date when the credit memo takes effect.  | [optional] [default to null]
**ExcludeFromAutoApplyRules** | **bool** | Whether the credit memo is excluded from the rule of automatically applying credit memos to invoices.  | [optional] [default to null]
**Items** | [**[]PutCreditMemoItemType**](PUTCreditMemoItemType.md) | Container for credit memo items.  | [optional] [default to null]
**ReasonCode** | **string** | A code identifying the reason for the transaction. The value must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code.  | [optional] [default to null]
**TransferredToAccounting** | **string** | Whether the credit memo is transferred to an external accounting system. Use this field for integration with accounting systems, such as NetSuite.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


