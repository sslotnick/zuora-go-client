# ProxyModifyRefund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCode** | **string** |  A code identifying the reason for the transaction. Must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code. **Character limit**: 32 **V****alues**: a valid reason code  | [optional] [default to null]
**Status** | **string** |  The status of the refund. **Character limit**: 10 **Values**: automatically generated:  - &#x60;Canceled&#x60; - &#x60;Error&#x60; - &#x60;Processed&#x60; - &#x60;Processing&#x60;  | [optional] [default to null]
**TransferredToAccounting** | **string** |  Specifies whether or not the object has been transferred to an external accounting system. Use this field for integrations with accounting systems such as NetSuite. **Character limit**: 10 **Values**: automatically generated:  - &#x60;Processing&#x60; - &#x60;Yes&#x60; - &#x60;Error&#x60; - &#x60;Ignore&#x60;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


