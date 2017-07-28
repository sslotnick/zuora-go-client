# GetSubscriptionType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpqBundleJsonIdQT** | **string** |  | [optional] [default to null]
**OpportunityCloseDateQT** | **string** | The closing date of the Opportunity. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**OpportunityNameQT** | **string** | The unique identifier of the Opportunity. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**QuoteBusinessTypeQT** | **string** | The specific identifier for the type of business transaction the Quote represents such as New, Upsell, Downsell, Renewal, or Churn. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**QuoteNumberQT** | **string** | The unique identifier of the Quote. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**QuoteTypeQT** | **string** | The Quote type that represents the subscription lifecycle stage such as New, Amendment, Renew or Cancel. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**AccountId** | **string** |  | [optional] [default to null]
**AccountName** | **string** |  | [optional] [default to null]
**AccountNumber** | **string** |  | [optional] [default to null]
**AutoRenew** | **bool** | If &#x60;true&#x60;, the subscription automatically renews at the end of the term. Default is &#x60;false&#x60;.  | [optional] [default to null]
**ContractEffectiveDate** | [**time.Time**](time.Time.md) | Effective contract date for this subscription, as yyyy-mm-dd.  | [optional] [default to null]
**ContractedMrr** | **string** | Monthly recurring revenue of the subscription.  | [optional] [default to null]
**CurrentTerm** | **int64** | The length of the period for the current subscription term.  | [optional] [default to null]
**CurrentTermPeriodType** | **string** | The period type for the current subscription term.  Values are:  * &#x60;Month&#x60; (default) * &#x60;Year&#x60; * &#x60;Day&#x60; * &#x60;Week&#x60;  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**CustomerAcceptanceDate** | [**time.Time**](time.Time.md) | The date on which the services or products within a subscription have been accepted by the customer, as yyyy-mm-dd.  | [optional] [default to null]
**Id** | **string** | Subscription ID.  | [optional] [default to null]
**InitialTerm** | **int64** | The length of the period for the first subscription term.  | [optional] [default to null]
**InitialTermPeriodType** | **string** | The period type for the first subscription term.  Values are:  * &#x60;Month&#x60; (default) * &#x60;Year&#x60; * &#x60;Day&#x60; * &#x60;Week&#x60;  | [optional] [default to null]
**InvoiceOwnerAccountId** | **string** |  | [optional] [default to null]
**InvoiceOwnerAccountName** | **string** |  | [optional] [default to null]
**InvoiceOwnerAccountNumber** | **string** |  | [optional] [default to null]
**InvoiceSeparately** | **string** | Separates a single subscription from other subscriptions and creates an invoice for the subscription.   If the value is &#x60;true&#x60;, the subscription is billed separately from other subscriptions. If the value is &#x60;false&#x60;, the subscription is included with other subscriptions in the account invoice.  | [optional] [default to null]
**Notes** | **string** | A string of up to 65,535 characters.  | [optional] [default to null]
**RatePlans** | [**[]GetSubscriptionRatePlanType**](GETSubscriptionRatePlanType.md) | Container for rate plans.  | [optional] [default to null]
**RenewalSetting** | **string** | Specifies whether a termed subscription will remain &#x60;TERMED&#x60; or change to &#x60;EVERGREEN&#x60; when it is renewed.   Values are:  * &#x60;RENEW_WITH_SPECIFIC_TERM&#x60; (default) * &#x60;RENEW_TO_EVERGREEN&#x60;  | [optional] [default to null]
**RenewalTerm** | **int64** | The length of the period for the subscription renewal term.  | [optional] [default to null]
**RenewalTermPeriodType** | **string** | The period type for the subscription renewal term.  Values are:  * &#x60;Month&#x60; (default) * &#x60;Year&#x60; * &#x60;Day&#x60; * &#x60;Week&#x60;  | [optional] [default to null]
**ServiceActivationDate** | [**time.Time**](time.Time.md) | The date on which the services or products within a subscription have been activated and access has been provided to the customer, as yyyy-mm-dd  | [optional] [default to null]
**Status** | **string** | Subscription status; possible values are:  * &#x60;Draft&#x60; * &#x60;PendingActivation&#x60; * &#x60;PendingAcceptance&#x60; * &#x60;Active&#x60; * &#x60;Cancelled&#x60; * &#x60;Suspended&#x60; (This value is in Limited Availability.)  | [optional] [default to null]
**SubscriptionNumber** | **string** |  | [optional] [default to null]
**SubscriptionStartDate** | [**time.Time**](time.Time.md) | Date the subscription becomes effective.  | [optional] [default to null]
**TermEndDate** | [**time.Time**](time.Time.md) | Date the subscription term ends. If the subscription is evergreen, this is null or is the cancellation date (if one has been set).  | [optional] [default to null]
**TermStartDate** | [**time.Time**](time.Time.md) | Date the subscription term begins. If this is a renewal subscription, this date is different from the subscription start date.  | [optional] [default to null]
**TermType** | **string** | Possible values are: &#x60;TERMED&#x60;, &#x60;EVERGREEN&#x60;.  | [optional] [default to null]
**TotalContractedValue** | **string** | Total contracted value of the subscription.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


