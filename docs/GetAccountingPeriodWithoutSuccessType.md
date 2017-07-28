# GetAccountingPeriodWithoutSuccessType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | **string** | ID of the user who created the accounting period.  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | Date and time when the accounting period was created.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The end date of the accounting period.  | [optional] [default to null]
**FileIds** | [**[]GetAccountingPeriodFileIdsType**](GETAccountingPeriodFileIdsType.md) | File IDs of the reports available for the accounting period. You can retrieve the reports by specifying the file ID in a [Get Files](https://knowledgecenter.zuora.com/DC_Developers/REST_API/B_REST_API_reference/Get_Files) REST API call.  | [optional] [default to null]
**FiscalYear** | **string** | Fiscal year of the accounting period.  | [optional] [default to null]
**FiscalQuarter** | **int64** |  | [optional] [default to null]
**Id** | **string** | ID of the accounting period.  | [optional] [default to null]
**Name** | **string** | Name of the accounting period.  | [optional] [default to null]
**Notes** | **string** | Any optional notes about the accounting period.  | [optional] [default to null]
**RunTrialBalanceEnd** | [**time.Time**](time.Time.md) | Date and time that the trial balance was completed. If the trial balance status is &#x60;Pending&#x60;, &#x60;Processing&#x60;, or &#x60;Error&#x60;, this field is &#x60;null&#x60;.  | [optional] [default to null]
**RunTrialBalanceErrorMessage** | **string** | If trial balance status is Error, an error message is returned in this field.  | [optional] [default to null]
**RunTrialBalanceStart** | [**time.Time**](time.Time.md) | Date and time that the trial balance was run. If the trial balance status is &#x60;Pending&#x60;, this field is &#x60;null&#x60;.  | [optional] [default to null]
**RunTrialBalanceStatus** | **string** | Status of the trial balance for the accounting period. Possible values:  * &#x60;Pending&#x60; * &#x60;Processing&#x60; * &#x60;Completed&#x60; * &#x60;Error&#x60;  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The start date of the accounting period.  | [optional] [default to null]
**Status** | **string** | Status of the accounting period. Possible values:  * &#x60;Open&#x60; * &#x60;PendingClose&#x60; * &#x60;Closed&#x60;  | [optional] [default to null]
**UpdatedBy** | **string** | D of the user who last updated the accounting period.  | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) | Date and time when the accounting period was last updated.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


