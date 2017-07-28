# GetCreditMemoType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the customer account associated with the credit memo.  | [optional] [default to null]
**Amount** | **float64** | The total amount of the credit memo.  | [optional] [default to null]
**AppliedAmount** | **float64** | The applied amount of the credit memo.  | [optional] [default to null]
**AutoApplyUponPosting** | **bool** | Whether the credit memo automatically applies to the invoice upon posting.  | [optional] [default to null]
**Comment** | **string** | Comments about the credit memo.  | [optional] [default to null]
**CreatedById** | **string** | The ID of the Zuora user who created the credit memo.  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date and time when the credit memo was created, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-01 15:31:10.  | [optional] [default to null]
**CreditMemoDate** | [**time.Time**](time.Time.md) | The date when the credit memo takes effect, in &#x60;yyyy-mm-dd&#x60; format. For example, 2017-05-20.  | [optional] [default to null]
**Currency** | **string** | A currency defined in the web-based UI administrative settings.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**ExcludeFromAutoApplyRules** | **bool** | Whether the credit memo is excluded from the rule of automatically applying credit memos to invoices.  | [optional] [default to null]
**Id** | **string** | The unique ID of the credit memo.  | [optional] [default to null]
**LatestPDFFileId** | **string** | The ID of the latest PDF file generated for the credit memo.  | [optional] [default to null]
**Number** | **string** | The unique identification number of the credit memo.  | [optional] [default to null]
**PostedById** | **string** | The ID of the Zuora user who posted the credit memo.  | [optional] [default to null]
**PostedOn** | [**time.Time**](time.Time.md) | The date and time when the credit memo was posted, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**ReasonCode** | **string** | A code identifying the reason for the transaction. The value must be an existing reason code or empty.  | [optional] [default to null]
**ReferredInvoiceId** | **string** | The ID of a referred invoice.  | [optional] [default to null]
**RefundAmount** | **float64** | The amount of the refund on the credit memo.  | [optional] [default to null]
**Status** | **string** | The status of the credit memo.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully. | [optional] [default to null]
**TargetDate** | [**time.Time**](time.Time.md) | The target date for the credit memo, in &#x60;yyyy-mm-dd&#x60; format. For example, 2017-07-20.  | [optional] [default to null]
**TaxAmount** | **float64** | The amount of taxation.  | [optional] [default to null]
**TotalTaxExemptAmount** | **float64** | The total amount of taxes or VAT for which the customer has an exemption.  | [optional] [default to null]
**TransferredToAccounting** | **string** | Whether the credit memo was transferred to an external accounting system.   | [optional] [default to null]
**UnappliedAmount** | **float64** | The unapplied amount of the credit memo.  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the Zuora user who last updated the credit memo.  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date and time when the credit memo was last updated, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-01 15:36:10.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


