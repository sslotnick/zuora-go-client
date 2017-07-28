# GetarPaymentTypewithSuccess

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the customer account that the payment is for.  | [optional] [default to null]
**Amount** | **float64** | The total amount of the payment.  | [optional] [default to null]
**AppliedAmount** | **float64** | The applied amount of the payment.  | [optional] [default to null]
**AuthTransactionId** | **string** | The authorization transaction ID from the payment gateway.  | [optional] [default to null]
**BankIdentificationNumber** | **string** | The first six digits of the credit card or debit card used for the payment, when applicable.  | [optional] [default to null]
**CancelledOn** | [**time.Time**](time.Time.md) | The date and time when the payment was cancelled, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**Comment** | **string** | Comments about the payment.  | [optional] [default to null]
**CreatedById** | **string** | The ID of the Zuora user who created the payment part.  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date and time when the payment was created, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-01 15:31:10.  | [optional] [default to null]
**CreditBalanceAmount** | **float64** | The amount that the payment transfers to the credit balance. The value is not &#x60;0&#x60; only for those payments that come from legacy payment operations performed without the Advanced AR Settlement feature.  | [optional] [default to null]
**Currency** | **string** | A currency defined in the web-based UI administrative settings.  For more information about the supported currencies and , see [ISO Currency Codes] (https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/D_Country%2C_State%2C_and_Province_Codes/D_Currencies_and_Their_3-Letter_Codes).  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | The date and time when the payment takes effect, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**FinanceInformation** | [**GetarPaymentTypeFinanceInformation**](GETARPaymentType_financeInformation.md) |  | [optional] [default to null]
**GatewayId** | **string** | The ID of the gateway instance that processes the payment.  | [optional] [default to null]
**GatewayOrderId** | **string** | A merchant-specified natural key value that can be passed to the electronic payment gateway when a payment is created.  | [optional] [default to null]
**GatewayResponse** | **string** | The message returned from the payment gateway for the payment. This message is gateway-dependent.  | [optional] [default to null]
**GatewayResponseCode** | **string** | The code returned from the payment gateway for the payment. This code is gateway-dependent.  | [optional] [default to null]
**GatewayState** | **string** | The status of the payment in the gateway; use for reconciliation.   | [optional] [default to null]
**Id** | **string** | The unique ID of the payment. For example, 4028905f5a87c0ff015a87eb6b75007f.  | [optional] [default to null]
**MarkedForSubmissionOn** | [**time.Time**](time.Time.md) | The date and time when a payment was marked and waiting for batch submission to the payment process, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**Number** | **string** | The unique identification number of the payment. For example, P-00000001.  | [optional] [default to null]
**PaymentMethodId** | **string** | The unique ID of the payment method that the customer used to make the payment.  | [optional] [default to null]
**PaymentMethodSnapshotId** | **string** | The unique ID of the payment method snapshot which is a copy of the particular Payment Method used in a transaction.  | [optional] [default to null]
**ReferenceId** | **string** | The transaction ID returned by the payment gateway. Use this field to reconcile payments between your gateway and Zuora Payments.  | [optional] [default to null]
**RefundAmount** | **float64** | The amount of the payment that is refunded.  | [optional] [default to null]
**SecondPaymentReferenceId** | **string** | The transaction ID returned by the payment gateway if there is an additional transaction for the payment. Use this field to reconcile payments between your gateway and Zuora Payments.  | [optional] [default to null]
**SettledOn** | [**time.Time**](time.Time.md) | The date and time when the payment was settled in the payment processor, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. This field is used by the Spectrum gateway only and not applicable to other gateways.  | [optional] [default to null]
**SoftDescriptor** | **string** | A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi.  | [optional] [default to null]
**SoftDescriptorPhone** | **string** | A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi.  | [optional] [default to null]
**Status** | **string** | The status of the payment.  | [optional] [default to null]
**SubmittedOn** | [**time.Time**](time.Time.md) | The date and time when the payment was submitted, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**Type_** | **string** | The type of the payment.  | [optional] [default to null]
**UnappliedAmount** | **float64** | The unapplied amount of the payment.  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the Zuora user who last updated the payment.  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date and time when the payment was last updated, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-02 15:36:10.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


