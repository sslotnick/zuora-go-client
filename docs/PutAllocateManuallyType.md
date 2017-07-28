# PutAllocateManuallyType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EventType** | **string** | Label of the revenue event type. Revenue event type labels can be duplicated. You can configure your revenue event type labels by navigating to **Settings &gt; Finance &gt; Configure Revenue Event Types** in the Zuora UI.  Note that &#x60;Credit Memo Posted&#x60; and &#x60;Debit Memo Posted&#x60; are only available if you enable the Advanced AR Settlement feature.  | [optional] [default to null]
**EventTypeSystemId** | **string** | System ID of the revenue event type. Each eventType has a unique system ID. You can configure your revenue event type system IDs by navigating to **Settings &gt; Finance &gt; Configure Revenue Event Types** in the Zuora UI.  Cannot be duplicated.  | [optional] [default to null]
**Notes** | **string** | Additional information about this record.  | [optional] [default to null]
**RevenueDistributions** | [**[]PostDistributionItemType**](POSTDistributionItemType.md) | An array of revenue distributions. Represents how you want to distribute revenue for this revenue schedule. You can distribute revenue into a maximum of 250 accounting periods with one revenue schedule.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


