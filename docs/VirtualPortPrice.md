# VirtualPortPrice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Unique identifier assigned to the virtual port. Either the uuid or the remaining attributes must be supplied. | [optional] [default to null]
**Type_** | [***VirtualPortType**](VirtualPortType.md) |  | [optional] [default to null]
**Location** | [***VirtualPortLocation**](VirtualPortLocation.md) |  | [optional] [default to null]
**Lag** | [***LinkAggregationGroup**](LinkAggregationGroup.md) |  | [optional] [default to null]
**PhysicalPortsQuantity** | **int32** | Number of physical ports requested. The defaults is 1. | [optional] [default to 1]
**Bandwidth** | **int32** | Aggregated data transfer capacity,  expressed as follows &lt;br&gt; -&gt; Mbps, megabits (1 million bits) per second &lt;br&gt; -&gt; Gbps, gigabits (1 billion bits) per second &lt;br&gt; Bandwidth must be divisible by physicalPortsQuantity. | [optional] [default to null]
**Redundancy** | [***VirtualPortRedundancy**](VirtualPortRedundancy.md) |  | [optional] [default to null]
**ConnectivitySource** | [***ConnectivitySource**](ConnectivitySource.md) |  | [optional] [default to null]
**ServiceType** | [***VirtualPortServiceType**](VirtualPortServiceType.md) |  | [optional] [default to null]
**Settings** | [***VirtualPortConfiguration**](VirtualPortConfiguration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

