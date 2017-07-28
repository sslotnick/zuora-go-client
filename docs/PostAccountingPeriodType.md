# PostAccountingPeriodType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The end date of the accounting period in yyyy-mm-dd format, for example, \&quot;2016-02-19\&quot;.  | [default to null]
**FiscalYear** | **string** | Fiscal year of the accounting period in yyyy format, for example, \&quot;2016\&quot;.  | [default to null]
**FiscalQuarter** | **int64** |  | [optional] [default to null]
**Name** | **string** | Name of the accounting period.  Accounting period name must be unique. Maximum of 100 characters.  | [default to null]
**Notes** | **string** | Notes about the accounting period.  Maximum of 255 characters.  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The start date of the accounting period in yyyy-mm-dd format, for example, \&quot;2016-02-19\&quot;.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


