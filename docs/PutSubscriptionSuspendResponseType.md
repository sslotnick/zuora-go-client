# PutSubscriptionSuspendResponseType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Invoice ID, if an invoice is generated during the subscription process.  | [optional] [default to null]
**PaidAmount** | **string** | Payment amount, if a payment is collected.  | [optional] [default to null]
**PaymentId** | **string** | Payment ID, if a payment is collected.  | [optional] [default to null]
**ResumeDate** | [**time.Time**](time.Time.md) | The date when subscription resumption takes effect, in the format yyyy-mm-dd.  | [optional] [default to null]
**SubscriptionId** | **string** | The subscription ID.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**SuspendDate** | [**time.Time**](time.Time.md) | The date when subscription suspension takes effect, in the format yyyy-mm-dd.  | [optional] [default to null]
**TermEndDate** | [**time.Time**](time.Time.md) | The date when the new subscription term ends, in the format yyyy-mm-dd.  | [optional] [default to null]
**TotalDeltaTcv** | **string** | Change in the total contracted value of the subscription as a result of the update.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


