# ProxyModifyPayment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  The unique account ID for the customer that the payment is for. **Character limit**: 32 **Values**: a valid account ID  | [optional] [default to null]
**AccountingCode** | **string** |  The Chart of Accounts  | [optional] [default to null]
**Amount** | **float64** |  The amount of the payment. **Character limit**: 16 **Values**: a valid currency value  | [optional] [default to null]
**Comment** | **string** |  Additional information related to the payment. **Character limit**: 255 **Values**: a string of 255 characters or fewer  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) |  The date when the payment takes effect, in &#x60;yyyy-mm-dd&#x60; format. **Character limit**: 29  | [optional] [default to null]
**PaymentMethodId** | **string** |  The ID of the payment method used for the payment. Required for Create. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**ReferenceId** | **string** |  The transaction ID returned by the payment gateway. Use this field to reconcile payments between your gateway and Zuora Payments. **Character limit**: 60 **Values**: a string of 60 characters or fewer  | [optional] [default to null]
**Status** | **string** |  The status of the payment in Zuora. The value depends on the type of payment. **Character limit**: 11 **Values**: one of the following:  -  Electronic payments: &#x60;Processed&#x60;, &#x60;Error&#x60;, &#x60;Voided&#x60;  -  External payments: &#x60;Processed&#x60;, &#x60;Canceled&#x60;  * Update of status can change value from &#x60;Processed&#x60; to &#x60;Canceled&#x60; when the payment type is external.  | [optional] [default to null]
**TransferredToAccounting** | **string** |  Indicates if the payment was transferred to an external accounting system. Use this field for integration with accounting systems, such as NetSuite. **Character limit**: 11 **Values**: &#x60;Processing&#x60;, &#x60;Yes&#x60;, &#x60;Error&#x60;, &#x60;Ignore&#x60;  | [optional] [default to null]
**Type_** | **string** |  Indicates if the payment is external or electronic. **Character limit**: 10 **Values**: &#x60;External&#x60;, &#x60;Electronic&#x60;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


