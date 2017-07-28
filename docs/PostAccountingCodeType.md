# PostAccountingCodeType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**GlAccountName** | **string** | Name of the account in your general ledger.  Field only available if you have Zuora Finance enabled. Maximum of 255 characters.  | [optional] [default to null]
**GlAccountNumber** | **string** | Account number in your general ledger.  Field only available if you have Zuora Finance enabled. Maximum of 255 characters.  | [optional] [default to null]
**Name** | **string** | Name of the accounting code.  Accounting code name must be unique. Maximum of 100 characters.  | [default to null]
**Notes** | **string** | Maximum of 2,000 characters.  | [optional] [default to null]
**Type_** | **string** | Accounting code type. You cannot create new accounting codes of type &#x60;AccountsReceivable&#x60;.   Note that &#x60;On-Account Receivable&#x60; is only available if you enable the Advanced AR Settlement feature.   | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


