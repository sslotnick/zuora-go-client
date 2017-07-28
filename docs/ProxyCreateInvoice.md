# ProxyCreateInvoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | [default to null]
**AdjustmentAmount** | **float64** |  The amount of the invoice adjustments associated with the invoice. **Character limi**t: 16 **Values**: a valid currency amount  | [default to null]
**Body** | **string** |  Required  | [default to null]
**CreditBalanceAdjustmentAmount** | **float64** |  The currency amount of the adjustment applied to the customer&#39;s credit balance. **Character limit**: 16 **Values**: a valid currency amount This field is only available if the [Zuora Global Support](http://support.zuora.com/) to enable this feature.    | [default to null]
**IncludesOneTime** | **bool** |  Specifies whether the invoice includes one-time charges. You can use this field only with the Generate call for the Invoice object. **Character limit**: 5 **Values**: automatically generated from one of the following: &#x60;True&#x60; (default), &#x60;False&#x60;  | [optional] [default to null]
**IncludesRecurring** | **bool** |  Specifies whether the invoice includes recurring charges. You can use this field only with the Generate call for the Invoice object. **Character limit**: 5 **Values**: automatically generated from one of the following: &#x60;True&#x60; (default), &#x60;False&#x60;  | [optional] [default to null]
**IncludesUsage** | **bool** |  Specifies whether the invoice includes usage charges. You can use this field only with the Generate call for the Invoice object. **Character limit**: 5 **Values**: automatically generated from one of the following: &#x60;True &#x60;(default), &#x60;False&#x60;  | [optional] [default to null]
**InvoiceDate** | [**time.Time**](time.Time.md) |  Specifies the date on which to generate the invoice. **Character limit**: 29 **Version notes**: --  | [optional] [default to null]
**PaymentAmount** | **float64** |  The amount of payments applied to the invoice. **Character limit**: 16 **Value**s: automatically generated  | [default to null]
**RefundAmount** | **float64** |  Specifies the amount of a refund that was applied against an earlier payment on the invoice. **Character limit**: 16 **Values**: automatically generated  | [default to null]
**TargetDate** | [**time.Time**](time.Time.md) |  This date is used to determine which charges are to be billed. All charges that are to be billed on this date or prior will be included in this bill run. **Character limit**: 29 **Version notes**: --  | [optional] [default to null]
**TaxAmount** | **float64** |  The total amount of the taxes applied to the invoice. **Character limit**: 16 **Values**: automatically generated  | [default to null]
**TaxExemptAmount** | **float64** |  The total amount of the invoice that is exempt from taxation. **Character limit**: 16 **Values**: automatically generated  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


