package models

type VPData struct {
	// Common
	RadiatorReturn      *int64 `json:"0001"`
	HeatCarrierForward  *int64 `json:"0004"`
	BrineInCondensor    *int64 `json:"0005"`
	BrineOutEvaporator  *int64 `json:"0006"`
	Outdoor             *int64 `json:"0007"`
	Indoor              *int64 `json:"0008"`
	WarmWaterTop        *int64 `json:"0009"`
	HotGas              *int64 `json:"000B"`
	Pool                *int64 `json:"0011"`
	HeatingSetpoint     *int64 `json:"0107"`
	RoomTempSetpoint    *int64 `json:"0203"`
	WarmWaterStopTemp   *int64 `json:"0208"`
	WarmWaterStartTemp  *int64 `json:"0212"`
	Compressor          *int64 `json:"1A01"`
	PumpHeatCircuit     *int64 `json:"1A05"`
	PumpRadiator        *int64 `json:"1A06"`
	SwitchValve1        *int64 `json:"1A07"`
	SwitchValve2        *int64 `json:"1A08"`
	RoomSensorInfluence *int64 `json:"2204"`
	PowerConsumption    *int64 `json:"CFAA"`

	// LW
	WarmWaterMid    *int64 `json:"000A"`
	PumpColdCircuit *int64 `json:"1A04"`

	// AW
	AirIntake *int64 `json:"000E"`
	Fan       *int64 `json:"1A09"`

	// Styr2002
	SuctionGas        *int64 `json:"000C"`
	LiquidFlow        *int64 `json:"000D"`
	RadiatorForward2  *int64 `json:"0020"`
	RadiatorReturn2   *int64 `json:"0022"`
	AddHeatStep1      *int64 `json:"1A02"`
	AddHeatStep2      *int64 `json:"1A03"`
	OperatingMode     *int64 `json:"2201"`
	AlarmReset        *int64 `json:"22F2"`
	LoadL1            *int64 `json:"4101"`
	LoadL2            *int64 `json:"4102"`
	LoadL3            *int64 `json:"4103"`
	HeatSet3Parallel  *int64 `json:"2207"`
	HeatSet1CurveL2   *int64 `json:"2222"`
	HeatSet3Parall2   *int64 `json:"2224"`
	DegreeMinIntegral *int64 `json:"8105"`
	CompressorStarts  *int64 `json:"BC61"`
	CompressorRuntime *int64 `json:"6C60"`
	AuxRuntime        *int64 `json:"6C63"`
	WarmWaterRuntime  *int64 `json:"6C64"`

	// Rego
	HeatCarrierReturn *int64 `json:"0003"`
	HeatSet1CurveR    *int64 `json:"0206"`
	HighPressostat    *int64 `json:"1A0A"`
	LowPressostat     *int64 `json:"1A0B"`
	HeatingCable      *int64 `json:"1A0C"`
	CrankCaseHeater   *int64 `json:"1A0D"`
	AddHeatStatus     *int64 `json:"3104"`
	OutdoorTempOffset *int64 `json:"0217"`

	// Conflicting, these require special handling
	Alarm              *int64 `json:"2A20"`
	RegoAlarm          *int64 `json:"1A20"`
	HeatSet1CurveL     *int64 `json:"2205"`
	RegoHeatSet1CurveL *int64 `json:"0205"`
}
