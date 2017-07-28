# GetAccountTypeBillingAndPayment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEmailAddresses** | **[]string** | A list of additional email addresses to receive emailed invoices.  | [optional] [default to null]
**BillCycleDay** | **string** | Billing cycle day (BCD), the day of the month when a bill run generates invoices for the account.  | [optional] [default to null]
**Currency** | **string** | A currency defined in the web-based UI administrative settings.  | [optional] [default to null]
**InvoiceDeliveryPrefsEmail** | **bool** | Whether the customer wants to receive invoices through email.   | [optional] [default to null]
**InvoiceDeliveryPrefsPrint** | **bool** | Whether the customer wants to receive printed invoices, such as through postal mail.  | [optional] [default to null]
**PaymentGateway** | **string** | The name of the payment gateway instance. If null or left unassigned, the Account will use the Default Gateway.  | [optional] [default to null]
**PaymentTerm** | **string** | A payment-terms indicator defined in the web-based UI administrative settings, e.g., \&quot;Net 30\&quot;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


