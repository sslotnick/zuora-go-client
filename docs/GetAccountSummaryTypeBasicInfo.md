# GetAccountSummaryTypeBasicInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Account number.  | [optional] [default to null]
**AdditionalEmailAddresses** | **[]string** | A list of additional email addresses to receive emailed invoices.  | [optional] [default to null]
**Balance** | **string** | Current outstanding balance.  | [optional] [default to null]
**Batch** | **string** | The alias name given to a batch. A string of 50 characters or less.  | [optional] [default to null]
**BillCycleDay** | **string** | Billing cycle day (BCD), the day of the month when a bill run generates invoices for the account.  | [optional] [default to null]
**Currency** | **string** | A currency value.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**DefaultPaymentMethod** | [**GetAccountSummaryTypeBasicInfoDefaultPaymentMethod**](GETAccountSummaryType_basicInfo_defaultPaymentMethod.md) |  | [optional] [default to null]
**Id** | **string** | Account ID.  | [optional] [default to null]
**InvoiceDeliveryPrefsEmail** | **bool** | Whether the customer wants to receive invoices through email.   | [optional] [default to null]
**InvoiceDeliveryPrefsPrint** | **bool** | Whether the customer wants to receive printed invoices, such as through postal mail.  | [optional] [default to null]
**LastInvoiceDate** | [**time.Time**](time.Time.md) | Date of the most recent invoice for the account; null if no invoice has ever been generated.  | [optional] [default to null]
**LastPaymentAmount** | **string** | Amount of the most recent payment collected for the account; null if no payment has ever been collected.  | [optional] [default to null]
**LastPaymentDate** | [**time.Time**](time.Time.md) | Date of the most recent payment collected for the account. Null if no payment has ever been collected.  | [optional] [default to null]
**Name** | **string** | Account name.  | [optional] [default to null]
**Status** | **string** | Account status; possible values are: &#x60;Active&#x60;, &#x60;Draft&#x60;, &#x60;Canceled&#x60;.  | [optional] [default to null]
**Tags** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


