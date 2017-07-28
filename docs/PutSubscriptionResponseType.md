# PutSubscriptionResponseType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Invoice amount. Preview mode only.  | [optional] [default to null]
**AmountWithoutTax** | **string** | Invoice amount minus tax. Preview mode only.  | [optional] [default to null]
**ChargeMetrics** | [**PutSubscriptionResponseTypeChargeMetrics**](PUTSubscriptionResponseType_chargeMetrics.md) |  | [optional] [default to null]
**CreditMemo** | [**PutSubscriptionResponseTypeCreditMemo**](PUTSubscriptionResponseType_creditMemo.md) |  | [optional] [default to null]
**Invoice** | [**interface{}**](interface{}.md) | Container for invoices.    **Note:** This field is only available if you set the Zuora REST API minor version to 207.0 or later in the request header. Also, the response structure is changed and the following invoice related response fields are moved to this **invoice** container:       * amount    * amountWithoutTax    * taxAmount    * invoiceItems    * targetDate    * chargeMetrics  | [optional] [default to null]
**InvoiceId** | **string** | Invoice ID, if an invoice is generated during the update.  | [optional] [default to null]
**InvoiceItems** | [**[]PutSubscriptionPreviewInvoiceItemsType**](PUTSubscriptionPreviewInvoiceItemsType.md) | Container for invoice items.  | [optional] [default to null]
**InvoiceTargetDate** | [**time.Time**](time.Time.md) | Date through which charges are calculated on the invoice, as yyyy-mm-dd. Preview mode only.  **Note:** This field is only available if you do not specify the Zuora REST API minor version or specify the minor version to 186.0, 187.0, 188.0, 189.0, and 196.0. See [Zuora REST API Versions](https://www.zuora.com/developer/api-reference/#section/API-Versions) for more information.  | [optional] [default to null]
**PaidAmount** | **string** | Payment amount, if a payment is collected  | [optional] [default to null]
**PaymentId** | **string** | Payment ID, if a payment is collected.  | [optional] [default to null]
**PreviewChargeMetricsResponse** | **string** |  | [optional] [default to null]
**SubscriptionId** | **string** | The ID of the resulting new subscription.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**TargetDate** | [**time.Time**](time.Time.md) | Date through which to calculate charges if an invoice is generated, as yyyy-mm-dd. Default is current date.  **Note:** This field is only available if you set the Zuora REST API minor version to 207.0 or later in the request header. See [Zuora REST API Versions](https://www.zuora.com/developer/api-reference/#section/API-Versions) for more information.  | [optional] [default to null]
**TaxAmount** | **string** | Tax amount on the invoice.  | [optional] [default to null]
**TotalDeltaMrr** | **string** | Change in the subscription monthly recurring revenue as a result of the update.  | [optional] [default to null]
**TotalDeltaTcv** | **string** | Change in the total contracted value of the subscription as a result of the update.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


