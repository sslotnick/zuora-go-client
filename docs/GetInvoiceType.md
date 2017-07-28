# GetInvoiceType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Customer account ID.  | [optional] [default to null]
**AccountName** | **string** | Customer account name.  | [optional] [default to null]
**AccountNumber** | **string** | Customer account number.  | [optional] [default to null]
**Amount** | **string** | Amount of the invoice before adjustments, discounts, and similar items.  | [optional] [default to null]
**Balance** | **string** | Balance remaining due on the invoice (after adjustments, discounts, etc.)  | [optional] [default to null]
**Body** | **string** | The REST URL of the invoice PDF file.  | [optional] [default to null]
**CreatedBy** | **string** | User ID of the person who created the invoice. If a bill run generated the invoice, then this is the user ID of person who created the bill run.  | [optional] [default to null]
**CreditBalanceAdjustmentAmount** | **string** |  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**DueDate** | [**time.Time**](time.Time.md) | Payment due date as _yyyy-mm-dd_.  | [optional] [default to null]
**Id** | **string** | Invoice ID.  | [optional] [default to null]
**InvoiceDate** | [**time.Time**](time.Time.md) | Invoice date as _yyyy-mm-dd_  | [optional] [default to null]
**InvoiceFiles** | [**[]GetInvoiceFileType**](GETInvoiceFileType.md) | Information about the invoice PDF file:  | [optional] [default to null]
**InvoiceItems** | [**[]GetInvoicesInvoiceItemType**](GETInvoicesInvoiceItemType.md) | Information on one or more items on this invoice:  | [optional] [default to null]
**InvoiceNumber** | **string** | Unique invoice ID, returned as a string.  | [optional] [default to null]
**InvoiceTargetDate** | [**time.Time**](time.Time.md) | Date through which charges on this invoice are calculated, as _yyyy-mm-dd_.  | [optional] [default to null]
**Status** | **string** | Status of the invoice in the system - not the payment status, but the status of the invoice itself. Possible values are: &#x60;Posted&#x60;, &#x60;Draft&#x60;, &#x60;Canceled&#x60;, &#x60;Error&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


