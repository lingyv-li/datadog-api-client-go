# MonitorGroupSearchResponseCounts

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Status** | Pointer to [**[]MonitorGroupSearchResponseCountsStatus**](MonitorGroupSearchResponseCountsStatus.md) | Search facets. | [optional] 
**Type** | Pointer to [**[]MonitorGroupSearchResponseCountsStatus**](MonitorGroupSearchResponseCountsStatus.md) | Search facets. | [optional] 

## Methods

### NewMonitorGroupSearchResponseCounts

`func NewMonitorGroupSearchResponseCounts() *MonitorGroupSearchResponseCounts`

NewMonitorGroupSearchResponseCounts instantiates a new MonitorGroupSearchResponseCounts object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorGroupSearchResponseCountsWithDefaults

`func NewMonitorGroupSearchResponseCountsWithDefaults() *MonitorGroupSearchResponseCounts`

NewMonitorGroupSearchResponseCountsWithDefaults instantiates a new MonitorGroupSearchResponseCounts object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetStatus

`func (o *MonitorGroupSearchResponseCounts) GetStatus() []MonitorGroupSearchResponseCountsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorGroupSearchResponseCounts) GetStatusOk() (*[]MonitorGroupSearchResponseCountsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorGroupSearchResponseCounts) SetStatus(v []MonitorGroupSearchResponseCountsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitorGroupSearchResponseCounts) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MonitorGroupSearchResponseCounts) GetType() []MonitorGroupSearchResponseCountsStatus`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorGroupSearchResponseCounts) GetTypeOk() (*[]MonitorGroupSearchResponseCountsStatus, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorGroupSearchResponseCounts) SetType(v []MonitorGroupSearchResponseCountsStatus)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorGroupSearchResponseCounts) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


