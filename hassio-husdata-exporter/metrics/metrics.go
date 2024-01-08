package metrics

import (
	"fmt"
	"husdata_exporter/localization"
	"husdata_exporter/models"

	"github.com/VictoriaMetrics/metrics"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Update is called periodically to update metrics with new values
func Update(vpdata *models.VPData, lang language.Tag) {
	p := message.NewPrinter(lang)

	// Common Metrics
	createGaugeIfNotNil(vpdata.RadiatorReturn, "sensor_temperature_celsius", "sensor", "RadiatorReturn", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatCarrierForward, "sensor_temperature_celsius", "sensor", "HeatCarrierForward", lang, p, true)
	createGaugeIfNotNil(vpdata.BrineInCondensor, "sensor_temperature_celsius", "sensor", "BrineInCondensor", lang, p, true)
	createGaugeIfNotNil(vpdata.BrineOutEvaporator, "sensor_temperature_celsius", "sensor", "BrineOutEvaporator", lang, p, true)
	createGaugeIfNotNil(vpdata.Outdoor, "sensor_temperature_celsius", "sensor", "Outdoor", lang, p, true)
	createGaugeIfNotNil(vpdata.Indoor, "sensor_temperature_celsius", "sensor", "Indoor", lang, p, true)
	createGaugeIfNotNil(vpdata.WarmWaterTop, "sensor_temperature_celsius", "sensor", "WarmWaterTop", lang, p, true)
	createGaugeIfNotNil(vpdata.HotGas, "sensor_temperature_celsius", "sensor", "HotGas", lang, p, true)
	createGaugeIfNotNil(vpdata.Pool, "sensor_temperature_celsius", "sensor", "Pool", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatingSetpoint, "heating_set_point", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.RoomTempSetpoint, "settings_room_temp", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.WarmWaterStopTemp, "settings_hot_water_stop_temperature_celsius", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.WarmWaterStartTemp, "settings_hot_water_start_temperature_celsius", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.Compressor, "unit_on", "unit", "Compressor", lang, p, false)
	createGaugeIfNotNil(vpdata.PumpHeatCircuit, "unit_on", "unit", "PumpHeatCircuit", lang, p, false)
	createGaugeIfNotNil(vpdata.PumpRadiator, "unit_on", "unit", "PumpRadiator", lang, p, false)
	createGaugeIfNotNil(vpdata.SwitchValve1, "unit_on", "unit", "SwitchValve1", lang, p, false)
	createGaugeIfNotNil(vpdata.SwitchValve2, "unit_on", "unit", "SwitchValve2", lang, p, false)
	createGaugeIfNotNil(vpdata.RoomSensorInfluence, "settings_room_compensation", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.PowerConsumption, "sensor_power_watt", "", "", lang, p, false)

	// LW Specific Metrics
	createGaugeIfNotNil(vpdata.WarmWaterMid, "sensor_temperature_celsius", "sensor", "WarmWaterMid", lang, p, true)
	createGaugeIfNotNil(vpdata.PumpColdCircuit, "unit_on", "unit", "PumpColdCircuit", lang, p, false)

	// AW Specific Metrics
	createGaugeIfNotNil(vpdata.AirIntake, "sensor_temperature_celsius", "sensor", "AirIntake", lang, p, true)
	createGaugeIfNotNil(vpdata.Fan, "unit_on", "unit", "Fan", lang, p, false)

	// Styr2002 Specific Metrics
	createGaugeIfNotNil(vpdata.SuctionGas, "sensor_temperature_celsius", "sensor", "SuctionGas", lang, p, true)
	createGaugeIfNotNil(vpdata.LiquidFlow, "sensor_temperature_celsius", "sensor", "LiquidFlow", lang, p, true)
	createGaugeIfNotNil(vpdata.RadiatorForward2, "sensor_temperature_celsius", "sensor", "RadiatorForward2", lang, p, true)
	createGaugeIfNotNil(vpdata.RadiatorReturn2, "sensor_temperature_celsius", "sensor", "RadiatorReturn2", lang, p, true)
	createGaugeIfNotNil(vpdata.AddHeatStep1, "unit_on", "unit", "AddHeatStep1", lang, p, false)
	createGaugeIfNotNil(vpdata.AddHeatStep2, "unit_on", "unit", "AddHeatStep2", lang, p, false)
	createGaugeIfNotNil(vpdata.OperatingMode, "mode", "", "", lang, p, true)
	createGaugeIfNotNil(vpdata.AlarmReset, "alarm_reset", "", "", lang, p, false)
	createGaugeIfNotNil(vpdata.LoadL1, "sensor_current_ampere", "sensor", "LoadL1", lang, p, true)
	createGaugeIfNotNil(vpdata.LoadL2, "sensor_current_ampere", "sensor", "LoadL2", lang, p, true)
	createGaugeIfNotNil(vpdata.LoadL3, "sensor_current_ampere", "sensor", "LoadL3", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet3Parallel, "settings_curve_offset", "supply", "HeatSet3Parallel", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet1CurveL2, "settings_curve_slope", "supply", "HeatSet1CurveL2", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet3Parall2, "settings_curve_offset", "supply", "HeatSet3Parall2", lang, p, true)
	createGaugeIfNotNil(vpdata.DegreeMinIntegral, "degree_minutes_celsius", "", "", lang, p, true)

	// Handling Counters
	createCounterIfNotNil(vpdata.CompressorStarts, "unit_starts_total", "unit", "CompressorStarts", lang, p, true)
	createCounterIfNotNil(vpdata.CompressorRuntime, "unit_runtime_seconds_total", "unit", "CompressorRuntime", lang, p, true)
	createCounterIfNotNil(vpdata.AuxRuntime, "unit_runtime_seconds_total", "unit", "AuxRuntime", lang, p, true)
	createCounterIfNotNil(vpdata.WarmWaterRuntime, "unit_runtime_seconds_total", "unit", "WarmWaterRuntime", lang, p, true)

	// Rego Specific Metrics
	createGaugeIfNotNil(vpdata.HeatCarrierReturn, "sensor_temperature_celsius", "sensor", "HeatCarrierReturn", lang, p, true)
	createGaugeIfNotNil(vpdata.HeatSet1CurveR, "settings_curve_slope", "supply", "HeatSet1CurveR", lang, p, true)
	createGaugeIfNotNil(vpdata.HighPressostat, "unit_on", "unit", "HighPressostat", lang, p, false)
	createGaugeIfNotNil(vpdata.LowPressostat, "unit_on", "unit", "LowPressostat", lang, p, false)
	createGaugeIfNotNil(vpdata.HeatingCable, "unit_on", "", "HeatingCable", lang, p, false)
	createGaugeIfNotNil(vpdata.CrankCaseHeater, "unit_on", "", "CrankCaseHeater", lang, p, false)
	createGaugeIfNotNil(vpdata.AddHeatStatus, "sensor_add_heat_status", "sensor", "AddHeatStatus", lang, p, true)

	// Conflicting Metrics Handling
	if vpdata.Alarm != nil {
		createGaugeIfNotNil(vpdata.Alarm, "alarm", "", "", lang, p, false)
	} else if vpdata.RegoAlarm != nil {
		createGaugeIfNotNil(vpdata.RegoAlarm, "alarm", "", "", lang, p, false)
	}

	if vpdata.HeatSet1CurveL != nil {
		createGaugeIfNotNil(vpdata.HeatSet1CurveL, "settings_curve_slope", "supply", "HeatSet1CurveL", lang, p, true)
	} else if vpdata.RegoHeatSet1CurveL != nil {
		createGaugeIfNotNil(vpdata.RegoHeatSet1CurveL, "settings_curve_slope", "supply", "HeatSet1CurveL", lang, p, true)
	}

}

// createGaugeIfNotNil creates a guage if the metric was provided from the husdata gateway
func createGaugeIfNotNil(metricValue *int64, metricName string, labelKey string, labelValue string, lang language.Tag, printer *message.Printer, applyDivision bool) {
	if metricValue != nil {
		var fullMetricName string
		if labelKey != "" && labelValue != "" {
			translatedLabelValue := localization.GetTranslation(lang, labelValue)
			fullMetricName = fmt.Sprintf(`heatpump_%s{%s="%s"}`, metricName, labelKey, translatedLabelValue)
		} else {
			fullMetricName = fmt.Sprintf(`heatpump_%s`, metricName)
		}

		metrics.GetOrCreateGauge(printer.Sprintf(fullMetricName), func() float64 {
			if applyDivision {
				return float64(*metricValue) / 10
			}
			return float64(*metricValue)
		})
	}
}

// createGaugeIfNotNil creates a counter if the metric was provided from the husdata gateway
func createCounterIfNotNil(metricValue *int64, metricName string, labelKey string, labelValue string, lang language.Tag, printer *message.Printer, applyDivision bool) {
	if metricValue != nil {
		var fullMetricName string
		if labelKey != "" && labelValue != "" {
			translatedLabelValue := localization.GetTranslation(lang, labelValue)
			fullMetricName = fmt.Sprintf(`heatpump_%s{%s="%s"}`, metricName, labelKey, translatedLabelValue)
		} else {
			fullMetricName = fmt.Sprintf(`heatpump_%s`, metricName)
		}

		if applyDivision {
			metrics.GetOrCreateCounter(printer.Sprintf(fullMetricName)).Set(uint64(*metricValue / 10))
		} else {
			metrics.GetOrCreateCounter(printer.Sprintf(fullMetricName)).Set(uint64(*metricValue))
		}
	}
}
