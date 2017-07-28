# GetcmTaxItemType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAmount** | **float64** | The applied amount of the credit memo taxation item.  | [optional] [default to null]
**ExemptAmount** | **float64** | The amount of taxes or VAT for which the customer has an exemption.  | [optional] [default to null]
**FinanceInformation** | [**GetcmTaxItemTypeFinanceInformation**](GETCMTaxItemType_financeInformation.md) |  | [optional] [default to null]
**Id** | **string** | The ID of the credit memo taxation item.  | [optional] [default to null]
**Jurisdiction** | **string** | The jurisdiction that applies the tax or VAT. This value is typically a state, province, county, or city.  | [optional] [default to null]
**LocationCode** | **string** | The identifier for the location based on the value of the &#x60;taxCode&#x60; field.  | [optional] [default to null]
**Name** | **string** | The name of the credit memo taxation item.  | [optional] [default to null]
**RefundAmount** | **float64** | The amount of the refund on the credit memo taxation item.  | [optional] [default to null]
**SourceTaxItemId** | **string** | The ID of the source taxation item.  | [optional] [default to null]
**TaxAmount** | **float64** | The amount of taxation.  | [optional] [default to null]
**TaxCode** | **string** | The tax code identifies which tax rules and tax rates to apply to a specific credit memo.  | [optional] [default to null]
**TaxCodeDescription** | **string** | The description of the tax code.  | [optional] [default to null]
**TaxDate** | [**time.Time**](time.Time.md) | The date that the tax is applied to the credit memo, in &#x60;yyyy-mm-dd&#x60; format.  | [optional] [default to null]
**TaxRate** | **float64** | The tax rate applied to the credit memo.  | [optional] [default to null]
**TaxRateDescription** | **string** | The description of the tax rate.  | [optional] [default to null]
**TaxRateType** | **string** | The type of the tax rate.  | [optional] [default to null]
**UnappliedAmount** | **float64** | The unapplied amount of the credit memo taxation item.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


