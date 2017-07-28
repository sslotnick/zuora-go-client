# ProxyGetUsage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  The ID of the account associated with the usage data. This field is required if no value is specified for the &#x60;AccountNumber&#x60; field. **Character limit**: 32 **Values**: a valid account ID  | [optional] [default to null]
**AccountNumber** | **string** |  The number of the account associated with the usage data. This field is required if no value is specified for the &#x60;AccountId&#x60; field. **Character limit**: 50 **Values**: a valid account number  | [optional] [default to null]
**ChargeId** | **string** |  The OrginalId of the rate plan charge related to the usage record, e.g., &#x60;2c9081a03c63c94c013c6873357a0117&#x60; **Character limit**: 32 **Values**: a valid rate plan charge OriginalID  | [optional] [default to null]
**CreatedById** | **string** |  The user ID of the person who uploaded the usage records. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the invoice was generated. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**EndDateTime** | [**time.Time**](time.Time.md) |  The end date and time of a range of time when usage is tracked. Use this field for reporting; this field doesn&#39;t affect usage calculation. **Character limit**: 29 **Values**: a valid date and time value  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**Quantity** | **float64** |  Indicates the number of units used. **Character limit**: 16 **Values**: a valid decimal amount equal to or greater than 0  | [optional] [default to null]
**RbeStatus** | **string** |  Indicates if the rating and billing engine (RBE) processed usage data for an invoice. **Character limit**: 9 **Values**: automatically generated to be one of the following values: &#x60;Importing&#x60;, &#x60;Pending&#x60;, &#x60;Processed&#x60;  | [optional] [default to null]
**SourceType** | **string** |  Indicates if the usage records were imported from the web-based UI or the API. **Character limit**: 6 **Values**: automatically generated to be one of the following values: &#x60;API&#x60;, &#x60;Import&#x60;  | [optional] [default to null]
**StartDateTime** | [**time.Time**](time.Time.md) |  The start date and time of a range of time when usage is tracked. Zuora uses this field value to determine the usage date. Unlike the &#x60;EndDateTime&#x60;, the &#x60;StartDateTime&#x60; field does affect usage calculation. **Character limit**: 29 **Values**: a valid date and time value  | [optional] [default to null]
**SubmissionDateTime** | [**time.Time**](time.Time.md) |  The date when usage was submitted. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**SubscriptionId** | **string** |  The ID of the subscription that contains the fees related to the usage data. **Character limit**: 32 **Values**: a valid subscription ID  | [optional] [default to null]
**UOM** | **string** |  Specifies the units to measure usage. Units of measure are configured in the web-based UI. Your values depend on your configuration in **Billing Settings**. **Character limit**: **Values**: a valid unit of measure  | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the user who last updated the usage upload. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the usage upload was last updated. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


