# GetUsageType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Customer account ID.  | [optional] [default to null]
**AccountName** | **string** | Customer account name.  | [optional] [default to null]
**AccountNumber** | **string** | Customer account number.  | [optional] [default to null]
**ChargeNumber** | **string** | Number of the rate-plan charge that pays for this usage.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Id** | **string** | Unique ID for the usage item.  | [optional] [default to null]
**Quantity** | **string** | Number of units used.  | [optional] [default to null]
**SourceName** | **string** | Source of the usage data. Possible values are: &#x60;Import&#x60;, &#x60;API&#x60;.  | [optional] [default to null]
**StartDateTime** | [**time.Time**](time.Time.md) | Start date of the time period in which usage is tracked. Zuora uses this field value to determine the usage date.  | [optional] [default to null]
**Status** | **string** | Possible values are: &#x60;Importing&#x60;, &#x60;Pending&#x60;, &#x60;Processed&#x60;.  | [optional] [default to null]
**SubmissionDateTime** | [**time.Time**](time.Time.md) | Date when usage was submitted.  | [optional] [default to null]
**SubscriptionNumber** | **string** | Number of the subscription covering this usage.  | [optional] [default to null]
**UnitOfMeasure** | **string** | Unit used to measure consumption.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


