# PatchShipNav200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nav** | [**ShipNav**](ShipNav.md) |  | 
**Fuel** | [**ShipFuel**](ShipFuel.md) |  | 
**Events** | [**[]ShipConditionEvent**](ShipConditionEvent.md) |  | 

## Methods

### NewPatchShipNav200ResponseData

`func NewPatchShipNav200ResponseData(nav ShipNav, fuel ShipFuel, events []ShipConditionEvent, ) *PatchShipNav200ResponseData`

NewPatchShipNav200ResponseData instantiates a new PatchShipNav200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchShipNav200ResponseDataWithDefaults

`func NewPatchShipNav200ResponseDataWithDefaults() *PatchShipNav200ResponseData`

NewPatchShipNav200ResponseDataWithDefaults instantiates a new PatchShipNav200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNav

`func (o *PatchShipNav200ResponseData) GetNav() ShipNav`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *PatchShipNav200ResponseData) GetNavOk() (*ShipNav, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *PatchShipNav200ResponseData) SetNav(v ShipNav)`

SetNav sets Nav field to given value.


### GetFuel

`func (o *PatchShipNav200ResponseData) GetFuel() ShipFuel`

GetFuel returns the Fuel field if non-nil, zero value otherwise.

### GetFuelOk

`func (o *PatchShipNav200ResponseData) GetFuelOk() (*ShipFuel, bool)`

GetFuelOk returns a tuple with the Fuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuel

`func (o *PatchShipNav200ResponseData) SetFuel(v ShipFuel)`

SetFuel sets Fuel field to given value.


### GetEvents

`func (o *PatchShipNav200ResponseData) GetEvents() []ShipConditionEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PatchShipNav200ResponseData) GetEventsOk() (*[]ShipConditionEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PatchShipNav200ResponseData) SetEvents(v []ShipConditionEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


