# GetAccountSummaryType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicInfo** | [**GetAccountSummaryTypeBasicInfo**](GETAccountSummaryType_basicInfo.md) |  | [optional] [default to null]
**BillToContact** | [**GetAccountSummaryTypeBillToContact**](GETAccountSummaryType_billToContact.md) |  | [optional] [default to null]
**Invoices** | [**[]GetAccountSummaryInvoiceType**](GETAccountSummaryInvoiceType.md) | Container for invoices. Only returns the last 6 invoices.  | [optional] [default to null]
**Payments** | [**[]GetAccountSummaryPaymentType**](GETAccountSummaryPaymentType.md) | Container for payments. Only returns the last 6 payments.  | [optional] [default to null]
**SoldToContact** | [**GetAccountSummaryTypeSoldToContact**](GETAccountSummaryType_soldToContact.md) |  | [optional] [default to null]
**Subscriptions** | [**[]GetAccountSummarySubscriptionType**](GETAccountSummarySubscriptionType.md) | Container for subscriptions.  | [optional] [default to null]
**Success** | **bool** | Returns &#x60;true&#x60; if the request was processed successfully.  | [optional] [default to null]
**TaxInfo** | [**GetAccountSummaryTypeTaxInfo**](GETAccountSummaryType_taxInfo.md) |  | [optional] [default to null]
**Usage** | [**[]GetAccountSummaryUsageType**](GETAccountSummaryUsageType.md) | Container for usage data. Only returns the last 6 months of usage.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


