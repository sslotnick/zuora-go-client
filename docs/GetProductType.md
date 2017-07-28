# GetProductType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | Category of the product. Used by Zuora Quotes Guided Product Selector.  Possible values are:   - Base Products   - Add On Services   - Miscellaneous Products  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Description** | **string** | Optional product description.  | [optional] [default to null]
**EffectiveEndDate** | [**time.Time**](time.Time.md) | The date when the product expires and cannot be subscribed to anymore, as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**EffectiveStartDate** | [**time.Time**](time.Time.md) | The date when the product becomes available and can be subscribed to, as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**Id** | **string** | Product ID.  | [optional] [default to null]
**Name** | **string** | Product name, up to 100 characters.  | [optional] [default to null]
**ProductFeatures** | [**[]GetProductFeatureType**](GetProductFeatureType.md) | Container for one or more product features. Only available when the following settings are enabled: - The Entitlements feature in your tenant - The Enable Feature Specification in Product and Subscriptions setting in Settings &gt; Billing  | [optional] [default to null]
**ProductRatePlans** | [**[]GetProductRatePlanType**](GETProductRatePlanType.md) | Container for one or more product rate plans:  | [optional] [default to null]
**Sku** | **string** | Unique product SKU, up to 50 characters.  | [optional] [default to null]
**Tags** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


