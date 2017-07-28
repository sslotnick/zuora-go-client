# ProxyCreateContact

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  The Zuora account ID associated with this contact. This field is not required when you use the Subscribe call. This field is required for all other calls. **Character limit: **32 **Values: **a valid account ID  | [default to null]
**Address1** | **string** |  The first line of the contact&#39;s address, which is often a street address or business name. **Character limit**: 255 **Values**: a string of 255 characters or fewer  | [optional] [default to null]
**Address2** | **string** |  The second line of the contact&#39;s address. **Character limit**: 255 **Values**: a string of 255 characters or fewer  | [optional] [default to null]
**City** | **string** |  The city of the contact&#39;s address. **Character limit**: 40 **Values: **a string of 40 characters or fewer  | [optional] [default to null]
**Country** | **string** |  The country of the contact&#39;s address.  | [optional] [default to null]
**County** | **string** |  The country. May optionally be used by Zuora Tax to calculate county tax. **Character limit**: 32 **Values**: a string of 32 characters or fewer  | [optional] [default to null]
**Description** | **string** |  A description for the contact. **Character limit**: 100 **Values**: a string of 100 characters or fewer  | [optional] [default to null]
**Fax** | **string** |  The contact&#39;s fax number. **Character limit**: 40 **Values**: a string of 40 characters or fewer  | [optional] [default to null]
**FirstName** | **string** |  The contact&#39;s first name. **Character limit**: 100 **Values**: a string of the contact&#39;s first name  | [default to null]
**HomePhone** | **string** |  The contact&#39;s home phone number. **Character limit**: 40 **Values**: a string of 40 characters or fewer  | [optional] [default to null]
**LastName** | **string** |  The contact&#39;s last name. **Character limit**: 100 **Values**: a string of 100 characters or fewer  | [default to null]
**MobilePhone** | **string** |  The contact&#39;s mobile phone number. **Character limit**: 40 **Values**: a string of 40 characters or fewer  | [optional] [default to null]
**NickName** | **string** |  A nickname for the contact. **Character limit**: 100 **Values**: a string of 100 characters or fewer  | [optional] [default to null]
**OtherPhone** | **string** |  An additional phone number for the contact. **Character limit**: 40 **Values**: a string of 40 characters or fewer  | [optional] [default to null]
**OtherPhoneType** | **string** | The type of the &#x60;OtherPhone&#x60;. **Character limit**: 20 **Values**: &#x60;Work&#x60;, &#x60;Mobile&#x60;, &#x60;Home&#x60;, &#x60;Other&#x60;  | [optional] [default to null]
**PersonalEmail** | **string** |  The contact&#39;s personal email address. **Character limit**: 80 **Values**: a string of 80 characters or fewer  | [optional] [default to null]
**PostalCode** | **string** |  The zip code for the contact&#39;s address. **Character limit:** 20 **Values: **a string of 20 characters or fewer  | [optional] [default to null]
**State** | **string** |  The state or province of the contact&#39;s address.  | [optional] [default to null]
**TaxRegion** | **string** | If using Zuora Tax rules  | [optional] [default to null]
**WorkEmail** | **string** |  The contact&#39;s business email address. **Character limit**: 80 **Values**: a string of 80 characters or fewer  | [optional] [default to null]
**WorkPhone** | **string** |  The contact&#39;s business phone number. **Character limit**: 40 **notes**: -- **Values**: a string of 40 characters or fewer  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


