# ProxyCreateBillRun

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | ID of the account used for single account bill run.  This field is only used for creating ad hoc bill run for a single customer account.  **Character limit:** 32  | [optional] [default to null]
**AutoEmail** | **bool** | Determines whether to auto send email or not by this bill run once the bill run completes.  **Note:** You must enable the [Support Bill Run Auto-Post Billing](https://knowledgecenter.zuora.com/CB_Billing/Billing_Settings/Define_Billing_Rules) rule to pass this field.  | [optional] [default to null]
**AutoPost** | **bool** | Determines whether to auto post bill run or not once the bill run completes.  **Note:** You must enable the [Support Bill Run Auto-Post Billing](https://knowledgecenter.zuora.com/CB_Billing/Billing_Settings/Define_Billing_Rules) rule to pass this field.  | [optional] [default to null]
**AutoRenewal** | **bool** | Determines whether to auto renew subscription or not by this bill run once the bill run completes.  | [optional] [default to null]
**Batch** | **string** | Batch of accounts for this bill run.   This field is only used for creating ad hoc bill run for multiple customer accounts.  **Character limit:** 20  **Values:** AllBatches or Batchn where n is a number between 1 and 50.  | [optional] [default to null]
**BillCycleDay** | **string** | The day of the bill cycle.  This field is only used for creating ad hoc bill run for multiple customer accounts.  **Character limit:** 32  **Values:** &#x60;AllBillCycleDays&#x60; or 01 - 31.  | [optional] [default to null]
**ChargeTypeToExclude** | **string** | Charge type or types to be excluded, separated with comma.  **Possible Values:** &#x60;OneTime&#x60;, &#x60;Recurring&#x60;, &#x60;Usage&#x60;, or a combination of these values.   **Character limit:** 50  | [optional] [default to null]
**InvoiceDate** | [**time.Time**](time.Time.md) | Invoice date for this bill run.  **Character limit:** 29  | [default to null]
**NoEmailForZeroAmountInvoice** | **bool** | Determines whether to suppress email for invoices with zero total or not for this bill run once the bill run completes. (Do not email invoices with 0 Invoice Total).  | [optional] [default to null]
**TargetDate** | [**time.Time**](time.Time.md) | Target date for this bill run. See [Create Bill Run](https://knowledgecenter.zuora.com/CB_Billing/J_Billing_Operations/G_Bill_Runs/Creating_Bill_Runs) for more information.  **Character limit:** 29  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


