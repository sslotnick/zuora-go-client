# PutAccountTypeSoldToContact

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | **string** | First address line, 255 characters or less.  | [optional] [default to null]
**Address2** | **string** | Second address line, 255 characters or less.  | [optional] [default to null]
**City** | **string** | City, 40 characters or less.  | [optional] [default to null]
**Country** | **string** | Country; must be a valid country name or abbreviation. If using Zuora Tax, you must specify a country in the sold-to contact to calculate tax. A bill-to contact may be used if no sold-to contact is provided.  | [optional] [default to null]
**County** | **string** | County; 32 characters or less. May optionally be used by Zuora Tax to calculate county tax.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Fax** | **string** | Fax phone number, 40 characters or less.  | [optional] [default to null]
**FirstName** | **string** | First name, 100 characters or less.  | [default to null]
**HomePhone** | **string** | Home phone number, 40 characters or less.  | [optional] [default to null]
**LastName** | **string** | Last name, 100 characters or less.  | [default to null]
**MobilePhone** | **string** | Mobile phone number, 40 characters or less.  | [optional] [default to null]
**Nickname** | **string** | Nickname for this contact  | [optional] [default to null]
**OtherPhone** | **string** | Other phone number, 40 characters or less.  | [optional] [default to null]
**OtherPhoneType** | **string** | Possible values are: &#x60;Work&#x60;, &#x60;Mobile&#x60;, &#x60;Home&#x60;, &#x60;Other&#x60;.  | [optional] [default to null]
**PersonalEmail** | **string** | Personal email address, 80 characters or less.  | [optional] [default to null]
**State** | **string** | State; must be a valid state or province name or 2-character abbreviation. If using Zuora Tax, be aware that Zuora Tax  requires a state (in the US) or province (in Canada) in this field for the sold-to contact to calculate tax, and that a bill-to contact may be used if no sold-to contact is provided.  | [optional] [default to null]
**TaxRegion** | **string** | If using Zuora Tax, a region string as optionally defined in your tax rules. Not required.  | [optional] [default to null]
**WorkEmail** | **string** | Work email address, 80 characters or less.  | [optional] [default to null]
**WorkPhone** | **string** | Work phone number, 40 characters or less.  | [optional] [default to null]
**ZipCode** | **string** | Zip code, 20 characters or less.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


