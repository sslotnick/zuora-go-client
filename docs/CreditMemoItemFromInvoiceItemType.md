# CreditMemoItemFromInvoiceItemType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount of the invoice item.  | [default to null]
**Comment** | **string** | Comments about the invoice item.  | [optional] [default to null]
**FinanceInformation** | [**CreditMemoItemFromInvoiceItemTypeFinanceInformation**](CreditMemoItemFromInvoiceItemType_financeInformation.md) |  | [optional] [default to null]
**InvoiceItemId** | **string** | The ID of the invoice item.  | [optional] [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) | The service end date of the invoice item.   | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) | The service start date of the invoice item.   | [optional] [default to null]
**SkuName** | **string** | The name of the SKU.  | [default to null]
**TaxItems** | [**[]CreditMemoTaxItemFromInvoiceTaxItemType**](CreditMemoTaxItemFromInvoiceTaxItemType.md) | Container for taxation items.  | [optional] [default to null]
**UnitOfMeasure** | **string** | The definable unit that you measure when determining charges.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


