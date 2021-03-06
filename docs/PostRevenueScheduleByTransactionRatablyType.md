# PostRevenueScheduleByTransactionRatablyType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldsValues** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**DistributionType** | **string** | How you want to distribute the revenue.  * Daily Distribution: Distributes revenue evenly across each day between the recognitionStart and recognitionEnd dates. * Monthly Distribution (Back Load): Back loads the revenue so you distribute the monthly amount in the partial month in the end only. * Monthly Distribution (Front Load): Front loads the revenue so you distribute the monthly amount in the partial month in the beginning only. * Monthly Distribution (Proration by Days): Splits the revenue amount between the two partial months.  **Note:** To use any of the Monthly Distribution options, you must have the \&quot;Monthly recognition over time\&quot; model enabled in **Settings &gt; Finance &gt; Manage Revenue Recognition Models** in the Zuora UI.  | [optional] [default to null]
**Notes** | **string** | Additional information about this record.  | [optional] [default to null]
**RecognitionEnd** | [**time.Time**](time.Time.md) | The end date of a recognition period in &#x60;yyyy-mm-dd&#x60; format.  The maximum difference between the &#x60;recognitionStart&#x60; and &#x60;recognitionEnd&#x60; date fields is equal to 250 multiplied by the length of an accounting period.  | [default to null]
**RecognitionStart** | [**time.Time**](time.Time.md) | The start date of a recognition period in &#x60;yyyy-mm-dd&#x60; format.  If there is a closed accounting period between the &#x60;recognitionStart&#x60; and &#x60;recognitionEnd&#x60; dates, the revenue that would be placed in the closed accounting period is instead placed in the next open accounting period.  | [default to null]
**RevenueEvent** | [**PostRevenueScheduleByTransactionRatablyTypeRevenueEvent**](POSTRevenueScheduleByTransactionRatablyType_revenueEvent.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


