# PutSpecificDateAllocationType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The revenue schedule amount, which is the sum of all revenue items. This field cannot be null and must be formatted based on the currency, such as &#x60;JPY 30&#x60; or &#x60;USD 30.15&#x60;. Test out the currency to ensure you are using the proper formatting otherwise, the response will fail and this error message is returned: &#x60;Allocation amount with wrong decimal places.&#x60;  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**DistributeOn** | [**time.Time**](time.Time.md) | The recognition date on which to distribute revenue for milestone-based recognition.  Type: date in &#x60;yyyy-mm-dd&#x60; format.  | [default to null]
**DistributionType** | **string** | How to distribute the revenue for milestone-based recognition.  | [default to null]
**EventType** | **string** | Label of the revenue event type. Revenue event type labels can be duplicated. You can configure your revenue event type labels by navigating to **Settings &gt; Finance &gt; Configure Revenue Event Types** in the Zuora UI.  Note that &#x60;Credit Memo Posted&#x60; and &#x60;Debit Memo Posted&#x60; are only available if you enable the Advanced AR Settlement feature.  | [optional] [default to null]
**EventTypeSystemId** | **string** | System ID of the revenue event type. Each eventType has a unique system ID. You can configure your revenue event type system IDs by navigating to **Settings &gt; Finance &gt; Configure Revenue Event Types** in the Zuora UI.  | [optional] [default to null]
**Notes** | **string** | Additional information about this record.  | [optional] [default to null]
**Percentage** | **string** | Percentage of the total recognition amount or total undistributed to distribute.  Required if distributionType is either: * specific date (delta percent total) * specific date (delta percent undistributed)  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


