# ProxyGetInvoicePayment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** |  The amount of the payment. **Character limit**: 16 **Values**: a valid currency amount  | [optional] [default to null]
**CreatedById** | **string** |  The user ID of the person who created the invoice payment. **Character limit**: 32 **V****alues**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the invoice payment was generated. **Character limit**: 29 **V****alues**: automatically generated  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**InvoiceId** | **string** |  The unique ID of the invoice associated with this invoice payment. **Character limit**: 32 **Values**: a valid invoice ID  | [optional] [default to null]
**PaymentId** | **string** |  The unique ID of the payment associated with this invoice payment. **Character limit**: 32 **V****alues**: a valid payment ID  | [optional] [default to null]
**RefundAmount** | **float64** | Specifies the amount of a refund applied against this InvoicePayment. **Character limit**: 16 **Values**: automatically generated  | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the user who last updated the invoice payment. **Character limit**: 32 **V****alues**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the invoice payment was last updated. **Character limit**: 29 **V****alues**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


