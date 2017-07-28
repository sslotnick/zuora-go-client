# ProxyGetPaymentMethodSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the customer account associated with this payment method. | [optional] [default to null]
**AchAbaCode** | **string** | The nine-digit routing number or ABA number used by banks. Applicable to ACH payment methods. | [optional] [default to null]
**AchAccountName** | **string** | The name of the account holder, which can be either a person or a company. Applicable to ACH payment methods. | [optional] [default to null]
**AchAccountNumberMask** | **string** | This is a masked displayable version of the ACH account number, used for security purposes. For example: &#x60;XXXXXXXXX54321&#x60;. | [optional] [default to null]
**AchAccountType** | **string** | The type of bank account associated with the ACH payment. | [optional] [default to null]
**AchBankName** | **string** | The name of the bank where the ACH payment account is held. | [optional] [default to null]
**BankBranchCode** | **string** | The branch code of the bank used for direct debit. | [optional] [default to null]
**BankCheckDigit** | **string** | The check digit in the international bank account number, which confirms the validity of the account. Applicable to direct debit payment methods. | [optional] [default to null]
**BankCity** | **string** | The city of the direct debit bank. | [optional] [default to null]
**BankCode** | **string** | The sort code or number that identifies the bank. This is also known as the sort code. | [optional] [default to null]
**BankIdentificationNumber** | **string** | The first six digits of the payment method&#39;s number, such as the credit card number or account number. Banks use this number to identify a payment method. | [optional] [default to null]
**BankName** | **string** | The name of the direct debit bank. | [optional] [default to null]
**BankPostalCode** | **string** | The zip code or postal code of the direct debit bank. | [optional] [default to null]
**BankStreetName** | **string** | The name of the street of the direct debit bank. | [optional] [default to null]
**BankStreetNumber** | **string** | The number of the direct debit bank. | [optional] [default to null]
**BankTransferAccountName** | **string** | The name on the direct debit bank account. | [optional] [default to null]
**BankTransferAccountNumberMask** | **string** | This is a masked displayable version of the bank account number, used for security purposes. For example: &#x60;XXXXXXXXX54321&#x60;. | [optional] [default to null]
**BankTransferAccountType** | **string** | The type of the customer&#39;s bank account. Applicable to direct debit payment methods. | [optional] [default to null]
**BankTransferType** | **string** | Specifies the type of direct debit transfer. The value of this field is dependent on the country of the user.  Possible Values:    * &#x60;AutomatischIncasso&#x60; (NL)  * &#x60;LastschriftDE&#x60; (Germany)  * &#x60;LastschriftAT&#x60; (Austria)  * &#x60;DemandeDePrelevement&#x60; (FR)  * &#x60;DirectDebitUK&#x60; (UK)  * &#x60;Domicil&#x60; (Belgium)  * &#x60;LastschriftCH&#x60; (CH)  * &#x60;RID&#x60; (Italy)  * &#x60;OrdenDeDomiciliacion&#x60; (Spain)  | [optional] [default to null]
**BusinessIdentificationCode** | **string** | The business identification code for Swiss direct payment methods that use the Global Collect payment gateway. Only applicable to direct debit payments in Switzerland with Global Collect. | [optional] [default to null]
**City** | **string** | The city of the customer&#39;s address. Applicable to debit payment methods. | [optional] [default to null]
**Country** | **string** | The two-letter country code of the customer&#39;s address. Applicable to direct debit payment methods. | [optional] [default to null]
**CreditCardAddress1** | **string** | The first line of the card holder&#39;s address, which is often a street address or business name. Applicable to credit card and direct debit payment methods. | [optional] [default to null]
**CreditCardAddress2** | **string** | The second line of the card holder&#39;s address. Applicable to credit card and direct debit payment methods. | [optional] [default to null]
**CreditCardCity** | **string** | The city of the card holder&#39;s address. Applicable to credit card and direct debit payment methods. | [optional] [default to null]
**CreditCardCountry** | **string** | The country of the card holder&#39;s address. | [optional] [default to null]
**CreditCardExpirationMonth** | **int32** | The expiration month of the credit card or debit card. Applicable to credit card and direct debit payment methods. | [optional] [default to null]
**CreditCardExpirationYear** | **int32** | The expiration month of the credit card or debit card. Applicable to credit card and direct debit payment methods. | [optional] [default to null]
**CreditCardHolderName** | **string** | The full name of the card holder. Applicable to credit card and direct debit payment methods. | [optional] [default to null]
**CreditCardMaskNumber** | **string** | A masked version of the credit or debit card number. | [optional] [default to null]
**CreditCardPostalCode** | **string** | The billing address&#39;s zip code. | [optional] [default to null]
**CreditCardState** | **string** | The billing address&#39;s state. Applicable if &#x60;CreditCardCountry&#x60; is either Canada or the US. | [optional] [default to null]
**CreditCardType** | **string** | The type of credit card or debit card. | [optional] [default to null]
**DeviceSessionId** | **string** | The session ID of the user when the &#x60;PaymentMethod&#x60; was created or updated. | [optional] [default to null]
**Email** | **string** | An email address for the payment method in addition to the bill to contact email address. | [optional] [default to null]
**ExistingMandate** | **string** | Indicates if the customer has an existing mandate or a new mandate. Only applicable to direct debit payment methods. | [optional] [default to null]
**FirstName** | **string** | The customer&#39;s first name. Only applicable to direct debit payment methods. | [optional] [default to null]
**IBAN** | **string** | The International Bank Account Number. Only applicable to direct debit payment methods. | [optional] [default to null]
**IPAddress** | **string** | The IP address of the user when the payment method was created or updated. | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**LastFailedSaleTransactionDate** | [**time.Time**](time.Time.md) | The date of the last failed attempt to collect payment with this payment method. | [optional] [default to null]
**LastName** | **string** | The customer&#39;s last name. Only applicable to direct debit payment methods. | [optional] [default to null]
**LastTransactionDateTime** | [**time.Time**](time.Time.md) | The date of the most recent transaction. | [optional] [default to null]
**LastTransactionStatus** | **string** | The status of the most recent transaction. | [optional] [default to null]
**MandateCreationDate** | [**time.Time**](time.Time.md) | The date when the mandate was created, in &#x60;yyyy-mm-dd&#x60; format. A mandate is a signed authorization for UK and NL customers. Only applicable to direct debit payment methods. | [optional] [default to null]
**MandateID** | **string** | The ID of the mandate. A mandate is a signed authorization for UK and NL customers. Only applicable to direct debit payment methods. | [optional] [default to null]
**MandateReceived** | **string** | Indicates if  the mandate was received. A mandate is a signed authorization for UK and NL customers. Only applicable to direct debit payment methods. | [optional] [default to null]
**MandateUpdateDate** | [**time.Time**](time.Time.md) | The date when the mandate was last updated, in &#x60;yyyy-mm-dd&#x60; format. A mandate is a signed authorization for UK and NL customers. Only applicable to direct debit payment methods. | [optional] [default to null]
**MaxConsecutivePaymentFailures** | **int32** | The number of allowable consecutive failures Zuora attempts with the payment method before stopping. | [optional] [default to null]
**Name** | **string** | The name of the payment method. | [optional] [default to null]
**NumConsecutiveFailures** | **int32** | The number of consecutive failed payment for the payment method. | [optional] [default to null]
**PaymentMethodId** | **string** | Object identifier of the payment method. | [optional] [default to null]
**PaymentMethodStatus** | **string** | Specifies the status of the payment method. | [optional] [default to null]
**PaymentRetryWindow** | **int32** | The retry interval setting, which prevents making a payment attempt if the last failed attempt was within the last specified number of hours. | [optional] [default to null]
**PaypalBaid** | **string** | The PayPal billing agreement ID, which is a contract between two PayPal accounts. | [optional] [default to null]
**PaypalEmail** | **string** | The email address associated with the account holder&#39;s PayPal account or of the PayPal account of the person paying for the service. | [optional] [default to null]
**PaypalPreapprovalKey** | **string** | PayPal&#39;s Adaptive Payments API key. | [optional] [default to null]
**PaypalType** | **string** | Specifies the PayPal gateway: PayFlow Pro (Express Checkout) or Adaptive Payments. | [optional] [default to null]
**Phone** | **string** | The phone number that the account holder registered with the bank. This field is used for credit card validation when passing to a gateway. | [optional] [default to null]
**PostalCode** | **string** | The zip code of the customer&#39;s address. Only applicable to direct debit payment methods. | [optional] [default to null]
**SecondTokenId** | **string** | A gateway unique identifier that replaces sensitive payment method data. Applicable to CC Reference Transaction payment methods. | [optional] [default to null]
**State** | **string** | The state of the customer&#39;s address. Only applicable to direct debit payment methods. | [optional] [default to null]
**StreetName** | **string** | The street name of the customer&#39;s address. Only applicable to direct debit payment methods. | [optional] [default to null]
**StreetNumber** | **string** | The street number of the customer&#39;s address. Only applicable to direct debit payment methods. | [optional] [default to null]
**TokenId** | **string** | A gateway unique identifier that replaces sensitive payment method data or represents a gateway&#39;s unique customer profile. Applicable to CC Reference Transaction payment methods. | [optional] [default to null]
**TotalNumberOfErrorPayments** | **int32** | The number of error payments that used this payment method. | [optional] [default to null]
**TotalNumberOfProcessedPayments** | **int32** | The number of successful payments that used this payment method. | [optional] [default to null]
**Type_** | **string** | The type of payment method. | [optional] [default to null]
**UseDefaultRetryRule** | **bool** | Determines whether to use the default retry rules configured in the Zuora Payments settings. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

