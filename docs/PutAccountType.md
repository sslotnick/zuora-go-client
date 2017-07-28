# PutAccountType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEmailAddresses** | **[]string** | A list of additional email addresses to receive emailed invoices. Use a comma to separate each email address.  **Note:** Invoices are emailed to the email addresses specified in this field only when the &#x60;invoiceDeliveryPrefsEmail&#x60; field is &#x60;true&#x60;.  | [optional] [default to null]
**AutoPay** | **bool** | Specifies whether future payments are to be automatically billed when they are due. Possible values are: &#x60;true&#x60;, &#x60;false&#x60;.  | [optional] [default to null]
**Batch** | **string** | The alias name given to a batch. A string of 50 characters or less.  | [optional] [default to null]
**BillToContact** | [**PutAccountTypeBillToContact**](PUTAccountType_billToContact.md) |  | [optional] [default to null]
**CommunicationProfileId** | **string** | The ID of a communication profile.  | [optional] [default to null]
**CrmId** | **string** | CRM account ID for the account, up to 100 characters.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**InvoiceDeliveryPrefsEmail** | **bool** | Whether the customer wants to receive invoices through email.   The default value is &#x60;false&#x60;.  | [optional] [default to null]
**InvoiceDeliveryPrefsPrint** | **bool** | Whether the customer wants to receive printed invoices, such as through postal mail.  The default value is &#x60;false&#x60;.  | [optional] [default to null]
**InvoiceTemplateId** | **string** | Invoice template ID, configured in Billing Settings in the Zuora UI.  | [optional] [default to null]
**Name** | **string** | Account name, up to 255 characters.  | [optional] [default to null]
**Notes** | **string** | A string of up to 65,535 characters.  | [optional] [default to null]
**PaymentGateway** | **string** | The name of the payment gateway instance. If null or left unassigned, the Account will use the Default Gateway.  | [optional] [default to null]
**SoldToContact** | [**PutAccountTypeSoldToContact**](PUTAccountType_soldToContact.md) |  | [optional] [default to null]
**Tagging** | **string** |  | [optional] [default to null]
**TaxInfo** | [**PostAccountTypeTaxInfo**](POSTAccountType_taxInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


