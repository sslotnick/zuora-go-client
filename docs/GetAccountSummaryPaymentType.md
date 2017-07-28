# GetAccountSummaryPaymentType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveDate** | [**time.Time**](time.Time.md) | Effective date as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**Id** | **string** | Payment ID.  | [optional] [default to null]
**PaidInvoices** | [**[]GetAccountSummaryPaymentInvoiceType**](GETAccountSummaryPaymentInvoiceType.md) | Container for paid invoices for this subscription.  | [optional] [default to null]
**PaymentNumber** | **string** | Payment number.  | [optional] [default to null]
**PaymentType** | **string** | Payment type; possible values are: &#x60;External&#x60;, &#x60;Electronic&#x60;.  | [optional] [default to null]
**Status** | **string** | Payment status. Possible values are: &#x60;Draft&#x60;, &#x60;Processing&#x60;, &#x60;Processed&#x60;, &#x60;Error&#x60;, &#x60;Voided&#x60;, &#x60;Canceled&#x60;, &#x60;Posted&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


