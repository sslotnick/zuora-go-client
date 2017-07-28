# PutDebitMemoItemType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount of the debit memo item.  | [optional] [default to null]
**Comment** | **string** | Comments about the debit memo item.  | [optional] [default to null]
**FinanceInformation** | [**PutDebitMemoItemTypeFinanceInformation**](PUTDebitMemoItemType_financeInformation.md) |  | [optional] [default to null]
**Id** | **string** | The ID of the debit memo item.  | [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) | The service end date of the debit memo item.  | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) | The service start date of the debit memo item.   | [optional] [default to null]
**SkuName** | **string** | The name of the SKU.  | [optional] [default to null]
**TaxItems** | [**[]PutDebitMemoTaxItemType**](PutDebitMemoTaxItemType.md) | Container for debit memo taxation items.  | [optional] [default to null]
**UnitOfMeasure** | **string** | The definable unit that you measure when determining charges.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


