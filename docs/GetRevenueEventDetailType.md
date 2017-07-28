# GetRevenueEventDetailType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | An account ID.  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | The date when the record was created in YYYY-MM-DD HH:MM:SS format.  | [optional] [default to null]
**Currency** | **string** | The type of currency used.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EventType** | **string** | Label of the revenue event type. Revenue event type labels can be duplicated. You can configure your revenue event type labels by navigating to **Settings &gt; Finance &gt; Configure Revenue Event Types** in the Zuora UI.  Note that &#x60;Credit Memo Posted&#x60; and &#x60;Debit Memo Posted&#x60; are only available if you enable the Advanced AR Settlement feature.  | [optional] [default to null]
**Notes** | **string** | Additional information about this record.  | [optional] [default to null]
**Number** | **string** | The revenue event number created when a revenue event occurs.  | [optional] [default to null]
**RecognitionEnd** | [**time.Time**](time.Time.md) | The end date of a recognition period in YYYY-MM-DD format.   The maximum difference of the recognitionStart and recognitionEnd date fields is equal to 250 multiplied by the length of an accounting period.  | [optional] [default to null]
**RecognitionStart** | [**time.Time**](time.Time.md) | The start date of a recognition period in YYYY-MM-DD format.  | [optional] [default to null]
**RevenueItems** | [**[]GetRevenueItemType**](GETRevenueItemType.md) | Revenue items are listed in ascending order by the accounting period start date.  | [optional] [default to null]
**SubscriptionChargeId** | **string** | The original subscription charge ID.  | [optional] [default to null]
**SubscriptionId** | **string** | The original subscription ID.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


