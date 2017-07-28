# ProxyGetInvoiceSplit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedById** | **string** | The ID of the Zuora user who created the InvoiceSplit object. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date when the InvoiceSplit object was created. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**InvoiceId** | **string** |  The ID of the original invoice that the InvoiceSplit object splits. This field becomes read-only after the InvoiceSplit object is created. **Character limit**: 32 **Values**: a valid invoice ID  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the Zuora user who last updated the invoice split. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date when the invoice split was last updated. **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


