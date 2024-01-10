package metrics

import (
	"fmt"
	"husdata_exporter/localization"
	"husdata_exporter/models"

	"github.com/VictoriaMetrics/metrics"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Constants
const divisionFactor = 10

// Update is called periodically to update metrics with new values
func Update(vpdata *models.VPData, lang language.Tag) {
	p := message.NewPrinter(lang)

	// Common Metrics
	createGaugeIfNotNil(vpdata.RadiatorReturn, "sensor_temperature_celsius", "sensor", "Radiator Return", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatCarrierForward, "sensor_temperature_celsius", "sensor", "Heat Carrier Forward", lang, p, true)
	createGaugeIfNotNil(vpdata.BrineInCondensor, "sensor_temperature_celsius", "sensor", "Brine In / Condensor", lang, p, true)
	createGaugeIfNotNil(vpdata.BrineOutEvaporator, "sensor_temperature_celsius", "sensor", "Brine Out / Evaporator", lang, p, true)
	createGaugeIfNotNil(vpdata.Outdoor, "sensor_temperature_celsius", "sensor", "Outdoor", lang, p, true)
	createGaugeIfNotNil(vpdata.Indoor, "sensor_temperature_celsius", "sensor", "Indoor", lang, p, true)
	createGaugeIfNotNil(vpdata.WarmWaterTop, "sensor_temperature_celsius", "sensor", "Hot Water Top", lang, p, true)
	createGaugeIfNotNil(vpdata.HotGas, "sensor_temperature_celsius", "sensor", "Hot Gas", lang, p, true)
	createGaugeIfNotNil(vpdata.Pool, "sensor_temperature_celsius", "sensor", "Pool", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatingSetpoint, "heating_set_point_celsius", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.RoomTempSetpoint, "settings_room_temperature_celsius", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.WarmWaterStopTemp, "settings_hot_water_stop_temperature_celsius", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.WarmWaterStartTemp, "settings_hot_water_start_temperature_celsius", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.Compressor, "unit_on", "unit", "Compressor", lang, p, false)
	createGaugeIfNotNil(vpdata.PumpHeatCircuit, "unit_on", "unit", "Pump Heat Circuit", lang, p, false)
	createGaugeIfNotNil(vpdata.PumpRadiator, "unit_on", "unit", "Pump Radiator", lang, p, false)
	createGaugeIfNotNil(vpdata.SwitchValve1, "unit_on", "unit", "Switch Valve 1", lang, p, false)
	createGaugeIfNotNil(vpdata.SwitchValve2, "unit_on", "unit", "Switch Valve 2", lang, p, false)
	createGaugeIfNotNil(vpdata.RoomSensorInfluence, "settings_room_compensation", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.PowerConsumption, "sensor_power_watt", "", "", lang, p, false)

	// LW Specific Metrics
	createGaugeIfNotNil(vpdata.WarmWaterMid, "sensor_temperature_celsius", "sensor", "Hot Water Mid", lang, p, true)
	createGaugeIfNotNil(vpdata.PumpColdCircuit, "unit_on", "unit", "Pump Cold Circuit", lang, p, false)

	// AW Specific Metrics
	createGaugeIfNotNil(vpdata.AirIntake, "sensor_temperature_celsius", "sensor", "Air Intake", lang, p, true)
	createGaugeIfNotNil(vpdata.Fan, "unit_on", "unit", "Fan", lang, p, false)

	// Styr2002 Specific Metrics
	createGaugeIfNotNil(vpdata.SuctionGas, "sensor_temperature_celsius", "sensor", "Suction Gas", lang, p, true)
	createGaugeIfNotNil(vpdata.LiquidFlow, "sensor_temperature_celsius", "sensor", "Liquid Flow", lang, p, true)
	createGaugeIfNotNil(vpdata.RadiatorForward2, "sensor_temperature_celsius", "sensor", "Radiator Forward 2", lang, p, true)
	createGaugeIfNotNil(vpdata.RadiatorReturn2, "sensor_temperature_celsius", "sensor", "Radiator Return 2", lang, p, true)
	createGaugeIfNotNil(vpdata.AddHeatStep1, "unit_on", "unit", "Add Heat Step 1", lang, p, false)
	createGaugeIfNotNil(vpdata.AddHeatStep2, "unit_on", "unit", "Add Heat Step 2", lang, p, false)
	createGaugeIfNotNil(vpdata.OperatingMode, "settings_mode", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.AlarmReset, "settings_alarm_reset", "", "", lang, p, false)
	createGaugeIfNotNil(vpdata.LoadL1, "sensor_current_ampere", "sensor", "L1", lang, p, true)
	createGaugeIfNotNil(vpdata.LoadL2, "sensor_current_ampere", "sensor", "L2", lang, p, true)
	createGaugeIfNotNil(vpdata.LoadL3, "sensor_current_ampere", "sensor", "L3", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet3Parallel, "settings_curve_offset", "supply", "Heat Set 3 Parallel", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet1CurveL2, "settings_curve_slope", "supply", "Heat Set 1 CurveL 2", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet3Parall2, "settings_curve_offset", "supply", "Heat Set 3 Parallel 2", lang, p, true)
	createGaugeIfNotNil(vpdata.DegreeMinIntegral, "degree_minutes_celsius", "", "", lang, p, true)

	// Handling Counters
	createCounterIfNotNil(vpdata.CompressorStarts, "unit_starts_total", "unit", "Compressor", lang, p, true)
	createCounterIfNotNil(vpdata.CompressorRuntime, "unit_runtime_seconds_total", "unit", "Compressor", lang, p, true)
	createCounterIfNotNil(vpdata.AuxRuntime, "unit_runtime_seconds_total", "unit", "Add Heat", lang, p, true)
	createCounterIfNotNil(vpdata.WarmWaterRuntime, "unit_runtime_seconds_total", "unit", "Hot Water", lang, p, true)

	// Rego Specific Metrics
	createGaugeIfNotNil(vpdata.HeatCarrierReturn, "sensor_temperature_celsius", "sensor", "Heat Carrier Return", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet1CurveR, "settings_curve_slope", "supply", "Heat Set 1 CurveR", lang, p, true)
	createGaugeIfNotNil(vpdata.HighPressostat, "unit_on", "unit", "High Pressostat", lang, p, false)
	createGaugeIfNotNil(vpdata.LowPressostat, "unit_on", "unit", "Low Pressostat", lang, p, false)
	createGaugeIfNotNil(vpdata.HeatingCable, "unit_on", "unit", "Heating Cable", lang, p, false)
	createGaugeIfNotNil(vpdata.CrankCaseHeater, "unit_on", "unit", "Crank Case Heater", lang, p, false)
	createGaugeIfNotNil(vpdata.AddHeatStatus, "sensor_add_heat_percent", "", "", lang, p, true)
	createCounterIfNotNil(vpdata.OutdoorTempOffset, "settings_outdoor_temperature_offset_celsius", "", "", lang, p, true)

	// Conflicting Metrics Handling
	if vpdata.Alarm != nil {
		createGaugeIfNotNil(vpdata.Alarm, "alarm", "", "", lang, p, false)
	} else if vpdata.RegoAlarm != nil {
		createGaugeIfNotNil(vpdata.RegoAlarm, "alarm", "", "", lang, p, false)
	}

	if vpdata.HeatSet1CurveL != nil {
		createGaugeIfNotNil(vpdata.HeatSet1CurveL, "settings_curve_slope", "supply", "Heat Set 1 CurveL", lang, p, true)
	} else if vpdata.RegoHeatSet1CurveL != nil {
		createGaugeIfNotNil(vpdata.RegoHeatSet1CurveL, "settings_curve_slope", "supply", "Heat Set 1 CurveL", lang, p, true)
	}

}

// createMetricName creates a formatted metric name based on provided parameters
func createMetricName(metricName, labelKey, labelValue string, lang language.Tag, printer *message.Printer) string {
	if labelKey != "" && labelValue != "" {
		var translatedLabelValue string
		if lang != language.English {
			translatedLabelValue = localization.GetTranslation(lang, labelValue)
		} else {
			translatedLabelValue = labelValue
		}
		return fmt.Sprintf(`heatpump_%s{%s="%s"}`, metricName, labelKey, translatedLabelValue)
	}
	return fmt.Sprintf(`heatpump_%s`, metricName)
}

// createGaugeIfNotNil creates a gauge if the metric was provided from the husdata gateway
func createGaugeIfNotNil(metricValue *int64, metricName, labelKey, labelValue string, lang language.Tag, printer *message.Printer, applyDivision bool) {
	if metricValue != nil {
		fullMetricName := createMetricName(metricName, labelKey, labelValue, lang, printer)

		metrics.GetOrCreateGauge(printer.Sprintf(fullMetricName), func() float64 {
			if applyDivision {
				return float64(*metricValue) / divisionFactor
			}
			return float64(*metricValue)
		})
	}
}

// createCounterIfNotNil creates a counter if the metric was provided from the husdata gateway
func createCounterIfNotNil(metricValue *int64, metricName, labelKey, labelValue string, lang language.Tag, printer *message.Printer, applyDivision bool) {
	if metricValue != nil {
		fullMetricName := createMetricName(metricName, labelKey, labelValue, lang, printer)

		if applyDivision {
			metrics.GetOrCreateCounter(printer.Sprintf(fullMetricName)).Set(uint64(*metricValue / divisionFactor))
		} else {
			metrics.GetOrCreateCounter(printer.Sprintf(fullMetricName)).Set(uint64(*metricValue))
		}
	}
}
