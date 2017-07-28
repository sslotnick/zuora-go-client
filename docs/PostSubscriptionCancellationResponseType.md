# PostSubscriptionCancellationResponseType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelledDate** | [**time.Time**](time.Time.md) | The date that the subscription was canceled.  | [optional] [default to null]
**InvoiceId** | **string** | ID of the invoice, if one is generated.  | [optional] [default to null]
**PaidAmount** | **string** | Amount paid.  | [optional] [default to null]
**PaymentId** | **string** | ID of the payment, if a payment is collected.  | [optional] [default to null]
**SubscriptionId** | **string** | The subscription ID.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**TotalDeltaMrr** | **string** | Change in the subscription monthly recurring revenue as a result of the update.  | [optional] [default to null]
**TotalDeltaTcv** | **string** | Change in the total contracted value of the subscription as a result of the update.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


