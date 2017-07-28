# ProxyGetImport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedById** | **string** |  The user ID of the person who created the import.  **Character limit**: 32  **Values**: automatically generated  | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) |  The date when the import was created.  **Character limit**: 29  **Values**: automatically generated  | [optional] [default to null]
**Id** | **string** | Object identifier. | [optional] [default to null]
**ImportType** | **string** |  The type of item imported.  **Character limit**: 7  **Values**: Usage  | [optional] [default to null]
**ImportedCount** | **int32** | The number of records successfully imported.  **Values**: automatically generated  | [optional] [default to null]
**Md5** | **string** |  A check to validate the import file&#39;s integrity.  **Character limit:** 32  **System-generated:** no  **Values**: a string of 32 characters or fewer  | [optional] [default to null]
**Name** | **string** |  A descriptive name for the import.  **Character limit:** 100  **Values:** one of the following:  - a string of 100 characters or fewer - if NULL default is: &#x60;import &lt;ImportType_value&gt;&#x60;  | [optional] [default to null]
**OriginalResourceUrl** | **string** |  The URL for your import file, which contains your records for upload. When you upload the file, Zuora assigns it to this address.  **Values:** automatic dynamically-generated URL  | [optional] [default to null]
**ResultResourceUrl** | **string** |  The URL for the import result file, which is a zipped CSV file.  **Values**: automatic dynamically-generated URL  | [optional] [default to null]
**Status** | **string** | The status of the import process.  **Values**: automatically generated using one of the following values:  - Pending - Processing - Completed - Failed  | [optional] [default to null]
**StatusReason** | **string** |  The reason for the system-generated status. Use this information if the import fails.  **Character limit**: 2000  **Values**: automatically generated error message  | [optional] [default to null]
**TotalCount** | **int32** |  The number of records in the import file.  **Character limit**:  **Values**: automatically generated  | [optional] [default to null]
**UpdatedById** | **string** |  The ID of the user who last updated the import.  **Character limit**: 32  **Values**: automatically generated  | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) |  The date when the import was last updated. **Character limit**: 29 **Values**: automatically generated  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


