# GetRefundPaymentType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account associated with this refund. Zuora associates the refund automatically with the account from the associated payment.  | [optional] [default to null]
**Amount** | **float64** | The total amount of the refund.  | [optional] [default to null]
**CancelledOn** | [**time.Time**](time.Time.md) | The date and time when the refund was cancelled, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.   | [optional] [default to null]
**Comment** | **string** | Comments about the refund.  | [optional] [default to null]
**CreatedById** | **string** | The ID of the Zuora user who created the refund.  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date and time when the refund was created, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-01 15:31:10.   | [optional] [default to null]
**CreditMemoId** | **string** | The ID of the credit memo associated with the refund.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**FinanceInformation** | [**GetRefundCreditMemoTypeFinanceInformation**](GETRefundCreditMemoType_financeInformation.md) |  | [optional] [default to null]
**GatewayId** | **string** | The ID of the gateway instance that processes the refund.  | [optional] [default to null]
**GatewayResponse** | **string** | The message returned from the payment gateway for the refund. This message is gateway-dependent.  | [optional] [default to null]
**GatewayResponseCode** | **string** | The code returned from the payment gateway for the refund. This code is gateway-dependent.  | [optional] [default to null]
**GatewayState** | **string** | The status of the refund in the gateway.  | [optional] [default to null]
**Id** | **string** | The ID of the created refund.  | [optional] [default to null]
**MarkedForSubmissionOn** | [**time.Time**](time.Time.md) | The date and time when a refund was marked and waiting for batch submission to the payment process, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**MethodType** | **string** | How an external refund was issued to a customer.  | [optional] [default to null]
**Number** | **string** | The unique identification number of the refund.  | [optional] [default to null]
**PaymentId** | **string** | The ID of the payment that is refunded.  | [optional] [default to null]
**PaymentMethodId** | **string** | The unique ID of the payment method that the customer used to make the refund.   | [optional] [default to null]
**PaymentMethodSnapshotId** | **string** | The unique ID of the payment method snapshot, which is a copy of the particular payment method used in a transaction.  | [optional] [default to null]
**ReasonCode** | **string** | A code identifying the reason for the transaction.   | [optional] [default to null]
**ReferenceId** | **string** | The transaction ID returned by the payment gateway for an electronic refund. Use this field to reconcile refunds between your gateway and Zuora Payments.  | [optional] [default to null]
**RefundDate** | [**time.Time**](time.Time.md) | The date when the refund takes effect, in &#x60;yyyy-mm-dd&#x60; format.  | [optional] [default to null]
**RefundTransactionTime** | [**time.Time**](time.Time.md) | The date and time when the refund was issued, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**SecondRefundReferenceId** | **string** | The transaction ID returned by the payment gateway if there is an additional transaction for the refund. Use this field to reconcile payments between your gateway and Zuora Payments.  | [optional] [default to null]
**SettledOn** | [**time.Time**](time.Time.md) | The date and time when the refund was settled in the payment processor, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. This field is used by the Spectrum gateway only and not applicable to other gateways.  | [optional] [default to null]
**SoftDescriptor** | **string** | A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi.  | [optional] [default to null]
**SoftDescriptorPhone** | **string** | A payment gateway-specific field that maps to Zuora for the gateways, Orbital, Vantiv and Verifi.  | [optional] [default to null]
**Status** | **string** | The status of the refund.  | [optional] [default to null]
**SubmittedOn** | [**time.Time**](time.Time.md) | The date and time when the refund was submitted, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully. | [optional] [default to null]
**Type_** | **string** | The type of the refund.  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the the Zuora user who last updated the refund.  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date and time when the refund was last updated, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format. For example, 2017-03-02 15:36:10.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


