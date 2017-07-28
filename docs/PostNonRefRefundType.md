# PostNonRefRefundType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | Comments about the refund.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**FinanceInformation** | [**PostNonRefRefundTypeFinanceInformation**](PostNonRefRefundType_financeInformation.md) |  | [optional] [default to null]
**Items** | [**[]RefundCreditMemoItemType**](RefundCreditMemoItemType.md) | Container for credit memo items.  **Note:** The Invoice Item Settlement feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**MethodType** | **string** | How an external refund was issued to a customer. This field is required for an external refund and must be left empty for an electronic refund. You can issue an external refund on a credit memo.  | [optional] [default to null]
**PaymentMethodId** | **string** | The ID of the payment method used for the refund. This field is required for an electronic refund, and the value must be an electronic payment method ID. This field must be left empty for an external refund.   | [optional] [default to null]
**ReasonCode** | **string** | A code identifying the reason for the transaction. The value must be an existing reason code or empty. If you do not specify a value, Zuora uses the default reason code.  | [optional] [default to null]
**ReferenceId** | **string** | The transaction ID returned by the payment gateway for an electronic refund. Use this field to reconcile refunds between your gateway and Zuora Payments.  | [optional] [default to null]
**RefundDate** | [**time.Time**](time.Time.md) | The date when the refund takes effect, in &#x60;yyyy-mm-dd&#x60; format. The date of the refund cannot be before the credit memo date. Specify this field only for external refunds. Zuora automatically generates this field for electronic refunds.  | [optional] [default to null]
**SecondRefundReferenceId** | **string** | The transaction ID returned by the payment gateway if there is an additional transaction for the refund. Use this field to reconcile payments between your gateway and Zuora Payments.  | [optional] [default to null]
**TotalAmount** | **float64** | The total amount of the refund. The amount cannot exceed the unapplied amount of the associated credit memo. If the original credit memo was applied to one or more invoices or debit memos, you have to unapply a full or partial credit memo from the invoices or debit memos, and then refund the full or partial unapplied credit memo to your customers.  | [default to null]
**Type_** | **string** | The type of the refund.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


