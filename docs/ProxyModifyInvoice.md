# ProxyModifyInvoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  The status of the invoice in the system. This status is not the status of the payment of the invoice, just the status of the invoice itself. **Character limit**: 8 **Values**: one of the following:  -  Draft (default, automatically set upon invoice creation)  -  Posted  -  Canceled   | [optional] [default to null]
**TransferredToAccounting** | **string** |  Specifies whether or not the invoice was transferred to an external accounting system, such as NetSuite. **Character limit**: 10 **Values**: Processing, Yes, Error, Ignore  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


