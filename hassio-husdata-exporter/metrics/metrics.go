package metrics

import (
	"husdata_exporter/localization"
	"husdata_exporter/models"

	"github.com/VictoriaMetrics/metrics"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Update(vpdata *models.VPData, lang language.Tag) {
	p := message.NewPrinter(lang)

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "RadiatorReturn")), func() float64 {
		return float64(vpdata.RadiatorReturn) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "RadiatorReturn2")), func() float64 {
		return float64(vpdata.RadiatorReturn2) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "HeatCarrierForward")), func() float64 {
		return float64(vpdata.HeatCarrierForward) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "BrineInCondensor")), func() float64 {
		return float64(vpdata.BrineInCondensor) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "BrineOutEvaporator")), func() float64 {
		return float64(vpdata.BrineOutEvaporator) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "Outdoor")), func() float64 {
		return float64(vpdata.Outdoor) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "Indoor")), func() float64 {
		return float64(vpdata.Indoor) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "WarmWaterTop")), func() float64 {
		return float64(vpdata.WarmWaterTop) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "WarmWaterMid")), func() float64 {
		return float64(vpdata.WarmWaterMid) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "HotGas")), func() float64 {
		return float64(vpdata.HotGas) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "SuctionGas")), func() float64 {
		return float64(vpdata.SuctionGas) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "LiquidFlow")), func() float64 {
		return float64(vpdata.LiquidFlow) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "Pool")), func() float64 {
		return float64(vpdata.Pool) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_temperature_celsius{sensor="%s"}`, localization.GetTranslation(lang, "RadiatorForward2")), func() float64 {
		return float64(vpdata.RadiatorForward2) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_current_ampere{sensor="%s"}`, localization.GetTranslation(lang, "LoadL1")), func() float64 {
		return float64(vpdata.LoadL1) / 10
	})
	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_current_ampere{sensor="%s"}`, localization.GetTranslation(lang, "LoadL2")), func() float64 {
		return float64(vpdata.LoadL2) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_sensor_current_ampere{sensor="%s"}`, localization.GetTranslation(lang, "LoadL3")), func() float64 {
		return float64(vpdata.LoadL3) / 10
	})

	metrics.GetOrCreateCounter(p.Sprintf(`heatpump_unit_starts_total{unit="%s"}`, localization.GetTranslation(lang, "CompressorStarts"))).Set(uint64(vpdata.CompressorStarts / 10))

	metrics.GetOrCreateCounter(p.Sprintf(`heatpump_unit_runtime_seconds_total{unit="%s"}`, localization.GetTranslation(lang, "CompressorRuntime"))).Set(uint64(vpdata.CompressorRuntime * 3600 / 10))

	metrics.GetOrCreateCounter(p.Sprintf(`heatpump_unit_runtime_seconds_total{unit="%s"}`, localization.GetTranslation(lang, "WarmWaterRuntime"))).Set(uint64(vpdata.WarmWaterRuntime * 3600 / 10))

	metrics.GetOrCreateCounter(p.Sprintf(`heatpump_unit_runtime_seconds_total{unit="%s"}`, localization.GetTranslation(lang, "AuxRuntime"))).Set(uint64(vpdata.AuxRuntime * 3600 / 10))

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_unit_on{unit="%s"}`, localization.GetTranslation(lang, "Compressor")), func() float64 {
		return float64(vpdata.Compressor)
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_unit_on{unit="%s"}`, localization.GetTranslation(lang, "AddHeatStep1")), func() float64 {
		return float64(vpdata.AddHeatStep1)
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_unit_on{unit="%s"}`, localization.GetTranslation(lang, "AddHeatStep2")), func() float64 {
		return float64(vpdata.AddHeatStep2)
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_unit_on{unit="%s"}`, localization.GetTranslation(lang, "PumpHeatCircuit")), func() float64 {
		return float64(vpdata.PumpHeatCircuit)
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_unit_on{unit="%s"}`, localization.GetTranslation(lang, "PumpColdCircuit")), func() float64 {
		return float64(vpdata.PumpColdCircuit)
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_unit_on{unit="%s"}`, localization.GetTranslation(lang, "PumpRadiator")), func() float64 {
		return float64(vpdata.PumpRadiator)
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_unit_on{unit="%s"}`, localization.GetTranslation(lang, "SwitchValve1")), func() float64 {
		return float64(vpdata.SwitchValve1)
	})

	metrics.GetOrCreateGauge(`heatpump_settings_hot_water_start_temperature_celsius`, func() float64 {
		return float64(vpdata.WarmWaterStartTemp) / 10
	})

	metrics.GetOrCreateGauge(`heatpump_settings_hot_water_stop_temperature_celsius`, func() float64 {
		return float64(vpdata.WarmWaterStopTemp) / 10
	})

	metrics.GetOrCreateGauge(`heatpump_settings_room_compensation`, func() float64 {
		return float64(vpdata.RoomSensorInfluence) / 10
	})

	metrics.GetOrCreateGauge(`heatpump_settings_room_temp`, func() float64 {
		return float64(vpdata.RoomTempSetpoint) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_settings_curve_slope{supply="%s"}`, localization.GetTranslation(lang, "HeatSet1CurveL")), func() float64 {
		return float64(vpdata.HeatSet1CurveL) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_settings_curve_slope{supply="%s"}`, localization.GetTranslation(lang, "HeatSet1CurveL2")), func() float64 {
		return float64(vpdata.HeatSet1CurveL2) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_settings_curve_offset{supply="%s"}`, localization.GetTranslation(lang, "HeatSet3Parallel")), func() float64 {
		return float64(vpdata.HeatSet3Parallel) / 10
	})

	metrics.GetOrCreateGauge(p.Sprintf(`heatpump_settings_curve_offset{supply="%s"}`, localization.GetTranslation(lang, "HeatSet3Parall2")), func() float64 {
		return float64(vpdata.HeatSet3Parall2) / 10
	})

	metrics.GetOrCreateGauge(`heatpump_degree_minutes_celsius`, func() float64 {
		return float64(vpdata.DegreeMinIntegral) / 10
	})

	metrics.GetOrCreateGauge(`heatpump_heating_set_point`, func() float64 {
		return float64(vpdata.HeatingSetpoint) / 10
	})

	metrics.GetOrCreateGauge(`heatpump_mode`, func() float64 {
		return float64(vpdata.OperatingMode) / 10
	})

	metrics.GetOrCreateGauge(`heatpump_alarm`, func() float64 {
		return float64(vpdata.Alarm)
	})

	metrics.GetOrCreateGauge(`heatpump_alarm_reset`, func() float64 {
		return float64(vpdata.AlarmReset)
	})

	metrics.GetOrCreateGauge(`heatpump_sensor_power_watt`, func() float64 {
		return float64(vpdata.PowerConsumption)
	})
}
