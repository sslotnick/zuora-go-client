# PostAccountTypeTaxInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VATId** | **string** | EU Value Added Tax ID.   **Note:** This feature is in Limited Availability. If you wish to have access to the feature, submit a request at [Zuora Global Support](https://support.zuora.com).  | [optional] [default to null]
**CompanyCode** | **string** | Unique code that identifies a company account in Avalara. Use this field to calculate taxes based on origin and sold-to addresses in Avalara.  **Note:** This feature is in Limited Availability. If you wish to have access to the feature, submit a request at [Zuora Global Support](https://support.zuora.com).   | [optional] [default to null]
**ExemptCertificateId** | **string** | ID of the customer tax exemption certificate. Requires Zuora Tax.  | [optional] [default to null]
**ExemptCertificateType** | **string** | Type of tax exemption certificate that the customer holds. Requires Zuora Tax.  | [optional] [default to null]
**ExemptDescription** | **string** | Description of the tax exemption certificate that the customer holds. Requires Zuora Tax.  | [optional] [default to null]
**ExemptEffectiveDate** | [**time.Time**](time.Time.md) | Date when the customer tax exemption starts. Requires Zuora Tax.  Format: &#x60;yyyy-mm-dd&#x60;. Defaults to the current date.  | [optional] [default to null]
**ExemptExpirationDate** | [**time.Time**](time.Time.md) | Date when the customer tax exemption expires. Requires Zuora Tax.  Format: &#x60;yyyy-mm-dd&#x60;. Defaults to the current date.  | [optional] [default to null]
**ExemptIssuingJurisdiction** | **string** | Jurisdiction in which the customer tax exemption certificate was issued.  | [optional] [default to null]
**ExemptStatus** | **string** | Status of the account tax exemption. Requires Zuora Tax.  Required if you use Zuora Tax. This field is unavailable if Zuora Tax is not used.  Values: &#x60;Yes&#x60;, &#x60;No&#x60;, &#x60;pendingVerification&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


