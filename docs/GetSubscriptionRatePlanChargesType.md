# GetSubscriptionRatePlanChargesType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyDiscountTo** | **string** | Specifies the type of charges a specific discount applies to.   This field is only used when applied to a discount charge model. If you are not using a discount charge model, the value is null.  Possible values:  * &#x60;RECURRING&#x60; * &#x60;USAGE&#x60; * &#x60;ONETIMERECURRING&#x60; * &#x60;ONETIMEUSAGE&#x60; * &#x60;RECURRINGUSAGE&#x60; * &#x60;ONETIMERECURRINGUSAGE&#x60;  | [optional] [default to null]
**BillingDay** | **string** | Billing cycle day (BCD), which is when bill runs generate invoices for charges associated with the product rate plan charge or the account.    Values:  * &#x60;DefaultFromCustomer&#x60; * &#x60;SpecificDayofMonth(#)&#x60; * &#x60;SubscriptionStartDay&#x60; * &#x60;ChargeTriggerDay&#x60; * &#x60;SpecificDayOfWeek/dayofweek&#x60;: in which dayofweek is the day in the week you define your billing periods to start.  In the response data, a day-of-the-month value (&#x60;1&#x60;-&#x60;31&#x60;) appears in place of the hash sign above (\&quot;#\&quot;). If this value exceeds the number of days in a particular month, the last day of the month is used as the BCD.  | [optional] [default to null]
**BillingPeriod** | **string** | Allows billing period to be overridden on the rate plan charge.  | [optional] [default to null]
**BillingPeriodAlignment** | **string** | Possible values:  * &#x60;AlignToCharge&#x60; * &#x60;AlignToSubscriptionStart&#x60; * &#x60;AlignToTermStart&#x60;  | [optional] [default to null]
**BillingTiming** | **string** | The billing timing for the charge. This field is only used if the &#x60;ratePlanChargeType&#x60; value is &#x60;Recurring&#x60;.  Possible values are:  * &#x60;In Advance&#x60; * &#x60;In Arrears&#x60;  **Note:** This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**ChargedThroughDate** | [**time.Time**](time.Time.md) | The date through which a customer has been billed for the charge.  | [optional] [default to null]
**Currency** | **string** | Currency used by the account. For example, &#x60;USD&#x60; or &#x60;EUR&#x60;.  | [optional] [default to null]
**CustomFieldC** | **string** | Any custom fields defined for this object. The custom field name is case-sensitive.  | [optional] [default to null]
**Description** | **string** | Description of the rate plan charge.  | [optional] [default to null]
**DiscountAmount** | **string** | The amount of the discount.  | [optional] [default to null]
**DiscountApplyDetails** | [**[]GetDiscountApplyDetailsType**](GETDiscountApplyDetailsType.md) | Container for the application details about a discount rate plan charge.   Only discount rate plan charges have values in this field.  | [optional] [default to null]
**DiscountClass** | **string** | The class that the discount belongs to. The discount class defines the order in which discount rate plan charges are applied.  For more information, see [Manage Discount Classes](https://knowledgecenter.zuora.com/BC_Subscription_Management/Product_Catalog/B_Charge_Models/Manage_Discount_Classes).  | [optional] [default to null]
**DiscountLevel** | **string** | The level of the discount. Values: &#x60;RatePlan&#x60;, &#x60;Subscription&#x60;, &#x60;Account&#x60;.  | [optional] [default to null]
**DiscountPercentage** | **string** | The amount of the discount as a percentage.  | [optional] [default to null]
**Dmrc** | **string** | The change (delta) of monthly recurring charge exists when the change in monthly recurring revenue caused by an amendment or a new subscription.  | [optional] [default to null]
**Done** | **bool** | A value of &#x60;true&#x60; indicates that an invoice for a charge segment has been completed. A value of &#x60;false&#x60; indicates that an invoice has not bee completed for the charge segment.  | [optional] [default to null]
**Dtcv** | **string** | After an amendment or an AutomatedPriceChange event, &#x60;dtcv&#x60; displays the change (delta) for the total contract value (TCV) amount for this charge, compared with its previous value with recurring charge types.  | [optional] [default to null]
**EffectiveEndDate** | [**time.Time**](time.Time.md) | The effective end date of the rate plan charge.  | [optional] [default to null]
**EffectiveStartDate** | [**time.Time**](time.Time.md) | The effective start date of the rate plan charge.  | [optional] [default to null]
**EndDateCondition** | **string** | Defines when the charge ends after the charge trigger date.  If the subscription ends before the charge end date, the charge ends when the subscription ends. But if the subscription end date is subsequently changed through a Renewal, or Terms and Conditions amendment, the charge will end on the charge end date.  Values:  * &#x60;Subscription_End&#x60; * &#x60;Fixed_Period&#x60; * &#x60;Specific_End_Date&#x60;  | [optional] [default to null]
**Id** | **string** | Rate plan charge ID.  | [optional] [default to null]
**IncludedUnits** | **string** | Specifies the number of units in the base set of units.  | [optional] [default to null]
**ListPriceBase** | **string** | List price base; possible values are:  * &#x60;Per_Billing_Period&#x60; * &#x60;Per_Month&#x60; * &#x60;Per_Week&#x60;  | [optional] [default to null]
**Model** | **string** | Charge model; possible values are:  * &#x60;FlatFee&#x60; * &#x60;PerUnit&#x60; * &#x60;Overage&#x60; * &#x60;Volume&#x60; * &#x60;Tiered&#x60; * &#x60;TieredWithOverage&#x60; * &#x60;DiscountFixedAmount&#x60; * &#x60;DiscountPercentage&#x60;  | [optional] [default to null]
**Mrr** | **string** | Monthly recurring revenue of the rate plan charge.  | [optional] [default to null]
**Name** | **string** | Charge name.  | [optional] [default to null]
**Number** | **string** | Charge number.  | [optional] [default to null]
**NumberOfPeriods** | **int64** | Specifies the number of periods to use when calculating charges in an overage smoothing charge model.  | [optional] [default to null]
**OriginalChargeId** | **string** | The original ID of the rate plan charge.  | [optional] [default to null]
**OverageCalculationOption** | **string** | Determines when to calculate overage charges.  | [optional] [default to null]
**OveragePrice** | **string** | The price for units over the allowed amount.  | [optional] [default to null]
**OverageUnusedUnitsCreditOption** | **string** | Determines whether to credit the customer with unused units of usage.  | [optional] [default to null]
**Price** | **string** | The price associated with the rate plan charge expressed as a decimal.  | [optional] [default to null]
**PriceChangeOption** | **string** | When the following is true:  1. AutomatedPriceChange setting is on  2. Charge type is not one-time  3. Charge model is not discount percentage  Then an automatic price change can have a value for when a termed subscription is renewed.   Values (one of the following):  * &#x60;NoChange&#x60; (default) * &#x60;SpecificPercentageValue&#x60; * &#x60;UseLatestProductCatalogPricing&#x60;  | [optional] [default to null]
**PriceIncreasePercentage** | **string** | A planned future price increase amount as a percentage.  | [optional] [default to null]
**PricingSummary** | **string** | Concise description of rate plan charge model.  | [optional] [default to null]
**ProcessedThroughDate** | [**time.Time**](time.Time.md) | The date until when charges have been processed. When billing in arrears, such as usage, this field value is the the same as the &#x60;ChargedThroughDate&#x60; value. This date is the earliest date when a charge can be amended.  | [optional] [default to null]
**ProductRatePlanChargeId** | **string** |  | [optional] [default to null]
**Quantity** | **string** | The quantity of units, such as the number of authors in a hosted wiki service. Valid for all charge models except for Flat Fee pricing.  | [optional] [default to null]
**RatingGroup** | **string** | Specifies a rating group based on which usage records are rated.   Possible values are:  * &#x60;ByBillingPeriod&#x60; (default) * &#x60;ByUsageStartDate&#x60; * &#x60;ByUsageRecord&#x60; * &#x60;ByUsageUpload&#x60;  **Note:** This field is only used for per unit, volume pricing, and tiered pricing charge models. Use this field only for Usage charges. One-Time Charges and Recurring Charges return &#x60;NULL&#x60;.  **Note:** This feature is in **Limited Availability**. If you wish to have access to the feature, submit a request at [Zuora Global Support](http://support.zuora.com/).  | [optional] [default to null]
**Segment** | **int64** | The identifying number of the subscription rate plan segment. Segments are numbered sequentially, starting with 1.  | [optional] [default to null]
**SmoothingModel** | **string** | Specifies when revenue recognition begins. When charge model is &#x60;Overage&#x60; or &#x60;TieredWithOverage&#x60;, &#x60;smoothingModel&#x60; will be one of the following values:  * &#x60;ContractEffectiveDate&#x60; * &#x60;ServiceActivationDate&#x60; * &#x60;CustomerAcceptanceDate&#x60;  | [optional] [default to null]
**SpecificBillingPeriod** | **int64** | Customizes the number of month or week for the charges billing period. This field is required if you set the value of the &#x60;BillingPeriod&#x60; field to &#x60;Specific_Months&#x60; or &#x60;Specific_Weeks&#x60;.  | [optional] [default to null]
**SpecificEndDate** | [**time.Time**](time.Time.md) | The specific date on which the charge ends. If the subscription ends before the specific end date, the charge ends when the subscription ends. But if the subscription end date is subsequently changed through a Renewal, or Terms and Conditions amendment, the charge will end on the specific end date.  | [optional] [default to null]
**Tcv** | **string** | The total contract value.  | [optional] [default to null]
**Tiers** | [**[]GetTierType**](GETTierType.md) | One or many defined ranges with distinct pricing.  | [optional] [default to null]
**TriggerDate** | [**time.Time**](time.Time.md) | The date that the rate plan charge will be triggered.  | [optional] [default to null]
**TriggerEvent** | **string** | The event that will cause the rate plan charge to be triggered.  Possible values:   * &#x60;ContractEffective&#x60; * &#x60;ServiceActivation&#x60; * &#x60;CustomerAcceptance&#x60; * &#x60;SpecificDate&#x60;  | [optional] [default to null]
**Type_** | **string** | Charge type. Possible values are: &#x60;OneTime&#x60;, &#x60;Recurring&#x60;, &#x60;Usage&#x60;.  | [optional] [default to null]
**UnusedUnitsCreditRates** | **string** | Specifies the rate to credit a customer for unused units of usage. This field is applicable only for overage charge models when the  &#x60;OverageUnusedUnitsCreditOption&#x60; field value is &#x60;CreditBySpecificRate&#x60;.  | [optional] [default to null]
**Uom** | **string** | Specifies the units to measure usage.   | [optional] [default to null]
**UpToPeriods** | **string** | Specifies the length of the period during which the charge is active. If this period ends before the subscription ends, the charge ends when this period ends.  If the subscription end date is subsequently changed through a Renewal, or Terms and Conditions amendment, the charge end date will change accordingly up to the original period end.  | [optional] [default to null]
**UpToPeriodsType** | **string** | The period type used to define when the charge ends.   Values:  * &#x60;Billing_Periods&#x60; * &#x60;Days&#x60; * &#x60;Weeks&#x60; * &#x60;Months&#x60; * &#x60;Years&#x60;  | [optional] [default to null]
**UsageRecordRatingOption** | **string** | Determines how Zuora processes usage records for per-unit usage charges.   | [optional] [default to null]
**Version** | **int64** | Rate plan charge revision number.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

