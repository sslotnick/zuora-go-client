/*
 * Zuora API Reference
 *
 *   # Introduction Welcome to the reference for the Zuora REST API!  <a href=\"http://en.wikipedia.org/wiki/REST_API\" target=\"_blank\">REST</a> is a web-service protocol that lends itself to rapid development by using everyday HTTP and JSON technology.  The Zuora REST API provides a broad set of operations and resources that:    * Enable Web Storefront integration from your website.   * Support self-service subscriber sign-ups and account management.   * Process revenue schedules through custom revenue rule models.   * Enable manipulation of most objects in the Zuora Object Model.      ## Endpoints      The Zuora REST API is provided via the following endpoints.   | Tenant              | Base URL for REST Endpoints |   |-------------------------|-------------------------|   |US Production | https://rest.zuora.com   |   |US API Sandbox    | https://rest.apisandbox.zuora.com|   |US Performance Test | https://rest.pt1.zuora.com |   |EU Production | https://rest.eu.zuora.com |   |EU Sandbox | https://rest.sandbox.eu.zuora.com |      The production endpoint provides access to your live user data. The API Sandbox tenant is a good place to test your code without affecting real-world data. To use it, you must be provisioned with an API Sandbox tenant - your Zuora representative can help you if needed.      ## Access to the API      If you have a Zuora tenant, you already have access to the API.      If you don't have a Zuora tenant, go to <a href=\" https://www.zuora.com/resource/zuora-test-drive\" target=\"_blank\">https://www.zuora.com/resource/zuora-test-drive</a> and sign up for a Production Test Drive tenant. The tenant comes with seed data, such as a sample product catalog.  We recommend that you <a href=\"https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/Manage_Users/Create_an_API_User\" target=\"_blank\">create an API user</a> specifically for making API calls. Don't log in to the Zuora UI with this account. Logging in to the UI enables a security feature that periodically expires the account's password, which may eventually cause authentication failures with the API. Note that a user role does not have write access to Zuora REST services unless it has the API Write Access permission as described in those instructions.   # API Changelog You can find the <a href=\"https://community.zuora.com/t5/Developers/API-Changelog/gpm-p/18092\" target=\"_blank\">Changelog</a> of the API Reference in the Zuora Community.   # Authentication  There are three ways to authenticate:    * Use username and password. Include authentication with each request in the header:         * `apiAccessKeyId`      * `apiSecretAccessKey`     * `entityId` or `entityName` (Only for [Zuora Multi-entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity \"Multi-entity\"). See \"Entity Id and Entity Name\" below for more information.)   * Use an authorization cookie. The cookie authorizes the user to make calls to the REST API for the duration specified in  **Administration > Security Policies > Session timeout**. The cookie expiration time is reset with this duration after every call to the REST API. To obtain a cookie, call the REST  `connections` resource with the following API user information:         *   ID         *   password     *   entity Id or entity name (Only for [Zuora Multi-entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity \"Multi-entity\"). See \"Entity Id and Entity Name\" below for more information.)         * For CORS-enabled APIs only: Include a 'single-use' token in the request header, which re-authenticates the user with each request. See below for more details.  ## Entity Id and Entity Name  The `entityId` and `entityName` parameters are only used for [Zuora Multi-entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity \"Zuora Multi-entity\").   The  `entityId` and `entityName` parameters specify the Id and the [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name \"Introduction to Entity and Entity Hierarchy\") that you want to access, respectively. Note that you must have permission to access the entity.   You can specify either the `entityId` or `entityName` parameter in the authentication to access and view an entity.    * If both `entityId` and `entityName` are specified in the authentication, an error occurs.    * If neither `entityId` nor `entityName` is specified in the authentication, you will log in to the entity in which your user account is created.      To get the entity Id and entity name, you can use the GET Entities REST call. For more information, see [API User Authentication](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/A_Overview_of_Multi-entity#API_User_Authentication \"API User Authentication\").      ## Token Authentication for CORS-Enabled APIs      The CORS mechanism enables REST API calls to Zuora to be made directly from your customer's browser, with all credit card and security information transmitted directly to Zuora. This minimizes your PCI compliance burden, allows you to implement advanced validation on your payment forms, and  makes your payment forms look just like any other part of your website.    For security reasons, instead of using cookies, an API request via CORS uses **tokens** for authentication.  The token method of authentication is only designed for use with requests that must originate from your customer's browser; **it should  not be considered a replacement to the existing cookie authentication** mechanism.  See [Zuora CORS REST](https://knowledgecenter.zuora.com/DC_Developers/REST_API/A_REST_basics/G_CORS_REST \"Zuora CORS REST\") for details on how CORS works and how you can begin to implement customer calls to the Zuora REST APIs. See  [HMAC Signatures](https://www.zuora.com/developer/API-Reference/#operation/POSTHMACSignature \"HMAC Signatures\") for details on the HMAC method that returns the authentication token.  # Requests and Responses  ## Request IDs  As a general rule, when asked to supply a \"key\" for an account or subscription (accountKey, account-key, subscriptionKey, subscription-key), you can provide either the actual ID or  the number of the entity.  ## HTTP Request Body  Most of the parameters and data accompanying your requests will be contained in the body of the HTTP request.   The Zuora REST API accepts JSON in the HTTP request body. No other data format (e.g., XML) is supported.  ### Data Type  ([Actions](https://www.zuora.com/developer/api-reference/#tag/Actions) and CRUD operations only) We recommend that you do not specify the decimal values with quotation marks, commas, and spaces. Use characters of `+-0-9.eE`, for example, `5`, `1.9`, `-8.469`, and `7.7e2`. Also, Zuora does not convert currencies for decimal values.  ## Testing a Request  Use a third party client, such as [curl](https://curl.haxx.se \"curl\"), [Postman](https://www.getpostman.com \"Postman\"), or [Advanced REST Client](https://advancedrestclient.com \"Advanced REST Client\"), to test the Zuora REST API.  You can test the Zuora REST API from the Zuora API Sandbox or Production tenants. If connecting to Production, bear in mind that you are working with your live production data, not sample data or test data.  ## Testing with Credit Cards  Sooner or later it will probably be necessary to test some transactions that involve credit cards. For suggestions on how to handle this, see [Going Live With Your Payment Gateway](https://knowledgecenter.zuora.com/CB_Billing/M_Payment_Gateways/C_Managing_Payment_Gateways/B_Going_Live_Payment_Gateways#Testing_with_Credit_Cards \"C_Zuora_User_Guides/A_Billing_and_Payments/M_Payment_Gateways/C_Managing_Payment_Gateways/B_Going_Live_Payment_Gateways#Testing_with_Credit_Cards\" ).  ## Concurrent Request Limits  Zuora enforces tenant-level concurrent request limits. See <a href=\"https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Policies/Concurrent_Request_Limits\" target=\"_blank\">Concurrent Request Limits</a> for more information.    ## Error Handling  Responses and error codes are detailed in [Responses and errors](https://knowledgecenter.zuora.com/DC_Developers/REST_API/A_REST_basics/3_Responses_and_errors \"Responses and errors\").  # Pagination  When retrieving information (using GET methods), the optional `pageSize` query parameter sets the maximum number of rows to return in a response. The maximum is `40`; larger values are treated as `40`. If this value is empty or invalid, `pageSize` typically defaults to `10`.  The default value for the maximum number of rows retrieved can be overridden at the method level.  If more rows are available, the response will include a `nextPage` element, which contains a URL for requesting the next page.  If this value is not provided, no more rows are available. No \"previous page\" element is explicitly provided; to support backward paging, use the previous call.  ## Array Size  For data items that are not paginated, the REST API supports arrays of up to 300 rows.  Thus, for instance, repeated pagination can retrieve thousands of customer accounts, but within any account an array of no more than 300 rate plans is returned.  # API Versions  The Zuora REST API are version controlled. Versioning ensures that Zuora REST API changes are backward compatible. Zuora uses a major and minor version nomenclature to manage changes. By specifying a version in a REST request, you can get expected responses regardless of future changes to the API.  ## Major Version  The major version number of the REST API appears in the REST URL. Currently, Zuora only supports the **v1** major version. For example, `POST https://rest.zuora.com/v1/subscriptions`.  ## Minor Version  Zuora uses minor versions for the REST API to control small changes. For example, a field in a REST method is deprecated and a new field is used to replace it.   Some fields in the REST methods are supported as of minor versions. If a field is not noted with a minor version, this field is available for all minor versions. If a field is noted with a minor version, this field is in version control. You must specify the supported minor version in the request header to process without an error.   If a field is in version control, it is either with a minimum minor version or a maximum minor version, or both of them. You can only use this field with the minor version between the minimum and the maximum minor versions. For example, the `invoiceCollect` field in the POST Subscription method is in version control and its maximum minor version is 189.0. You can only use this field with the minor version 189.0 or earlier.  If you specify a version number in the request header that is not supported, Zuora will use the minimum minor version of the REST API. In our REST API documentation, if a field or feature requires a minor version number, we note that in the field description.  You only need to specify the version number when you use the fields require a minor version. To specify the minor version, set the `zuora-version` parameter to the minor version number in the request header for the request call. For example, the `collect` field is in 196.0 minor version. If you want to use this field for the POST Subscription method, set the  `zuora-version` parameter to `196.0` in the request header. The `zuora-version` parameter is case sensitive.  For all the REST API fields, by default, if the minor version is not specified in the request header, Zuora will use the minimum minor version of the REST API to avoid breaking your integration.   ### Minor Version History  The supported minor versions are not serial. This section documents the changes made to each Zuora REST API minor version.  The following table lists the supported versions and the fields that have a Zuora REST API minor version.  | Fields         | Minor Version      | REST Methods    | Description | |:--------|:--------|:--------|:--------| | invoiceCollect | 189.0 and earlier  | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Generates an invoice and collects a payment for a subscription. | | collect        | 196.0 and later    | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Collects an automatic payment for a subscription. | | invoice | 196.0 and 207.0| [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Generates an invoice for a subscription. | | invoiceTargetDate | 196.0 and earlier  | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\") |Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | invoiceTargetDate | 207.0 and earlier  | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | targetDate | 207.0 and later | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\") |Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | targetDate | 211.0 and later | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | includeExisting DraftInvoiceItems | 196.0 and earlier| [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | Specifies whether to include draft invoice items in subscription previews. Specify it to be `true` (default) to include draft invoice items in the preview result. Specify it to be `false` to excludes draft invoice items in the preview result. | | includeExisting DraftDocItems | 207.0 and later  | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | Specifies whether to include draft invoice items in subscription previews. Specify it to be `true` (default) to include draft invoice items in the preview result. Specify it to be `false` to excludes draft invoice items in the preview result. | | previewType | 196.0 and earlier| [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | The type of preview you will receive. The possible values are `InvoiceItem`(default), `ChargeMetrics`, and `InvoiceItemChargeMetrics`. | | previewType | 207.0 and later  | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | The type of preview you will receive. The possible values are `LegalDoc`(default), `ChargeMetrics`, and `LegalDocChargeMetrics`. | | runBilling  | 211.0 and later  | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Generates an invoice or credit memo for a subscription. **Note:** Credit memos are only available if you have the Advanced AR Settlement feature enabled. |  #### Version 207.0 and Later  The response structure of the [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\") and [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") methods are changed. The following invoice related response fields are moved to the invoice container:    * amount   * amountWithoutTax   * taxAmount   * invoiceItems   * targetDate   * chargeMetrics  # Zuora Object Model  The following diagram presents a high-level view of the key Zuora objects. Click the image to open it in a new tab to resize it.  <a href=\"https://www.zuora.com/wp-content/uploads/2017/01/ZuoraERD.jpeg\" target=\"_blank\"><img src=\"https://www.zuora.com/wp-content/uploads/2017/01/ZuoraERD.jpeg\" alt=\"Zuora Object Model Diagram\"></a>  You can use the [Describe object](https://www.zuora.com/developer/api-reference/#operation/GET_Describe) operation to list the fields of each Zuora object that is available in your tenant. When you call the operation, you must specify the API name of the Zuora object.  The following table provides the API name of each Zuora object:  | Object                                        | API Name                                   | |-----------------------------------------------|--------------------------------------------| | Account                                       | `Account`                                  | | Accounting Code                               | `AccountingCode`                           | | Accounting Period                             | `AccountingPeriod`                         | | Amendment                                     | `Amendment`                                | | Application Group                             | `ApplicationGroup`                         | | Contact                                       | `Contact`                                  | | Contact Snapshot                              | `ContactSnapshot`                          | | Billing Run                                   | `BillingRun`                               | | Credit Balance Adjustment                     | `CreditBalanceAdjustment`                  | | Credit Memo                                   | `CreditMemo`                               | | Credit Memo Application                       | `CreditMemoApplication`                    | | Credit Memo Application Item                  | `CreditMemoApplicationItem`                | | Credit Memo Item                              | `CreditMemoItem`                           | | Credit Memo Part                              | `CreditMemoPart`                           | | Credit Memo Part Item                         | `CreditMemoPartItem`                       | | Credit Taxation Item                          | `CreditTaxationItem`                       | | Custom Exchange Rate                          | `FXCustomRate`                             | | Debit Memo                                    | `DebitMemo`                                | | Debit Memo Item                               | `DebitMemoItem`                            | | Debit Taxation Item                           | `DebitTaxationItem`                        | | Payment Method                                | `PaymentMethod`                            | | Entity                                        | `Tenant`                                   | | Gateway Reconciliation Event                  | `PaymentGatewayReconciliationEventLog`     | | Gateway Reconciliation Job                    | `PaymentReconciliationJob`                 | | Gateway Reconciliation Log                    | `PaymentReconciliationLog`                 | | Invoice                                       | `Invoice`                                  | | Invoice Adjustment                            | `InvoiceAdjustment`                        | | Invoice Item                                  | `InvoiceItem`                              | | Invoice Item Adjustment                       | `InvoiceItemAdjustment`                    | | Invoice Payment                               | `InvoicePayment`                           | | Journal Entry                                 | `JournalEntry`                             | | Journal Entry Item                            | `JournalEntryItem`                         | | Journal Run                                   | `JournalRun`                               | | Payment                                       | `Payment`                                  | | Payment Application                           | `PaymentApplication`                       | | Payment Application Item                      | `PaymentApplicationItem`                   | | Payment Method Snapshot                       | `PaymentMethodSnapshot`                    | | Payment Method Transaction Log                | `PaymentMethodTransactionLog`              | | Payment Method Update                         | `UpdaterDetail`                            | | Payment Part                                  | `PaymentPart`                              | | Payment Part Item                             | `PaymentPartItem`                          | | Payment Transaction Log                       | `PaymentTransactionLog`                    | | Processed Usage                               | `ProcessedUsage`                           | | Product                                       | `Product`                                  | | Product Rate Plan                             | `ProductRatePlan`                          | | Product Rate Plan Charge                      | `ProductRatePlanCharge`                    | | Product Rate Plan Charge Tier                 | `ProductRatePlanChargeTier`                | | Rate Plan                                     | `RatePlan`                                 | | Rate Plan Charge                              | `RatePlanCharge`                           | | Rate Plan Charge Tier                         | `RatePlanChargeTier`                       | | Refund                                        | `Refund`                                   | | Refund Application                            | `RefundApplication`                        | | Refund Application Item                       | `RefundApplicationItem`                    | | Refund Invoice Payment                        | `RefundInvoicePayment`                     | | Refund Part                                   | `RefundPart`                               | | Refund Part Item                              | `RefundPartItem`                           | | Refund Transaction Log                        | `RefundTransactionLog`                     | | Revenue Charge Summary                        | `RevenueChargeSummary`                     | | Revenue Charge Summary Item                   | `RevenueChargeSummaryItem`                 | | Revenue Event                                 | `RevenueEvent`                             | | Revenue Event Credit Memo Item                | `RevenueEventCreditMemoItem`               | | Revenue Event Debit Memo Item                 | `RevenueEventDebitMemoItem`                | | Revenue Event Invoice Item                    | `RevenueEventInvoiceItem`                  | | Revenue Event Invoice Item Adjustment         | `RevenueEventInvoiceItemAdjustment`        | | Revenue Event Item                            | `RevenueEventItem`                         | | Revenue Event Item Credit Memo Item           | `RevenueEventItemCreditMemoItem`           | | Revenue Event Item Debit Memo Item            | `RevenueEventItemDebitMemoItem`            | | Revenue Event Item Invoice Item               | `RevenueEventItemInvoiceItem`              | | Revenue Event Item Invoice Item Adjustment    | `RevenueEventItemInvoiceItemAdjustment`    | | Revenue Event Type                            | `RevenueEventType`                         | | Revenue Schedule                              | `RevenueSchedule`                          | | Revenue Schedule Credit Memo Item             | `RevenueScheduleCreditMemoItem`            | | Revenue Schedule Debit Memo Item              | `RevenueScheduleDebitMemoItem`             | | Revenue Schedule Invoice Item                 | `RevenueScheduleInvoiceItem`               | | Revenue Schedule Invoice Item Adjustment      | `RevenueScheduleInvoiceItemAdjustment`     | | Revenue Schedule Item                         | `RevenueScheduleItem`                      | | Revenue Schedule Item Credit Memo Item        | `RevenueScheduleItemCreditMemoItem`        | | Revenue Schedule Item Debit Memo Item         | `RevenueScheduleItemDebitMemoItem`         | | Revenue Schedule Item Invoice Item            | `RevenueScheduleItemInvoiceItem`           | | Revenue Schedule Item Invoice Item Adjustment | `RevenueScheduleItemInvoiceItemAdjustment` | | Subscription                                  | `Subscription`                             | | Taxation Item                                 | `TaxationItem`                             | | Updater Batch                                 | `UpdaterBatch`                             | | Usage                                         | `Usage`                                    |
 *
 * OpenAPI spec version: 2017-07-12
 * Contact: docs@zuora.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package zuora

import (
	"time"
)

//  This is the object defining the payment details for the Account. The Account will be updated with this payment as the default payment method. Use this field if you are associating an electronic payment method with the account. A payment gateway must be enabled. Values: A valid electronic PaymentMethod.
type SubscribeRequestPaymentMethod struct {

	//  The ID of the customer account associated with this payment method.
	AccountId string `json:"AccountId,omitempty"`

	//  The nine-digit routing number or ABA number used by banks. Use this field for ACH payment methods.   **Character limit**: 9   **Values**: a string of 9 characters or fewer
	AchAbaCode string `json:"AchAbaCode,omitempty"`

	//  The name of the account holder, which can be either a person or a company. Use this field for ACH payment methods.   **Character limit**: 70   **Values**: a string of 70 characters or fewer
	AchAccountName string `json:"AchAccountName,omitempty"`

	//  The bank account number associated with the ACH payment. Use this field for ACH payment methods.   **Character limit**: 30   **Values**: a string of 30 numeric characters or fewer
	AchAccountNumber string `json:"AchAccountNumber,omitempty"`

	//  This is a masked displayable version of the ACH account number, used for security purposes. For example: `XXXXXXXXX54321`. Use this field for ACH payment methods.   **Character limit**: 32   **Values**: automatically generated
	AchAccountNumberMask string `json:"AchAccountNumberMask,omitempty"`

	//  The type of bank account associated with the ACH payment. Use this field for ACH payment methods.   **Character limit**: 16   **Values**:  - `BusinessChecking` - `Checking` - `Saving`
	AchAccountType string `json:"AchAccountType,omitempty"`

	//  Line 1 for the ACH address. Required on create for the Vantiv payment gateway. Optional for other gateways. **Character limit:** **Values:** an address
	AchAddress1 string `json:"AchAddress1,omitempty"`

	//  Line 2 for the ACH address. Required on create for the Vantiv payment gateway. Optional for other gateways. **Character limit:** **Values:** an address
	AchAddress2 string `json:"AchAddress2,omitempty"`

	//  The name of the bank where the ACH payment account is held. Use this field for ACH payment methods.   **Character limit**: 70   **Values**: a string of 70 characters or fewer
	AchBankName string `json:"AchBankName,omitempty"`

	//  Specifies whether a payment method is available in Zuora. The default value is `false`.   **Character limit**: 5   **Values**: `true`, `false`
	Active bool `json:"Active,omitempty"`

	//  The branch code of the bank used for direct debit. Use this field for direct debit payment methods.   **Character limit**: 10   **Values**:  string of 10 characters or fewer
	BankBranchCode string `json:"BankBranchCode,omitempty"`

	// The check digit in the international bank account number, which confirms the validity of the account. Use this field for direct debit payment methods.  **Character limit**: 4   **Values**:  string of 4 characters or fewer
	BankCheckDigit string `json:"BankCheckDigit,omitempty"`

	//  The city of the direct debit bank. Use this field for direct debit payment methods.   **Character limit**:70   **Values**:  string of 70 characters or fewer
	BankCity string `json:"BankCity,omitempty"`

	//  The sort code or number that identifies the bank. This is also known as the sort code. This field is required for direct debit payment methods.   **Character limit**: 18   **Values**:  string of 18 characters or fewer
	BankCode string `json:"BankCode,omitempty"`

	//  The first six digits of the payment method's number, such as the credit card number or account number. Banks use this number to identify a payment method.   **Character limit**: 6   **Values**:  string of 6 characters or fewer
	BankIdentificationNumber string `json:"BankIdentificationNumber,omitempty"`

	//  The name of the direct debit bank. Use this field for direct debit payment methods.   **Character limit**:80   **Values**:  string of 80 characters or fewer
	BankName string `json:"BankName,omitempty"`

	//  The zip code or postal code of the direct debit bank. Use this field for direct debit payment methods.   **Character limit**:20   **Values**:  string of 20 characters or fewer
	BankPostalCode string `json:"BankPostalCode,omitempty"`

	//  The name of the street of the direct debit bank. Use this field for direct debit payment methods.   **Character limit**:60   **Values**:  string of 60 characters or fewer
	BankStreetName string `json:"BankStreetName,omitempty"`

	//  The number of the direct debit bank. Use this field for direct debit payment methods.   **Character limit**:10   **Values**:  string of 10 characters or fewer
	BankStreetNumber string `json:"BankStreetNumber,omitempty"`

	//  The name on the direct debit bank account. Use this field for direct debit payment methods.   **Character limit**: 60   **Values**:  string of 60 characters or fewer
	BankTransferAccountName string `json:"BankTransferAccountName,omitempty"`

	//  The number of the customer's bank account. Use this field for direct debit payment methods.   **Character limit**:30   **Values**:  string of 30 characters or fewer
	BankTransferAccountNumber string `json:"BankTransferAccountNumber,omitempty"`

	//  This is a masked displayable version of the ACH account number, used for security purposes. For example: `XXXXXXXXX54321`.   **Character limit**: 32   **Values**: automatically generated
	BankTransferAccountNumberMask string `json:"BankTransferAccountNumberMask,omitempty"`

	//  The type of the customer's bank account. Use this field for direct debit payment methods.   **Character limit**: 11   **Values**: `DirectDebit`
	BankTransferAccountType string `json:"BankTransferAccountType,omitempty"`

	//  Specifies the type of direct debit transfer. The value of this field is dependent on the country of the user. Use this field is used for direct debit payment methods.   **Character limit**: 20   **Values**:  - `AutomatischIncasso` (NL) - `LastschriftDE` (Germany) - `LastschriftAT` (Austria) - `DemandeDePrelevement` (FR) - `DirectDebitUK` (UK) - `Domicil` (Belgium) - `LastschriftCH` (CH) - `RID` (Italy) - `OrdenDeDomiciliacion` (Spain)
	BankTransferType string `json:"BankTransferType,omitempty"`

	//  The business identification code for Swiss direct payment methods that use the Global Collect payment gateway. Use this field only for direct debit payments in Switzerland with Global Collect.   **Character limit**: 11   **Values**: string of 11 characters or fewer
	BusinessIdentificationCode string `json:"BusinessIdentificationCode,omitempty"`

	//  The city of the customer's address. Use this field for direct debit payment methods.   **Character limit**:80   **Values**:  string of 80 characters or fewer
	City string `json:"City,omitempty"`

	//  The two-letter country code of the customer's address. Use this field for direct debit payment methods.   **Character limit**: 2   **Values**: a valid country code
	Country string `json:"Country,omitempty"`

	//  The user ID of the person who created the `PaymentMethod` object when there is a login user in the user session. In Hosted Payment Method and Z-Checkout pages, this field is set to 3 as there is no login user to initiate a user session.   **Character limit**: 32   **Values**: automatically generated
	CreatedById string `json:"CreatedById,omitempty"`

	//  The date when the `PaymentMethod` object was created in the Zuora system.   **Character limit**: 29   **Values**: automatically generated
	CreatedDate time.Time `json:"CreatedDate,omitempty"`

	//  The first line of the card holder's address, which is often a street address or business name. Use this field for credit card and direct debit payment methods.   **Character limit**: 255   **Values**: a string of 255 characters or fewer
	CreditCardAddress1 string `json:"CreditCardAddress1,omitempty"`

	//  The second line of the card holder's address. Use this field for credit card and direct debit payment methods.   **Character limit**: 255   **Values**: a string of 255 characters or fewer
	CreditCardAddress2 string `json:"CreditCardAddress2,omitempty"`

	//  The city of the card holder's address. Use this field for credit card and direct debit payment methods  **Character limit**: 40   **Values**: a string of 40 characters or fewer
	CreditCardCity string `json:"CreditCardCity,omitempty"`

	//  The country of the card holder's address.
	CreditCardCountry string `json:"CreditCardCountry,omitempty"`

	//  The expiration month of the credit card or debit card. Use this field for credit card and direct debit payment methods.   **Character limit**: 2   **Values**: a two-digit number, 01 - 12
	CreditCardExpirationMonth int32 `json:"CreditCardExpirationMonth,omitempty"`

	//  The expiration month of the credit card or debit card. Use this field for credit card and direct debit payment methods.   **Character limit**: 4   **Values**: a four-digit number
	CreditCardExpirationYear int32 `json:"CreditCardExpirationYear,omitempty"`

	//  The full name of the card holder. Use this field for credit card and direct debit payment methods.   **Character limit**: 50   **Values**: a string of 50 characters or fewer
	CreditCardHolderName string `json:"CreditCardHolderName,omitempty"`

	//  A masked version of the credit or debit card number.   **Character limit**: 32   **Values**: automatically generated
	CreditCardMaskNumber string `json:"CreditCardMaskNumber,omitempty"`

	//  The credit card or debit card number. This is an insert-only field; it cannot be updated nor queried for security purposes. This field is required only when you define a debit card or credit card payment.   **Character limit**: 16   **Values**: a string of 16 characters or fewer
	CreditCardNumber string `json:"CreditCardNumber,omitempty"`

	//  The billing address's zip code. This field is required only when you define a debit card or credit card payment.   **Character limit**: 20   **Values**: a string of 20 characters or fewer
	CreditCardPostalCode string `json:"CreditCardPostalCode,omitempty"`

	//  The CVV or CVV2 security code. See [How do I control what information Zuora sends over to the Payment Gateway?](https://knowledgecenter.zuora.com/kb/How_do_I_control_information_sent_to_payment_gateways_when_verifying_payment_methods%3F) for more information. To ensure PCI compliance, this value is not stored and cannot be queried.   **Character limit**:   **Values**: a valid CVV or CVV2 security code
	CreditCardSecurityCode string `json:"CreditCardSecurityCode,omitempty"`

	//  The billing address's state. Use this field is if the `CreditCardCountry` value is either Canada or the US. State names must be spelled in full.
	CreditCardState string `json:"CreditCardState,omitempty"`

	//  The type of credit card or debit card. This field is required only when you define a debit card or credit card payment.   **Character limit**: 32   **Values**: `AmericanExpress`, `Discover`, `MasterCard`, `Visa`
	CreditCardType string `json:"CreditCardType,omitempty"`

	//  The session ID of the user when the `PaymentMethod` was created or updated. Some gateways use this field for fraud prevention. If this field is passed to Zuora, then Zuora passes this field to supported gateways. Currently only Verifi supports this field.   **Character limit**: 255   **Values**:
	DeviceSessionId string `json:"DeviceSessionId,omitempty"`

	//  An email address for the payment method in addition to the bill to contact email address.   **Character limit**: 80   **Values**: a string of 80 characters or fewer
	Email string `json:"Email,omitempty"`

	//  Indicates if the customer has an existing mandate or a new mandate. A mandate is a signed authorization for UK and NL customers. When you are migrating mandates from another system, be sure to set this field correctly. If you indicate that a new mandate is an existing mandate or vice-versa, then transactions fail. This field is used only for the direct debit payment method.   **Character limit**: 3   **Values**: `Yes`, `No`
	ExistingMandate string `json:"ExistingMandate,omitempty"`

	//  The customer's first name. This field is used only for the direct debit payment method.   **Character limit**: 30   **Values**: a string of 30 characters or fewer
	FirstName string `json:"FirstName,omitempty"`

	GatewayOptionData SubscribeRequestPaymentMethodGatewayOptionData `json:"GatewayOptionData,omitempty"`

	//  The International Bank Account Number. This field is used only for the direct debit payment method.   **Character limit**: 42   **Values**: a string of 42 characters or fewer
	IBAN string `json:"IBAN,omitempty"`

	//  The IP address of the user when the payment method was created or updated. Some gateways use this field for fraud prevention. If this field is passed to Zuora, then Zuora passes this field to supported gateways. Currently PayPal, CyberSource, Authorize.Net, and Verifi support this field.   **Character limit**: 15   **Values**: a string of 15 characters or fewer
	IPAddress string `json:"IPAddress,omitempty"`

	//  The ID of this object. Upon creation, the ID of this object is `PaymentMethodId`.   **Character limit**: 32   **Values**: automatically generated
	Id string `json:"Id,omitempty"`

	//  The date of the last failed attempt to collect payment with this payment method.   **Character limit**: 29   **Values**: automatically generated
	LastFailedSaleTransactionDate time.Time `json:"LastFailedSaleTransactionDate,omitempty"`

	//  The customer's last name. This field is used only for the direct debit payment method.   **Character limit**: 70   **Values**: a string of 70 characters or fewer
	LastName string `json:"LastName,omitempty"`

	//  The date of the most recent transaction.   **Character limit**: 29   **Values**: a valid date and time value
	LastTransactionDateTime time.Time `json:"LastTransactionDateTime,omitempty"`

	//  The status of the most recent transaction.   **Character limit**: 39   **Values**: automatically generated
	LastTransactionStatus string `json:"LastTransactionStatus,omitempty"`

	//  The date when the mandate was created, in `yyyy-mm-dd` format. A mandate is a signed authorization for UK and NL customers. This field is used only for the direct debit payment method.   **Character limit**: 29
	MandateCreationDate time.Time `json:"MandateCreationDate,omitempty"`

	//  The ID of the mandate. A mandate is a signed authorization for UK and NL customers. This field is used only for the direct debit payment method.   **Character limit**: 36   **Values**: a string of 36 characters or fewer
	MandateID string `json:"MandateID,omitempty"`

	//  Indicates if  the mandate was received. A mandate is a signed authorization for UK and NL customers. This field is used only for the direct debit payment method.   **Character limit**: 3   **Values**: `Yes`, `No `(case-sensitive)
	MandateReceived string `json:"MandateReceived,omitempty"`

	//  The date when the mandate was last updated, in `yyyy-mm-dd` format. A mandate is a signed authorization for UK and NL customers. This field is used only for the direct debit payment method.   **Character limit**: 29
	MandateUpdateDate time.Time `json:"MandateUpdateDate,omitempty"`

	//  Specifies the number of allowable consecutive failures Zuora attempts with the payment method before stopping.    **Values**: a valid number
	MaxConsecutivePaymentFailures int32 `json:"MaxConsecutivePaymentFailures,omitempty"`

	//  Create Query Delete Filter
	Name string `json:"Name,omitempty"`

	//  The number of consecutive failed payment for this payment method. It is reset to 0 upon successful payment. You can use the API to update the field value to 0.   **Character limit**:   **Values**: a positive whole number
	NumConsecutiveFailures int32 `json:"NumConsecutiveFailures,omitempty"`

	//  Specifies the status of the payment method. It is set to Active on creation.   **Character limit**: 6   **Values**: `Active` or `Closed` PaymentMethodStatus should not be used in the `create ` call. You can only set this field to **Closed** via the `update ` call.
	PaymentMethodStatus string `json:"PaymentMethodStatus,omitempty"`

	//  The retry interval setting, which prevents making a payment attempt if the last failed attempt was within the last specified number of hours. This field is required if the `UseDefaultRetryRule` field value is set to `false`.   **Character limit**: 4   **Values**: a whole number between 1 and 1000, exclusive
	PaymentRetryWindow int32 `json:"PaymentRetryWindow,omitempty"`

	//  The PayPal billing agreement ID, which is a contract between two PayPal accounts. Typically, the selling party initiates a request to create a BAID, and sends it to buying party for acceptance. The seller can keep track of the BAID and use it for future charges against the buyer. This field is required when defining a PayPal payment method.   **Character limit**: 64   **Values**: a string of 64 characters or fewer
	PaypalBaid string `json:"PaypalBaid,omitempty"`

	//  The email address associated with the account holder's PayPal account or of the PayPal account of the person paying for the service. This field is required only when you define a PayPal payment method.   **Character limit**: 80   **Values**: a string of 80 characters or fewer
	PaypalEmail string `json:"PaypalEmail,omitempty"`

	//  PayPal's Adaptive Payments API key. Zuora does not create this key, nor does it call PayPal to generate it. You must use PayPal's Adaptive Payments' API to generate this key, and then pass it to Zuora. Zuora uses this key to authorize future payments to PayPal's Adaptive Payments API. This field is required when you use PayPal Adaptive Payments gateway.   **Character limit**: 32   **Values**: a valid PayPal Adaptive Payment pre-approval key
	PaypalPreapprovalKey string `json:"PaypalPreapprovalKey,omitempty"`

	//  Specifies the PayPal gateway: PayFlow Pro (Express Checkout) or Adaptive Payments. This field is required when you use PayPal Adaptive Payments or Payflow Pro (Express Checkout) gateways.   **Character limit**: 32   **Values**: `ExpressCheckout`, `AdaptivePayments`
	PaypalType string `json:"PaypalType,omitempty"`

	//  The phone number that the account holder registered with the bank. This field is used for credit card validation when passing to a gateway.   **Character limit**: 40   **Values**: a string of 40 characters or fewer
	Phone string `json:"Phone,omitempty"`

	//  The zip code of the customer's address. This field is used only for the direct debit payment method.   **Character limit**: 20   **Values**: a string of 20 characters or fewer
	PostalCode string `json:"PostalCode,omitempty"`

	//  A gateway unique identifier that replaces sensitive payment method data. SecondTokenId is conditionally required only when TokenID is being used to represent a gateway customer profile. SecondTokenId is used in the CC Reference Transaction payment method.  **Character limit**: 64   **Values**: a string of 64 characters or fewer
	SecondTokenId string `json:"SecondTokenId,omitempty"`

	//  Creates the payment method even if authorization fails with the payment gateway.   **Character limit**: 5   **Values**: `t``rue`, `false`
	SkipValidation bool `json:"SkipValidation,omitempty"`

	//  The state of the customer's address. This field is used only for the direct debit payment method.   **Character limit**: 70   **Values**: a string of 70 characters or fewer
	State string `json:"State,omitempty"`

	//  The street name of the customer's address. This field is used only for the direct debit payment method.   **Character limit**: 100   **Values**: a string of 100 characters or fewer
	StreetName string `json:"StreetName,omitempty"`

	//  The street number of the customer's address. This field is used only for the direct debit payment method.   **Character limit**: 30   **Values**: a string of 30 characters or fewer
	StreetNumber string `json:"StreetNumber,omitempty"`

	//  A gateway unique identifier that replaces sensitive payment method data or represents a gateway's unique customer profile. When TokenId is used to represent a customer profile, then SecondTokenId is conditionally required for representing the underlying tokenized payment method. TokenId is required for the CC Reference Transaction payment method.   **Character limit**: 255   **Values**: a string of 255 characters or fewer
	TokenId string `json:"TokenId,omitempty"`

	//  The number of error payments that used this payment method.   **Character limit**:   **Values**: automatically generated
	TotalNumberOfErrorPayments int32 `json:"TotalNumberOfErrorPayments,omitempty"`

	//  The number of successful payments that used this payment method.   **Character limit**:   **Values**: automatically generated
	TotalNumberOfProcessedPayments int32 `json:"TotalNumberOfProcessedPayments,omitempty"`

	//  Create Query Update Delete Filter
	Type_ string `json:"Type,omitempty"`

	//  The ID of the user who last updated the payment method.   **Character limit**: 32   **Values**: automatically generated
	UpdatedById string `json:"UpdatedById,omitempty"`

	//  The date when the payment method was last updated.   **Character limit**: 29   **Values**: automatically generated
	UpdatedDate time.Time `json:"UpdatedDate,omitempty"`

	//  Determines whether to use the default retry rules configured in the [Z-Payments settings](https://knowledgecenter.zuora.com/CB_Billing/L_Payment_Methods/H_Configure_Payment_Method_Retry_Rules). Set this to `true` to use the default retry rules. Set this to `false` to set the specific rules for this payment method. If you set this value to `false`, then the fields, `PaymentRetryWindow` and `MaxConsecutivePaymentFailures`, are required.   **Character limit**: 5   **Values**: `t``rue`, `false`
	UseDefaultRetryRule bool `json:"UseDefaultRetryRule,omitempty"`
}
