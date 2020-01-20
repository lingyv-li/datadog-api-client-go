# MetricMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Metric description | [optional] 
**Integration** | Pointer to **string** | Name of the integration that sent the metric if applicable | [optional] [readonly] 
**PerUnit** | Pointer to **string** | Per unit of the metric such as &#x60;second&#x60; in &#x60;bytes per second&#x60; | [optional] 
**ShortName** | Pointer to **string** | A more human-readable and abbreviated version of the metric name | [optional] 
**StatsdInterval** | Pointer to **int64** | Statsd flush interval of the metric in seconds if applicable | [optional] 
**Type** | Pointer to **string** | Metric type such as &#x60;gauge&#x60; or &#x60;rate&#x60; | 
**Unit** | Pointer to **string** | Primary unit of the metric such as &#x60;byte&#x60; or &#x60;operation&#x60; | [optional] 

## Methods

### GetDescription

`func (o *MetricMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricMetadata) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MetricMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MetricMetadata) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetIntegration

`func (o *MetricMetadata) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *MetricMetadata) GetIntegrationOk() (string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegration

`func (o *MetricMetadata) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegration

`func (o *MetricMetadata) SetIntegration(v string)`

SetIntegration gets a reference to the given string and assigns it to the Integration field.

### GetPerUnit

`func (o *MetricMetadata) GetPerUnit() string`

GetPerUnit returns the PerUnit field if non-nil, zero value otherwise.

### GetPerUnitOk

`func (o *MetricMetadata) GetPerUnitOk() (string, bool)`

GetPerUnitOk returns a tuple with the PerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPerUnit

`func (o *MetricMetadata) HasPerUnit() bool`

HasPerUnit returns a boolean if a field has been set.

### SetPerUnit

`func (o *MetricMetadata) SetPerUnit(v string)`

SetPerUnit gets a reference to the given string and assigns it to the PerUnit field.

### GetShortName

`func (o *MetricMetadata) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *MetricMetadata) GetShortNameOk() (string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShortName

`func (o *MetricMetadata) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortName

`func (o *MetricMetadata) SetShortName(v string)`

SetShortName gets a reference to the given string and assigns it to the ShortName field.

### GetStatsdInterval

`func (o *MetricMetadata) GetStatsdInterval() int64`

GetStatsdInterval returns the StatsdInterval field if non-nil, zero value otherwise.

### GetStatsdIntervalOk

`func (o *MetricMetadata) GetStatsdIntervalOk() (int64, bool)`

GetStatsdIntervalOk returns a tuple with the StatsdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatsdInterval

`func (o *MetricMetadata) HasStatsdInterval() bool`

HasStatsdInterval returns a boolean if a field has been set.

### SetStatsdInterval

`func (o *MetricMetadata) SetStatsdInterval(v int64)`

SetStatsdInterval gets a reference to the given int64 and assigns it to the StatsdInterval field.

### GetType

`func (o *MetricMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricMetadata) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *MetricMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *MetricMetadata) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetUnit

`func (o *MetricMetadata) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricMetadata) GetUnitOk() (string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnit

`func (o *MetricMetadata) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnit

`func (o *MetricMetadata) SetUnit(v string)`

SetUnit gets a reference to the given string and assigns it to the Unit field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

