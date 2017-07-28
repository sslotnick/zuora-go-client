# GetPaymentType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | **string** | Customer account ID.  | [optional] [default to null]
**AccountName** | **string** | Customer account name.  | [optional] [default to null]
**AccountNumber** | **string** | Customer account number.  | [optional] [default to null]
**Amount** | **string** | Payment amount.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | Effective payment date as _yyyy-mm-dd_.  | [optional] [default to null]
**GatewayTransactionNumber** | **string** | Transaction ID from payment gateway.  | [optional] [default to null]
**Id** | **string** | PaymentID.  | [optional] [default to null]
**PaidInvoices** | [**[]GetPaidInvoicesType**](GETPaidInvoicesType.md) | Information about one or more invoices to which this payment was applied:  | [optional] [default to null]
**PaymentMethodID** | **string** | Payment method.  | [optional] [default to null]
**PaymentNumber** | **string** | Unique payment number.  | [optional] [default to null]
**Status** | **string** | Possible values are: &#x60;Draft&#x60;, &#x60;Processing&#x60;, &#x60;Processed&#x60;, &#x60;Error&#x60;, &#x60;Voided&#x60;, &#x60;Canceled&#x60;, &#x60;Posted.  | [optional] [default to null]
**Type_** | **string** | Possible values are: &#x60;External&#x60;, &#x60;Electronic&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


