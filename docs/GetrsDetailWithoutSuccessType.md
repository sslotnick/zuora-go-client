# GetrsDetailWithoutSuccessType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | An account ID.  | [optional] [default to null]
**Amount** | **string** | The revenue schedule amount, which is the sum of all revenue items. This field cannot be null and must be formatted based on the currency, such as &#x60;JPY 30&#x60; or &#x60;USD 30.15&#x60;. Test out the currency to ensure you are using the proper formatting otherwise, the response will fail and this error message is returned: &#x60;Allocation amount with wrong decimal places.&#x60;  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | The date and time when the record was created, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**Currency** | **string** | The type of currency used.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**LinkedTransactionId** | **string** | The linked transaction ID for billing transactions. This field is used for all rules except for the custom unlimited or manual recognition rule models. If using the custom unlimited rule model, then the field value must be null. If the field is not null, then the referenceId field must be null.  | [optional] [default to null]
**LinkedTransactionNumber** | **string** | The number for the linked invoice item or invoice item adjustment transaction. This field is used for all rules except for the custom unlimited or manual recognition rule models. If using the custom unlimited or manual recognition rule models, then the field value is null.  | [optional] [default to null]
**LinkedTransactionType** | **string** | The type of linked transaction for billing transactions, which can be invoice item or invoice item adjustment. This field is used for all rules except for the custom unlimited or manual recognition rule models.  | [optional] [default to null]
**Notes** | **string** | Additional information about this record.  | [optional] [default to null]
**Number** | **string** | Revenue schedule number. The revenue schedule number is always prefixed with \&quot;RS\&quot;, for example, \&quot;RS-00000001\&quot;.  | [optional] [default to null]
**RecognitionRuleName** | **string** | The name of the recognition rule.  | [optional] [default to null]
**RecognizedRevenue** | **string** | The revenue that was distributed in a closed accounting period.  | [optional] [default to null]
**ReferenceId** | **string** | Reference ID is used only in the custom unlimited rule to create a revenue schedule. In this scenario, the revenue schedule is not linked to an invoice item or invoice item adjustment.  | [optional] [default to null]
**RevenueItems** | [**[]GetRsRevenueItemType**](GETRsRevenueItemType.md) | Revenue items are listed in ascending order by the accounting period start date.  | [optional] [default to null]
**RevenueScheduleDate** | [**time.Time**](time.Time.md) | The effective date of the revenue schedule. For example, the revenue schedule date for bookings-based revenue recognition is typically set to the order date or contract date.  The date cannot be in a closed accounting period. The date must be in the &#x60;yyyy-mm-dd&#x60; format.  | [optional] [default to null]
**SubscriptionChargeId** | **string** | The original subscription charge ID.  | [optional] [default to null]
**SubscriptionId** | **string** | The original subscription ID.  | [optional] [default to null]
**UndistributedUnrecognizedRevenue** | **string** | Revenue in the open-ended accounting period.  | [optional] [default to null]
**UnrecognizedRevenue** | **string** | Revenue distributed in all open accounting periods, which includes the open-ended accounting period.  | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) | The date when the revenue automation start date was set, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


