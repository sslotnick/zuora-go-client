# CreatePaymentType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the customer account that the payment is created for.  | [optional] [default to null]
**Amount** | **float64** | The total amount of the payment.  | [default to null]
**Comment** | **string** | Additional information related to the payment.  | [optional] [default to null]
**Currency** | **string** | A currency defined in the web-based UI administrative settings.  | [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**DebitMemos** | [**[]PaymentDebitMemoApplicationCreateRequestType**](PaymentDebitMemoApplicationCreateRequestType.md) | Container for debit memos.  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The date when the payment takes effect, in &#x60;yyyy-mm-dd&#x60; format.  | [optional] [default to null]
**FinanceInformation** | [**CreatePaymentTypeFinanceInformation**](CreatePaymentType_financeInformation.md) |  | [optional] [default to null]
**GatewayId** | **string** | The ID of the gateway instance that processes the payment. When creating a payment, the ID must be a valid gateway instance name and this gateway must support the specific payment method. If not specified, the default gateway on the Account will be used.  | [optional] [default to null]
**Invoices** | [**[]PaymentInvoiceApplicationCreateRequestType**](PaymentInvoiceApplicationCreateRequestType.md) | Container for invoices.  | [optional] [default to null]
**PaymentMethodId** | **string** | The unique ID of the payment method that the customer used to make the payment.  | [default to null]
**ReferenceId** | **string** | The transaction ID returned by the payment gateway. Use this field to reconcile payments between your gateway and Zuora Payments.  | [optional] [default to null]
**Type_** | **string** | The type of the payment.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


