# RatePlanChargeTier

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedById** | **string** | The ID of the Zuora user who created the RatePlanChargeTier object.  **Character limit**: 32   **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the RatePlanChargeTier object was created.   **Character limit**: 29   **Values**: automatically generated  | [optional] [default to null]
**EndingUnit** | **float64** |  The end number of a range of units for the tier.   **Character limit**: 16   **Values**: any positive decimal value  | [optional] [default to null]
**IsOveragePrice** | **bool** |  Indicates if the price is an overage price. An overage occurs when usage surpasses the last defined tier. This field is applicable only to tier pricing and volume pricing models.    **Values**: true, false  | [optional] [default to null]
**Price** | **float64** |  The price of the tier if the charge is a flat fee, or the price of each unit in the tier if the change model is tiered pricing.   **Character limit**: 16   **Values**: any positive decimal value  | [optional] [default to null]
**PriceFormat** | **string** |  Indicates if the price is a flat fee or is per unit.   **Character limit**: 8   **Values**: &#x60;Flat Fee&#x60;, &#x60;Per Unit&#x60;  | [optional] [default to null]
**RatePlanChargeId** | **string** |  The ID of the subscription or amendment rate plan charge associated with this tier. You can&#39;t create an unassociated tier.   **Character limit**: 32   **Values**: inherited from &#x60;RatePlanCharge&#x60;.&#x60;Id&#x60;  | [optional] [default to null]
**StartingUnit** | **float64** |  The start number of a range of units for the tier.   **Character limit**: 16   **Values**: any positive decimal value  | [optional] [default to null]
**Tier** | **int32** |  A unique number that identifies the tier that the price applies to.   **Character limit**: 20   **Values**: automatically generated  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the last user to update the object.  **Character limit**: 32   **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the object was last updated.   **Character limit**: 29   **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


