# GetCreditMemoItemType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The total amount of the credit memo item.  | [optional] [default to null]
**AppliedAmount** | **float64** | The applied amount of the credit memo item.  | [optional] [default to null]
**Comment** | **string** | Comments about the credit memo item.  | [optional] [default to null]
**CreatedById** | **string** | The ID of the Zuora user who created the credit memo item.  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date and time when the credit memo item was created, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-01 15:31:10.  | [optional] [default to null]
**CreditTaxItems** | [**[]GetcmTaxItemType**](GETCMTaxItemType.md) | Container for credit memo taxation items.  | [optional] [default to null]
**FinanceInformation** | [**GetCreditMemoItemTypeFinanceInformation**](GETCreditMemoItemType_financeInformation.md) |  | [optional] [default to null]
**Id** | **string** | The ID of the credit memo item.  | [optional] [default to null]
**RefundAmount** | **float64** | The amount of the refund on the credit memo item.  | [optional] [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) | The service end date of the credit memo item.  | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) | The service start date of the credit memo item. If the associated charge is a one-time fee, this date is the date of that charge.  | [optional] [default to null]
**Sku** | **string** | The SKU for the product associated with the credit memo item.  | [optional] [default to null]
**SkuName** | **string** | The name of the SKU.  | [optional] [default to null]
**SourceItemId** | **string** | The ID of the source item.  | [optional] [default to null]
**SourceItemType** | **string** | The type of the source item.  | [optional] [default to null]
**SubscriptionId** | **string** | The ID of the subscription associated with the credit memo item.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully. | [optional] [default to null]
**UnappliedAmount** | **float64** | The unapplied amount of the credit memo item.  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the Zuora user who last updated the credit memo item.  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date and time when the credit memo item was last updated, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-02 15:36:10.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


