# PostSubscriptionCancellationType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpportunityCloseDateQT** | **string** | The closing date of the Opportunity. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**OpportunityNameQT** | **string** | The unique identifier of the Opportunity. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**QuoteBusinessTypeQT** | **string** | The specific identifier for the type of business transaction the Quote represents such as New, Upsell, Downsell, Renewal, or Churn. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**QuoteNumberQT** | **string** | The unique identifier of the Quote. This field is populated when the subscription originates from Zuora Quotes.  This field is used only for reporting subscription metrics.    | [optional] [default to null]
**QuoteTypeQT** | **string** | The Quote type that represents the subscription lifecycle stage such as New, Amendment, Renew or Cancel. This field is populated when the subscription originates from Zuora Quotes.   | [optional] [default to null]
**ApplyCreditBalance** | **bool** | Applies a credit balance to an invoice.  If the value is &#x60;true&#x60;, the credit balance is applied to the invoice. If the value is &#x60;false&#x60;, no action is taken.   To view the credit balance adjustment, retrieve the details of the invoice using the Get Invoices method.  Prerequisite: &#x60;invoice&#x60; must be &#x60;true&#x60;.   **Note:**    - If you are using the field &#x60;invoiceCollect&#x60; rather than the field &#x60;invoice&#x60;, the &#x60;invoiceCollect&#x60; value must be &#x60;true&#x60;.   - This field is deprecated if you have the Advanced AR Settlement feature enabled.  | [optional] [default to null]
**CancellationEffectiveDate** | [**time.Time**](time.Time.md) | Date the cancellation takes effect, in the format yyyy-mm-dd.  Use only if &#x60;cancellationPolicy&#x60; is &#x60;SpecificDate&#x60;. Should not be earlier than the subscription contract-effective date, later than the subscription term-end date, or within a period for which the customer has been invoiced.  | [optional] [default to null]
**CancellationPolicy** | **string** | Cancellation method. Possible values are: &#x60;EndOfCurrentTerm&#x60;, &#x60;EndOfLastInvoicePeriod&#x60;, &#x60;SpecificDate&#x60;. If using &#x60;SpecificDate&#x60;, the &#x60;cancellationEffectiveDate&#x60; field is required.  | [default to null]
**Collect** | **bool** | Collects an automatic payment for a subscription. The collection generated in this operation is only for this subscription, not for the entire customer account.  If the value is &#x60;true&#x60;, the automatic payment is collected. If the value is &#x60;false&#x60;, no action is taken.  The default value is &#x60;false&#x60;.  This field is in Zuora REST API version control. Supported minor versions are 196.0 or later. To use this field in the method, you must set the &#x60;zuora-version&#x60; parameter to the minor version number in the request header.   Prerequisite: &#x60;invoice&#x60; must be &#x60;true&#x60;.   | [optional] [default to null]
**Invoice** | **bool** | **Note:** This field has been replaced by the &#x60;runBilling&#x60; field. The &#x60;invoice&#x60; field is only available for backward compatibility.   Creates an invoice for a subscription. The invoice generated in this operation is only for this subscription, not for the entire customer account.   If the value is &#x60;true&#x60;, an invoice is created. If the value is &#x60;false&#x60;, no action is taken. The default value is &#x60;false&#x60;.    This field is in Zuora REST API version control. Supported minor versions are &#x60;196.0&#x60; and &#x60;207.0&#x60;. To use this field in the method, you must set the zuora-version parameter to the minor version number in the request header.   | [optional] [default to null]
**InvoiceCollect** | **bool** | This field has been replaced by the &#x60;invoice&#x60; field and the &#x60;collect&#x60; field. &#x60;invoiceCollect&#x60; is available only for backward compatibility.  If &#x60;true&#x60;, an invoice is generated and payment automatically collected. Default is &#x60;false&#x60;.  This field is in Zuora REST API version control. Supported minor versions are &#x60;186.0&#x60;, &#x60;187.0&#x60;, &#x60;188.0&#x60;, &#x60;189.0&#x60;, and &#x60;196.0&#x60;.  | [default to null]
**InvoiceTargetDate** | [**time.Time**](time.Time.md) | **Note:** This field has been replaced by the &#x60;targetDate&#x60; field. The &#x60;invoiceTargetDate&#x60; field is only available for backward compatibility.   Date through which to calculate charges if an invoice is generated, as yyyy-mm-dd. Default is current date.   This field is in Zuora REST API version control. Supported minor versions are &#x60;207.0&#x60; and earlier.    | [optional] [default to null]
**RunBilling** | **bool** | Creates an invoice for a subscription. If you have the Advanced AR Settlement feature enabled, a credit memo might also be created based on the [invoice and credit memo generation rule](https://knowledgecenter.zuora.com/CB_Billing/Advanced_AR_Settlement/Credit_and_Debit_Memos/Rules_for_Generating_Invoices_and_Credit_Memos).     The billing documents generated in this operation is only for this subscription, not for the entire customer account.   Possible values:  - &#x60;true&#x60;: An invoice is created. If you have the Advanced AR Settlement feature enabled, a credit memo might also be created.   - &#x60;false&#x60;: No invoice is created.   **Note:** This field is in Zuora REST API version control. Supported minor versions are &#x60;211.0&#x60; or later. To use this field in the method, you must set the &#x60;zuora-version&#x60; parameter to the minor version number in the request header.  | [optional] [default to null]
**TargetDate** | [**time.Time**](time.Time.md) | Date through which to calculate charges if an invoice or a credit memo is generated, as yyyy-mm-dd. Default is current date.   **Note:** The credit memo is only available if you have the Avdanced AR Settlement feature enabled.   This field is in Zuora REST API version control. Supported minor versions are &#x60;211.0&#x60; and later. To use this field in the method, you must set the  &#x60;zuora-version&#x60; parameter to the minor version number in the request header.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

