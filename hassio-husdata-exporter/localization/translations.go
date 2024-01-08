package localization

import "golang.org/x/text/language"

var (
	SwedishCatalog = map[string]string{
		"OperatingMode":       "Läge",
		"RoomSensorInfluence": "Rumskompensering",
		"HeatSet1CurveL":      "Kurvlutning",
		"HeatSet3Parallel":    "Förskjutning värmekurva",
		"HeatSet1CurveL2":     "Kurvlutning 2",
		"HeatSet3Parall2":     "Förskjutning värmekurva 2",
		"AddHeatStatus":       "Tillskott",
		"LoadL1":              "Ström L1",
		"LoadL2":              "Ström L2",
		"LoadL3":              "Ström L3",
		"DegreeMinIntegral":   "Gradminuter",
		"RadiatorReturn":      "Returledning",
		"HeatCarrierForward":  "Framledning",
		"BrineInCondensor":    "Köldbärare In",
		"BrineOutEvaporator":  "Köldbärare Ut",
		"Outdoor":             "Ute",
		"Indoor":              "Rum",
		"WarmWaterTop":        "Varmvatten Topp",
		"WarmWaterMid":        "Varmvatten Laddning",
		"HotGas":              "Hetgas",
		"SuctionGas":          "Suggas",
		"LiquidFlow":          "Vätskeledning",
		"Pool":                "Pool",
		"RadiatorForward2":    "Framledning 2",
		"RadiatorReturn2":     "Returledning 2",
		"RoomTempSetpoint":    "Rumstemp",
		"WarmWaterStartTemp":  "Varmvatten Start",
		"WarmWaterStopTemp":   "Varmvatten Stop",
		"AlarmReset":          "Alarm Reset",
		"HeatingSetpoint":     "Börvärde",
		"Compressor":          "Kompressor",
		"AddHeatStep1":        "Tillskott I",
		"AddHeatStep2":        "Tillskott II",
		"PumpColdCircuit":     "Köldbärarpump",
		"PumpHeatCircuit":     "Värmebärarpump",
		"PumpRadiator":        "Cirkulationspump",
		"SwitchValve1":        "Varmvattenladdning",
		"Alarm":               "Alarm",
		"CompressorStarts":    "Kompressorstarter",
		"CompressorRuntime":   "Kompressortimmar",
		"AuxRuntime":          "Tillskottstimmar",
		"WarmWaterRuntime":    "Varmvattentimmar",
		"PowerConsumption":    "Effekt",
		"HeatCarrierReturn":   "Returledning intern",
		"HighPressostat":      "Pressostat hög",
		"LowPressostat":       "Pressostat låg",
		"HeatingCable":        "Värmekabel",
		"CrankCaseHeater":     "Vevhusvärmare",
		"AirIntake":           "Luftintag",
		"Fan":                 "Fläkt",
	}

	EnglishCatalog = map[string]string{
		"OperatingMode":       "Operating Mode",
		"RoomSensorInfluence": "Room sensor influence",
		"HeatSet1CurveL":      "Heat set 1, CurveL",
		"HeatSet3Parallel":    "Heat set 3, Parallel",
		"HeatSet1CurveL2":     "Heat set 1, CurveL2",
		"HeatSet3Parall2":     "Heat set 3, Parall2",
		"AddHeatStatus":       "Add heat status",
		"LoadL1":              "Load L1",
		"LoadL2":              "Load L2",
		"LoadL3":              "Load L3",
		"DegreeMinIntegral":   "Degree min / Integral",
		"RadiatorReturn":      "Radiator return",
		"HeatCarrierForward":  "Heat carrier forward",
		"BrineInCondensor":    "Brine in / Condensor",
		"BrineOutEvaporator":  "Brine out / Evaporator",
		"Outdoor":             "Outdoor",
		"Indoor":              "Indoor",
		"WarmWaterTop":        "Hot water top",
		"WarmWaterMid":        "Hot water mid",
		"HotGas":              "Hot gas",
		"SuctionGas":          "Suction gas",
		"LiquidFlow":          "Liquid flow",
		"Pool":                "Pool",
		"RadiatorForward2":    "Radiator forward 2",
		"RadiatorReturn2":     "Radiator return 2",
		"RoomTempSetpoint":    "RoomTempSetpoint",
		"WarmWaterStartTemp":  "Hot water start temp",
		"WarmWaterStopTemp":   "Hot water stop temp",
		"AlarmReset":          "Alarm reset",
		"HeatingSetpoint":     "Heating setpoint",
		"Compressor":          "Compressor",
		"AddHeatStep1":        "Add heat step 1",
		"AddHeatStep2":        "Add heat step 2",
		"PumpColdCircuit":     "Pump cold circuit",
		"PumpHeatCircuit":     "Pump heat circuit",
		"PumpRadiator":        "Pump radiator",
		"SwitchValve1":        "Switch valve 1",
		"Alarm":               "Alarm",
		"CompressorStarts":    "Compressor starts",
		"CompressorRuntime":   "Compressor runtime",
		"AuxRuntime":          "Aux runtime",
		"WarmWaterRuntime":    "Hot water runtime",
		"PowerConsumption":    "Power consumption",
		"HeatCarrierReturn":   "Intern returledning",
		"HighPressostat":      "High pressostat",
		"LowPressostat":       "Low pressostat",
		"HeatingCable":        "Heating cable",
		"CrankCaseHeater":     "Crankcase heater",
		"AirIntake":           "Air intake",
		"Fan":                 "Fan",
	}
)

func GetTranslation(lang language.Tag, key string) string {
	var catalog map[string]string
	switch lang {
	case language.English:
		catalog = EnglishCatalog
	case language.Swedish:
		catalog = SwedishCatalog
	default:
		catalog = SwedishCatalog
	}

	if msg, ok := catalog[key]; ok {
		return msg
	}

	return key
}
