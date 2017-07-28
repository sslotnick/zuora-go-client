# ProxyGetPayment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  The unique account ID for the customer that the payment is for. **Character limit**: 32 **Values**: a valid account ID  | [optional] [default to null]
**AccountingCode** | **string** |  The Chart of Accounts  | [optional] [default to null]
**Amount** | **float64** |  The amount of the payment. **Character limit**: 16 **Values**: a valid currency value  | [optional] [default to null]
**AppliedCreditBalanceAmount** | **float64** |  The amount of the payment to apply to a credit balance. This field is required in the Create call when the &#x60;AppliedInvoiceAmount&#x60; field value is null. **Character limit**: 16 **Values**: a valid currency value  | [optional] [default to null]
**AuthTransactionId** | **string** |  The authorization transaction ID from the payment gateway. Use this field for electronic payments, such as credit cards. **Character limit**: 50 **Values**: a string of 50 characters or fewer  | [optional] [default to null]
**BankIdentificationNumber** | **string** |  The first six digits of the credit card or debit card used for the payment, when applicable. **Character limit**: 6 **Values**: automatically generated  | [optional] [default to null]
**CancelledOn** | [**time.Time**](time.Time.md) |  The date when the payment was canceled. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**Comment** | **string** |  Additional information related to the payment. **Character limit**: 255 **Values**: a string of 255 characters or fewer  | [optional] [default to null]
**CreatedById** | **string** |  The user ID of the person who created the &#x60;Payment&#x60; object. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the Payment object was created in the Zuora system. **Character limit**: 29 **Values** automatically generated  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) |  The date when the payment takes effect, in &#x60;yyyy-mm-dd&#x60; format. **Character limit**: 29  | [optional] [default to null]
**Gateway** | **string** |  Name of the gateway instance that processes the payment. When creating a Payment, this must be a valid gateway instance name and this gateway must support the specific payment method. If not specified, the default gateway on the Account will be used.  | [optional] [default to null]
**GatewayOrderId** | **string** |  A merchant-specified natural key value that can be passed to the electronic payment gateway when a payment is created. If not specified, the PaymentNumber will be passed in instead. **Character limit**: 70 **Values**: a string of 70 characters or fewer  | [optional] [default to null]
**GatewayResponse** | **string** |  The message returned from the payment gateway for the payment. This message is gateway-dependent. **Character limit**: 500 **Values**: automatically generated  | [optional] [default to null]
**GatewayResponseCode** | **string** |  The code returned from the payment gateway for the payment. This code is gateway-dependent. **Character limit**: 20 **Values**: automatically generated  | [optional] [default to null]
**GatewayState** | **string** |  The status of the payment in the gateway; use for reconciliation. **Character limit**: 19 **Values**: automatically generated  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**MarkedForSubmissionOn** | [**time.Time**](time.Time.md) |  The date when a payment was marked and waiting for batch submission to the payment process. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**PaymentMethodId** | **string** |  The ID of the payment method used for the payment. Required for Create. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**PaymentMethodSnapshotId** | **string** |  The unique ID of the payment method snapshot which is a copy of the particular Payment Method used in a transaction. **Character limit**: 32 **Values**: a valid payment method snapshot ID  | [optional] [default to null]
**PaymentNumber** | **string** |  The unique identification number of a payment. For example: P-00000028. **Character limit**: 32 **Values**: automatically generated  | [optional] [default to null]
**ReferenceId** | **string** |  The transaction ID returned by the payment gateway. Use this field to reconcile payments between your gateway and Z-Payments. **Character limit**: 60 **Values**: a string of 60 characters or fewer  | [optional] [default to null]
**RefundAmount** | **float64** |  The amount of the payment that is refunded. This field is null if no refund is made against the payment. **Character limit**: 16 **Values**: automatically generated  | [optional] [default to null]
**SecondPaymentReferenceId** | **string** |  The transaction ID returned by the payment gateway if there is an additional transaction for the payment. Use this field to reconcile payments between your gateway and Zuora Payments. **Character limit**: 60 **Values**: a string of 60 characters or fewer  | [optional] [default to null]
**SettledOn** | [**time.Time**](time.Time.md) |  The date when the payment was settled in the payment processor. This field is used by the Spectrum gateway only and not applicable to other gateways. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]
**SoftDescriptor** | **string** |  A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi. **Character limit**: 35 **Values**: &#x60;[SDMerchantName]*[SDProductionInfo]&#x60;  | [optional] [default to null]
**SoftDescriptorPhone** | **string** |  A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi. **Character limit**: 20 **Values**: &#x60;[SDPhone]&#x60;  | [optional] [default to null]
**Source** | **string** |  Indicates how the payment was created, whether through API, manually, import, or payment run. **Character limit**: **Values**: Payment Run, Import, Manually, and API  | [optional] [default to null]
**SourceName** | **string** |  Name of the source. **Character limit**: **Values**: Payment Run number or a file name.  | [optional] [default to null]
**Status** | **string** |  The status of the payment in Zuora. The value depends on the type of payment. **Character limit**: 11 **Values**: one of the following:  -  Electronic payments: &#x60;Processed&#x60;, &#x60;Error&#x60;, &#x60;Voided&#x60;  -  External payments: &#x60;Processed&#x60;, &#x60;Canceled&#x60; * Update of status can change value from &#x60;Processed&#x60; to &#x60;Canceled&#x60; when the payment type is external.  | [optional] [default to null]
**SubmittedOn** | [**time.Time**](time.Time.md) |  The date when the payment was submitted. **Character limit:** 29 **Values**: automatically generated  | [optional] [default to null]
**TransferredToAccounting** | **string** |  Indicates if the payment was transferred to an external accounting system. Use this field for integration with accounting systems, such as NetSuite. **Character limit**: 11 **Values**: &#x60;Processing&#x60;, &#x60;Yes&#x60;, &#x60;Error&#x60;, &#x60;Ignore&#x60;  | [optional] [default to null]
**Type_** | **string** |  Indicates if the payment is external or electronic. **Character limit**: 10 **Values**: &#x60;External&#x60;, &#x60;Electronic&#x60;  | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the user who last updated the payment. **Character limit**: **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the payment was last updated. **Character limit**: **Values** **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


