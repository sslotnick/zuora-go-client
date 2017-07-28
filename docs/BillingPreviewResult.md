# BillingPreviewResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | ID of the customer account to which the billing preview applies.  | [optional] [default to null]
**CreditMemoItems** | [**[]PostBillingPreviewCreditMemoItem**](POSTBillingPreviewCreditMemoItem.md) | An array of credit memo items returned as the result of the billing preivew request.  **Note:** The credit memo items are only available if you have Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**InvoiceItems** | [**[]PostBillingPreviewInvoiceItem**](POSTBillingPreviewInvoiceItem.md) | An array of invoice items returned as the result of the billing preview request.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


