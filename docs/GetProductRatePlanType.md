# GetProductRatePlanType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Description** | **string** | Rate plan description.  | [optional] [default to null]
**EffectiveEndDate** | [**time.Time**](time.Time.md) | Final date the rate plan is active, as &#x60;yyyy-mm-dd&#x60;. After this date, the rate plan status is &#x60;Expired&#x60;.  | [optional] [default to null]
**EffectiveStartDate** | [**time.Time**](time.Time.md) | First date the rate plan is active (i.e., available to be subscribed to), as &#x60;yyyy-mm-dd&#x60;.  Before this date, the status is &#x60;NotStarted&#x60;.  | [optional] [default to null]
**Id** | **string** | Unique product rate-plan charge ID.  | [optional] [default to null]
**Name** | **string** | Name of the product rate-plan charge. (Not required to be unique.)  | [optional] [default to null]
**ProductRatePlanCharges** | [**[]GetProductRatePlanChargeType**](GETProductRatePlanChargeType.md) | Field attributes describing the product rate plan charges:  | [optional] [default to null]
**Status** | **string** | Possible vales are: &#x60;Active&#x60;, &#x60;Expired&#x60;, &#x60;NotStarted&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


