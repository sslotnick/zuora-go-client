# PutRenewSubscriptionResponseType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Invoice ID, if one is generated.  | [optional] [default to null]
**PaidAmount** | **string** | Payment amount, if payment is collected.  | [optional] [default to null]
**PaymentId** | **string** | Payment ID, if payment is collected.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**TermEndDate** | [**time.Time**](time.Time.md) | Date the new subscription term ends, as yyyy-mm-dd.  | [optional] [default to null]
**TermStartDate** | [**time.Time**](time.Time.md) | Date the new subscription term begins, as yyyy-mm-dd.  | [optional] [default to null]
**TotalDeltaMrr** | **string** | Change in the subscription monthly recurring revenue as a result of the update. For a renewal, this is the MRR of the subscription in the new term.  | [optional] [default to null]
**TotalDeltaTcv** | **string** | Change in the total contracted value of the subscription as a result of the update. For a renewal, this is the TCV of the subscription in the new term.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


