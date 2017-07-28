# ProxyCreatePayment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  The unique account ID for the customer that the payment is for. **Character limit**: 32 **Values**: a valid account ID  | [default to null]
**AccountingCode** | **string** |  The Chart of Accounts  | [optional] [default to null]
**Amount** | **float64** |  The amount of the payment. **Character limit**: 16 **Values**: a valid currency value  | [default to null]
**AppliedCreditBalanceAmount** | **float64** |  The amount of the payment to apply to a credit balance. This field is required in the Create call when the &#x60;AppliedInvoiceAmount&#x60; field value is null. **Character limit**: 16 **Values**: a valid currency value  | [default to null]
**AuthTransactionId** | **string** |  The authorization transaction ID from the payment gateway. Use this field for electronic payments, such as credit cards. **Character limit**: 50 **Values**: a string of 50 characters or fewer  | [optional] [default to null]
**Comment** | **string** |  Additional information related to the payment. **Character limit**: 255 **Values**: a string of 255 characters or fewer  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) |  The date when the payment takes effect, in &#x60;yyyy-mm-dd&#x60; format. **Character limit**: 29  | [default to null]
**Gateway** | **string** |  Name of the gateway instance that processes the payment. When creating a Payment, this must be a valid gateway instance name and this gateway must support the specific payment method. If not specified, the default gateway on the Account will be used.  | [optional] [default to null]
**GatewayOrderId** | **string** |  A merchant-specified natural key value that can be passed to the electronic payment gateway when a payment is created. If not specified, the PaymentNumber will be passed in instead. **Character limit**: 70 **Values**: a string of 70 characters or fewer  | [optional] [default to null]
**GatewayResponse** | **string** |  The message returned from the payment gateway for the payment. This message is gateway-dependent. **Character limit**: 500 **Values**: automatically generated  | [default to null]
**GatewayResponseCode** | **string** |  The code returned from the payment gateway for the payment. This code is gateway-dependent. **Character limit**: 20 **Values**: automatically generated  | [default to null]
**GatewayState** | **string** |  The status of the payment in the gateway; use for reconciliation. **Character limit**: 19 **Values**: automatically generated  | [default to null]
**PaymentMethodId** | **string** |  The ID of the payment method used for the payment. Required for Create. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**PaymentNumber** | **string** |  The unique identification number of a payment. For example: P-00000028. **Character limit**: 32 **Values**: automatically generated  | [default to null]
**ReferenceId** | **string** |  The transaction ID returned by the payment gateway. Use this field to reconcile payments between your gateway and Z-Payments. **Character limit**: 60 **Values**: a string of 60 characters or fewer  | [optional] [default to null]
**SoftDescriptor** | **string** |  A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi. **Character limit**: 35 **Values**: &#x60;[SDMerchantName]*[SDProductionInfo]&#x60;  | [optional] [default to null]
**SoftDescriptorPhone** | **string** |  A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi. **Character limit**: 20 **Values**: &#x60;[SDPhone]&#x60;  | [optional] [default to null]
**Status** | **string** |  The status of the payment in Zuora. The value depends on the type of payment. **Character limit**: 11 **Values**: one of the following:  -  Electronic payments: &#x60;Processed&#x60;, &#x60;Error&#x60;, &#x60;Voided&#x60;  -  External payments: &#x60;Processed&#x60;, &#x60;Canceled&#x60;  * Update of status can change value from &#x60;Processed&#x60; to &#x60;Canceled&#x60; when the payment type is external.  | [default to null]
**Type_** | **string** |  Indicates if the payment is external or electronic. **Character limit**: 10 **Values**: &#x60;External&#x60;, &#x60;Electronic&#x60;  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


