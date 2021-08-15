# DashboardSearchSummary

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Total** | Pointer to **int32** |  | [optional] 
**Dashboards** | Pointer to [**[]DashboardSearchSummaryDefinition**](DashboardSearchSummaryDefinition.md) | List of dashboard definitions. | [optional] 

## Methods

### NewDashboardSearchSummary

`func NewDashboardSearchSummary() *DashboardSearchSummary`

NewDashboardSearchSummary instantiates a new DashboardSearchSummary object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDashboardSearchSummaryWithDefaults

`func NewDashboardSearchSummaryWithDefaults() *DashboardSearchSummary`

NewDashboardSearchSummaryWithDefaults instantiates a new DashboardSearchSummary object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTotal

`func (o *DashboardSearchSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DashboardSearchSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DashboardSearchSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DashboardSearchSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetDashboards

`func (o *DashboardSearchSummary) GetDashboards() []DashboardSearchSummaryDefinition`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *DashboardSearchSummary) GetDashboardsOk() (*[]DashboardSearchSummaryDefinition, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *DashboardSearchSummary) SetDashboards(v []DashboardSearchSummaryDefinition)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *DashboardSearchSummary) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


