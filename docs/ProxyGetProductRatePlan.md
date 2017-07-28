# ProxyGetProductRatePlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedById** | **string** | The ID of the Zuora user who created the &#x60;ProductRatePlan&#x60; object. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the &#x60;ProductRatePlan&#x60; object was created. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**Description** | **string** | A description of the product rate plan. **Character limit**: 500 **Values**: a string of 500 characters or fewer  | [optional] [default to null]
**EffectiveEndDate** | [**time.Time**](time.Time.md) |  The date when the product rate plan expires and can&#39;t be subscribed to, in &#x60;yyyy-mm-dd&#x60; format. **Character limit**: 29  | [optional] [default to null]
**EffectiveStartDate** | [**time.Time**](time.Time.md) |  The date when the product rate plan becomes available and can be subscribed to, in &#x60;yyyy-mm-dd&#x60; format. **Character limit**: 29  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**Name** | **string** | The name of the product rate plan. The name doesn&#39;t have to be unique in a Product Catalog, but the name has to be unique within a product. **Character limit**: 100 **Values**: a string of 100 characters or fewer  | [optional] [default to null]
**ProductId** | **string** | The ID of the product that contains the product rate plan. **Character limit**: 32 **Values**: a string of 32 characters or fewer  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the last user to update the object. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date when the object was last updated. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


