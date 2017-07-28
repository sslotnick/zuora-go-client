# PostRevenueScheduleByDateRangeTypeRevenueEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EventType** | **string** | Label of the revenue event type. Revenue event type labels can be duplicated. You can configure your revenue event type labels by navigating to **Settings &gt; Finance &gt; Configure Revenue Event Types** in the Zuora UI.  Note that &#x60;Credit Memo Posted&#x60; and &#x60;Debit Memo Posted&#x60; are only available if you enable the Advanced AR Settlement feature.  | [optional] [default to null]
**EventTypeSystemId** | **string** | System ID of the revenue event type. Each eventType has a unique system ID. You can configure your revenue event type system IDs by navigating to **Settings &gt; Finance &gt; Configure Revenue Event Types** in the Zuora UI.  Required only if there is more than one revenue event type with the same label.  | [optional] [default to null]
**Notes** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


