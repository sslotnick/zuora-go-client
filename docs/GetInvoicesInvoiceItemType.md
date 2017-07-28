# GetInvoicesInvoiceItemType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeAmount** | **string** | The amount of the charge. This amount doesn&#39;t include taxes regardless if the charge&#39;s tax mode is inclusive or exclusive.  | [optional] [default to null]
**ChargeDescription** | **string** | Description of the charge.  | [optional] [default to null]
**ChargeId** | **string** | ID of the charge.  | [optional] [default to null]
**ChargeName** | **string** | Name of the charge.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Id** | **string** | Item ID.  | [optional] [default to null]
**ProductName** | **string** | Name of the product associated with this item.  | [optional] [default to null]
**Quantity** | **string** | Quantity of this item, in the configured unit of measure for the charge.  | [optional] [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) | End date of the service period for this item, i.e., the last day of the service period, as _yyyy-mm-dd_.  | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) | Start date of the service period for this item, as _yyyy-mm-dd_. For a one-time fee item, the date of the charge.  | [optional] [default to null]
**SubscriptionId** | **string** | ID of the subscription for this item.  | [optional] [default to null]
**SubscriptionName** | **string** | Name of the subscription for this item.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**TaxAmount** | **string** | Tax applied to the charge.  | [optional] [default to null]
**UnitOfMeasure** | **string** | Unit used to measure consumption.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


