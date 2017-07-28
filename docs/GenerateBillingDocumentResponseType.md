# GenerateBillingDocumentResponseType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditMemos** | [**[]CreditMemoResponseType**](CreditMemoResponseType.md) | Container for generated credit memos.  **Note:** This container is only available if you have the Advanced AR Settlement feature enabled. The Advanced AR Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**Invoices** | [**[]InvoiceResponseType**](InvoiceResponseType.md) | Container for generated invoics.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


