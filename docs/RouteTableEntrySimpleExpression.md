# RouteTableEntrySimpleExpression

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** | Possible field names to use on filters:  * &#x60;/type&#x60; - Route table entry type  * &#x60;/prefix&#x60; - Route table entry type  * &#x60;/nextHop&#x60; - Route table entry type  * &#x60;/state&#x60; - Route table entry type  * &#x60;/_*&#x60; - all-category search  | [optional] [default to null]
**Operator** | **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;!&#x3D;&#x60; - not equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;[NOT] BETWEEN&#x60; - (not) between  * &#x60;[NOT] LIKE&#x60; - (not) like  * &#x60;[NOT] IN&#x60; - (not) in  * &#x60;~*&#x60; - case-insensitive like  | [optional] [default to null]
**Values** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

