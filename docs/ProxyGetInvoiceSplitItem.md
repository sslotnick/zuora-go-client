# ProxyGetInvoiceSplitItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedById** | **string** |  The ID of the Zuora user who created the InvoiceSplitItem object. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the InvoiceSplitItem was created. **Values**: automatically generated  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**InvoiceDate** | [**time.Time**](time.Time.md) |  The generation date of the new split invoice, in &#x60;yyyy-mm-dd&#x60; format. **Character limit**: 29  | [optional] [default to null]
**InvoiceId** | **string** |  The new invoice after the split. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**InvoiceSplitId** | **string** |  The ID of the invoice split associated with the individual invoice split item. **Character limit**: 32 **Values**: a valid invoice split ID  | [optional] [default to null]
**PaymentTerm** | **string** |  Indicates when the customer pays the split invoice. **Values**: a valid, active payment term  | [optional] [default to null]
**SplitPercentage** | **float64** |  Specifies the percentage of the original invoice that you want to be the balance of the split invoice. The total of the SplitPercentage field values for all of the InvoiceSplitItem objects in an InvoiceSplit object must equal 100. **Values**:  | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the Zuora user who last updated the invoice split. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the invoice split was last updated. **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


