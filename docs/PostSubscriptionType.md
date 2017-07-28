# PostSubscriptionType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpqBundleJsonIdQT** | **string** |  | [optional] [default to null]
**OpportunityCloseDateQT** | **string** |  | [optional] [default to null]
**OpportunityNameQT** | **string** |  | [optional] [default to null]
**QuoteBusinessTypeQT** | **string** |  | [optional] [default to null]
**QuoteNumberQT** | **string** |  | [optional] [default to null]
**QuoteTypeQT** | **string** |  | [optional] [default to null]
**AccountKey** | **string** | Customer account number or ID  | [default to null]
**ApplyCreditBalance** | **bool** | Applies a credit balance to an invoice.  If the value is &#x60;true&#x60;, the credit balance is applied to the invoice. If the value is &#x60;false&#x60;, no action is taken.   To view the credit balance adjustment, retrieve the details of the invoice using the Get Invoices method.  Prerequisite: &#x60;invoice&#x60; must be &#x60;true&#x60;.   **Note:**    - If you are using the field &#x60;invoiceCollect&#x60; rather than the field &#x60;invoice&#x60;, the &#x60;invoiceCollect&#x60; value must be &#x60;true&#x60;.   - This field is deprecated if you have the Advanced AR Settlement feature enabled.  | [optional] [default to null]
**Collect** | **bool** | Collects an automatic payment for a subscription. The collection generated in this operation is only for this subscription, not for the entire customer account. If the value is &#x60;true&#x60;, the automatic payment is collected. If the value is &#x60;false&#x60;, no action is taken.  The default value is &#x60;true&#x60;.  **Prerequisite:** The invoice field must be &#x60;true&#x60;.   **Note:** This field is in Zuora REST API version control. Supported minor versions are 196.0 or later. To use this field in the method, you must set the &#x60;zuora-version&#x60; field to the minor version number in the request header.  | [optional] [default to null]
**ContractEffectiveDate** | [**time.Time**](time.Time.md) | Effective contract date for this subscription, as yyyy-mm-dd  | [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**CustomerAcceptanceDate** | [**time.Time**](time.Time.md) | The date on which the services or products within a subscription have been accepted by the customer, as yyyy-mm-dd.  Default value is dependent on the value of other fields. See **Notes** section for more details.  | [optional] [default to null]
**InitialTerm** | **int64** | The length of the period for the first subscription term. Default is &#x60;0&#x60;. If &#x60;termType&#x60; is &#x60;TERMED&#x60;, then this field is required, and the value must be greater than &#x60;0&#x60;. If &#x60;termType&#x60; is &#x60;EVERGREEN&#x60;, this field is ignored.  | [optional] [default to null]
**InitialTermPeriodType** | **string** | The period type for the first subscription term.  This field is used with the &#x60;InitialTerm&#x60; field to specify the initial subscription term.  Values are:  * &#x60;Month&#x60; (default) * &#x60;Year&#x60; * &#x60;Day&#x60; * &#x60;Week&#x60;  | [optional] [default to null]
**Invoice** | **bool** | **Note:** This field has been replaced by the &#x60;runBilling&#x60; field. The &#x60;invoice&#x60; field is only available for backward compatibility.   Creates an invoice for a subscription. The invoice generated in this operation is only for this subscription, not for the entire customer account.   If the value is &#x60;true&#x60;, an invoice is created. If the value is &#x60;false&#x60;, no action is taken. The default value is &#x60;true&#x60;.    This field is in Zuora REST API version control. Supported minor versions are &#x60;196.0&#x60; and &#x60;207.0&#x60;. To use this field in the method, you must set the zuora-version parameter to the minor version number in the request header.   | [optional] [default to null]
**InvoiceCollect** | **bool** | **Note:** This field has been replaced by the invoice field and the collect field. invoiceCollect is available only for backward compatibility.  If &#x60;true&#x60; (default), an invoice is generated and payment collected automatically during the subscription process. If &#x60;false&#x60;, no invoicing or payment takes place. The invoice generated in this operation is only for this subscription, not for the entire customer account.  This field is in Zuora REST API version control. Supported minor versions are &#x60;186.0&#x60;, &#x60;187.0&#x60;, &#x60;188.0&#x60;, &#x60;189.0&#x60;, and &#x60;196.0&#x60;.  | [optional] [default to null]
**InvoiceOwnerAccountKey** | **string** | Invoice owner account number or ID.  **Note:** This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**InvoiceSeparately** | **bool** | Separates a single subscription from other subscriptions and invoices the charge independently.   If the value is &#x60;true&#x60;, the subscription is billed separately from other subscriptions. If the value is &#x60;false&#x60;, the subscription is included with other subscriptions in the account invoice.  The default value is &#x60;false&#x60;.  Prerequisite: The default subscription setting Enable Subscriptions to be Invoiced Separately must be set to Yes.  | [optional] [default to null]
**InvoiceTargetDate** | [**time.Time**](time.Time.md) | **Note:** This field has been replaced by the &#x60;targetDate&#x60; field. The &#x60;invoiceTargetDate&#x60; field is only available for backward compatibility.   Date through which to calculate charges if an invoice is generated, as yyyy-mm-dd. Default is current date.   This field is in Zuora REST API version control. Supported minor versions are &#x60;207.0&#x60; and earlier.    | [optional] [default to null]
**Notes** | **string** | String of up to 500 characters.  | [optional] [default to null]
**RenewalSetting** | **string** | Specifies whether a termed subscription will remain termed or change to evergreen when it is renewed.  Values:  * &#x60;RENEW_WITH_SPECIFIC_TERM&#x60; (default) * &#x60;RENEW_TO_EVERGREEN&#x60;  | [optional] [default to null]
**RenewalTerm** | **int64** | The length of the period for the subscription renewal term. Default is &#x60;0&#x60;.  | [default to null]
**RenewalTermPeriodType** | **string** | The period type for the subscription renewal term.  This field is used with the &#x60;renewalTerm&#x60; field to specify the subscription renewal term.  Values are:  * &#x60;Month&#x60; (default) * &#x60;Year&#x60; * &#x60;Day&#x60; * &#x60;Week&#x60;  | [optional] [default to null]
**RunBilling** | **bool** | Creates an invoice for a subscription. If you have the Advanced AR Settlement feature enabled, a credit memo might also be created based on the [invoice and credit memo generation rule](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/Credit_and_Debit_Memos/Rules_for_Generating_Invoices_and_Credit_Memos).     The billing documents generated in this operation is only for this subscription, not for the entire customer account.   Possible values:  - &#x60;true&#x60;: An invoice is created. If you have the Advanced AR Settlement feature enabled, a credit memo might also be created.   - &#x60;false&#x60;: No invoice is created.   **Note:** This field is in Zuora REST API version control. Supported minor versions are &#x60;211.0&#x60; or later. To use this field in the method, you must set the &#x60;zuora-version&#x60; parameter to the minor version number in the request header.  | [optional] [default to null]
**ServiceActivationDate** | [**time.Time**](time.Time.md) | The date on which the services or products within a subscription have been activated and access has been provided to the customer, as yyyy-mm-dd.  Default value is dependent on the value of other fields. See **Notes** section for more details.  | [optional] [default to null]
**SubscribeToRatePlans** | [**[]PostSrpCreateType**](POSTSrpCreateType.md) | Container for one or more rate plans for this subscription.  | [default to null]
**SubscriptionNumber** | **string** | Subscription Number. The value can be up to 1000 characters.  If you do not specify a subscription number when creating a subscription, Zuora will generate a subscription number automatically.  If the account is created successfully, the subscription number is returned in the &#x60;subscriptionNumber&#x60; response field.  | [optional] [default to null]
**TargetDate** | [**time.Time**](time.Time.md) | Date through which to calculate charges if an invoice or a credit memo is generated, as yyyy-mm-dd. Default is current date.   **Note:** The credit memo is only available if you have the Avdanced AR Settlement feature enabled.   This field is in Zuora REST API version control. Supported minor versions are &#x60;211.0&#x60; and later. To use this field in the method, you must set the  &#x60;zuora-version&#x60; parameter to the minor version number in the request header.  | [optional] [default to null]
**TermStartDate** | [**time.Time**](time.Time.md) | The date on which the subscription term begins, as yyyy-mm-dd. If this is a renewal subscription, this date is different from the subscription start date.  | [optional] [default to null]
**TermType** | **string** | Possible values are: &#x60;TERMED&#x60;, &#x60;EVERGREEN&#x60;.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

