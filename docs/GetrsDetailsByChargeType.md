# GetrsDetailsByChargeType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPage** | [**ParametersGlobalResponseNextPage**](#/parameters/GLOBAL_RESPONSE_nextPage.md) |  | [optional] [default to null]
**RevenueSchedules** | [**[]GetrsDetailWithoutSuccessType**](GETRSDetailWithoutSuccessType.md) | Represents how revenue will be recognized over time.  This contains the details of a revenue schedule. If you do not specify the &#x60;pageSize&#x60; variable, the default number of revenue schedules returned per invocation is 8, and if there are more than 8 revenue schedules to be returned, the &#x60;nextPage&#x60; field will provide a hyperlink to view the next page(s) of revenue events. The order of revenue schedules is descending by the &#x60;updatedOn&#x60; field.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


