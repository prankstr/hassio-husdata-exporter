package models

type VPData struct {
	// Common
	RadiatorReturn      *int64 `json:"0001,omitempty"`
	HeatCarrierForward  *int64 `json:"0004,omitempty"`
	BrineInCondensor    *int64 `json:"0005,omitempty"`
	BrineOutEvaporator  *int64 `json:"0006,omitempty"`
	Outdoor             *int64 `json:"0007,omitempty"`
	Indoor              *int64 `json:"0008,omitempty"`
	WarmWaterTop        *int64 `json:"0009,omitempty"`
	HotGas              *int64 `json:"000B,omitempty"`
	Pool                *int64 `json:"0011,omitempty"`
	HeatingSetpoint     *int64 `json:"0107,omitempty"`
	RoomTempSetpoint    *int64 `json:"0203,omitempty"`
	WarmWaterStopTemp   *int64 `json:"0208,omitempty"`
	WarmWaterStartTemp  *int64 `json:"0212,omitempty"`
	Compressor          *int64 `json:"1A01,omitempty"`
	PumpHeatCircuit     *int64 `json:"1A05,omitempty"`
	PumpRadiator        *int64 `json:"1A06,omitempty"`
	SwitchValve1        *int64 `json:"1A07,omitempty"`
	SwitchValve2        *int64 `json:"1A08,omitempty"`
	RoomSensorInfluence *int64 `json:"2204,omitempty"`
	PowerConsumption    *int64 `json:"CFAA,omitempty"`

	// LW
	WarmWaterMid    *int64 `json:"000A,omitempty"`
	PumpColdCircuit *int64 `json:"1A04,omitempty"`

	// AW
	AirIntake *int64 `json:"000E,omitempty"`
	Fan       *int64 `json:"1A09,omitempty"`

	// Styr2002
	SuctionGas        *int64 `json:"000C,omitempty"`
	LiquidFlow        *int64 `json:"000D,omitempty"`
	RadiatorForward2  *int64 `json:"0020,omitempty"`
	RadiatorReturn2   *int64 `json:"0022,omitempty"`
	AddHeatStep1      *int64 `json:"1A02,omitempty"`
	AddHeatStep2      *int64 `json:"1A03,omitempty"`
	OperatingMode     *int64 `json:"2201,omitempty"`
	AlarmReset        *int64 `json:"22F2,omitempty"`
	LoadL1            *int64 `json:"4101,omitempty"`
	LoadL2            *int64 `json:"4102,omitempty"`
	LoadL3            *int64 `json:"4103,omitempty"`
	HeatSet3Parallel  *int64 `json:"2207,omitempty"`
	HeatSet1CurveL2   *int64 `json:"2222,omitempty"`
	HeatSet3Parall2   *int64 `json:"2224,omitempty"`
	DegreeMinIntegral *int64 `json:"8105,omitempty"`
	CompressorStarts  *int64 `json:"BC61,omitempty"`
	CompressorRuntime *int64 `json:"6C60,omitempty"`
	AuxRuntime        *int64 `json:"6C63,omitempty"`
	WarmWaterRuntime  *int64 `json:"6C64,omitempty"`

	// Rego
	HeatCarrierReturn *int64 `json:"0003,omitempty"`
	HeatSet1CurveR    *int64 `json:"0206,omitempty"`
	HighPressostat    *int64 `json:"1A0A,omitempty"`
	LowPressostat     *int64 `json:"1A0B,omitempty"`
	HeatingCable      *int64 `json:"1A0C,omitempty"`
	CrankCaseHeater   *int64 `json:"1A0D,omitempty"`
	AddHeatStatus     *int64 `json:"3104,omitempty"`

	// Conflicting, these require special handling
	Alarm              *int64 `json:"1A20,omitempty"`
	RegoAlarm          *int64 `json:"2A20,omitempty"`
	RegoHeatSet1CurveL *int64 `json:"0205,omitempty"`
	HeatSet1CurveL     *int64 `json:"2205,omitempty"`
}
