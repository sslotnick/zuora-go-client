# PostBillingPreviewInvoiceItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeAmount** | **string** | The amount of the charge. This amount doesn&#39;t include taxes regardless if the charge&#39;s tax mode is inclusive or exclusive.  | [optional] [default to null]
**ChargeDate** | [**time.Time**](time.Time.md) | The date when the invoice item was created.  | [optional] [default to null]
**ChargeDescription** | **string** | Description of the charge.  | [optional] [default to null]
**ChargeId** | **string** | Id of the charge.  | [optional] [default to null]
**ChargeName** | **string** | Name of the charge.  | [optional] [default to null]
**ChargeNumber** | **string** | Number of the charge.  | [optional] [default to null]
**ChargeType** | **string** | The type of charge.   Possible values are &#x60;OneTime&#x60;, &#x60;Recurring&#x60;, and &#x60;Usage&#x60;.  | [optional] [default to null]
**Id** | **string** | Invoice item ID.  | [optional] [default to null]
**ProcessingType** | **float64** | Identifies the kind of charge.   Possible values: * 0: charge * 1: discount * 2: prepayment * 3: tax | [optional] [default to null]
**ProductName** | **string** | Name of the product associated with this item.  | [optional] [default to null]
**Quantity** | **string** | Quantity of this item, in the configured unit of measure for the charge.  | [optional] [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) | End date of the service period for this item, i.e., the last day of the service period, in &#x60;yyyy-mm-dd&#x60; format.  | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) | Start date of the service period for this item, in &#x60;yyyy-mm-dd&#x60; format. If the charge is a one-time fee, this is the date of that charge.  | [optional] [default to null]
**SubscriptionId** | **string** | ID of the subscription associated with this item.  | [optional] [default to null]
**SubscriptionName** | **string** | Name of the subscription associated with this item.  | [optional] [default to null]
**SubscriptionNumber** | **string** | Number of the subscription associated with this item.  | [optional] [default to null]
**TaxAmount** | **string** | Tax applied to the charge.  | [optional] [default to null]
**UnitOfMeasure** | **string** | Unit used to measure consumption.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


