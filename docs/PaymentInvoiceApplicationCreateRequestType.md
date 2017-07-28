# PaymentInvoiceApplicationCreateRequestType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | The amount of the payment associated with the invoice.  | [default to null]
**InvoiceId** | **string** | The unique ID of the invoice that the payment is created on.  | [optional] [default to null]
**Items** | [**[]PaymentInvoiceApplicationItemCreateRequestType**](PaymentInvoiceApplicationItemCreateRequestType.md) | Container for invoice items.  **Note:** The Invoice Item Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


