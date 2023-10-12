# PtpAdvanceConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeScale** | **string** | Time Scale value, ARB denotes Arbitrary and PTP denotes Precision Time Protocol. | [optional] [default to null]
**Domain** | **int32** |  | [optional] [default to null]
**Priority1** | **int32** |  | [optional] [default to null]
**Priority2** | **int32** |  | [optional] [default to null]
**LogAnnounceInterval** | **int32** | The mean time interval between Announce messages. A shorter interval makes ptp4l react faster to the changes in the master-slave hierarchy. The interval should be the same in the whole domain. It&#x27;s specified as a power of two in seconds. The default is 1 (2 seconds). | [optional] [default to null]
**LogSyncInterval** | **int32** | The mean time interval between Sync messages. A shorter interval may improve accuracy of the local clock. It&#x27;s specified as a power of two in seconds. The default is 0 (1 second). | [optional] [default to null]
**LogDelayReqInterval** | **int32** |  | [optional] [default to null]
**TransportMode** | **string** |  | [optional] [default to null]
**GrantTime** | **int32** | Unicast Grant Time in seconds. For Multicast and Hybrid transport modes, grant time defaults to 300 seconds. For Unicast mode, grant time can be between 30 to 7200. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

