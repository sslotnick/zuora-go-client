# PostQuoteDocType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apiuser** | **string** | If not using Salesforce locale, this API Zuora user will be used to retrieve the locale from Zuora.  | [optional] [default to null]
**DocumentType** | **string** | Type of the document to generate: &#x60;PDF&#x60; or &#x60;DOC&#x60;.  | [default to null]
**Locale** | **string** | Salesforce locale value to use.  | [optional] [default to null]
**Password** | **string** |  | [optional] [default to null]
**QuoteId** | **string** | ｜ Id of the quote。 | [default to null]
**Sandbox** | **string** |  | [optional] [default to null]
**ServerUrl** | **string** | SOAP URL used to login to Salesforce to get data. You can get the value with the following code in a Visualforce page: &#x60;{!$Api.Partner_Server_URL_100}&#x60;  | [default to null]
**SessionId** | **string** | Salesforce session id used to log in to Salesforce to get data. You can get the value with the following code in a Visualforce page: *{!$Api.Session_ID}*  | [default to null]
**TemplateId** | **string** | Id of the quote template in Zuora.  | [default to null]
**Token** | **string** |  | [optional] [default to null]
**UseSFDCLocale** | **string** | If using Salesforce org locale, set this to a value that is not null.  | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]
**ZquotesMajorVersion** | **string** | The major version number of Zuora Quotes you are generating the quote document in. You can use a quote template with hierarchy sizes bigger than 3 if this is set to 7 or higher.  | [optional] [default to null]
**ZquotesMinorVersion** | **string** | The minor version number of Zuora Quotes you are generating the quote document in.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


