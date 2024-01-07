package models

type VPData struct {
	OperatingMode       int64 `json:"2201"`
	RoomSensorInfluence int64 `json:"2204"`
	HeatSet1CurveL      int64 `json:"2205"`
	HeatSet3Parallel    int64 `json:"2207"`
	HeatSet1CurveL2     int64 `json:"2222"`
	HeatSet3Parall2     int64 `json:"2224"`
	AddHeatStatus       int64 `json:"3104"`
	LoadL1              int64 `json:"4101"`
	LoadL2              int64 `json:"4102"`
	LoadL3              int64 `json:"4103"`
	DegreeMinIntegral   int64 `json:"8105"`
	RadiatorReturn      int64 `json:"0001"`
	HeatCarrierForward  int64 `json:"0004"`
	BrineInCondensor    int64 `json:"0005"`
	BrineOutEvaporator  int64 `json:"0006"`
	Outdoor             int64 `json:"0007"`
	Indoor              int64 `json:"0008"`
	WarmWaterTop        int64 `json:"0009"`
	WarmWaterMid        int64 `json:"000A"`
	HotGas              int64 `json:"000B"`
	SuctionGas          int64 `json:"000C"`
	LiquidFlow          int64 `json:"000D"`
	Pool                int64 `json:"0011"`
	RadiatorForward2    int64 `json:"0020"`
	RadiatorReturn2     int64 `json:"0022"`
	RoomTempSetpoint    int64 `json:"0203"`
	WarmWaterStartTemp  int64 `json:"0212"`
	WarmWaterStopTemp   int64 `json:"0208"`
	AlarmReset          int64 `json:"22F2"`
	HeatingSetpoint     int64 `json:"0107"`
	Compressor          int64 `json:"1A01"`
	AddHeatStep1        int64 `json:"1A02"`
	AddHeatStep2        int64 `json:"1A03"`
	PumpColdCircuit     int64 `json:"1A04"`
	PumpHeatCircuit     int64 `json:"1A05"`
	PumpRadiator        int64 `json:"1A06"`
	SwitchValve1        int64 `json:"1A07"`
	Alarm               int64 `json:"2A20"`
	CompressorStarts    int64 `json:"BC61"`
	CompressorRuntime   int64 `json:"6C60"`
	AuxRuntime          int64 `json:"6C63"`
	WarmWaterRuntime    int64 `json:"6C64"`
	PowerConsumption    int64 `json:"CFAA"`
}
