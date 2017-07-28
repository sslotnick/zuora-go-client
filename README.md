# Go API client for zuora

  # Introduction Welcome to the reference for the Zuora REST API!  <a href=\"http://en.wikipedia.org/wiki/REST_API\" target=\"_blank\">REST</a> is a web-service protocol that lends itself to rapid development by using everyday HTTP and JSON technology.  The Zuora REST API provides a broad set of operations and resources that:    * Enable Web Storefront integration from your website.   * Support self-service subscriber sign-ups and account management.   * Process revenue schedules through custom revenue rule models.   * Enable manipulation of most objects in the Zuora Object Model.      ## Endpoints      The Zuora REST API is provided via the following endpoints.   | Tenant              | Base URL for REST Endpoints |   |-------------------------|-------------------------|   |US Production | https://rest.zuora.com   |   |US API Sandbox    | https://rest.apisandbox.zuora.com|   |US Performance Test | https://rest.pt1.zuora.com |   |EU Production | https://rest.eu.zuora.com |   |EU Sandbox | https://rest.sandbox.eu.zuora.com |      The production endpoint provides access to your live user data. The API Sandbox tenant is a good place to test your code without affecting real-world data. To use it, you must be provisioned with an API Sandbox tenant - your Zuora representative can help you if needed.      ## Access to the API      If you have a Zuora tenant, you already have access to the API.      If you don't have a Zuora tenant, go to <a href=\" https://www.zuora.com/resource/zuora-test-drive\" target=\"_blank\">https://www.zuora.com/resource/zuora-test-drive</a> and sign up for a Production Test Drive tenant. The tenant comes with seed data, such as a sample product catalog.  We recommend that you <a href=\"https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/Manage_Users/Create_an_API_User\" target=\"_blank\">create an API user</a> specifically for making API calls. Don't log in to the Zuora UI with this account. Logging in to the UI enables a security feature that periodically expires the account's password, which may eventually cause authentication failures with the API. Note that a user role does not have write access to Zuora REST services unless it has the API Write Access permission as described in those instructions.   # API Changelog You can find the <a href=\"https://community.zuora.com/t5/Developers/API-Changelog/gpm-p/18092\" target=\"_blank\">Changelog</a> of the API Reference in the Zuora Community.   # Authentication  There are three ways to authenticate:    * Use username and password. Include authentication with each request in the header:         * `apiAccessKeyId`      * `apiSecretAccessKey`     * `entityId` or `entityName` (Only for [Zuora Multi-entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity \"Multi-entity\"). See \"Entity Id and Entity Name\" below for more information.)   * Use an authorization cookie. The cookie authorizes the user to make calls to the REST API for the duration specified in  **Administration > Security Policies > Session timeout**. The cookie expiration time is reset with this duration after every call to the REST API. To obtain a cookie, call the REST  `connections` resource with the following API user information:         *   ID         *   password     *   entity Id or entity name (Only for [Zuora Multi-entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity \"Multi-entity\"). See \"Entity Id and Entity Name\" below for more information.)         * For CORS-enabled APIs only: Include a 'single-use' token in the request header, which re-authenticates the user with each request. See below for more details.  ## Entity Id and Entity Name  The `entityId` and `entityName` parameters are only used for [Zuora Multi-entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity \"Zuora Multi-entity\").   The  `entityId` and `entityName` parameters specify the Id and the [name of the entity](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/B_Introduction_to_Entity_and_Entity_Hierarchy#Name_and_Display_Name \"Introduction to Entity and Entity Hierarchy\") that you want to access, respectively. Note that you must have permission to access the entity.   You can specify either the `entityId` or `entityName` parameter in the authentication to access and view an entity.    * If both `entityId` and `entityName` are specified in the authentication, an error occurs.    * If neither `entityId` nor `entityName` is specified in the authentication, you will log in to the entity in which your user account is created.      To get the entity Id and entity name, you can use the GET Entities REST call. For more information, see [API User Authentication](https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Multi-entity/A_Overview_of_Multi-entity#API_User_Authentication \"API User Authentication\").      ## Token Authentication for CORS-Enabled APIs      The CORS mechanism enables REST API calls to Zuora to be made directly from your customer's browser, with all credit card and security information transmitted directly to Zuora. This minimizes your PCI compliance burden, allows you to implement advanced validation on your payment forms, and  makes your payment forms look just like any other part of your website.    For security reasons, instead of using cookies, an API request via CORS uses **tokens** for authentication.  The token method of authentication is only designed for use with requests that must originate from your customer's browser; **it should  not be considered a replacement to the existing cookie authentication** mechanism.  See [Zuora CORS REST](https://knowledgecenter.zuora.com/DC_Developers/REST_API/A_REST_basics/G_CORS_REST \"Zuora CORS REST\") for details on how CORS works and how you can begin to implement customer calls to the Zuora REST APIs. See  [HMAC Signatures](https://www.zuora.com/developer/API-Reference/#operation/POSTHMACSignature \"HMAC Signatures\") for details on the HMAC method that returns the authentication token.  # Requests and Responses  ## Request IDs  As a general rule, when asked to supply a \"key\" for an account or subscription (accountKey, account-key, subscriptionKey, subscription-key), you can provide either the actual ID or  the number of the entity.  ## HTTP Request Body  Most of the parameters and data accompanying your requests will be contained in the body of the HTTP request.   The Zuora REST API accepts JSON in the HTTP request body. No other data format (e.g., XML) is supported.  ### Data Type  ([Actions](https://www.zuora.com/developer/api-reference/#tag/Actions) and CRUD operations only) We recommend that you do not specify the decimal values with quotation marks, commas, and spaces. Use characters of `+-0-9.eE`, for example, `5`, `1.9`, `-8.469`, and `7.7e2`. Also, Zuora does not convert currencies for decimal values.  ## Testing a Request  Use a third party client, such as [curl](https://curl.haxx.se \"curl\"), [Postman](https://www.getpostman.com \"Postman\"), or [Advanced REST Client](https://advancedrestclient.com \"Advanced REST Client\"), to test the Zuora REST API.  You can test the Zuora REST API from the Zuora API Sandbox or Production tenants. If connecting to Production, bear in mind that you are working with your live production data, not sample data or test data.  ## Testing with Credit Cards  Sooner or later it will probably be necessary to test some transactions that involve credit cards. For suggestions on how to handle this, see [Going Live With Your Payment Gateway](https://knowledgecenter.zuora.com/CB_Billing/M_Payment_Gateways/C_Managing_Payment_Gateways/B_Going_Live_Payment_Gateways#Testing_with_Credit_Cards \"C_Zuora_User_Guides/A_Billing_and_Payments/M_Payment_Gateways/C_Managing_Payment_Gateways/B_Going_Live_Payment_Gateways#Testing_with_Credit_Cards\" ).  ## Concurrent Request Limits  Zuora enforces tenant-level concurrent request limits. See <a href=\"https://knowledgecenter.zuora.com/BB_Introducing_Z_Business/Policies/Concurrent_Request_Limits\" target=\"_blank\">Concurrent Request Limits</a> for more information.    ## Error Handling  Responses and error codes are detailed in [Responses and errors](https://knowledgecenter.zuora.com/DC_Developers/REST_API/A_REST_basics/3_Responses_and_errors \"Responses and errors\").  # Pagination  When retrieving information (using GET methods), the optional `pageSize` query parameter sets the maximum number of rows to return in a response. The maximum is `40`; larger values are treated as `40`. If this value is empty or invalid, `pageSize` typically defaults to `10`.  The default value for the maximum number of rows retrieved can be overridden at the method level.  If more rows are available, the response will include a `nextPage` element, which contains a URL for requesting the next page.  If this value is not provided, no more rows are available. No \"previous page\" element is explicitly provided; to support backward paging, use the previous call.  ## Array Size  For data items that are not paginated, the REST API supports arrays of up to 300 rows.  Thus, for instance, repeated pagination can retrieve thousands of customer accounts, but within any account an array of no more than 300 rate plans is returned.  # API Versions  The Zuora REST API are version controlled. Versioning ensures that Zuora REST API changes are backward compatible. Zuora uses a major and minor version nomenclature to manage changes. By specifying a version in a REST request, you can get expected responses regardless of future changes to the API.  ## Major Version  The major version number of the REST API appears in the REST URL. Currently, Zuora only supports the **v1** major version. For example, `POST https://rest.zuora.com/v1/subscriptions`.  ## Minor Version  Zuora uses minor versions for the REST API to control small changes. For example, a field in a REST method is deprecated and a new field is used to replace it.   Some fields in the REST methods are supported as of minor versions. If a field is not noted with a minor version, this field is available for all minor versions. If a field is noted with a minor version, this field is in version control. You must specify the supported minor version in the request header to process without an error.   If a field is in version control, it is either with a minimum minor version or a maximum minor version, or both of them. You can only use this field with the minor version between the minimum and the maximum minor versions. For example, the `invoiceCollect` field in the POST Subscription method is in version control and its maximum minor version is 189.0. You can only use this field with the minor version 189.0 or earlier.  If you specify a version number in the request header that is not supported, Zuora will use the minimum minor version of the REST API. In our REST API documentation, if a field or feature requires a minor version number, we note that in the field description.  You only need to specify the version number when you use the fields require a minor version. To specify the minor version, set the `zuora-version` parameter to the minor version number in the request header for the request call. For example, the `collect` field is in 196.0 minor version. If you want to use this field for the POST Subscription method, set the  `zuora-version` parameter to `196.0` in the request header. The `zuora-version` parameter is case sensitive.  For all the REST API fields, by default, if the minor version is not specified in the request header, Zuora will use the minimum minor version of the REST API to avoid breaking your integration.   ### Minor Version History  The supported minor versions are not serial. This section documents the changes made to each Zuora REST API minor version.  The following table lists the supported versions and the fields that have a Zuora REST API minor version.  | Fields         | Minor Version      | REST Methods    | Description | |:--------|:--------|:--------|:--------| | invoiceCollect | 189.0 and earlier  | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Generates an invoice and collects a payment for a subscription. | | collect        | 196.0 and later    | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Collects an automatic payment for a subscription. | | invoice | 196.0 and 207.0| [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Generates an invoice for a subscription. | | invoiceTargetDate | 196.0 and earlier  | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\") |Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | invoiceTargetDate | 207.0 and earlier  | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | targetDate | 207.0 and later | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\") |Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | targetDate | 211.0 and later | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Date through which charges are calculated on the invoice, as `yyyy-mm-dd`. | | includeExisting DraftInvoiceItems | 196.0 and earlier| [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | Specifies whether to include draft invoice items in subscription previews. Specify it to be `true` (default) to include draft invoice items in the preview result. Specify it to be `false` to excludes draft invoice items in the preview result. | | includeExisting DraftDocItems | 207.0 and later  | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | Specifies whether to include draft invoice items in subscription previews. Specify it to be `true` (default) to include draft invoice items in the preview result. Specify it to be `false` to excludes draft invoice items in the preview result. | | previewType | 196.0 and earlier| [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | The type of preview you will receive. The possible values are `InvoiceItem`(default), `ChargeMetrics`, and `InvoiceItemChargeMetrics`. | | previewType | 207.0 and later  | [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") | The type of preview you will receive. The possible values are `LegalDoc`(default), `ChargeMetrics`, and `LegalDocChargeMetrics`. | | runBilling  | 211.0 and later  | [Create Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_Subscription \"Create Subscription\"); [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\"); [Renew Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_RenewSubscription \"Renew Subscription\"); [Cancel Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_CancelSubscription \"Cancel Subscription\"); [Suspend Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_SuspendSubscription \"Suspend Subscription\"); [Resume Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_ResumeSubscription \"Resume Subscription\"); [Create Account](https://www.zuora.com/developer/API-Reference/#operation/POST_Account \"Create Account\")|Generates an invoice or credit memo for a subscription. **Note:** Credit memos are only available if you have the Advanced AR Settlement feature enabled. |  #### Version 207.0 and Later  The response structure of the [Preview Subscription](https://www.zuora.com/developer/api-reference/#operation/POST_SubscriptionPreview \"Preview Subscription\") and [Update Subscription](https://www.zuora.com/developer/api-reference/#operation/PUT_Subscription \"Update Subscription\") methods are changed. The following invoice related response fields are moved to the invoice container:    * amount   * amountWithoutTax   * taxAmount   * invoiceItems   * targetDate   * chargeMetrics  # Zuora Object Model  The following diagram presents a high-level view of the key Zuora objects. Click the image to open it in a new tab to resize it.  <a href=\"https://www.zuora.com/wp-content/uploads/2017/01/ZuoraERD.jpeg\" target=\"_blank\"><img src=\"https://www.zuora.com/wp-content/uploads/2017/01/ZuoraERD.jpeg\" alt=\"Zuora Object Model Diagram\"></a>  You can use the [Describe object](https://www.zuora.com/developer/api-reference/#operation/GET_Describe) operation to list the fields of each Zuora object that is available in your tenant. When you call the operation, you must specify the API name of the Zuora object.  The following table provides the API name of each Zuora object:  | Object                                        | API Name                                   | |-----------------------------------------------|--------------------------------------------| | Account                                       | `Account`                                  | | Accounting Code                               | `AccountingCode`                           | | Accounting Period                             | `AccountingPeriod`                         | | Amendment                                     | `Amendment`                                | | Application Group                             | `ApplicationGroup`                         | | Contact                                       | `Contact`                                  | | Contact Snapshot                              | `ContactSnapshot`                          | | Billing Run                                   | `BillingRun`                               | | Credit Balance Adjustment                     | `CreditBalanceAdjustment`                  | | Credit Memo                                   | `CreditMemo`                               | | Credit Memo Application                       | `CreditMemoApplication`                    | | Credit Memo Application Item                  | `CreditMemoApplicationItem`                | | Credit Memo Item                              | `CreditMemoItem`                           | | Credit Memo Part                              | `CreditMemoPart`                           | | Credit Memo Part Item                         | `CreditMemoPartItem`                       | | Credit Taxation Item                          | `CreditTaxationItem`                       | | Custom Exchange Rate                          | `FXCustomRate`                             | | Debit Memo                                    | `DebitMemo`                                | | Debit Memo Item                               | `DebitMemoItem`                            | | Debit Taxation Item                           | `DebitTaxationItem`                        | | Payment Method                                | `PaymentMethod`                            | | Entity                                        | `Tenant`                                   | | Gateway Reconciliation Event                  | `PaymentGatewayReconciliationEventLog`     | | Gateway Reconciliation Job                    | `PaymentReconciliationJob`                 | | Gateway Reconciliation Log                    | `PaymentReconciliationLog`                 | | Invoice                                       | `Invoice`                                  | | Invoice Adjustment                            | `InvoiceAdjustment`                        | | Invoice Item                                  | `InvoiceItem`                              | | Invoice Item Adjustment                       | `InvoiceItemAdjustment`                    | | Invoice Payment                               | `InvoicePayment`                           | | Journal Entry                                 | `JournalEntry`                             | | Journal Entry Item                            | `JournalEntryItem`                         | | Journal Run                                   | `JournalRun`                               | | Payment                                       | `Payment`                                  | | Payment Application                           | `PaymentApplication`                       | | Payment Application Item                      | `PaymentApplicationItem`                   | | Payment Method Snapshot                       | `PaymentMethodSnapshot`                    | | Payment Method Transaction Log                | `PaymentMethodTransactionLog`              | | Payment Method Update                         | `UpdaterDetail`                            | | Payment Part                                  | `PaymentPart`                              | | Payment Part Item                             | `PaymentPartItem`                          | | Payment Transaction Log                       | `PaymentTransactionLog`                    | | Processed Usage                               | `ProcessedUsage`                           | | Product                                       | `Product`                                  | | Product Rate Plan                             | `ProductRatePlan`                          | | Product Rate Plan Charge                      | `ProductRatePlanCharge`                    | | Product Rate Plan Charge Tier                 | `ProductRatePlanChargeTier`                | | Rate Plan                                     | `RatePlan`                                 | | Rate Plan Charge                              | `RatePlanCharge`                           | | Rate Plan Charge Tier                         | `RatePlanChargeTier`                       | | Refund                                        | `Refund`                                   | | Refund Application                            | `RefundApplication`                        | | Refund Application Item                       | `RefundApplicationItem`                    | | Refund Invoice Payment                        | `RefundInvoicePayment`                     | | Refund Part                                   | `RefundPart`                               | | Refund Part Item                              | `RefundPartItem`                           | | Refund Transaction Log                        | `RefundTransactionLog`                     | | Revenue Charge Summary                        | `RevenueChargeSummary`                     | | Revenue Charge Summary Item                   | `RevenueChargeSummaryItem`                 | | Revenue Event                                 | `RevenueEvent`                             | | Revenue Event Credit Memo Item                | `RevenueEventCreditMemoItem`               | | Revenue Event Debit Memo Item                 | `RevenueEventDebitMemoItem`                | | Revenue Event Invoice Item                    | `RevenueEventInvoiceItem`                  | | Revenue Event Invoice Item Adjustment         | `RevenueEventInvoiceItemAdjustment`        | | Revenue Event Item                            | `RevenueEventItem`                         | | Revenue Event Item Credit Memo Item           | `RevenueEventItemCreditMemoItem`           | | Revenue Event Item Debit Memo Item            | `RevenueEventItemDebitMemoItem`            | | Revenue Event Item Invoice Item               | `RevenueEventItemInvoiceItem`              | | Revenue Event Item Invoice Item Adjustment    | `RevenueEventItemInvoiceItemAdjustment`    | | Revenue Event Type                            | `RevenueEventType`                         | | Revenue Schedule                              | `RevenueSchedule`                          | | Revenue Schedule Credit Memo Item             | `RevenueScheduleCreditMemoItem`            | | Revenue Schedule Debit Memo Item              | `RevenueScheduleDebitMemoItem`             | | Revenue Schedule Invoice Item                 | `RevenueScheduleInvoiceItem`               | | Revenue Schedule Invoice Item Adjustment      | `RevenueScheduleInvoiceItemAdjustment`     | | Revenue Schedule Item                         | `RevenueScheduleItem`                      | | Revenue Schedule Item Credit Memo Item        | `RevenueScheduleItemCreditMemoItem`        | | Revenue Schedule Item Debit Memo Item         | `RevenueScheduleItemDebitMemoItem`         | | Revenue Schedule Item Invoice Item            | `RevenueScheduleItemInvoiceItem`           | | Revenue Schedule Item Invoice Item Adjustment | `RevenueScheduleItemInvoiceItemAdjustment` | | Subscription                                  | `Subscription`                             | | Taxation Item                                 | `TaxationItem`                             | | Updater Batch                                 | `UpdaterBatch`                             | | Usage                                         | `Usage`                                    | 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2017-07-12
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./zuora"
```

## Documentation for API Endpoints

All URIs are relative to *https://rest.zuora.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountingCodesApi* | [**DELETEAccountingCode**](docs/AccountingCodesApi.md#deleteaccountingcode) | **Delete** /v1/accounting-codes/{ac-id} | Delete accounting code
*AccountingCodesApi* | [**GETAccountingCode**](docs/AccountingCodesApi.md#getaccountingcode) | **Get** /v1/accounting-codes/{ac-id} | Query an accounting code
*AccountingCodesApi* | [**GETAllAccountingCodes**](docs/AccountingCodesApi.md#getallaccountingcodes) | **Get** /v1/accounting-codes | Get all accounting codes
*AccountingCodesApi* | [**POSTAccountingCode**](docs/AccountingCodesApi.md#postaccountingcode) | **Post** /v1/accounting-codes | Create accounting code
*AccountingCodesApi* | [**PUTAccountingCode**](docs/AccountingCodesApi.md#putaccountingcode) | **Put** /v1/accounting-codes/{ac-id} | Update an accounting code
*AccountingCodesApi* | [**PUTActivateAccountingCode**](docs/AccountingCodesApi.md#putactivateaccountingcode) | **Put** /v1/accounting-codes/{ac-id}/activate | Activate accounting code
*AccountingCodesApi* | [**PUTDeactivateAccountingCode**](docs/AccountingCodesApi.md#putdeactivateaccountingcode) | **Put** /v1/accounting-codes/{ac-id}/deactivate | Deactivate accounting code
*AccountingPeriodsApi* | [**DELETEAccountingPeriod**](docs/AccountingPeriodsApi.md#deleteaccountingperiod) | **Delete** /v1/accounting-periods/{ap-id} | Delete accounting period
*AccountingPeriodsApi* | [**GETAccountingPeriod**](docs/AccountingPeriodsApi.md#getaccountingperiod) | **Get** /v1/accounting-periods/{ap-id} | Get accounting period
*AccountingPeriodsApi* | [**GETAllAccountingPeriods**](docs/AccountingPeriodsApi.md#getallaccountingperiods) | **Get** /v1/accounting-periods | Get all accounting periods
*AccountingPeriodsApi* | [**POSTAccountingPeriod**](docs/AccountingPeriodsApi.md#postaccountingperiod) | **Post** /v1/accounting-periods | Create accounting period
*AccountingPeriodsApi* | [**PUTCloseAccountingPeriod**](docs/AccountingPeriodsApi.md#putcloseaccountingperiod) | **Put** /v1/accounting-periods/{ap-id}/close | Close accounting period
*AccountingPeriodsApi* | [**PUTPendingCloseAccountingPeriod**](docs/AccountingPeriodsApi.md#putpendingcloseaccountingperiod) | **Put** /v1/accounting-periods/{ap-id}/pending-close | Set accounting period to pending close
*AccountingPeriodsApi* | [**PUTReopenAccountingPeriod**](docs/AccountingPeriodsApi.md#putreopenaccountingperiod) | **Put** /v1/accounting-periods/{ap-id}/reopen | Re-open accounting period
*AccountingPeriodsApi* | [**PUTRunTrialBalance**](docs/AccountingPeriodsApi.md#putruntrialbalance) | **Put** /v1/accounting-periods/{ap-id}/run-trial-balance | Run trial balance
*AccountingPeriodsApi* | [**PUTUpdateAccountingPeriod**](docs/AccountingPeriodsApi.md#putupdateaccountingperiod) | **Put** /v1/accounting-periods/{ap-id} | Update accounting period
*AccountsApi* | [**GETAccount**](docs/AccountsApi.md#getaccount) | **Get** /v1/accounts/{account-key} | Get account
*AccountsApi* | [**GETAccountSummary**](docs/AccountsApi.md#getaccountsummary) | **Get** /v1/accounts/{account-key}/summary | Get account summary
*AccountsApi* | [**ObjectDELETEAccount**](docs/AccountsApi.md#objectdeleteaccount) | **Delete** /v1/object/account/{id} | CRUD: Delete Account
*AccountsApi* | [**ObjectGETAccount**](docs/AccountsApi.md#objectgetaccount) | **Get** /v1/object/account/{id} | CRUD: Retrieve Account
*AccountsApi* | [**ObjectPOSTAccount**](docs/AccountsApi.md#objectpostaccount) | **Post** /v1/object/account | CRUD: Create Account
*AccountsApi* | [**ObjectPUTAccount**](docs/AccountsApi.md#objectputaccount) | **Put** /v1/object/account/{id} | CRUD: Update Account
*AccountsApi* | [**POSTAccount**](docs/AccountsApi.md#postaccount) | **Post** /v1/accounts | Create account
*AccountsApi* | [**PUTAccount**](docs/AccountsApi.md#putaccount) | **Put** /v1/accounts/{account-key} | Update account
*ActionsApi* | [**ActionPOSTamend**](docs/ActionsApi.md#actionpostamend) | **Post** /v1/action/amend | Amend
*ActionsApi* | [**ActionPOSTcreate**](docs/ActionsApi.md#actionpostcreate) | **Post** /v1/action/create | Create
*ActionsApi* | [**ActionPOSTdelete**](docs/ActionsApi.md#actionpostdelete) | **Post** /v1/action/delete | Delete
*ActionsApi* | [**ActionPOSTexecute**](docs/ActionsApi.md#actionpostexecute) | **Post** /v1/action/execute | Execute
*ActionsApi* | [**ActionPOSTgenerate**](docs/ActionsApi.md#actionpostgenerate) | **Post** /v1/action/generate | Generate
*ActionsApi* | [**ActionPOSTquery**](docs/ActionsApi.md#actionpostquery) | **Post** /v1/action/query | Query
*ActionsApi* | [**ActionPOSTqueryMore**](docs/ActionsApi.md#actionpostquerymore) | **Post** /v1/action/queryMore | QueryMore
*ActionsApi* | [**ActionPOSTsubscribe**](docs/ActionsApi.md#actionpostsubscribe) | **Post** /v1/action/subscribe | Subscribe
*ActionsApi* | [**ActionPOSTupdate**](docs/ActionsApi.md#actionpostupdate) | **Post** /v1/action/update | Update
*AmendmentsApi* | [**GETAmendmentsByKey**](docs/AmendmentsApi.md#getamendmentsbykey) | **Get** /v1/amendments/{amendment-key} | Get amendments by key
*AmendmentsApi* | [**GETAmendmentsBySubscriptionID**](docs/AmendmentsApi.md#getamendmentsbysubscriptionid) | **Get** /v1/amendments/subscriptions/{subscription-id} | Get amendments by subscription ID
*AmendmentsApi* | [**ObjectDELETEAmendment**](docs/AmendmentsApi.md#objectdeleteamendment) | **Delete** /v1/object/amendment/{id} | CRUD: Delete Amendment
*AmendmentsApi* | [**ObjectGETAmendment**](docs/AmendmentsApi.md#objectgetamendment) | **Get** /v1/object/amendment/{id} | CRUD: Retrieve Amendment
*AmendmentsApi* | [**ObjectPOSTAmendment**](docs/AmendmentsApi.md#objectpostamendment) | **Post** /v1/object/amendment | CRUD: Create Amendment
*AmendmentsApi* | [**ObjectPUTAmendment**](docs/AmendmentsApi.md#objectputamendment) | **Put** /v1/object/amendment/{id} | CRUD: Update Amendment
*AttachmentsApi* | [**DELETEAttachments**](docs/AttachmentsApi.md#deleteattachments) | **Delete** /v1/attachments/{attachment-id} | Delete attachments
*AttachmentsApi* | [**GETAttachments**](docs/AttachmentsApi.md#getattachments) | **Get** /v1/attachments/{attachment-id} | View attachments
*AttachmentsApi* | [**GETAttachmentsList**](docs/AttachmentsApi.md#getattachmentslist) | **Get** /v1/attachments/{object-type}/{object-key} | View attachments list
*AttachmentsApi* | [**POSTAttachments**](docs/AttachmentsApi.md#postattachments) | **Post** /v1/attachments | Add attachments
*AttachmentsApi* | [**PUTAttachments**](docs/AttachmentsApi.md#putattachments) | **Put** /v1/attachments/{attachment-id} | Edit attachments
*BillRunApi* | [**ObjectDELETEBillRun**](docs/BillRunApi.md#objectdeletebillrun) | **Delete** /v1/object/bill-run/{id} | CRUD: Delete Bill Run
*BillRunApi* | [**ObjectGETBillRun**](docs/BillRunApi.md#objectgetbillrun) | **Get** /v1/object/bill-run/{id} | CRUD: Retrieve Bill Run
*BillRunApi* | [**ObjectPOSTBillRun**](docs/BillRunApi.md#objectpostbillrun) | **Post** /v1/object/bill-run | CRUD: Create Bill Run
*BillRunApi* | [**ObjectPUTBillRun**](docs/BillRunApi.md#objectputbillrun) | **Put** /v1/object/bill-run/{id} | CRUD: Post or Cancel Bill Run
*BillRunApi* | [**POSTEmailBillingDocumentsfromBillRun**](docs/BillRunApi.md#postemailbillingdocumentsfrombillrun) | **Post** /v1/bill-runs/{billRunId}/emails | Email billing documents generated from bill run
*BillingDocumentsApi* | [**POSTGenerateBillingDocuments**](docs/BillingDocumentsApi.md#postgeneratebillingdocuments) | **Post** /v1/accounts/{id}/billing-documents/generate | Generate billing documents by account
*BillingPreviewRunApi* | [**GETBillingPreviewRun**](docs/BillingPreviewRunApi.md#getbillingpreviewrun) | **Get** /v1/billing-preview-runs/{billingPreviewRunId} | Get Billing Preview Run
*BillingPreviewRunApi* | [**POSTBillingPreviewRun**](docs/BillingPreviewRunApi.md#postbillingpreviewrun) | **Post** /v1/billing-preview-runs | Create Billing Preview Run
*CatalogApi* | [**GETCatalog**](docs/CatalogApi.md#getcatalog) | **Get** /v1/catalog/products | Get product catalog
*CatalogApi* | [**PUTCatalog**](docs/CatalogApi.md#putcatalog) | **Put** /v1/catalog/products/{product-id}/share | Multi-entity: Share a Product with an Entity
*ChargeRevenueSummariesApi* | [**GETCRSByCRSNumber**](docs/ChargeRevenueSummariesApi.md#getcrsbycrsnumber) | **Get** /v1/charge-revenue-summaries/{crs-number} | Get charge summary details by CRS number
*ChargeRevenueSummariesApi* | [**GETCRSByChargeID**](docs/ChargeRevenueSummariesApi.md#getcrsbychargeid) | **Get** /v1/charge-revenue-summaries/subscription-charges/{charge-key} | Get charge summary details by charge ID
*CommunicationProfilesApi* | [**ObjectGETCommunicationProfile**](docs/CommunicationProfilesApi.md#objectgetcommunicationprofile) | **Get** /v1/object/communication-profile/{id} | CRUD: Retrieve CommunicationProfile
*ConnectionsApi* | [**POSTConnections**](docs/ConnectionsApi.md#postconnections) | **Post** /v1/connections | Establish connection to Zuora REST API service
*ContactsApi* | [**ObjectDELETEContact**](docs/ContactsApi.md#objectdeletecontact) | **Delete** /v1/object/contact/{id} | CRUD: Delete Contact
*ContactsApi* | [**ObjectGETContact**](docs/ContactsApi.md#objectgetcontact) | **Get** /v1/object/contact/{id} | CRUD: Retrieve Contact
*ContactsApi* | [**ObjectPOSTContact**](docs/ContactsApi.md#objectpostcontact) | **Post** /v1/object/contact | CRUD: Create Contact
*ContactsApi* | [**ObjectPUTContact**](docs/ContactsApi.md#objectputcontact) | **Put** /v1/object/contact/{id} | CRUD: Update Contact
*CreditBalanceAdjustmentsApi* | [**ObjectGETCreditBalanceAdjustment**](docs/CreditBalanceAdjustmentsApi.md#objectgetcreditbalanceadjustment) | **Get** /v1/object/credit-balance-adjustment/{id} | CRUD: Retrieve CreditBalanceAdjustment
*CreditMemosApi* | [**DELETECreditMemo**](docs/CreditMemosApi.md#deletecreditmemo) | **Delete** /v1/creditmemos/{creditMemoId} | Delete credit memo
*CreditMemosApi* | [**GETCreditMemo**](docs/CreditMemosApi.md#getcreditmemo) | **Get** /v1/creditmemos/{creditMemoId} | Get credit memo
*CreditMemosApi* | [**GETCreditMemoItem**](docs/CreditMemosApi.md#getcreditmemoitem) | **Get** /v1/creditmemos/{creditMemoId}/items/{cmitemid} | Get credit memo item
*CreditMemosApi* | [**GETCreditMemoItemPart**](docs/CreditMemosApi.md#getcreditmemoitempart) | **Get** /v1/creditmemos/{creditMemoId}/parts/{partid}/itemparts/{itempartid} | Get credit memo part item
*CreditMemosApi* | [**GETCreditMemoItemParts**](docs/CreditMemosApi.md#getcreditmemoitemparts) | **Get** /v1/creditmemos/{creditMemoId}/parts/{partid}/itemparts | Get credit memo part items
*CreditMemosApi* | [**GETCreditMemoItems**](docs/CreditMemosApi.md#getcreditmemoitems) | **Get** /v1/creditmemos/{creditMemoId}/items | Get credit memo items
*CreditMemosApi* | [**GETCreditMemoPart**](docs/CreditMemosApi.md#getcreditmemopart) | **Get** /v1/creditmemos/{creditMemoId}/parts/{partid} | Get credit memo part
*CreditMemosApi* | [**GETCreditMemoParts**](docs/CreditMemosApi.md#getcreditmemoparts) | **Get** /v1/creditmemos/{creditMemoId}/parts | Get credit memo parts
*CreditMemosApi* | [**GETCreditMemos**](docs/CreditMemosApi.md#getcreditmemos) | **Get** /v1/creditmemos | Get credit memos
*CreditMemosApi* | [**POSTApplyCreditMemo**](docs/CreditMemosApi.md#postapplycreditmemo) | **Post** /v1/creditmemos/{creditMemoId}/apply | Apply credit memo
*CreditMemosApi* | [**POSTCancelCreditMemo**](docs/CreditMemosApi.md#postcancelcreditmemo) | **Post** /v1/creditmemos/{creditMemoId}/cancel | Cancel credit memo
*CreditMemosApi* | [**POSTCreditMemoFromInvoice**](docs/CreditMemosApi.md#postcreditmemofrominvoice) | **Post** /v1/creditmemos/invoice/{invoiceId} | Create credit memo from invoice
*CreditMemosApi* | [**POSTCreditMemoFromPrpc**](docs/CreditMemosApi.md#postcreditmemofromprpc) | **Post** /v1/creditmemos | Create credit memo from charge
*CreditMemosApi* | [**POSTCreditMemoPDF**](docs/CreditMemosApi.md#postcreditmemopdf) | **Post** /v1/creditmemos/{creditMemoId}/pdfs | Create credit memo PDF
*CreditMemosApi* | [**POSTEmailCreidtMemo**](docs/CreditMemosApi.md#postemailcreidtmemo) | **Post** /v1/creditmemos/{creditMemoId}/emails | Email credit memo
*CreditMemosApi* | [**POSTPostCreditMemo**](docs/CreditMemosApi.md#postpostcreditmemo) | **Post** /v1/creditmemos/{creditMemoId}/post | Post credit memo
*CreditMemosApi* | [**POSTQueryCreditMemos**](docs/CreditMemosApi.md#postquerycreditmemos) | **Post** /v1/creditmemos/query | Query credit memos by account
*CreditMemosApi* | [**POSTRefundCreditMemo**](docs/CreditMemosApi.md#postrefundcreditmemo) | **Post** /v1/creditmemos/{creditmemoId}/refund | Refund credit memo
*CreditMemosApi* | [**POSTUnapplyCreditMemo**](docs/CreditMemosApi.md#postunapplycreditmemo) | **Post** /v1/creditmemos/{creditMemoId}/unapply | Unapply credit memo
*CreditMemosApi* | [**PUTUpdateCreditMemo**](docs/CreditMemosApi.md#putupdatecreditmemo) | **Put** /v1/creditmemos/{creditMemoId} | Update credit memo
*CustomExchangeRatesApi* | [**GETCustomExchangeRates**](docs/CustomExchangeRatesApi.md#getcustomexchangerates) | **Get** /v1/custom-exchange-rates/{currency} | Get custom foreign currency exchange rates
*DebitMemosApi* | [**DELETEDebitMemo**](docs/DebitMemosApi.md#deletedebitmemo) | **Delete** /v1/debitmemos/{debitMemoId} | Delete debit memo
*DebitMemosApi* | [**GETDebitMemo**](docs/DebitMemosApi.md#getdebitmemo) | **Get** /v1/debitmemos/{debitMemoId} | Get debit memo
*DebitMemosApi* | [**GETDebitMemoItem**](docs/DebitMemosApi.md#getdebitmemoitem) | **Get** /v1/debitmemos/{debitMemoId}/items/{dmitemid} | Get debit memo item
*DebitMemosApi* | [**GETDebitMemoItems**](docs/DebitMemosApi.md#getdebitmemoitems) | **Get** /v1/debitmemos/{debitMemoId}/items | Get debit memo items
*DebitMemosApi* | [**GETDebitMemos**](docs/DebitMemosApi.md#getdebitmemos) | **Get** /v1/debitmemos | Get debit memos
*DebitMemosApi* | [**POSTCancelDebitMemo**](docs/DebitMemosApi.md#postcanceldebitmemo) | **Post** /v1/debitmemos/{debitMemoId}/cancel | Cancel debit memo
*DebitMemosApi* | [**POSTDebitMemoFromInvoice**](docs/DebitMemosApi.md#postdebitmemofrominvoice) | **Post** /v1/debitmemos/invoice/{invoiceId} | Create debit memo from invoice
*DebitMemosApi* | [**POSTDebitMemoFromPrpc**](docs/DebitMemosApi.md#postdebitmemofromprpc) | **Post** /v1/debitmemos | Create debit memo from charge
*DebitMemosApi* | [**POSTDebitMemoPDF**](docs/DebitMemosApi.md#postdebitmemopdf) | **Post** /v1/debitmemos/{debitMemoId}/pdfs | Create debit memo PDF
*DebitMemosApi* | [**POSTEmailDebitMemo**](docs/DebitMemosApi.md#postemaildebitmemo) | **Post** /v1/debitmemos/{debitMemoId}/emails | Email debit memo
*DebitMemosApi* | [**POSTPostDebitMemo**](docs/DebitMemosApi.md#postpostdebitmemo) | **Post** /v1/debitmemos/{debitMemoId}/post | Post debit memo
*DebitMemosApi* | [**PUTDebitMemo**](docs/DebitMemosApi.md#putdebitmemo) | **Put** /v1/debitmemos/{debitMemoId} | Update debit memo
*DescribeApi* | [**GETDescribe**](docs/DescribeApi.md#getdescribe) | **Get** /v1/describe/{object} | Describe object
*EntitiesApi* | [**DELETEEntities**](docs/EntitiesApi.md#deleteentities) | **Delete** /v1/entities/{id} | Multi-entity: Delete entity
*EntitiesApi* | [**GETEntities**](docs/EntitiesApi.md#getentities) | **Get** /v1/entities | Multi-entity: Get entities
*EntitiesApi* | [**GETEntityById**](docs/EntitiesApi.md#getentitybyid) | **Get** /v1/entities/{id} | Multi-entity: Get entity by Id
*EntitiesApi* | [**POSTEntities**](docs/EntitiesApi.md#postentities) | **Post** /v1/entities | Multi-entity: Create entity
*EntitiesApi* | [**PUTEntities**](docs/EntitiesApi.md#putentities) | **Put** /v1/entities/{id} | Multi-entity: Update entity
*EntitiesApi* | [**PUTProvisionEntity**](docs/EntitiesApi.md#putprovisionentity) | **Put** /v1/entities/{id}/provision | Multi-entity: Provision entity
*EntityConnectionsApi* | [**GETEntityConnections**](docs/EntityConnectionsApi.md#getentityconnections) | **Get** /v1/entity-connections | Multi-entity: Get connections
*EntityConnectionsApi* | [**POSTEntityConnections**](docs/EntityConnectionsApi.md#postentityconnections) | **Post** /v1/entity-connections | Multi-entity: Initiate connection
*EntityConnectionsApi* | [**PUTEntityConnectionsAccept**](docs/EntityConnectionsApi.md#putentityconnectionsaccept) | **Put** /v1/entity-connections/{connection-id}/accept | Multi-entity: Accept connection
*EntityConnectionsApi* | [**PUTEntityConnectionsDeny**](docs/EntityConnectionsApi.md#putentityconnectionsdeny) | **Put** /v1/entity-connections/{connection-id}/deny | Multi-entity: Deny connection
*EntityConnectionsApi* | [**PUTEntityConnectionsDisconnect**](docs/EntityConnectionsApi.md#putentityconnectionsdisconnect) | **Put** /v1/entity-connections/{connection-id}/disconnect | Multi-entity: Disconnect connection
*ExportsApi* | [**ObjectGETExport**](docs/ExportsApi.md#objectgetexport) | **Get** /v1/object/export/{id} | CRUD: Retrieve Export
*ExportsApi* | [**ObjectPOSTExport**](docs/ExportsApi.md#objectpostexport) | **Post** /v1/object/export | CRUD: Create Export
*FeaturesApi* | [**ObjectDELETEFeature**](docs/FeaturesApi.md#objectdeletefeature) | **Delete** /v1/object/feature/{id} | CRUD: Delete Feature
*FeaturesApi* | [**ObjectGETFeature**](docs/FeaturesApi.md#objectgetfeature) | **Get** /v1/object/feature/{id} | CRUD: Retrieve Feature
*GetFilesApi* | [**GETFiles**](docs/GetFilesApi.md#getfiles) | **Get** /v1/files/{file-id} | Get files
*HMACSignaturesApi* | [**POSTHMACSignatures**](docs/HMACSignaturesApi.md#posthmacsignatures) | **Post** /v1/hmac-signatures | Return HMAC signatures
*HostedPagesApi* | [**GetHostedPages**](docs/HostedPagesApi.md#gethostedpages) | **Get** /v1/hostedpages | Return hosted pages
*ImportsApi* | [**ObjectGETImport**](docs/ImportsApi.md#objectgetimport) | **Get** /v1/object/import/{id} | CRUD: Retrieve Import
*InvoiceAdjustmentsApi* | [**ObjectDELETEInvoiceAdjustment**](docs/InvoiceAdjustmentsApi.md#objectdeleteinvoiceadjustment) | **Delete** /v1/object/invoice-adjustment/{id} | CRUD: Delete InvoiceAdjustment
*InvoiceAdjustmentsApi* | [**ObjectGETInvoiceAdjustment**](docs/InvoiceAdjustmentsApi.md#objectgetinvoiceadjustment) | **Get** /v1/object/invoice-adjustment/{id} | CRUD: Retrieve InvoiceAdjustment
*InvoiceAdjustmentsApi* | [**ObjectPOSTInvoiceAdjustment**](docs/InvoiceAdjustmentsApi.md#objectpostinvoiceadjustment) | **Post** /v1/object/invoice-adjustment | CRUD: Create InvoiceAdjustment
*InvoiceAdjustmentsApi* | [**ObjectPUTInvoiceAdjustment**](docs/InvoiceAdjustmentsApi.md#objectputinvoiceadjustment) | **Put** /v1/object/invoice-adjustment/{id} | CRUD: Update InvoiceAdjustment
*InvoiceItemAdjustmentsApi* | [**ObjectDELETEInvoiceItemAdjustment**](docs/InvoiceItemAdjustmentsApi.md#objectdeleteinvoiceitemadjustment) | **Delete** /v1/object/invoice-item-adjustment/{id} | CRUD: Delete InvoiceItemAdjustment
*InvoiceItemAdjustmentsApi* | [**ObjectGETInvoiceItemAdjustment**](docs/InvoiceItemAdjustmentsApi.md#objectgetinvoiceitemadjustment) | **Get** /v1/object/invoice-item-adjustment/{id} | CRUD: Retrieve InvoiceItemAdjustment
*InvoiceItemsApi* | [**ObjectGETInvoiceItem**](docs/InvoiceItemsApi.md#objectgetinvoiceitem) | **Get** /v1/object/invoice-item/{id} | CRUD: Retrieve InvoiceItem
*InvoicePaymentsApi* | [**ObjectGETInvoicePayment**](docs/InvoicePaymentsApi.md#objectgetinvoicepayment) | **Get** /v1/object/invoice-payment/{id} | CRUD: Retrieve InvoicePayment
*InvoicePaymentsApi* | [**ObjectPOSTInvoicePayment**](docs/InvoicePaymentsApi.md#objectpostinvoicepayment) | **Post** /v1/object/invoice-payment | CRUD: Create InvoicePayment
*InvoicePaymentsApi* | [**ObjectPUTInvoicePayment**](docs/InvoicePaymentsApi.md#objectputinvoicepayment) | **Put** /v1/object/invoice-payment/{id} | CRUD: Update InvoicePayment
*InvoiceSplitItemsApi* | [**ObjectGETInvoiceSplitItem**](docs/InvoiceSplitItemsApi.md#objectgetinvoicesplititem) | **Get** /v1/object/invoice-split-item/{id} | CRUD: Retrieve InvoiceSplitItem
*InvoiceSplitsApi* | [**ObjectGETInvoiceSplit**](docs/InvoiceSplitsApi.md#objectgetinvoicesplit) | **Get** /v1/object/invoice-split/{id} | CRUD: Retrieve InvoiceSplit
*InvoicesApi* | [**ObjectDELETEInvoice**](docs/InvoicesApi.md#objectdeleteinvoice) | **Delete** /v1/object/invoice/{id} | CRUD: Delete Invoice
*InvoicesApi* | [**ObjectGETInvoice**](docs/InvoicesApi.md#objectgetinvoice) | **Get** /v1/object/invoice/{id} | CRUD: Retrieve Invoice
*InvoicesApi* | [**ObjectPOSTInvoice**](docs/InvoicesApi.md#objectpostinvoice) | **Post** /v1/object/invoice | CRUD: Create Invoice
*InvoicesApi* | [**ObjectPUTInvoice**](docs/InvoicesApi.md#objectputinvoice) | **Put** /v1/object/invoice/{id} | CRUD: Update Invoice
*InvoicesApi* | [**POSTEmailInvoice**](docs/InvoicesApi.md#postemailinvoice) | **Post** /v1/invoices/{invoiceId}/emails | Email invoice
*InvoicesApi* | [**PUTReverseInvoice**](docs/InvoicesApi.md#putreverseinvoice) | **Put** /v1/invoices/{invoiceId}/reverse | Reverse invoice
*JournalRunsApi* | [**DELETEJournalRun**](docs/JournalRunsApi.md#deletejournalrun) | **Delete** /v1/journal-runs/{jr-number} | Delete journal run
*JournalRunsApi* | [**GETJournalRun**](docs/JournalRunsApi.md#getjournalrun) | **Get** /v1/journal-runs/{jr-number} | Get journal run
*JournalRunsApi* | [**POSTJournalRun**](docs/JournalRunsApi.md#postjournalrun) | **Post** /v1/journal-runs | Create journal run
*JournalRunsApi* | [**PUTJournalRun**](docs/JournalRunsApi.md#putjournalrun) | **Put** /v1/journal-runs/{jr-number}/cancel | Cancel journal run
*MassUpdaterApi* | [**GETMassUpdater**](docs/MassUpdaterApi.md#getmassupdater) | **Get** /v1/bulk/{bulk-key} | Get mass action result
*MassUpdaterApi* | [**POSTMassUpdater**](docs/MassUpdaterApi.md#postmassupdater) | **Post** /v1/bulk | Perform mass action
*MassUpdaterApi* | [**PUTMassUpdater**](docs/MassUpdaterApi.md#putmassupdater) | **Put** /v1/bulk/{bulk-key}/stop | Stop mass action
*NotificationHistoryApi* | [**GETCalloutHistory**](docs/NotificationHistoryApi.md#getcallouthistory) | **Get** /v1/notification-history/callout | Get callout notification histories
*NotificationHistoryApi* | [**GETEmailHistory**](docs/NotificationHistoryApi.md#getemailhistory) | **Get** /v1/notification-history/email | Get email notification histories
*OAuthApi* | [**CreateToken**](docs/OAuthApi.md#createtoken) | **Post** /oauth/token | Generate an OAuth token
*OperationsApi* | [**POSTBillingPreview**](docs/OperationsApi.md#postbillingpreview) | **Post** /v1/operations/billing-preview | Create billing preview
*OperationsApi* | [**POSTTransactionInvoicePayment**](docs/OperationsApi.md#posttransactioninvoicepayment) | **Post** /v1/operations/invoice-collect | Invoice and collect
*PaymentGatewaysApi* | [**GETPaymentgateways**](docs/PaymentGatewaysApi.md#getpaymentgateways) | **Get** /v1/paymentgateways | Get payment gateways
*PaymentMethodSnapshotsApi* | [**ObjectGETPaymentMethodSnapshot**](docs/PaymentMethodSnapshotsApi.md#objectgetpaymentmethodsnapshot) | **Get** /v1/object/payment-method-snapshot/{id} | CRUD: Retrieve PaymentMethodSnapshot
*PaymentMethodTransactionLogsApi* | [**ObjectDELETEPaymentMethodTransactionLog**](docs/PaymentMethodTransactionLogsApi.md#objectdeletepaymentmethodtransactionlog) | **Delete** /v1/object/payment-method-transaction-log/{id} | CRUD: Delete PaymentMethodTransactionLog
*PaymentMethodTransactionLogsApi* | [**ObjectGETPaymentMethodTransactionLog**](docs/PaymentMethodTransactionLogsApi.md#objectgetpaymentmethodtransactionlog) | **Get** /v1/object/payment-method-transaction-log/{id} | CRUD: Retrieve PaymentMethodTransactionLog
*PaymentMethodsApi* | [**DELETEPaymentMethods**](docs/PaymentMethodsApi.md#deletepaymentmethods) | **Delete** /v1/payment-methods/{payment-method-id} | Delete payment method
*PaymentMethodsApi* | [**GETPaymentMethods**](docs/PaymentMethodsApi.md#getpaymentmethods) | **Get** /v1/payment-methods/credit-cards/accounts/{account-key} | Get payment methods
*PaymentMethodsApi* | [**ObjectDELETEPaymentMethod**](docs/PaymentMethodsApi.md#objectdeletepaymentmethod) | **Delete** /v1/object/payment-method/{id} | CRUD: Delete PaymentMethod
*PaymentMethodsApi* | [**ObjectGETPaymentMethod**](docs/PaymentMethodsApi.md#objectgetpaymentmethod) | **Get** /v1/object/payment-method/{id} | CRUD: Retrieve PaymentMethod
*PaymentMethodsApi* | [**ObjectPOSTPaymentMethod**](docs/PaymentMethodsApi.md#objectpostpaymentmethod) | **Post** /v1/object/payment-method | CRUD: Create PaymentMethod
*PaymentMethodsApi* | [**ObjectPUTPaymentMethod**](docs/PaymentMethodsApi.md#objectputpaymentmethod) | **Put** /v1/object/payment-method/{id} | CRUD: Update PaymentMethod
*PaymentMethodsApi* | [**POSTPaymentMethods**](docs/PaymentMethodsApi.md#postpaymentmethods) | **Post** /v1/payment-methods/credit-cards | Create payment method
*PaymentMethodsApi* | [**POSTPaymentMethodsDecryption**](docs/PaymentMethodsApi.md#postpaymentmethodsdecryption) | **Post** /v1/payment-methods/decryption | Create payment method decryption
*PaymentMethodsApi* | [**PUTPaymentMethods**](docs/PaymentMethodsApi.md#putpaymentmethods) | **Put** /v1/payment-methods/credit-cards/{payment-method-id} | Update payment method
*PaymentTransactionLogsApi* | [**ObjectDELETEPaymentTransactionLog**](docs/PaymentTransactionLogsApi.md#objectdeletepaymenttransactionlog) | **Delete** /v1/object/payment-transaction-log/{id} | CRUD: Delete PaymentTransactionLog
*PaymentTransactionLogsApi* | [**ObjectGETPaymentTransactionLog**](docs/PaymentTransactionLogsApi.md#objectgetpaymenttransactionlog) | **Get** /v1/object/payment-transaction-log/{id} | CRUD: Retrieve PaymentTransactionLog
*PaymentsApi* | [**DELETEPayment**](docs/PaymentsApi.md#deletepayment) | **Delete** /v1/payments/{paymentId} | Delete payment
*PaymentsApi* | [**GETPayemntItemParts**](docs/PaymentsApi.md#getpayemntitemparts) | **Get** /v1/payments/{paymentId}/parts/{partid}/itemparts | Get payment part items
*PaymentsApi* | [**GETPayment**](docs/PaymentsApi.md#getpayment) | **Get** /v1/payments/{paymentId} | Get payment
*PaymentsApi* | [**GETPaymentItemPart**](docs/PaymentsApi.md#getpaymentitempart) | **Get** /v1/payments/{paymentId}/parts/{partid}/itemparts/{itempartid} | Get payment part item
*PaymentsApi* | [**GETPaymentPart**](docs/PaymentsApi.md#getpaymentpart) | **Get** /v1/payments/{paymentId}/parts/{partid} | Get payment part
*PaymentsApi* | [**GETPaymentParts**](docs/PaymentsApi.md#getpaymentparts) | **Get** /v1/payments/{paymentId}/parts | Get payment parts
*PaymentsApi* | [**GETRetrieveAllPayments**](docs/PaymentsApi.md#getretrieveallpayments) | **Get** /v1/payments | Get all payments
*PaymentsApi* | [**ObjectDELETEPayment**](docs/PaymentsApi.md#objectdeletepayment) | **Delete** /v1/object/payment/{id} | CRUD: Delete Payment
*PaymentsApi* | [**ObjectGETPayment**](docs/PaymentsApi.md#objectgetpayment) | **Get** /v1/object/payment/{id} | CRUD: Retrieve Payment
*PaymentsApi* | [**ObjectPOSTPayment**](docs/PaymentsApi.md#objectpostpayment) | **Post** /v1/object/payment | CRUD: Create Payment
*PaymentsApi* | [**ObjectPUTPayment**](docs/PaymentsApi.md#objectputpayment) | **Put** /v1/object/payment/{id} | CRUD: Update Payment
*PaymentsApi* | [**POSTApplyPayment**](docs/PaymentsApi.md#postapplypayment) | **Post** /v1/payments/{paymentId}/apply | Apply payment
*PaymentsApi* | [**POSTCancelPayment**](docs/PaymentsApi.md#postcancelpayment) | **Post** /v1/payments/{paymentId}/cancel | Cancel payment
*PaymentsApi* | [**POSTCreatePayment**](docs/PaymentsApi.md#postcreatepayment) | **Post** /v1/payments | Create payment
*PaymentsApi* | [**POSTRefundPayment**](docs/PaymentsApi.md#postrefundpayment) | **Post** /v1/payments/{paymentId}/refund | Refund payment
*PaymentsApi* | [**POSTTransferPayemnt**](docs/PaymentsApi.md#posttransferpayemnt) | **Post** /v1/payments/{paymentId}/accounts | Transfer payment
*PaymentsApi* | [**POSTUNAPPLYPAYMENT**](docs/PaymentsApi.md#postunapplypayment) | **Post** /v1/payments/{paymentId}/unapply | Unapply payment
*PaymentsApi* | [**PUTUpdatePayemnt**](docs/PaymentsApi.md#putupdatepayemnt) | **Put** /v1/payments/{paymentId} | Update payment
*ProductFeaturesApi* | [**ObjectDELETEProductFeature**](docs/ProductFeaturesApi.md#objectdeleteproductfeature) | **Delete** /v1/object/product-feature/{id} | CRUD: Delete ProductFeature
*ProductFeaturesApi* | [**ObjectGETProductFeature**](docs/ProductFeaturesApi.md#objectgetproductfeature) | **Get** /v1/object/product-feature/{id} | CRUD: Retrieve ProductFeature
*ProductRatePlanChargeTiersApi* | [**ObjectGETProductRatePlanChargeTier**](docs/ProductRatePlanChargeTiersApi.md#objectgetproductrateplanchargetier) | **Get** /v1/object/product-rate-plan-charge-tier/{id} | CRUD: Retrieve ProductRatePlanChargeTier
*ProductRatePlanChargesApi* | [**ObjectDELETEProductRatePlanCharge**](docs/ProductRatePlanChargesApi.md#objectdeleteproductrateplancharge) | **Delete** /v1/object/product-rate-plan-charge/{id} | CRUD: Delete ProductRatePlanCharge
*ProductRatePlanChargesApi* | [**ObjectGETProductRatePlanCharge**](docs/ProductRatePlanChargesApi.md#objectgetproductrateplancharge) | **Get** /v1/object/product-rate-plan-charge/{id} | CRUD: Retrieve ProductRatePlanCharge
*ProductRatePlansApi* | [**ObjectDELETEProductRatePlan**](docs/ProductRatePlansApi.md#objectdeleteproductrateplan) | **Delete** /v1/object/product-rate-plan/{id} | CRUD: Delete ProductRatePlan
*ProductRatePlansApi* | [**ObjectGETProductRatePlan**](docs/ProductRatePlansApi.md#objectgetproductrateplan) | **Get** /v1/object/product-rate-plan/{id} | CRUD: Retrieve ProductRatePlan
*ProductRatePlansApi* | [**ObjectPOSTProductRatePlan**](docs/ProductRatePlansApi.md#objectpostproductrateplan) | **Post** /v1/object/product-rate-plan | CRUD: Create ProductRatePlan
*ProductRatePlansApi* | [**ObjectPUTProductRatePlan**](docs/ProductRatePlansApi.md#objectputproductrateplan) | **Put** /v1/object/product-rate-plan/{id} | CRUD: Update ProductRatePlan
*ProductsApi* | [**ObjectDELETEProduct**](docs/ProductsApi.md#objectdeleteproduct) | **Delete** /v1/object/product/{id} | CRUD: Delete Product
*ProductsApi* | [**ObjectGETProduct**](docs/ProductsApi.md#objectgetproduct) | **Get** /v1/object/product/{id} | CRUD: Retrieve Product
*ProductsApi* | [**ObjectPOSTProduct**](docs/ProductsApi.md#objectpostproduct) | **Post** /v1/object/product | CRUD: Create Product
*ProductsApi* | [**ObjectPUTProduct**](docs/ProductsApi.md#objectputproduct) | **Put** /v1/object/product/{id} | CRUD: Update Product
*QuotesDocumentApi* | [**POSTQuotesDocument**](docs/QuotesDocumentApi.md#postquotesdocument) | **Post** /v1/quotes/document | Generate quotes document
*RSASignaturesApi* | [**POSTDecryptRSASignatures**](docs/RSASignaturesApi.md#postdecryptrsasignatures) | **Post** /v1/rsa-signatures/decrypt | Decrypt RSA signature
*RSASignaturesApi* | [**POSTRSASignatures**](docs/RSASignaturesApi.md#postrsasignatures) | **Post** /v1/rsa-signatures | Generate RSA signature
*RatePlanChargeTiersApi* | [**ObjectGETRatePlanChargeTier**](docs/RatePlanChargeTiersApi.md#objectgetrateplanchargetier) | **Get** /v1/object/rate-plan-charge-tier/{id} | CRUD: Retrieve RatePlanChargeTier
*RatePlanChargesApi* | [**ObjectGETRatePlanCharge**](docs/RatePlanChargesApi.md#objectgetrateplancharge) | **Get** /v1/object/rate-plan-charge/{id} | CRUD: Retrieve RatePlanCharge
*RatePlansApi* | [**ObjectGETRatePlan**](docs/RatePlansApi.md#objectgetrateplan) | **Get** /v1/object/rate-plan/{id} | CRUD: Retrieve RatePlan
*RefundInvoicePaymentsApi* | [**ObjectGETRefundInvoicePayment**](docs/RefundInvoicePaymentsApi.md#objectgetrefundinvoicepayment) | **Get** /v1/object/refund-invoice-payment/{id} | CRUD: Retrieve RefundInvoicePayment
*RefundTransactionLogsApi* | [**ObjectDELETERefundTransactionLog**](docs/RefundTransactionLogsApi.md#objectdeleterefundtransactionlog) | **Delete** /v1/object/refund-transaction-log/{id} | CRUD: Delete RefundTransactionLog
*RefundTransactionLogsApi* | [**ObjectGETRefundTransactionLog**](docs/RefundTransactionLogsApi.md#objectgetrefundtransactionlog) | **Get** /v1/object/refund-transaction-log/{id} | CRUD: Retrieve RefundTransactionLog
*RefundsApi* | [**DELETERefund**](docs/RefundsApi.md#deleterefund) | **Delete** /v1/refunds/{refundId} | Delete refund
*RefundsApi* | [**GETRefund**](docs/RefundsApi.md#getrefund) | **Get** /v1/refunds/{refundId} | Get refund
*RefundsApi* | [**GETRefundItemPart**](docs/RefundsApi.md#getrefunditempart) | **Get** /v1/refunds/{refundId}/parts/{refundpartid}/itemparts/{itempartid} | Get refund part item
*RefundsApi* | [**GETRefundItemParts**](docs/RefundsApi.md#getrefunditemparts) | **Get** /v1/refunds/{refundId}/parts/{refundpartid}/itemparts | Get refund part items
*RefundsApi* | [**GETRefundPart**](docs/RefundsApi.md#getrefundpart) | **Get** /v1/refunds/{refundId}/parts/{refundpartid} | Get refund part
*RefundsApi* | [**GETRefundParts**](docs/RefundsApi.md#getrefundparts) | **Get** /v1/refunds/{refundId}/parts | Get refund parts
*RefundsApi* | [**GETRefunds**](docs/RefundsApi.md#getrefunds) | **Get** /v1/refunds | Get all refunds
*RefundsApi* | [**ObjectDELETERefund**](docs/RefundsApi.md#objectdeleterefund) | **Delete** /v1/object/refund/{id} | CRUD: Delete Refund
*RefundsApi* | [**ObjectGETRefund**](docs/RefundsApi.md#objectgetrefund) | **Get** /v1/object/refund/{id} | CRUD: Retrieve Refund
*RefundsApi* | [**ObjectPOSTRefund**](docs/RefundsApi.md#objectpostrefund) | **Post** /v1/object/refund | CRUD: Create Refund
*RefundsApi* | [**ObjectPUTRefund**](docs/RefundsApi.md#objectputrefund) | **Put** /v1/object/refund/{id} | CRUD: Update Refund
*RefundsApi* | [**POSTCancelRefund**](docs/RefundsApi.md#postcancelrefund) | **Post** /v1/refunds/{refundId}/cancel | Cancel refund
*RefundsApi* | [**PUTUpdateRefund**](docs/RefundsApi.md#putupdaterefund) | **Put** /v1/refunds/{refundId} | Update refund
*RevenueEventsApi* | [**GETRevenueEventDetails**](docs/RevenueEventsApi.md#getrevenueeventdetails) | **Get** /v1/revenue-events/{event-number} | Get revenue event details
*RevenueEventsApi* | [**GETRevenueEventForRevenueSchedule**](docs/RevenueEventsApi.md#getrevenueeventforrevenueschedule) | **Get** /v1/revenue-events/revenue-schedules/{rs-number} | Get revenue events for a revenue schedule
*RevenueItemsApi* | [**GETRevenueItemsByChargeRevenueEventNumber**](docs/RevenueItemsApi.md#getrevenueitemsbychargerevenueeventnumber) | **Get** /v1/revenue-items/revenue-events/{event-number} | Get revenue items by revenue event number
*RevenueItemsApi* | [**GETRevenueItemsByChargeRevenueSummaryNumber**](docs/RevenueItemsApi.md#getrevenueitemsbychargerevenuesummarynumber) | **Get** /v1/revenue-items/charge-revenue-summaries/{crs-number} | Get revenue items by charge revenue summary number
*RevenueItemsApi* | [**GETRevenueItemsByRevenueSchedule**](docs/RevenueItemsApi.md#getrevenueitemsbyrevenueschedule) | **Get** /v1/revenue-items/revenue-schedules/{rs-number} | Get revenue items by revenue schedule
*RevenueItemsApi* | [**PUTCustomFieldsonRevenueItemsByRevenueEvent**](docs/RevenueItemsApi.md#putcustomfieldsonrevenueitemsbyrevenueevent) | **Put** /v1/revenue-items/revenue-events/{event-number} | Update custom fields on revenue items by revenue event number
*RevenueItemsApi* | [**PUTCustomFieldsonRevenueItemsByRevenueSchedule**](docs/RevenueItemsApi.md#putcustomfieldsonrevenueitemsbyrevenueschedule) | **Put** /v1/revenue-items/revenue-schedules/{rs-number} | Update custom fields on revenue items by revenue schedule number
*RevenueRulesApi* | [**GETRevenueRecRulebyProductRatePlanCharge**](docs/RevenueRulesApi.md#getrevenuerecrulebyproductrateplancharge) | **Get** /v1/revenue-recognition-rules/product-charges/{charge-key} | Get revenue recognition rule by product rate plan charge
*RevenueRulesApi* | [**GETRevenueRecRules**](docs/RevenueRulesApi.md#getrevenuerecrules) | **Get** /v1/revenue-recognition-rules/subscription-charges/{charge-key} | Get revenue recognition rule by subscription charge
*RevenueSchedulesApi* | [**DELETERS**](docs/RevenueSchedulesApi.md#deleters) | **Delete** /v1/revenue-schedules/{rs-number} | Delete revenue schedule
*RevenueSchedulesApi* | [**GETRS**](docs/RevenueSchedulesApi.md#getrs) | **Get** /v1/revenue-schedules/{rs-number} | Get revenue schedule details
*RevenueSchedulesApi* | [**GETRSbyCreditMemoItem**](docs/RevenueSchedulesApi.md#getrsbycreditmemoitem) | **Get** /v1/revenue-schedules/credit-memo-items/{cmi-id} | Get revenue schedule by credit memo item ID 
*RevenueSchedulesApi* | [**GETRSbyDebitMemoItem**](docs/RevenueSchedulesApi.md#getrsbydebitmemoitem) | **Get** /v1/revenue-schedules/debit-memo-items/{dmi-id} | Get revenue schedule by debit memo item ID 
*RevenueSchedulesApi* | [**GETRSbyInvoiceItem**](docs/RevenueSchedulesApi.md#getrsbyinvoiceitem) | **Get** /v1/revenue-schedules/invoice-items/{invoice-item-id} | Get revenue schedule by invoice item ID
*RevenueSchedulesApi* | [**GETRSbyInvoiceItemAdjustment**](docs/RevenueSchedulesApi.md#getrsbyinvoiceitemadjustment) | **Get** /v1/revenue-schedules/invoice-item-adjustments/{invoice-item-adj-id}/ | Get revenue schedule by invoice item adjustment
*RevenueSchedulesApi* | [**GETRSbyProductChargeAndBillingAccount**](docs/RevenueSchedulesApi.md#getrsbyproductchargeandbillingaccount) | **Get** /v1/revenue-schedules/product-charges/{charge-key}/{account-key} | Get all revenue schedules of product charge by charge ID and billing account ID 
*RevenueSchedulesApi* | [**GETRSforSubscCharge**](docs/RevenueSchedulesApi.md#getrsforsubsccharge) | **Get** /v1/revenue-schedules/subscription-charges/{charge-key} | Get revenue schedule by subscription charge
*RevenueSchedulesApi* | [**POSTRSforCrditMemoItemManualDistribution**](docs/RevenueSchedulesApi.md#postrsforcrditmemoitemmanualdistribution) | **Post** /v1/revenue-schedules/credit-memo-items/{cmi-id} | Create revenue schedule for credit memo item (manual distribution) 
*RevenueSchedulesApi* | [**POSTRSforCreditMemoItemDistributeByDateRange**](docs/RevenueSchedulesApi.md#postrsforcreditmemoitemdistributebydaterange) | **Post** /v1/revenue-schedules/credit-memo-items/{cmi-id}/distribute-revenue-with-date-range | Create revenue schedule for credit memo item (distribute by date range) 
*RevenueSchedulesApi* | [**POSTRSforDebitMemoItemDistributeByDateRange**](docs/RevenueSchedulesApi.md#postrsfordebitmemoitemdistributebydaterange) | **Post** /v1/revenue-schedules/debit-memo-items/{dmi-id}/distribute-revenue-with-date-range | Create revenue schedule for debit memo item (distribute by date range) 
*RevenueSchedulesApi* | [**POSTRSforDebitMemoItemManualDistribution**](docs/RevenueSchedulesApi.md#postrsfordebitmemoitemmanualdistribution) | **Post** /v1/revenue-schedules/debit-memo-items/{dmi-id} | Create revenue schedule for debit memo item (manual distribution) 
*RevenueSchedulesApi* | [**POSTRSforInvoiceItemAdjustmentDistributeByDateRange**](docs/RevenueSchedulesApi.md#postrsforinvoiceitemadjustmentdistributebydaterange) | **Post** /v1/revenue-schedules/invoice-item-adjustments/{invoice-item-adj-key}/distribute-revenue-with-date-range | Create revenue schedule for Invoice Item Adjustment (distribute by date range)
*RevenueSchedulesApi* | [**POSTRSforInvoiceItemAdjustmentManualDistribution**](docs/RevenueSchedulesApi.md#postrsforinvoiceitemadjustmentmanualdistribution) | **Post** /v1/revenue-schedules/invoice-item-adjustments/{invoice-item-adj-key} | Create revenue schedule for Invoice Item Adjustment (manual distribution)
*RevenueSchedulesApi* | [**POSTRSforInvoiceItemDistributeByDateRange**](docs/RevenueSchedulesApi.md#postrsforinvoiceitemdistributebydaterange) | **Post** /v1/revenue-schedules/invoice-items/{invoice-item-id}/distribute-revenue-with-date-range | Create revenue schedule for Invoice Item (distribute by date range)
*RevenueSchedulesApi* | [**POSTRSforInvoiceItemManualDistribution**](docs/RevenueSchedulesApi.md#postrsforinvoiceitemmanualdistribution) | **Post** /v1/revenue-schedules/invoice-items/{invoice-item-id} | Create revenue schedule for Invoice Item (manual distribution)
*RevenueSchedulesApi* | [**POSTRSforSubsCharge**](docs/RevenueSchedulesApi.md#postrsforsubscharge) | **Post** /v1/revenue-schedules/subscription-charges/{charge-key} | Create revenue schedule on subscription charge
*RevenueSchedulesApi* | [**PUTRSBasicInfo**](docs/RevenueSchedulesApi.md#putrsbasicinfo) | **Put** /v1/revenue-schedules/{rs-number}/basic-information | Update revenue schedule basic information
*RevenueSchedulesApi* | [**PUTRevenueAcrossAP**](docs/RevenueSchedulesApi.md#putrevenueacrossap) | **Put** /v1/revenue-schedules/{rs-number}/distribute-revenue-across-accounting-periods | Distribute revenue across accounting periods
*RevenueSchedulesApi* | [**PUTRevenueByRecognitionStartandEndDates**](docs/RevenueSchedulesApi.md#putrevenuebyrecognitionstartandenddates) | **Put** /v1/revenue-schedules/{rs-number}/distribute-revenue-with-date-range | Distribute revenue by recognition start and end dates
*RevenueSchedulesApi* | [**PUTRevenueSpecificDate**](docs/RevenueSchedulesApi.md#putrevenuespecificdate) | **Put** /v1/revenue-schedules/{rs-number}/distribute-revenue-on-specific-date | Distribute revenue on specific date
*SettingsApi* | [**GETRevenueAutomationStartDate**](docs/SettingsApi.md#getrevenueautomationstartdate) | **Get** /v1/settings/finance/revenue-automation-start-date | Get the revenue automation start date
*SubscriptionProductFeaturesApi* | [**ObjectDELETESubscriptionProductFeature**](docs/SubscriptionProductFeaturesApi.md#objectdeletesubscriptionproductfeature) | **Delete** /v1/object/subscription-product-feature/{id} | CRUD: Delete SubscriptionProductFeature
*SubscriptionProductFeaturesApi* | [**ObjectGETSubscriptionProductFeature**](docs/SubscriptionProductFeaturesApi.md#objectgetsubscriptionproductfeature) | **Get** /v1/object/subscription-product-feature/{id} | CRUD: Retrieve SubscriptionProductFeature
*SubscriptionsApi* | [**GETSubscriptionsByAccount**](docs/SubscriptionsApi.md#getsubscriptionsbyaccount) | **Get** /v1/subscriptions/accounts/{account-key} | Get subscriptions by account
*SubscriptionsApi* | [**GETSubscriptionsByKey**](docs/SubscriptionsApi.md#getsubscriptionsbykey) | **Get** /v1/subscriptions/{subscription-key} | Get subscriptions by key
*SubscriptionsApi* | [**GETSubscriptionsByKeyAndVersion**](docs/SubscriptionsApi.md#getsubscriptionsbykeyandversion) | **Get** /v1/subscriptions/{subscription-key}/versions/{version} | Get subscriptions by key and version
*SubscriptionsApi* | [**ObjectDELETESubscription**](docs/SubscriptionsApi.md#objectdeletesubscription) | **Delete** /v1/object/subscription/{id} | CRUD: Delete Subscription
*SubscriptionsApi* | [**ObjectGETSubscription**](docs/SubscriptionsApi.md#objectgetsubscription) | **Get** /v1/object/subscription/{id} | CRUD: Retrieve Subscription
*SubscriptionsApi* | [**ObjectPUTSubscription**](docs/SubscriptionsApi.md#objectputsubscription) | **Put** /v1/object/subscription/{id} | CRUD: Update Subscription
*SubscriptionsApi* | [**POSTPreviewSubscription**](docs/SubscriptionsApi.md#postpreviewsubscription) | **Post** /v1/subscriptions/preview | Preview subscription
*SubscriptionsApi* | [**POSTSubscription**](docs/SubscriptionsApi.md#postsubscription) | **Post** /v1/subscriptions | Create subscription
*SubscriptionsApi* | [**PUTCancelSubscription**](docs/SubscriptionsApi.md#putcancelsubscription) | **Put** /v1/subscriptions/{subscription-key}/cancel | Cancel subscription
*SubscriptionsApi* | [**PUTRenewSubscription**](docs/SubscriptionsApi.md#putrenewsubscription) | **Put** /v1/subscriptions/{subscription-key}/renew | Renew subscription
*SubscriptionsApi* | [**PUTResumeSubscription**](docs/SubscriptionsApi.md#putresumesubscription) | **Put** /v1/subscriptions/{subscription-key}/resume | Resume subscription
*SubscriptionsApi* | [**PUTSubscription**](docs/SubscriptionsApi.md#putsubscription) | **Put** /v1/subscriptions/{subscription-key} | Update subscription
*SubscriptionsApi* | [**PUTSuspendSubscription**](docs/SubscriptionsApi.md#putsuspendsubscription) | **Put** /v1/subscriptions/{subscription-key}/suspend | Suspend subscription
*SummaryJournalEntriesApi* | [**DELETESummaryJournalEntry**](docs/SummaryJournalEntriesApi.md#deletesummaryjournalentry) | **Delete** /v1/journal-entries/{je-number} | Delete summary journal entry
*SummaryJournalEntriesApi* | [**GETAllSummaryJournalEntries**](docs/SummaryJournalEntriesApi.md#getallsummaryjournalentries) | **Get** /v1/journal-entries/journal-runs/{jr-number} | Get all summary journal entries in a journal run
*SummaryJournalEntriesApi* | [**GETSummaryJournalEntry**](docs/SummaryJournalEntriesApi.md#getsummaryjournalentry) | **Get** /v1/journal-entries/{je-number} | Get summary journal entry
*SummaryJournalEntriesApi* | [**POSTSummaryJournalEntry**](docs/SummaryJournalEntriesApi.md#postsummaryjournalentry) | **Post** /v1/journal-entries | Create summary journal entry
*SummaryJournalEntriesApi* | [**PUTBasicSummaryJournalEntry**](docs/SummaryJournalEntriesApi.md#putbasicsummaryjournalentry) | **Put** /v1/journal-entries/{je-number}/basic-information | Update basic information of a summary journal entry
*SummaryJournalEntriesApi* | [**PUTSummaryJournalEntry**](docs/SummaryJournalEntriesApi.md#putsummaryjournalentry) | **Put** /v1/journal-entries/{je-number}/cancel | Cancel summary journal entry
*TaxationItemsApi* | [**DELETETaxationItem**](docs/TaxationItemsApi.md#deletetaxationitem) | **Delete** /v1/taxationitems/{id} | Delete taxation item
*TaxationItemsApi* | [**GETTaxationItem**](docs/TaxationItemsApi.md#gettaxationitem) | **Get** /v1/taxationitems/{id} | Get taxation item 
*TaxationItemsApi* | [**ObjectDELETETaxationItem**](docs/TaxationItemsApi.md#objectdeletetaxationitem) | **Delete** /v1/object/taxation-item/{id} | CRUD: Delete TaxationItem
*TaxationItemsApi* | [**ObjectGETTaxationItem**](docs/TaxationItemsApi.md#objectgettaxationitem) | **Get** /v1/object/taxation-item/{id} | CRUD: Retrieve TaxationItem
*TaxationItemsApi* | [**ObjectPOSTTaxationItem**](docs/TaxationItemsApi.md#objectposttaxationitem) | **Post** /v1/object/taxation-item | CRUD: Create TaxationItem
*TaxationItemsApi* | [**ObjectPUTTaxationItem**](docs/TaxationItemsApi.md#objectputtaxationitem) | **Put** /v1/object/taxation-item/{id} | CRUD: Update TaxationItem
*TaxationItemsApi* | [**POSTCMTaxationItems**](docs/TaxationItemsApi.md#postcmtaxationitems) | **Post** /v1/taxationitems/creditmemo/{memoId} | Create taxation items for credit memo
*TaxationItemsApi* | [**POSTDMTaxationItems**](docs/TaxationItemsApi.md#postdmtaxationitems) | **Post** /v1/taxationitems/debitmemo/{memoId} | Create taxation items for debit memo
*TaxationItemsApi* | [**PUTTaxationItem**](docs/TaxationItemsApi.md#puttaxationitem) | **Put** /v1/taxationitems/{id} | Update taxation item
*TransactionsApi* | [**GETTransactionInvoice**](docs/TransactionsApi.md#gettransactioninvoice) | **Get** /v1/transactions/invoices/accounts/{account-key} | Get invoices
*TransactionsApi* | [**GETTransactionPayment**](docs/TransactionsApi.md#gettransactionpayment) | **Get** /v1/transactions/payments/accounts/{account-key} | Get payments
*UnitOfMeasureApi* | [**ObjectDELETEUnitOfMeasure**](docs/UnitOfMeasureApi.md#objectdeleteunitofmeasure) | **Delete** /v1/object/unit-of-measure/{id} | CRUD: Delete UnitOfMeasure
*UnitOfMeasureApi* | [**ObjectGETUnitOfMeasure**](docs/UnitOfMeasureApi.md#objectgetunitofmeasure) | **Get** /v1/object/unit-of-measure/{id} | CRUD: Retrieve UnitOfMeasure
*UnitOfMeasureApi* | [**ObjectPOSTUnitOfMeasure**](docs/UnitOfMeasureApi.md#objectpostunitofmeasure) | **Post** /v1/object/unit-of-measure | CRUD: Create UnitOfMeasure
*UnitOfMeasureApi* | [**ObjectPUTUnitOfMeasure**](docs/UnitOfMeasureApi.md#objectputunitofmeasure) | **Put** /v1/object/unit-of-measure/{id} | CRUD: Update UnitOfMeasure
*UsageApi* | [**GETUsage**](docs/UsageApi.md#getusage) | **Get** /v1/usage/accounts/{account-key} | Get usage
*UsageApi* | [**ObjectDELETEUsage**](docs/UsageApi.md#objectdeleteusage) | **Delete** /v1/object/usage/{id} | CRUD: Delete Usage
*UsageApi* | [**ObjectGETUsage**](docs/UsageApi.md#objectgetusage) | **Get** /v1/object/usage/{id} | CRUD: Retrieve Usage
*UsageApi* | [**ObjectPOSTUsage**](docs/UsageApi.md#objectpostusage) | **Post** /v1/object/usage | CRUD: Create Usage
*UsageApi* | [**ObjectPUTUsage**](docs/UsageApi.md#objectputusage) | **Put** /v1/object/usage/{id} | CRUD: Update Usage
*UsageApi* | [**POSTUsage**](docs/UsageApi.md#postusage) | **Post** /v1/usage | Post usage
*UsersApi* | [**GETEntitiesUserAccessible**](docs/UsersApi.md#getentitiesuseraccessible) | **Get** /v1/users/{username}/accessible-entities | Multi-entity: Get entities that a user can access
*UsersApi* | [**PUTAcceptUserAccess**](docs/UsersApi.md#putacceptuseraccess) | **Put** /v1/users/{username}/accept-access | Multi-entity: Accept user access
*UsersApi* | [**PUTDenyUserAccess**](docs/UsersApi.md#putdenyuseraccess) | **Put** /v1/users/{username}/deny-access | Multi-entity: Deny user access
*UsersApi* | [**PUTSendUserAccessRequests**](docs/UsersApi.md#putsenduseraccessrequests) | **Put** /v1/users/{username}/request-access | Multi-entity: Send user access requests


## Documentation For Models

 - [AmendRequest](docs/AmendRequest.md)
 - [AmendRequestAmendOptions](docs/AmendRequestAmendOptions.md)
 - [AmendRequestPreviewOptions](docs/AmendRequestPreviewOptions.md)
 - [AmendResult](docs/AmendResult.md)
 - [Amendment](docs/Amendment.md)
 - [AmendmentRatePlanData](docs/AmendmentRatePlanData.md)
 - [ApplyCreditMemoType](docs/ApplyCreditMemoType.md)
 - [ApplyPaymentType](docs/ApplyPaymentType.md)
 - [BillingPreviewResult](docs/BillingPreviewResult.md)
 - [ChargeMetricsData](docs/ChargeMetricsData.md)
 - [CommonResponseType](docs/CommonResponseType.md)
 - [CreateEntityResponseType](docs/CreateEntityResponseType.md)
 - [CreateEntityType](docs/CreateEntityType.md)
 - [CreatePaymentType](docs/CreatePaymentType.md)
 - [CreatePaymentTypeFinanceInformation](docs/CreatePaymentTypeFinanceInformation.md)
 - [CreditMemoApplyDebitMemoItemRequestType](docs/CreditMemoApplyDebitMemoItemRequestType.md)
 - [CreditMemoApplyDebitMemoRequestType](docs/CreditMemoApplyDebitMemoRequestType.md)
 - [CreditMemoApplyInvoiceItemRequestType](docs/CreditMemoApplyInvoiceItemRequestType.md)
 - [CreditMemoApplyInvoiceRequestType](docs/CreditMemoApplyInvoiceRequestType.md)
 - [CreditMemoFromChargeDetailType](docs/CreditMemoFromChargeDetailType.md)
 - [CreditMemoFromChargeDetailTypeFinanceInformation](docs/CreditMemoFromChargeDetailTypeFinanceInformation.md)
 - [CreditMemoFromChargeType](docs/CreditMemoFromChargeType.md)
 - [CreditMemoFromInvoiceType](docs/CreditMemoFromInvoiceType.md)
 - [CreditMemoItemFromInvoiceItemType](docs/CreditMemoItemFromInvoiceItemType.md)
 - [CreditMemoItemFromInvoiceItemTypeFinanceInformation](docs/CreditMemoItemFromInvoiceItemTypeFinanceInformation.md)
 - [CreditMemoQueryType](docs/CreditMemoQueryType.md)
 - [CreditMemoResponseType](docs/CreditMemoResponseType.md)
 - [CreditMemoTaxItemFromInvoiceTaxItemType](docs/CreditMemoTaxItemFromInvoiceTaxItemType.md)
 - [CreditMemoTaxItemFromInvoiceTaxItemTypeFinanceInformation](docs/CreditMemoTaxItemFromInvoiceTaxItemTypeFinanceInformation.md)
 - [CreditMemoUnapplyDebitMemoItemRequestType](docs/CreditMemoUnapplyDebitMemoItemRequestType.md)
 - [CreditMemoUnapplyDebitMemoRequestType](docs/CreditMemoUnapplyDebitMemoRequestType.md)
 - [CreditMemoUnapplyInvoiceItemRequestType](docs/CreditMemoUnapplyInvoiceItemRequestType.md)
 - [CreditMemoUnapplyInvoiceRequestType](docs/CreditMemoUnapplyInvoiceRequestType.md)
 - [DebitMemoFromChargeDetailType](docs/DebitMemoFromChargeDetailType.md)
 - [DebitMemoFromChargeDetailTypeFinanceInformation](docs/DebitMemoFromChargeDetailTypeFinanceInformation.md)
 - [DebitMemoFromChargeType](docs/DebitMemoFromChargeType.md)
 - [DebitMemoFromInvoiceType](docs/DebitMemoFromInvoiceType.md)
 - [DebitMemoItemFromInvoiceItemType](docs/DebitMemoItemFromInvoiceItemType.md)
 - [DebitMemoItemFromInvoiceItemTypeFinanceInformation](docs/DebitMemoItemFromInvoiceItemTypeFinanceInformation.md)
 - [DebitMemoTaxItemFromInvoiceTaxItemType](docs/DebitMemoTaxItemFromInvoiceTaxItemType.md)
 - [DebitMemoTaxItemFromInvoiceTaxItemTypeFinanceInformation](docs/DebitMemoTaxItemFromInvoiceTaxItemTypeFinanceInformation.md)
 - [DeletEntityResponseType](docs/DeletEntityResponseType.md)
 - [DeleteResult](docs/DeleteResult.md)
 - [ElectronicPaymentOptions](docs/ElectronicPaymentOptions.md)
 - [EventRevenueItemType](docs/EventRevenueItemType.md)
 - [ExecuteResult](docs/ExecuteResult.md)
 - [ExternalPaymentOptions](docs/ExternalPaymentOptions.md)
 - [FinanceInformation](docs/FinanceInformation.md)
 - [GatewayOption](docs/GatewayOption.md)
 - [GenerateBillingDocumentResponseType](docs/GenerateBillingDocumentResponseType.md)
 - [GetAccountSummaryInvoiceType](docs/GetAccountSummaryInvoiceType.md)
 - [GetAccountSummaryPaymentInvoiceType](docs/GetAccountSummaryPaymentInvoiceType.md)
 - [GetAccountSummaryPaymentType](docs/GetAccountSummaryPaymentType.md)
 - [GetAccountSummarySubscriptionRatePlanType](docs/GetAccountSummarySubscriptionRatePlanType.md)
 - [GetAccountSummarySubscriptionType](docs/GetAccountSummarySubscriptionType.md)
 - [GetAccountSummaryType](docs/GetAccountSummaryType.md)
 - [GetAccountSummaryTypeBasicInfo](docs/GetAccountSummaryTypeBasicInfo.md)
 - [GetAccountSummaryTypeBasicInfoDefaultPaymentMethod](docs/GetAccountSummaryTypeBasicInfoDefaultPaymentMethod.md)
 - [GetAccountSummaryTypeBillToContact](docs/GetAccountSummaryTypeBillToContact.md)
 - [GetAccountSummaryTypeSoldToContact](docs/GetAccountSummaryTypeSoldToContact.md)
 - [GetAccountSummaryTypeTaxInfo](docs/GetAccountSummaryTypeTaxInfo.md)
 - [GetAccountSummaryUsageType](docs/GetAccountSummaryUsageType.md)
 - [GetAccountType](docs/GetAccountType.md)
 - [GetAccountTypeBasicInfo](docs/GetAccountTypeBasicInfo.md)
 - [GetAccountTypeBillToContact](docs/GetAccountTypeBillToContact.md)
 - [GetAccountTypeBillingAndPayment](docs/GetAccountTypeBillingAndPayment.md)
 - [GetAccountTypeMetrics](docs/GetAccountTypeMetrics.md)
 - [GetAccountTypeSoldToContact](docs/GetAccountTypeSoldToContact.md)
 - [GetAccountingCodeItemType](docs/GetAccountingCodeItemType.md)
 - [GetAccountingCodeItemWithoutSuccessType](docs/GetAccountingCodeItemWithoutSuccessType.md)
 - [GetAccountingCodesType](docs/GetAccountingCodesType.md)
 - [GetAccountingPeriodFileIdsType](docs/GetAccountingPeriodFileIdsType.md)
 - [GetAccountingPeriodType](docs/GetAccountingPeriodType.md)
 - [GetAccountingPeriodWithoutSuccessType](docs/GetAccountingPeriodWithoutSuccessType.md)
 - [GetAccountingPeriodsType](docs/GetAccountingPeriodsType.md)
 - [GetAmendmentType](docs/GetAmendmentType.md)
 - [GetAttachmentResponseType](docs/GetAttachmentResponseType.md)
 - [GetAttachmentResponseWithoutSuccessType](docs/GetAttachmentResponseWithoutSuccessType.md)
 - [GetAttachmentsResponseType](docs/GetAttachmentsResponseType.md)
 - [GetBillingPreviewRunResponse](docs/GetBillingPreviewRunResponse.md)
 - [GetCalloutHistoryVOsType](docs/GetCalloutHistoryVOsType.md)
 - [GetCalloutHistoryVoType](docs/GetCalloutHistoryVoType.md)
 - [GetCatalogType](docs/GetCatalogType.md)
 - [GetChargeRsDetailType](docs/GetChargeRsDetailType.md)
 - [GetCreditMemoCollectionType](docs/GetCreditMemoCollectionType.md)
 - [GetCreditMemoItemPartType](docs/GetCreditMemoItemPartType.md)
 - [GetCreditMemoItemPartTypewithSuccess](docs/GetCreditMemoItemPartTypewithSuccess.md)
 - [GetCreditMemoItemPartsCollectionType](docs/GetCreditMemoItemPartsCollectionType.md)
 - [GetCreditMemoItemType](docs/GetCreditMemoItemType.md)
 - [GetCreditMemoItemTypeFinanceInformation](docs/GetCreditMemoItemTypeFinanceInformation.md)
 - [GetCreditMemoItemTypewithSuccess](docs/GetCreditMemoItemTypewithSuccess.md)
 - [GetCreditMemoItemTypewithSuccessFinanceInformation](docs/GetCreditMemoItemTypewithSuccessFinanceInformation.md)
 - [GetCreditMemoItemsListType](docs/GetCreditMemoItemsListType.md)
 - [GetCreditMemoPartType](docs/GetCreditMemoPartType.md)
 - [GetCreditMemoPartTypewithSuccess](docs/GetCreditMemoPartTypewithSuccess.md)
 - [GetCreditMemoPartsCollectionType](docs/GetCreditMemoPartsCollectionType.md)
 - [GetCreditMemoType](docs/GetCreditMemoType.md)
 - [GetCreditMemoTypewithSuccess](docs/GetCreditMemoTypewithSuccess.md)
 - [GetCustomExchangeRatesDataType](docs/GetCustomExchangeRatesDataType.md)
 - [GetCustomExchangeRatesType](docs/GetCustomExchangeRatesType.md)
 - [GetDebitMemoCollectionType](docs/GetDebitMemoCollectionType.md)
 - [GetDebitMemoItemCollectionType](docs/GetDebitMemoItemCollectionType.md)
 - [GetDebitMemoItemType](docs/GetDebitMemoItemType.md)
 - [GetDebitMemoItemTypeFinanceInformation](docs/GetDebitMemoItemTypeFinanceInformation.md)
 - [GetDebitMemoItemTypewithSuccess](docs/GetDebitMemoItemTypewithSuccess.md)
 - [GetDebitMemoType](docs/GetDebitMemoType.md)
 - [GetDebitMemoTypewithSuccess](docs/GetDebitMemoTypewithSuccess.md)
 - [GetDiscountApplyDetailsType](docs/GetDiscountApplyDetailsType.md)
 - [GetEmailHistoryVOsType](docs/GetEmailHistoryVOsType.md)
 - [GetEmailHistoryVoType](docs/GetEmailHistoryVoType.md)
 - [GetEntitiesResponseType](docs/GetEntitiesResponseType.md)
 - [GetEntitiesResponseTypeWithId](docs/GetEntitiesResponseTypeWithId.md)
 - [GetEntitiesType](docs/GetEntitiesType.md)
 - [GetEntitiesUserAccessibleResponseType](docs/GetEntitiesUserAccessibleResponseType.md)
 - [GetEntityConnectionsArrayItemsType](docs/GetEntityConnectionsArrayItemsType.md)
 - [GetEntityConnectionsResponseType](docs/GetEntityConnectionsResponseType.md)
 - [GetHostedPageType](docs/GetHostedPageType.md)
 - [GetHostedPagesType](docs/GetHostedPagesType.md)
 - [GetInvoiceFileType](docs/GetInvoiceFileType.md)
 - [GetInvoiceFileWrapper](docs/GetInvoiceFileWrapper.md)
 - [GetInvoiceType](docs/GetInvoiceType.md)
 - [GetInvoicesInvoiceItemType](docs/GetInvoicesInvoiceItemType.md)
 - [GetJournalEntriesInJournalRunType](docs/GetJournalEntriesInJournalRunType.md)
 - [GetJournalEntryDetailType](docs/GetJournalEntryDetailType.md)
 - [GetJournalEntryDetailTypeWithoutSuccess](docs/GetJournalEntryDetailTypeWithoutSuccess.md)
 - [GetJournalEntryItemType](docs/GetJournalEntryItemType.md)
 - [GetJournalEntrySegmentType](docs/GetJournalEntrySegmentType.md)
 - [GetJournalRunTransactionType](docs/GetJournalRunTransactionType.md)
 - [GetJournalRunType](docs/GetJournalRunType.md)
 - [GetMassUpdateType](docs/GetMassUpdateType.md)
 - [GetPaidInvoicesType](docs/GetPaidInvoicesType.md)
 - [GetPaymentGatwaysResponse](docs/GetPaymentGatwaysResponse.md)
 - [GetPaymentItemPartCollectionType](docs/GetPaymentItemPartCollectionType.md)
 - [GetPaymentItemPartType](docs/GetPaymentItemPartType.md)
 - [GetPaymentItemPartTypewithSuccess](docs/GetPaymentItemPartTypewithSuccess.md)
 - [GetPaymentMethodType](docs/GetPaymentMethodType.md)
 - [GetPaymentMethodTypeCardHolderInfo](docs/GetPaymentMethodTypeCardHolderInfo.md)
 - [GetPaymentMethodsType](docs/GetPaymentMethodsType.md)
 - [GetPaymentPartType](docs/GetPaymentPartType.md)
 - [GetPaymentPartTypewithSuccess](docs/GetPaymentPartTypewithSuccess.md)
 - [GetPaymentPartsCollectionType](docs/GetPaymentPartsCollectionType.md)
 - [GetPaymentType](docs/GetPaymentType.md)
 - [GetPaymentsType](docs/GetPaymentsType.md)
 - [GetProductDiscountApplyDetailsType](docs/GetProductDiscountApplyDetailsType.md)
 - [GetProductFeatureType](docs/GetProductFeatureType.md)
 - [GetProductRatePlanChargePricingTierType](docs/GetProductRatePlanChargePricingTierType.md)
 - [GetProductRatePlanChargePricingType](docs/GetProductRatePlanChargePricingType.md)
 - [GetProductRatePlanChargeType](docs/GetProductRatePlanChargeType.md)
 - [GetProductRatePlanType](docs/GetProductRatePlanType.md)
 - [GetProductType](docs/GetProductType.md)
 - [GetRefundCollectionType](docs/GetRefundCollectionType.md)
 - [GetRefundCreditMemoType](docs/GetRefundCreditMemoType.md)
 - [GetRefundCreditMemoTypeFinanceInformation](docs/GetRefundCreditMemoTypeFinanceInformation.md)
 - [GetRefundItemPartCollectionType](docs/GetRefundItemPartCollectionType.md)
 - [GetRefundItemPartType](docs/GetRefundItemPartType.md)
 - [GetRefundItemPartTypewithSuccess](docs/GetRefundItemPartTypewithSuccess.md)
 - [GetRefundPartCollectionType](docs/GetRefundPartCollectionType.md)
 - [GetRefundPaymentType](docs/GetRefundPaymentType.md)
 - [GetRefundType](docs/GetRefundType.md)
 - [GetRefundTypewithSuccess](docs/GetRefundTypewithSuccess.md)
 - [GetRevenueEventDetailType](docs/GetRevenueEventDetailType.md)
 - [GetRevenueEventDetailWithoutSuccessType](docs/GetRevenueEventDetailWithoutSuccessType.md)
 - [GetRevenueEventDetailsType](docs/GetRevenueEventDetailsType.md)
 - [GetRevenueItemType](docs/GetRevenueItemType.md)
 - [GetRevenueItemsType](docs/GetRevenueItemsType.md)
 - [GetRevenueRecognitionRuleAssociationType](docs/GetRevenueRecognitionRuleAssociationType.md)
 - [GetRevenueStartDateSettingType](docs/GetRevenueStartDateSettingType.md)
 - [GetRsRevenueItemType](docs/GetRsRevenueItemType.md)
 - [GetRsRevenueItemsType](docs/GetRsRevenueItemsType.md)
 - [GetSubscriptionProductFeatureType](docs/GetSubscriptionProductFeatureType.md)
 - [GetSubscriptionRatePlanChargesType](docs/GetSubscriptionRatePlanChargesType.md)
 - [GetSubscriptionRatePlanType](docs/GetSubscriptionRatePlanType.md)
 - [GetSubscriptionType](docs/GetSubscriptionType.md)
 - [GetSubscriptionTypeWithSuccess](docs/GetSubscriptionTypeWithSuccess.md)
 - [GetSubscriptionWrapper](docs/GetSubscriptionWrapper.md)
 - [GetTaxationItemListType](docs/GetTaxationItemListType.md)
 - [GetTaxationItemType](docs/GetTaxationItemType.md)
 - [GetTaxationItemTypeFinanceInformation](docs/GetTaxationItemTypeFinanceInformation.md)
 - [GetTaxationItemTypewithSuccess](docs/GetTaxationItemTypewithSuccess.md)
 - [GetTierType](docs/GetTierType.md)
 - [GetUsageType](docs/GetUsageType.md)
 - [GetUsageWrapper](docs/GetUsageWrapper.md)
 - [GetaPaymentGatwayResponse](docs/GetaPaymentGatwayResponse.md)
 - [GetarPaymentType](docs/GetarPaymentType.md)
 - [GetarPaymentTypeFinanceInformation](docs/GetarPaymentTypeFinanceInformation.md)
 - [GetarPaymentTypewithSuccess](docs/GetarPaymentTypewithSuccess.md)
 - [GetcmTaxItemType](docs/GetcmTaxItemType.md)
 - [GetcmTaxItemTypeFinanceInformation](docs/GetcmTaxItemTypeFinanceInformation.md)
 - [GetdmTaxItemType](docs/GetdmTaxItemType.md)
 - [GetdmTaxItemTypeFinanceInformation](docs/GetdmTaxItemTypeFinanceInformation.md)
 - [GetrsDetailForProductChargeType](docs/GetrsDetailForProductChargeType.md)
 - [GetrsDetailType](docs/GetrsDetailType.md)
 - [GetrsDetailWithoutSuccessType](docs/GetrsDetailWithoutSuccessType.md)
 - [GetrsDetailsByChargeType](docs/GetrsDetailsByChargeType.md)
 - [GetrsDetailsByProductChargeType](docs/GetrsDetailsByProductChargeType.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Invoice](docs/Invoice.md)
 - [InvoiceData](docs/InvoiceData.md)
 - [InvoiceDataInvoice](docs/InvoiceDataInvoice.md)
 - [InvoiceItem](docs/InvoiceItem.md)
 - [InvoiceProcessingOptions](docs/InvoiceProcessingOptions.md)
 - [InvoiceResponseType](docs/InvoiceResponseType.md)
 - [ListOfExchangeRates](docs/ListOfExchangeRates.md)
 - [ModelError](docs/ModelError.md)
 - [NewChargeMetrics](docs/NewChargeMetrics.md)
 - [PaymentCollectionResponseType](docs/PaymentCollectionResponseType.md)
 - [PaymentDebitMemoApplicationApplyRequestType](docs/PaymentDebitMemoApplicationApplyRequestType.md)
 - [PaymentDebitMemoApplicationCreateRequestType](docs/PaymentDebitMemoApplicationCreateRequestType.md)
 - [PaymentDebitMemoApplicationItemApplyRequestType](docs/PaymentDebitMemoApplicationItemApplyRequestType.md)
 - [PaymentDebitMemoApplicationItemCreateRequestType](docs/PaymentDebitMemoApplicationItemCreateRequestType.md)
 - [PaymentDebitMemoApplicationItemUnapplyRequestType](docs/PaymentDebitMemoApplicationItemUnapplyRequestType.md)
 - [PaymentDebitMemoApplicationUnapplyRequestType](docs/PaymentDebitMemoApplicationUnapplyRequestType.md)
 - [PaymentInvoiceApplicationApplyRequestType](docs/PaymentInvoiceApplicationApplyRequestType.md)
 - [PaymentInvoiceApplicationCreateRequestType](docs/PaymentInvoiceApplicationCreateRequestType.md)
 - [PaymentInvoiceApplicationItemApplyRequestType](docs/PaymentInvoiceApplicationItemApplyRequestType.md)
 - [PaymentInvoiceApplicationItemCreateRequestType](docs/PaymentInvoiceApplicationItemCreateRequestType.md)
 - [PaymentInvoiceApplicationItemUnapplyRequestType](docs/PaymentInvoiceApplicationItemUnapplyRequestType.md)
 - [PaymentInvoiceApplicationUnapplyRequestType](docs/PaymentInvoiceApplicationUnapplyRequestType.md)
 - [PostAccountResponseType](docs/PostAccountResponseType.md)
 - [PostAccountType](docs/PostAccountType.md)
 - [PostAccountTypeBillToContact](docs/PostAccountTypeBillToContact.md)
 - [PostAccountTypeCreditCard](docs/PostAccountTypeCreditCard.md)
 - [PostAccountTypeCreditCardCardHolderInfo](docs/PostAccountTypeCreditCardCardHolderInfo.md)
 - [PostAccountTypeSoldToContact](docs/PostAccountTypeSoldToContact.md)
 - [PostAccountTypeSubscription](docs/PostAccountTypeSubscription.md)
 - [PostAccountTypeTaxInfo](docs/PostAccountTypeTaxInfo.md)
 - [PostAccountingCodeResponseType](docs/PostAccountingCodeResponseType.md)
 - [PostAccountingCodeType](docs/PostAccountingCodeType.md)
 - [PostAccountingPeriodResponseType](docs/PostAccountingPeriodResponseType.md)
 - [PostAccountingPeriodType](docs/PostAccountingPeriodType.md)
 - [PostAttachmentResponseType](docs/PostAttachmentResponseType.md)
 - [PostAttachmentType](docs/PostAttachmentType.md)
 - [PostBillingPreviewCreditMemoItem](docs/PostBillingPreviewCreditMemoItem.md)
 - [PostBillingPreviewInvoiceItem](docs/PostBillingPreviewInvoiceItem.md)
 - [PostBillingPreviewParam](docs/PostBillingPreviewParam.md)
 - [PostBillingPreviewRunParam](docs/PostBillingPreviewRunParam.md)
 - [PostCreditMemoEmailRequestType](docs/PostCreditMemoEmailRequestType.md)
 - [PostDebitMemoEmailType](docs/PostDebitMemoEmailType.md)
 - [PostDecryptResponseType](docs/PostDecryptResponseType.md)
 - [PostDecryptionType](docs/PostDecryptionType.md)
 - [PostDistributionItemType](docs/PostDistributionItemType.md)
 - [PostEntityConnectionsResponseType](docs/PostEntityConnectionsResponseType.md)
 - [PostEntityConnectionsType](docs/PostEntityConnectionsType.md)
 - [PostGenerateBillingDocumentType](docs/PostGenerateBillingDocumentType.md)
 - [PostInvoiceCollectInvoicesType](docs/PostInvoiceCollectInvoicesType.md)
 - [PostInvoiceCollectResponseType](docs/PostInvoiceCollectResponseType.md)
 - [PostInvoiceCollectType](docs/PostInvoiceCollectType.md)
 - [PostInvoiceEmailRequestType](docs/PostInvoiceEmailRequestType.md)
 - [PostJournalEntryItemType](docs/PostJournalEntryItemType.md)
 - [PostJournalEntryResponseType](docs/PostJournalEntryResponseType.md)
 - [PostJournalEntrySegmentType](docs/PostJournalEntrySegmentType.md)
 - [PostJournalEntryType](docs/PostJournalEntryType.md)
 - [PostJournalRunResponseType](docs/PostJournalRunResponseType.md)
 - [PostJournalRunTransactionType](docs/PostJournalRunTransactionType.md)
 - [PostJournalRunType](docs/PostJournalRunType.md)
 - [PostMassUpdateResponseType](docs/PostMassUpdateResponseType.md)
 - [PostMassUpdateType](docs/PostMassUpdateType.md)
 - [PostMassUpdateTypeParams](docs/PostMassUpdateTypeParams.md)
 - [PostMemoPdfResponse](docs/PostMemoPdfResponse.md)
 - [PostNonRefRefundType](docs/PostNonRefRefundType.md)
 - [PostNonRefRefundTypeFinanceInformation](docs/PostNonRefRefundTypeFinanceInformation.md)
 - [PostPaymentMethodDecryption](docs/PostPaymentMethodDecryption.md)
 - [PostPaymentMethodDecryptionCardHolderInfo](docs/PostPaymentMethodDecryptionCardHolderInfo.md)
 - [PostPaymentMethodResponseDecryption](docs/PostPaymentMethodResponseDecryption.md)
 - [PostPaymentMethodResponseType](docs/PostPaymentMethodResponseType.md)
 - [PostPaymentMethodType](docs/PostPaymentMethodType.md)
 - [PostPaymentMethodTypeCardHolderInfo](docs/PostPaymentMethodTypeCardHolderInfo.md)
 - [PostQuoteDocResponseType](docs/PostQuoteDocResponseType.md)
 - [PostQuoteDocType](docs/PostQuoteDocType.md)
 - [PostRefundType](docs/PostRefundType.md)
 - [PostRefundTypeFinanceInformation](docs/PostRefundTypeFinanceInformation.md)
 - [PostRevenueScheduleByChargeResponseType](docs/PostRevenueScheduleByChargeResponseType.md)
 - [PostRevenueScheduleByChargeType](docs/PostRevenueScheduleByChargeType.md)
 - [PostRevenueScheduleByChargeTypeRevenueEvent](docs/PostRevenueScheduleByChargeTypeRevenueEvent.md)
 - [PostRevenueScheduleByDateRangeType](docs/PostRevenueScheduleByDateRangeType.md)
 - [PostRevenueScheduleByDateRangeTypeRevenueEvent](docs/PostRevenueScheduleByDateRangeTypeRevenueEvent.md)
 - [PostRevenueScheduleByTransactionRatablyType](docs/PostRevenueScheduleByTransactionRatablyType.md)
 - [PostRevenueScheduleByTransactionRatablyTypeRevenueEvent](docs/PostRevenueScheduleByTransactionRatablyTypeRevenueEvent.md)
 - [PostRevenueScheduleByTransactionResponseType](docs/PostRevenueScheduleByTransactionResponseType.md)
 - [PostRevenueScheduleByTransactionType](docs/PostRevenueScheduleByTransactionType.md)
 - [PostRevenueScheduleByTransactionTypeRevenueEvent](docs/PostRevenueScheduleByTransactionTypeRevenueEvent.md)
 - [PostScCreateType](docs/PostScCreateType.md)
 - [PostSrpCreateType](docs/PostSrpCreateType.md)
 - [PostSubscriptionCancellationResponseType](docs/PostSubscriptionCancellationResponseType.md)
 - [PostSubscriptionCancellationType](docs/PostSubscriptionCancellationType.md)
 - [PostSubscriptionPreviewCreditMemoItemsType](docs/PostSubscriptionPreviewCreditMemoItemsType.md)
 - [PostSubscriptionPreviewInvoiceItemsType](docs/PostSubscriptionPreviewInvoiceItemsType.md)
 - [PostSubscriptionPreviewResponseType](docs/PostSubscriptionPreviewResponseType.md)
 - [PostSubscriptionPreviewResponseTypeChargeMetrics](docs/PostSubscriptionPreviewResponseTypeChargeMetrics.md)
 - [PostSubscriptionPreviewResponseTypeCreditMemo](docs/PostSubscriptionPreviewResponseTypeCreditMemo.md)
 - [PostSubscriptionPreviewType](docs/PostSubscriptionPreviewType.md)
 - [PostSubscriptionPreviewTypePreviewAccountInfo](docs/PostSubscriptionPreviewTypePreviewAccountInfo.md)
 - [PostSubscriptionPreviewTypePreviewAccountInfoBillToContact](docs/PostSubscriptionPreviewTypePreviewAccountInfoBillToContact.md)
 - [PostSubscriptionResponseType](docs/PostSubscriptionResponseType.md)
 - [PostSubscriptionType](docs/PostSubscriptionType.md)
 - [PostTaxationItemForCmType](docs/PostTaxationItemForCmType.md)
 - [PostTaxationItemForCmTypeFinanceInformation](docs/PostTaxationItemForCmTypeFinanceInformation.md)
 - [PostTaxationItemForDmType](docs/PostTaxationItemForDmType.md)
 - [PostTaxationItemForDmTypeFinanceInformation](docs/PostTaxationItemForDmTypeFinanceInformation.md)
 - [PostTaxationItemListForCmType](docs/PostTaxationItemListForCmType.md)
 - [PostTaxationItemListForDmType](docs/PostTaxationItemListForDmType.md)
 - [PostTierType](docs/PostTierType.md)
 - [PostUsageResponseType](docs/PostUsageResponseType.md)
 - [PosthmacSignatureResponseType](docs/PosthmacSignatureResponseType.md)
 - [PosthmacSignatureType](docs/PosthmacSignatureType.md)
 - [PostrsaSignatureResponseType](docs/PostrsaSignatureResponseType.md)
 - [PostrsaSignatureType](docs/PostrsaSignatureType.md)
 - [ProvisionEntityResponseType](docs/ProvisionEntityResponseType.md)
 - [ProxyActionamendRequest](docs/ProxyActionamendRequest.md)
 - [ProxyActionamendResponse](docs/ProxyActionamendResponse.md)
 - [ProxyActioncreateRequest](docs/ProxyActioncreateRequest.md)
 - [ProxyActioncreateResponse](docs/ProxyActioncreateResponse.md)
 - [ProxyActiondeleteRequest](docs/ProxyActiondeleteRequest.md)
 - [ProxyActiondeleteResponse](docs/ProxyActiondeleteResponse.md)
 - [ProxyActionexecuteRequest](docs/ProxyActionexecuteRequest.md)
 - [ProxyActionexecuteResponse](docs/ProxyActionexecuteResponse.md)
 - [ProxyActiongenerateRequest](docs/ProxyActiongenerateRequest.md)
 - [ProxyActiongenerateResponse](docs/ProxyActiongenerateResponse.md)
 - [ProxyActionqueryMoreRequest](docs/ProxyActionqueryMoreRequest.md)
 - [ProxyActionqueryMoreResponse](docs/ProxyActionqueryMoreResponse.md)
 - [ProxyActionqueryRequest](docs/ProxyActionqueryRequest.md)
 - [ProxyActionqueryResponse](docs/ProxyActionqueryResponse.md)
 - [ProxyActionsubscribeRequest](docs/ProxyActionsubscribeRequest.md)
 - [ProxyActionsubscribeResponse](docs/ProxyActionsubscribeResponse.md)
 - [ProxyActionupdateRequest](docs/ProxyActionupdateRequest.md)
 - [ProxyActionupdateResponse](docs/ProxyActionupdateResponse.md)
 - [ProxyBadRequestResponse](docs/ProxyBadRequestResponse.md)
 - [ProxyBadRequestResponseErrors](docs/ProxyBadRequestResponseErrors.md)
 - [ProxyCreateAccount](docs/ProxyCreateAccount.md)
 - [ProxyCreateAmendment](docs/ProxyCreateAmendment.md)
 - [ProxyCreateBillRun](docs/ProxyCreateBillRun.md)
 - [ProxyCreateContact](docs/ProxyCreateContact.md)
 - [ProxyCreateExport](docs/ProxyCreateExport.md)
 - [ProxyCreateInvoice](docs/ProxyCreateInvoice.md)
 - [ProxyCreateInvoiceAdjustment](docs/ProxyCreateInvoiceAdjustment.md)
 - [ProxyCreateInvoicePayment](docs/ProxyCreateInvoicePayment.md)
 - [ProxyCreateOrModifyResponse](docs/ProxyCreateOrModifyResponse.md)
 - [ProxyCreatePayment](docs/ProxyCreatePayment.md)
 - [ProxyCreatePaymentMethod](docs/ProxyCreatePaymentMethod.md)
 - [ProxyCreateProduct](docs/ProxyCreateProduct.md)
 - [ProxyCreateProductRatePlan](docs/ProxyCreateProductRatePlan.md)
 - [ProxyCreateRefund](docs/ProxyCreateRefund.md)
 - [ProxyCreateTaxationItem](docs/ProxyCreateTaxationItem.md)
 - [ProxyCreateUnitOfMeasure](docs/ProxyCreateUnitOfMeasure.md)
 - [ProxyCreateUsage](docs/ProxyCreateUsage.md)
 - [ProxyDeleteResponse](docs/ProxyDeleteResponse.md)
 - [ProxyGetAccount](docs/ProxyGetAccount.md)
 - [ProxyGetAmendment](docs/ProxyGetAmendment.md)
 - [ProxyGetBillRun](docs/ProxyGetBillRun.md)
 - [ProxyGetCommunicationProfile](docs/ProxyGetCommunicationProfile.md)
 - [ProxyGetContact](docs/ProxyGetContact.md)
 - [ProxyGetCreditBalanceAdjustment](docs/ProxyGetCreditBalanceAdjustment.md)
 - [ProxyGetExport](docs/ProxyGetExport.md)
 - [ProxyGetFeature](docs/ProxyGetFeature.md)
 - [ProxyGetImport](docs/ProxyGetImport.md)
 - [ProxyGetInvoice](docs/ProxyGetInvoice.md)
 - [ProxyGetInvoiceAdjustment](docs/ProxyGetInvoiceAdjustment.md)
 - [ProxyGetInvoiceItem](docs/ProxyGetInvoiceItem.md)
 - [ProxyGetInvoiceItemAdjustment](docs/ProxyGetInvoiceItemAdjustment.md)
 - [ProxyGetInvoicePayment](docs/ProxyGetInvoicePayment.md)
 - [ProxyGetInvoiceSplit](docs/ProxyGetInvoiceSplit.md)
 - [ProxyGetInvoiceSplitItem](docs/ProxyGetInvoiceSplitItem.md)
 - [ProxyGetPayment](docs/ProxyGetPayment.md)
 - [ProxyGetPaymentMethod](docs/ProxyGetPaymentMethod.md)
 - [ProxyGetPaymentMethodSnapshot](docs/ProxyGetPaymentMethodSnapshot.md)
 - [ProxyGetPaymentMethodTransactionLog](docs/ProxyGetPaymentMethodTransactionLog.md)
 - [ProxyGetPaymentTransactionLog](docs/ProxyGetPaymentTransactionLog.md)
 - [ProxyGetProduct](docs/ProxyGetProduct.md)
 - [ProxyGetProductFeature](docs/ProxyGetProductFeature.md)
 - [ProxyGetProductRatePlan](docs/ProxyGetProductRatePlan.md)
 - [ProxyGetProductRatePlanCharge](docs/ProxyGetProductRatePlanCharge.md)
 - [ProxyGetProductRatePlanChargeTier](docs/ProxyGetProductRatePlanChargeTier.md)
 - [ProxyGetRatePlan](docs/ProxyGetRatePlan.md)
 - [ProxyGetRatePlanCharge](docs/ProxyGetRatePlanCharge.md)
 - [ProxyGetRatePlanChargeTier](docs/ProxyGetRatePlanChargeTier.md)
 - [ProxyGetRefund](docs/ProxyGetRefund.md)
 - [ProxyGetRefundInvoicePayment](docs/ProxyGetRefundInvoicePayment.md)
 - [ProxyGetRefundTransactionLog](docs/ProxyGetRefundTransactionLog.md)
 - [ProxyGetSubscription](docs/ProxyGetSubscription.md)
 - [ProxyGetSubscriptionProductFeature](docs/ProxyGetSubscriptionProductFeature.md)
 - [ProxyGetTaxationItem](docs/ProxyGetTaxationItem.md)
 - [ProxyGetUnitOfMeasure](docs/ProxyGetUnitOfMeasure.md)
 - [ProxyGetUsage](docs/ProxyGetUsage.md)
 - [ProxyModifyAccount](docs/ProxyModifyAccount.md)
 - [ProxyModifyAmendment](docs/ProxyModifyAmendment.md)
 - [ProxyModifyBillRun](docs/ProxyModifyBillRun.md)
 - [ProxyModifyContact](docs/ProxyModifyContact.md)
 - [ProxyModifyInvoice](docs/ProxyModifyInvoice.md)
 - [ProxyModifyInvoiceAdjustment](docs/ProxyModifyInvoiceAdjustment.md)
 - [ProxyModifyInvoicePayment](docs/ProxyModifyInvoicePayment.md)
 - [ProxyModifyPayment](docs/ProxyModifyPayment.md)
 - [ProxyModifyPaymentMethod](docs/ProxyModifyPaymentMethod.md)
 - [ProxyModifyProduct](docs/ProxyModifyProduct.md)
 - [ProxyModifyProductRatePlan](docs/ProxyModifyProductRatePlan.md)
 - [ProxyModifyRefund](docs/ProxyModifyRefund.md)
 - [ProxyModifySubscription](docs/ProxyModifySubscription.md)
 - [ProxyModifyTaxationItem](docs/ProxyModifyTaxationItem.md)
 - [ProxyModifyUnitOfMeasure](docs/ProxyModifyUnitOfMeasure.md)
 - [ProxyModifyUsage](docs/ProxyModifyUsage.md)
 - [ProxyNoDataResponse](docs/ProxyNoDataResponse.md)
 - [PutAcceptUserAccessResponseType](docs/PutAcceptUserAccessResponseType.md)
 - [PutAccountType](docs/PutAccountType.md)
 - [PutAccountTypeBillToContact](docs/PutAccountTypeBillToContact.md)
 - [PutAccountTypeSoldToContact](docs/PutAccountTypeSoldToContact.md)
 - [PutAccountingCodeType](docs/PutAccountingCodeType.md)
 - [PutAccountingPeriodType](docs/PutAccountingPeriodType.md)
 - [PutAllocateManuallyType](docs/PutAllocateManuallyType.md)
 - [PutAttachmentType](docs/PutAttachmentType.md)
 - [PutBasicSummaryJournalEntryType](docs/PutBasicSummaryJournalEntryType.md)
 - [PutCatalogType](docs/PutCatalogType.md)
 - [PutCreditMemoItemType](docs/PutCreditMemoItemType.md)
 - [PutCreditMemoItemTypeFinanceInformation](docs/PutCreditMemoItemTypeFinanceInformation.md)
 - [PutCreditMemoTaxItemType](docs/PutCreditMemoTaxItemType.md)
 - [PutCreditMemoTaxItemTypeFinanceInformation](docs/PutCreditMemoTaxItemTypeFinanceInformation.md)
 - [PutCreditMemoType](docs/PutCreditMemoType.md)
 - [PutDebitMemoItemType](docs/PutDebitMemoItemType.md)
 - [PutDebitMemoItemTypeFinanceInformation](docs/PutDebitMemoItemTypeFinanceInformation.md)
 - [PutDebitMemoTaxItemType](docs/PutDebitMemoTaxItemType.md)
 - [PutDebitMemoTaxItemTypeFinanceInformation](docs/PutDebitMemoTaxItemTypeFinanceInformation.md)
 - [PutDebitMemoType](docs/PutDebitMemoType.md)
 - [PutDenyUserAccessResponseType](docs/PutDenyUserAccessResponseType.md)
 - [PutEntityConnectionsAcceptResponseType](docs/PutEntityConnectionsAcceptResponseType.md)
 - [PutEntityConnectionsDenyResponseType](docs/PutEntityConnectionsDenyResponseType.md)
 - [PutEntityConnectionsDisconnectResponseType](docs/PutEntityConnectionsDisconnectResponseType.md)
 - [PutEventRiDetailType](docs/PutEventRiDetailType.md)
 - [PutJournalEntryItemType](docs/PutJournalEntryItemType.md)
 - [PutPaymentMethodResponseType](docs/PutPaymentMethodResponseType.md)
 - [PutPaymentMethodType](docs/PutPaymentMethodType.md)
 - [PutRefundType](docs/PutRefundType.md)
 - [PutRefundTypeFinanceInformation](docs/PutRefundTypeFinanceInformation.md)
 - [PutRenewSubscriptionResponseType](docs/PutRenewSubscriptionResponseType.md)
 - [PutRenewSubscriptionType](docs/PutRenewSubscriptionType.md)
 - [PutRevenueScheduleResponseType](docs/PutRevenueScheduleResponseType.md)
 - [PutReverseInvoiceResponseType](docs/PutReverseInvoiceResponseType.md)
 - [PutReverseInvoiceResponseTypeCreditMemo](docs/PutReverseInvoiceResponseTypeCreditMemo.md)
 - [PutReverseInvoiceType](docs/PutReverseInvoiceType.md)
 - [PutScAddType](docs/PutScAddType.md)
 - [PutScUpdateType](docs/PutScUpdateType.md)
 - [PutScheduleRiDetailType](docs/PutScheduleRiDetailType.md)
 - [PutSendUserAccessRequestResponseType](docs/PutSendUserAccessRequestResponseType.md)
 - [PutSendUserAccessRequestType](docs/PutSendUserAccessRequestType.md)
 - [PutSpecificDateAllocationType](docs/PutSpecificDateAllocationType.md)
 - [PutSrpAddType](docs/PutSrpAddType.md)
 - [PutSrpRemoveType](docs/PutSrpRemoveType.md)
 - [PutSrpUpdateType](docs/PutSrpUpdateType.md)
 - [PutSubscriptionPreviewInvoiceItemsType](docs/PutSubscriptionPreviewInvoiceItemsType.md)
 - [PutSubscriptionResponseType](docs/PutSubscriptionResponseType.md)
 - [PutSubscriptionResponseTypeChargeMetrics](docs/PutSubscriptionResponseTypeChargeMetrics.md)
 - [PutSubscriptionResponseTypeCreditMemo](docs/PutSubscriptionResponseTypeCreditMemo.md)
 - [PutSubscriptionResumeResponseType](docs/PutSubscriptionResumeResponseType.md)
 - [PutSubscriptionResumeType](docs/PutSubscriptionResumeType.md)
 - [PutSubscriptionSuspendResponseType](docs/PutSubscriptionSuspendResponseType.md)
 - [PutSubscriptionSuspendType](docs/PutSubscriptionSuspendType.md)
 - [PutSubscriptionType](docs/PutSubscriptionType.md)
 - [PutTaxationItemType](docs/PutTaxationItemType.md)
 - [PutrsBasicInfoType](docs/PutrsBasicInfoType.md)
 - [PutrsTermType](docs/PutrsTermType.md)
 - [RatePlan](docs/RatePlan.md)
 - [RatePlanChargeData](docs/RatePlanChargeData.md)
 - [RatePlanChargeDataInRatePlanData](docs/RatePlanChargeDataInRatePlanData.md)
 - [RatePlanChargeDataInRatePlanDataRatePlanCharge](docs/RatePlanChargeDataInRatePlanDataRatePlanCharge.md)
 - [RatePlanChargeDataRatePlanCharge](docs/RatePlanChargeDataRatePlanCharge.md)
 - [RatePlanChargeTier](docs/RatePlanChargeTier.md)
 - [RatePlanData](docs/RatePlanData.md)
 - [RatePlanDataRatePlan](docs/RatePlanDataRatePlan.md)
 - [RatePlanDataSubscriptionProductFeatureList](docs/RatePlanDataSubscriptionProductFeatureList.md)
 - [RefundCreditMemoItemType](docs/RefundCreditMemoItemType.md)
 - [RefundPartResponseType](docs/RefundPartResponseType.md)
 - [RefundPartResponseTypewithSuccess](docs/RefundPartResponseTypewithSuccess.md)
 - [RevenueScheduleItemType](docs/RevenueScheduleItemType.md)
 - [SaveResult](docs/SaveResult.md)
 - [SubscribeRequest](docs/SubscribeRequest.md)
 - [SubscribeRequestAccount](docs/SubscribeRequestAccount.md)
 - [SubscribeRequestBillToContact](docs/SubscribeRequestBillToContact.md)
 - [SubscribeRequestPaymentMethod](docs/SubscribeRequestPaymentMethod.md)
 - [SubscribeRequestPaymentMethodGatewayOptionData](docs/SubscribeRequestPaymentMethodGatewayOptionData.md)
 - [SubscribeRequestPreviewOptions](docs/SubscribeRequestPreviewOptions.md)
 - [SubscribeRequestSoldToContact](docs/SubscribeRequestSoldToContact.md)
 - [SubscribeRequestSubscribeOptions](docs/SubscribeRequestSubscribeOptions.md)
 - [SubscribeRequestSubscribeOptionsElectronicPaymentOptions](docs/SubscribeRequestSubscribeOptionsElectronicPaymentOptions.md)
 - [SubscribeRequestSubscribeOptionsExternalPaymentOptions](docs/SubscribeRequestSubscribeOptionsExternalPaymentOptions.md)
 - [SubscribeRequestSubscribeOptionsSubscribeInvoiceProcessingOptions](docs/SubscribeRequestSubscribeOptionsSubscribeInvoiceProcessingOptions.md)
 - [SubscribeRequestSubscriptionData](docs/SubscribeRequestSubscriptionData.md)
 - [SubscribeRequestSubscriptionDataSubscription](docs/SubscribeRequestSubscriptionDataSubscription.md)
 - [SubscribeResult](docs/SubscribeResult.md)
 - [SubscribeResultChargeMetricsData](docs/SubscribeResultChargeMetricsData.md)
 - [SubscribeResultInvoiceResult](docs/SubscribeResultInvoiceResult.md)
 - [SubscriptionProductFeature](docs/SubscriptionProductFeature.md)
 - [SubscriptionProductFeatureList](docs/SubscriptionProductFeatureList.md)
 - [TokenResponse](docs/TokenResponse.md)
 - [TransferPaymentType](docs/TransferPaymentType.md)
 - [UnapplyCreditMemoType](docs/UnapplyCreditMemoType.md)
 - [UnapplyPaymentType](docs/UnapplyPaymentType.md)
 - [UpdateEntityResponseType](docs/UpdateEntityResponseType.md)
 - [UpdateEntityType](docs/UpdateEntityType.md)
 - [UpdatePaymentType](docs/UpdatePaymentType.md)
 - [ZObject](docs/ZObject.md)
 - [ZObjectUpdate](docs/ZObjectUpdate.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author

docs@zuora.com

