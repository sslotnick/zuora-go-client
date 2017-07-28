# GetDebitMemoItemType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount of the debit memo item.  | [optional] [default to null]
**Balance** | **float64** | The balance of the debit memo item.  | [optional] [default to null]
**BeAppliedAmount** | **float64** | The applied amount of the debit memo item.  | [optional] [default to null]
**Comment** | **string** | Comments about the debit memo item.  | [optional] [default to null]
**CreatedById** | **string** | The ID of the Zuora user who created the debit memo item.  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date and time when the debit memo item was created, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-01 15:31:10.  | [optional] [default to null]
**FinanceInformation** | [**GetDebitMemoItemTypeFinanceInformation**](GETDebitMemoItemType_financeInformation.md) |  | [optional] [default to null]
**Id** | **string** | The ID of the debit memo item.  | [optional] [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) | The end date of the service period associated with this debit memo item. Service ends one second before the date specified in this field.  | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) | The start date of the service period associated with this debit memo item. If the associated charge is a one-time fee, this date is the date of that charge.  | [optional] [default to null]
**Sku** | **string** | The SKU for the product associated with the debit memo item.  | [optional] [default to null]
**SkuName** | **string** | The name of the SKU.  | [optional] [default to null]
**SourceItemId** | **string** | The ID of the source item.  | [optional] [default to null]
**SourceItemType** | **string** | The type of the source item.   | [optional] [default to null]
**SubscriptionId** | **string** | The ID of the subscription associated with the debit memo item.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully. | [optional] [default to null]
**TaxItems** | [**[]GetdmTaxItemType**](GETDMTaxItemType.md) | Container for debit memo taxation items.  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the Zuora user who last updated the debit memo item.  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date and time when the debit memo item was last updated, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-02 15:36:10.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


