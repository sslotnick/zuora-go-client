# ProxyGetCreditBalanceAdjustment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  The account ID of the credit balance&#39;s account. Zuora generates this value from the source transaction. **Character limit**: 32 **V****alues**: automatically generated from:  - CreditBalanceAdjustment.SourceTransactionId or - CreditBalanceAdjustment.SourceTransactionNumber  | [optional] [default to null]
**AccountingCode** | **string** |  The Chart of Accounts  | [optional] [default to null]
**AdjustmentDate** | [**time.Time**](time.Time.md) |  The date when the credit balance adjustment is applied. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**Amount** | **float64** |  The amount of the adjustment. **Character limit**: 16 **Values**: a valid currency amount  | [optional] [default to null]
**CancelledOn** | [**time.Time**](time.Time.md) |  The date when the credit balance adjustment was canceled. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**Comment** | **string** |  Use this field to record comments about the credit balance adjustment. **Character limit**: 255 **Values**: a string of 255 characters or fewer  | [optional] [default to null]
**CreatedById** | **string** |  The user ID of the person who created the credit balance adjustment. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the credit balance adjustmentwas generated. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**IntegrationIdNS** | **string** |  | [optional] [default to null]
**IntegrationStatusNS** | **string** |  | [optional] [default to null]
**Number** | **string** |  A unique identifier for the credit balance adjustment. Zuora generates this number in the format, &lt;em&gt;CBA-xxxxxxxx&lt;/em&gt;, such as CBA-00375919. **Character limit**: 255 **Values**: automatically generated  | [optional] [default to null]
**ReasonCode** | **string** |  A code identifying the reason for the transaction. Must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code. **Character limit**: 32 **V****alues**: a valid reason code  | [optional] [default to null]
**ReferenceId** | **string** |  The ID of the payment that the credit balance adjustment is for. **Character limit**: 32 **Values**: a string of 60 characters or fewer  | [optional] [default to null]
**SourceTransactionId** | **string** |  The ID of the object that the credit balance adjustment is applied to. You must specify a value for either the &#x60;SourceTransactionId&#x60; field or the &#x60;SourceTransactionNumber&#x60; field. **Character limit**: 32 **Values**: one of the following:  - InvoiceId - PaymentId - RefundId  | [optional] [default to null]
**SourceTransactionNumber** | **string** |  The number of the object that the credit balance adjustment is applied to. You must specify a value for either the &#x60;SourceTransactionId&#x60; field or the &#x60;SourceTransactionNumber&#x60; field. **Character limit**: 50 **Values**: one of the following:  - InvoiceNumber - PaymentNumber - RefundNumber  | [optional] [default to null]
**SourceTransactionType** | **string** |  The source of the credit balance adjustment. **Character limit**: **Values**: automatically generated; one of the following:  - Invoice - Payment - Refund  | [optional] [default to null]
**Status** | **string** |  The status of the credit balance adjustment. **Character limit**: 9 **Values**: automatically generated; one of the following:  - Processed - Canceled  | [optional] [default to null]
**SyncDateNS** | **string** |  | [optional] [default to null]
**TransferredToAccounting** | **string** |  Indicates the status of the credit balance adjustment&#39;s transfer to an external accounting system, such as NetSuite. **Character limit**: **Values**: one of the following:  - Processing - Yes - Error - Ignore  | [optional] [default to null]
**Type_** | **string** | Create Query Filter | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the user who last updated the credit balance adjustment. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the credit balance adjustment was last updated. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


