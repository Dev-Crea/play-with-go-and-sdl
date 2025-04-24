# InstallShipModule201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**Agent**](Agent.md) |  | 
**Modules** | [**[]ShipModule**](ShipModule.md) |  | 
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 
**Transaction** | [**InstallShipModule201ResponseDataTransaction**](InstallShipModule201ResponseDataTransaction.md) |  | 

## Methods

### NewInstallShipModule201ResponseData

`func NewInstallShipModule201ResponseData(agent Agent, modules []ShipModule, cargo ShipCargo, transaction InstallShipModule201ResponseDataTransaction, ) *InstallShipModule201ResponseData`

NewInstallShipModule201ResponseData instantiates a new InstallShipModule201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallShipModule201ResponseDataWithDefaults

`func NewInstallShipModule201ResponseDataWithDefaults() *InstallShipModule201ResponseData`

NewInstallShipModule201ResponseDataWithDefaults instantiates a new InstallShipModule201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *InstallShipModule201ResponseData) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *InstallShipModule201ResponseData) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *InstallShipModule201ResponseData) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetModules

`func (o *InstallShipModule201ResponseData) GetModules() []ShipModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *InstallShipModule201ResponseData) GetModulesOk() (*[]ShipModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *InstallShipModule201ResponseData) SetModules(v []ShipModule)`

SetModules sets Modules field to given value.


### GetCargo

`func (o *InstallShipModule201ResponseData) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *InstallShipModule201ResponseData) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *InstallShipModule201ResponseData) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.


### GetTransaction

`func (o *InstallShipModule201ResponseData) GetTransaction() InstallShipModule201ResponseDataTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *InstallShipModule201ResponseData) GetTransactionOk() (*InstallShipModule201ResponseDataTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *InstallShipModule201ResponseData) SetTransaction(v InstallShipModule201ResponseDataTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


