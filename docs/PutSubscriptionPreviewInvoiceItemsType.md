# PutSubscriptionPreviewInvoiceItemsType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeAmount** | **string** | The amount of the charge. This amount doesn&#39;t include taxes unless the charge&#39;s tax mode is inclusive.  | [optional] [default to null]
**ChargeDescription** | **string** | Description of the charge.  | [optional] [default to null]
**ChargeName** | **string** | Name of the charge  | [optional] [default to null]
**ProductName** | **string** | Name of the product associated with this item.  | [optional] [default to null]
**ProductRatePlanChargeId** | **string** |  | [optional] [default to null]
**Quantity** | **string** | Quantity of this item.  | [optional] [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) | End date of the service period for this item, i.e., the last day of the period, as yyyy-mm-dd.  | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) | Service start date as yyyy-mm-dd. If the charge is a one-time fee, this is the date of that charge.  | [optional] [default to null]
**UnitOfMeasure** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


