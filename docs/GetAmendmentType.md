# GetAmendmentType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenew** | **bool** | Determines whether the subscription is automatically renewed, or whether it expires at the end of the term and needs to be manually renewed.   | [optional] [default to null]
**BaseRatePlanId** | **string** | The rate plan ID on which changes are made. Only the Update or Remove amendment returns a base rate plan ID.  | [optional] [default to null]
**BaseSubscriptionId** | **string** | The ID of the subscription based on which the amendment is created.  | [optional] [default to null]
**Code** | **string** | The amendment code.  | [optional] [default to null]
**ContractEffectiveDate** | [**time.Time**](time.Time.md) | The date when the amendment becomes effective for billing purposes, as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**CurrentTerm** | **int64** | The length of the period for the current subscription term.   | [optional] [default to null]
**CurrentTermPeriodType** | **string** | The period type for the current subscription term. Possible values are:  - Month - Year - Day - Week  | [optional] [default to null]
**CustomerAcceptanceDate** | [**time.Time**](time.Time.md) | The date when the customer accepts the amendment changes to the subscription, as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**Description** | **string** | Description of the amendment.  | [optional] [default to null]
**DestinationAccountId** | **string** | The ID of the account that the subscription is being transferred to.  | [optional] [default to null]
**DestinationInvoiceOwnerId** | **string** | The ID of the invoice that the subscription is being transferred to.  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The date when the amendment changes take effective.   | [optional] [default to null]
**Id** | **string** | The amendment ID.  | [optional] [default to null]
**Name** | **string** | The name of the amendment.  | [optional] [default to null]
**NewRatePlanId** | **string** | The ID of the rate plan charge on which amendment is made. Only the Add or Update amendment returns a new rate plan ID.  | [optional] [default to null]
**NewSubscriptionId** | **string** | The ID of the subscription that the amendment changes.  | [optional] [default to null]
**RenewalSetting** | **string** | Specifies whether a termed subscription will remain termed or change to evergreen when it is renewed. Possible values are:  - RENEW_WITH_SPECIFIC_TERM - RENEW_TO_EVERGREEN  | [optional] [default to null]
**RenewalTerm** | **int64** | The term of renewal for the amended subscription.  | [optional] [default to null]
**RenewalTermPeriodType** | **string** | The period type for the subscription renewal term. Possible values are:  - Month - Year - Day - Week  | [optional] [default to null]
**ResumeDate** | [**time.Time**](time.Time.md) | The date when the subscription resumption takes effect, as &#x60;yyyy-mm-dd&#x60;.  **Note:** This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**ServiceActivationDate** | [**time.Time**](time.Time.md) | The date when service is activated, as &#x60;yyyy-mm-dd&#x60;.  | [optional] [default to null]
**SpecificUpdateDate** | [**time.Time**](time.Time.md) | The date when the Update Product amendment takes effect.  Only for the Update Product amendments if there is already a future-dated Update Product amendment on the subscription.  | [optional] [default to null]
**Status** | **string** | The status of the amendment. Possible values are:  - Draft  - Pending Activation - Pending Acceptance - Completed  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**SuspendDate** | [**time.Time**](time.Time.md) | The date when the subscription suspension takes effect, as &#x60;yyyy-mm-dd&#x60;.  **Note:** This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**TermStartDate** | [**time.Time**](time.Time.md) | The date when the new terms and conditions take effect.  | [optional] [default to null]
**TermType** | **string** | Indicates if the subscription is &#x60;TERMED&#x60; or &#x60;EVERGREEN&#x60;.  | [optional] [default to null]
**Type_** | **string** | Type of the amendment. Possible values are:  - Cancellation - NewProduct - OwnerTransfer - RemoveProduct - Renewal - UpdateProduct - TermsAndConditions  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


