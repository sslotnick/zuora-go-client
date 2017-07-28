# PutSrpUpdateType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeUpdateDetails** | [**[]PutScUpdateType**](PUTScUpdateType.md) | Container for one or more product rate plan charges.   | [optional] [default to null]
**ContractEffectiveDate** | [**time.Time**](time.Time.md) | The date when the amendment changes take effect. The format of the date is yyyy-mm-dd.  If there is already a future-dated Update Product amendment on the subscription, the &#x60;specificUpdateDate&#x60; field will be used instead of this field to specify when the Update Product amendment takes effect.  | [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**CustomerAcceptanceDate** | [**time.Time**](time.Time.md) | The date when the customer accepts the contract in yyyy-mm-dd format.  If this field is not set:  * If the &#x60;serviceActivationDate&#x60; field is not set, the value of this field is set to be the contract effective date. * If the &#x60;serviceActivationDate&#x60; field is set, the value of this field is set to be the service activation date.  The billing trigger dates must follow this rule:  contractEffectiveDate &lt;&#x3D; serviceActivationDate &lt;&#x3D; contractAcceptanceDate  | [optional] [default to null]
**RatePlanId** | **string** | ID of a rate plan for this subscription.  | [default to null]
**ServiceActivationDate** | [**time.Time**](time.Time.md) | The date when the update amendment is activated in yyyy-mm-dd format.  You must specify a Service Activation date if the Customer Acceptance date is set. If the Customer Acceptance date is not set, the value of the &#x60;serviceActivationDate&#x60; field defaults to be the Contract Effective Date.  The billing trigger dates must follow this rule:  contractEffectiveDate &lt;&#x3D; serviceActivationDate &lt;&#x3D; contractAcceptanceDate  | [optional] [default to null]
**SpecificUpdateDate** | [**time.Time**](time.Time.md) | The date when the Update Product amendment takes effect. This field is only applicable if there is already a future-dated Update Product amendment on the subscription. The format of the date is yyyy-mm-dd.  Required only for Update Product amendments if there is already a future-dated Update Product amendment on the subscription.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


