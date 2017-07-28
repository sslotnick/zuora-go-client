# GetTaxationItemTypewithSuccess

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedById** | **string** | The ID of the Zuora user who created the taxation item.  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The date and time when the taxation item was created in the Zuora system, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]
**ExemptAmount** | **float64** | The amount of taxes or VAT for which the customer has an exemption.  | [optional] [default to null]
**FinanceInformation** | [**GetTaxationItemTypeFinanceInformation**](GETTaxationItemType_financeInformation.md) |  | [optional] [default to null]
**Id** | **string** | The ID of the taxation item.  | [optional] [default to null]
**Jurisdiction** | **string** | The jurisdiction that applies the tax or VAT. This value is typically a state, province, county, or city.  | [optional] [default to null]
**LocationCode** | **string** | The identifier for the location based on the value of the &#x60;taxCode&#x60; field.  | [optional] [default to null]
**MemoItemId** | **string** | The ID of the credit or debit memo associated with the taxation item.  | [optional] [default to null]
**Name** | **string** | The name of the taxation item.  | [optional] [default to null]
**SourceTaxItemId** | **string** | The ID of the taxation item of the invoice, which the credit or debit memo is created from.  | [optional] [default to null]
**TaxAmount** | **float64** | The amount of the tax applied to the credit or debit memo.  | [optional] [default to null]
**TaxCode** | **string** | The tax code identifies which tax rules and tax rates to apply to a specific credit or debit memo.  | [optional] [default to null]
**TaxCodeDescription** | **string** | The description of the tax code.  | [optional] [default to null]
**TaxDate** | [**time.Time**](time.Time.md) | The date when the tax is applied to the credit or debit memo.  | [optional] [default to null]
**TaxRate** | **float64** | The tax rate applied to the credit or debit memo.  | [optional] [default to null]
**TaxRateDescription** | **string** | The description of the tax rate.  | [optional] [default to null]
**TaxRateType** | **string** | The type of the tax rate applied to the credit or debit memo.  | [optional] [default to null]
**UpdatedById** | **string** | The ID of the Zuora user who last updated the taxation item.  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The date and time when the taxation item was last updated, in &#x60;yyyy-mm-dd hh:mm:ss&#x60; format.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


