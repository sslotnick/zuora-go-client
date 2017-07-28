# GetSubscriptionRatePlanType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Id** | **string** | Rate plan ID.  | [optional] [default to null]
**LastChangeType** | **string** | The last amendment on the rate plan.  Possible Values:  * &#x60;Add&#x60; * &#x60;Update&#x60; * &#x60;Remove&#x60;  | [optional] [default to null]
**ProductId** | **string** |  | [optional] [default to null]
**ProductName** | **string** |  | [optional] [default to null]
**ProductRatePlanId** | **string** |  | [optional] [default to null]
**ProductSku** | **string** | The unique SKU for the product.  | [optional] [default to null]
**RatePlanCharges** | [**[]GetSubscriptionRatePlanChargesType**](GETSubscriptionRatePlanChargesType.md) | Container for one or more charges.  | [optional] [default to null]
**RatePlanName** | **string** | Name of the rate plan.  | [optional] [default to null]
**SubscriptionProductFeatures** | [**[]GetSubscriptionProductFeatureType**](GETSubscriptionProductFeatureType.md) | Container for one or more features.   Only available when the following settings are enabled:  * The Entitlements feature in your tenant.  * The Enable Feature Specification in Product and Subscriptions setting in Zuora Billing Settings | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


