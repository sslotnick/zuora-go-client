# GetAccountSummaryInvoiceType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Invoice amount before adjustments, discounts, and similar items.  | [optional] [default to null]
**Balance** | **string** | Balance due on the invoice.  | [optional] [default to null]
**DueDate** | [**time.Time**](time.Time.md) | Due date as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**Id** | **string** | Invoice ID.  | [optional] [default to null]
**InvoiceDate** | [**time.Time**](time.Time.md) | Invoice date as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**InvoiceNumber** | **string** | Invoice number.  | [optional] [default to null]
**Status** | **string** | Invoice status - not the payment status of the invoice, just the status of the invoice itself. Possible values are: &#x60;Posted&#x60;, &#x60;Draft&#x60;, &#x60;Canceled&#x60;, &#x60;Error&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


