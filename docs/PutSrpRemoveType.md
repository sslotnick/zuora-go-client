# PutSrpRemoveType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractEffectiveDate** | [**time.Time**](time.Time.md) | Effective date of the new subscription, as yyyy-mm-dd.  | [default to null]
**CustomerAcceptanceDate** | [**time.Time**](time.Time.md) | The date when the customer accepts the contract in yyyy-mm-dd format.  If this field is not set:  * If the &#x60;serviceActivationDate&#x60; field is not set, the value of this field is set to be the contract effective date. * If the &#x60;serviceActivationDate&#x60; field is set, the value of this field is set to be the service activation date.  The billing trigger dates must follow this rule:  contractEffectiveDate &lt;&#x3D; serviceActivationDate &lt;&#x3D; contractAcceptanceDate  | [optional] [default to null]
**RatePlanId** | **string** | ID of a rate-plan charge for this subscription  | [default to null]
**ServiceActivationDate** | [**time.Time**](time.Time.md) | The date when the remove amendment is activated in yyyy-mm-dd format.  You must specify a Service Activation date if the Customer Acceptance date is set. If the Customer Acceptance date is not set, the value of the &#x60;serviceActivationDate&#x60; field defaults to be the Contract Effective Date.  The billing trigger dates must follow this rule:  contractEffectiveDate &lt;&#x3D; serviceActivationDate &lt;&#x3D; contractAcceptanceDate  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


