# GetAccountSummarySubscriptionType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpqBundleJsonIdQT** | **string** |  | [optional] [default to null]
**OpportunityCloseDateQT** | **string** |  | [optional] [default to null]
**OpportunityNameQT** | **string** |  | [optional] [default to null]
**QuoteBusinessTypeQT** | **string** |  | [optional] [default to null]
**QuoteNumberQT** | **string** |  | [optional] [default to null]
**QuoteTypeQT** | **string** |  | [optional] [default to null]
**AutoRenew** | **bool** | If &#x60;true&#x60;, auto-renew is enabled. If &#x60;false&#x60;, auto-renew is disabled.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Id** | **string** | Subscription ID.  | [optional] [default to null]
**InitialTerm** | **string** | Duration of the initial subscription term in whole months.   | [optional] [default to null]
**RatePlans** | [**[]GetAccountSummarySubscriptionRatePlanType**](GETAccountSummarySubscriptionRatePlanType.md) | Container for rate plans for this subscription.  | [optional] [default to null]
**RenewalTerm** | **string** | Duration of the renewal term in whole months.  | [optional] [default to null]
**Status** | **string** | Subscription status; possible values are: &#x60;Draft&#x60;, &#x60;PendingActivation&#x60;, &#x60;PendingAcceptance&#x60;, &#x60;Active&#x60;, &#x60;Cancelled&#x60;, &#x60;Expired&#x60;.  | [optional] [default to null]
**SubscriptionNumber** | **string** | Subscription Number.  | [optional] [default to null]
**SubscriptionStartDate** | [**time.Time**](time.Time.md) | Subscription start date.  | [optional] [default to null]
**TermEndDate** | [**time.Time**](time.Time.md) | End date of the subscription term. If the subscription is evergreen, this is either null or equal to the cancellation date, as appropriate.  | [optional] [default to null]
**TermStartDate** | [**time.Time**](time.Time.md) | Start date of the subscription term. If this is a renewal subscription, this date is different than the subscription start date.  | [optional] [default to null]
**TermType** | **string** | Possible values are: &#x60;TERMED&#x60;, &#x60;EVERGREEN&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


