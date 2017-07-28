# GetChargeRsDetailType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | An account ID.  | [optional] [default to null]
**Amount** | **string** | The revenue schedule amount, which is the sum of all revenue items. This field cannot be null and must be formatted based on the currency, such as *JPY 30* or USD *30.15*. Test out the currency to ensure you are using the proper formatting otherwise, the response will fail and this error message is returned:  *\&quot;Allocation amount with wrong decimal places.\&quot;*  | [optional] [default to null]
**Currency** | **string** | The type of currency used.   | [optional] [default to null]
**Notes** | **string** | Additional information about this record.  | [optional] [default to null]
**Number** | **string** | The charge revenue summary number.  | [optional] [default to null]
**RecognitionRuleName** | **string** | The name of the recognition rule.  | [optional] [default to null]
**RecognizedRevenue** | **string** | The revenue that was distributed in a closed accounting period.  | [optional] [default to null]
**RevenueItems** | [**[]GetRevenueItemType**](GETRevenueItemType.md) | Revenue items are listed in ascending order by the accounting period start date.  | [optional] [default to null]
**SubscriptionChargeId** | **string** | The original subscription charge ID.  | [optional] [default to null]
**SubscriptionId** | **string** | The original subscription ID.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**UndistributedUnrecognizedRevenue** | **string** | Revenue in the open-ended accounting period.  | [optional] [default to null]
**UnrecognizedRevenue** | **string** | Revenue distributed in all open accounting periods, which includes the open-ended accounting period.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


