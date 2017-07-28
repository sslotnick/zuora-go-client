# RatePlanDataRatePlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmendmentId** | **string** |  The ID of the amendment associated with the rate plan. This field only applies to amendment rate plans.   **Character limit**: 32   **Values**: a valid amendment ID  | [optional] [default to null]
**AmendmentSubscriptionRatePlanId** | **string** | The ID of the subscription rate plan modified by the amendment. This field only applies to amendment rate plans.  **Character limit**: 32   **Values**: a valid rate plan ID  | [optional] [default to null]
**AmendmentType** | **string** | The type of amendment associated with the rate plan. This field only applies to amendment rate plans.  **Character limit**: 18   **Values**: inherited from &#x60;Amendment.Type&#x60;  | [optional] [default to null]
**CreatedById** | **string** | The ID of the Zuora user who created the RatePlan object.  **Character limit**: 32   **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date when the &#x60;RatePlan&#x60; object was last updated.  **Character limit**: 29   **Values**: automatically generated  | [optional] [default to null]
**Name** | **string** | The name of the rate plan. Leave this null in a subscribe call to inherited the &#x60;ProductRatePlan.Name&#x60; field value.  **Character limit**: 100   **Values**: a string of 100 characters or fewer or inherited from ProductRatePlan.Name  | [optional] [default to null]
**ProductRatePlanId** | **string** | The ID of the associated product rate plan.  **Character limit**: 32   **Values**: a valid product rate plan ID  | [optional] [default to null]
**SubscriptionId** | **string** | The ID of the subscription that the rate plan belongs to.  **Character limit**: 32   **Values**: a valid subscription ID  | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the user who last updated the rate plan.   **Character limit**: 32   **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the rate plan was last updated.   **Character limit**: 29   **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


