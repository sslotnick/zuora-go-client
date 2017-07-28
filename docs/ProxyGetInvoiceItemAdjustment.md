# ProxyGetInvoiceItemAdjustment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  The ID of the account that owns the invoice. **Values**: inherited from &#x60;Account.ID&#x60; for the invoice owner  | [optional] [default to null]
**AccountingCode** | **string** |  The accounting code for the invoice item. Accounting codes group transactions that contain similar accounting attributes. **Character limit**: 100 **Values**: inherited from &#x60;InvoiceItem.AccountingCode&#x60;  | [optional] [default to null]
**AdjustmentDate** | [**time.Time**](time.Time.md) |  The date when the invoice item adjustment is applied, in &#x60;yyyy-mm-dd&#x60; format. This date must be the same as the invoice&#39;s date or later. **Character limit**: 29  | [optional] [default to null]
**AdjustmentNumber** | **string** |  A unique string to identify an individual invoice item adjustment. **Character limit**: 255 **Values**: automatically generated  | [optional] [default to null]
**Amount** | **float64** |  The amount of the invoice item adjustment. The value of Amount must be positive. Use the required parameter Type to either credit or charge (debit) this amount on the invoice. **Character limit**: 16 **Values**: a valid currency amount  | [optional] [default to null]
**CancelledById** | **string** |  The ID of the Zuora user who canceled the invoice item adjustment. Zuora generates this read-only field only if the adjustment is canceled. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CancelledDate** | [**time.Time**](time.Time.md) |  The date when the invoice item adjustment is canceled. Zuora generates this read-only field if this adjustment is canceled. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**Comment** | **string** |  Use this field to record comments about the invoice item adjustment. **Character limit**: 255 **Values**: a string of 255 characters or fewer  | [optional] [default to null]
**CreatedById** | **string** |  The user ID of the person who created the invoice item. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date the invoice item was created. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**CustomerName** | **string** |  The name of the account that owns the associated invoice.  **Character limit**: 50  **Values**: inherited from &#x60;Account.Name&#x60;  **Note**: This value changes if &#x60;Account.Name&#x60; is updated. The values of &#x60;UpdatedById&#x60; and &#x60;UpdatedDate&#x60; for the &#x60;InvoiceItemAdjustment&#x60; do not change when &#x60;Account.Name&#x60; is updated.  | [optional] [default to null]
**CustomerNumber** | **string** |  The unique account number of the customer&#39;s account.  **Character limit**: 50  **Values**: inherited from &#x60;Account.AccountNumber&#x60;  **Note**: This value changes if &#x60;Account.AccountNumber&#x60; is updated. The values of &#x60;UpdatedById&#x60; and &#x60;UpdatedDate&#x60; for the &#x60;InvoiceItemAdjustment&#x60; do not change when &#x60;Account.AccountNumber&#x60; is updated.  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**InvoiceId** | **string** |  The ID of the invoice associated with the adjustment. The adjustment invoice item is in this invoice. This field is optional if you specify a value for the &#x60;InvoiceNumber&#x60; field. **Character limit**: 3 **Values**: a valid invoice ID  | [optional] [default to null]
**InvoiceItemName** | **string** |  The name of the invoice item&#39;s charge. This field is required in the Query call, but is inherited in other calls. **Character limit**: 255 **Values**: inherited from &#x60;InvoiceItem.ChargeName&#x60;  | [optional] [default to null]
**InvoiceNumber** | **string** |  The unique identification number for the invoice that contains the invoice item. This field is optional if you specify a value for the &#x60;InvoiceId&#x60; field. **Character limit**: 32 **Values**: a valid invoice number  | [optional] [default to null]
**ReasonCode** | **string** |  A code identifying the reason for the transaction. Must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code. **Character limit**: 32 **V****alues**: a valid reason code  | [optional] [default to null]
**ReferenceId** | **string** |  A code to reference an object external to Zuora. For example, you can use this field to reference a case number in an external system. **Character limit**: 60 **Values**: a string of 60 characters or fewer  | [optional] [default to null]
**ServiceEndDate** | [**time.Time**](time.Time.md) |  The end date of the service period associated with the invoice item. Service ends one second before the date in this value.  **Character limit**: 29  | [optional] [default to null]
**ServiceStartDate** | [**time.Time**](time.Time.md) |  The start date of the service period associated with the invoice item. Service ends one second before the date in this value.  **Character limit**: 29  | [optional] [default to null]
**SourceId** | **string** |  The ID of the item specified in the SourceType field. **Character limit**: 32 **Values**: a valid invoice item ID or taxation item ID  | [optional] [default to null]
**SourceType** | **string** |  The type of adjustment. **Character limit**: 13 **Values**: InvoiceDetail, Tax  | [optional] [default to null]
**Status** | **string** |  The status of the invoice item adjustment. This field is required in the Query call, but is automatically generated in other calls. **Character limit**: 9 **Values**: Canceled, Processed  | [optional] [default to null]
**TransferredToAccounting** | **string** |  Indicates the status of the adjustment&#39;s transfer to an external accounting system, such as NetSuite. **Character limit**: 10 **Values**: Processing, Yes, Error, Ignore  | [optional] [default to null]
**Type_** | **string** |  Query Filter  | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the user who last updated the invoice item. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the invoice item was last updated. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


