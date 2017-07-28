# PostTaxationItemForCmType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExemptAmount** | **float64** | The amount of taxes or VAT for which the customer has an exemption.  | [optional] [default to null]
**FinanceInformation** | [**PostTaxationItemForCmTypeFinanceInformation**](POSTTaxationItemForCMType_financeInformation.md) |  | [optional] [default to null]
**Jurisdiction** | **string** | The jurisdiction that applies the tax or VAT. This value is typically a state, province, county, or city.  | [default to null]
**LocationCode** | **string** | The identifier for the location based on the value of the &#x60;taxCode&#x60; field.  | [optional] [default to null]
**MemoItemId** | **string** | The ID of the credit memo that the taxation item is created for.  | [optional] [default to null]
**Name** | **string** | The name of the taxation item.  | [default to null]
**SourceTaxItemId** | **string** | The ID of the taxation item of the invoice, which the credit memo is created from.   If you want to use this REST API to create taxation items for a credit memo created from an invoice, the taxation items of the invoice must be created or imported through the SOAP API call.  **Note:**    - This field is only used if the credit memo is created from an invoice.    - If you do not contain this field in the request body, Zuora will automatically set a value for the &#x60;sourceTaxItemId&#x60; field based on the tax location code, tax jurisdiction, and tax rate.  | [optional] [default to null]
**TaxAmount** | **float64** | The amount of the tax applied to the credit memo.  | [default to null]
**TaxCode** | **string** | The tax code identifies which tax rules and tax rates to apply to a specific credit memo.  | [optional] [default to null]
**TaxCodeDescription** | **string** | The description of the tax code.  | [optional] [default to null]
**TaxDate** | [**time.Time**](time.Time.md) | The date when the tax is applied to the credit memo.  | [optional] [default to null]
**TaxRate** | **float64** | The tax rate applied to the credit memo.  | [default to null]
**TaxRateDescription** | **string** | The description of the tax rate.  | [optional] [default to null]
**TaxRateType** | **string** | The type of the tax rate applied to the credit memo.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


