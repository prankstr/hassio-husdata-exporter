# Husdata Prometheus Exporter for Home Assistant

## Overview
Husdata Prometheus Exporter is a Home Assistant add-on that integrates with Husdata gateways(H60/H66). It polls and exports data from these Husdata gateways, presenting Husdata heat pump metrics in a format that is fully compatible with Prometheus, enhancing your ability to monitor and analyze your heat pump's performance effectively.

This add-on has been developed primarily for the Husdata H66 and Nibe ground source heat pumps equipped with the Styr2002 interface as well as the Bosch EHP AW pumps with the Rego800 interface. It offers mostly complete metrics coverage for these specific models, as supported by Husdata. While a lot of the metrics should work for other heat pumps as well the compatibility and functionality may vary.

## Features
- Periodic polling of data from Husdata gateways.
- Exports Husdata metrics in a prometheus format for easy integration and monitoring.
- Customizable polling intervals.
- Supports Swedish and English label values for sensors etc(see below).
- Supports Home Assistant Ingress and a toggleable external server for direct access.

## Installation

### Prerequisites
- A running Home Assistant installation with add-on support
- Heatpump with a Husdata H60 or H66 gateway connected and API enabled

### Steps
1. **Add Repository to Home Assistant:**
   [![Add repository to my Home Assistant](https://my.home-assistant.io/badges/supervisor_add_addon_repository.svg)](https://my.home-assistant.io/redirect/supervisor_add_addon_repository/?repository_url=https%3A%2F%2Fgithub.com%2Fprankstr%2Fhassio-husdata-exporter) 

    Or manually:
   - Navigate to the Add-on Store in your Home Assistant UI: `Settings` -> `Add-ons` -> `Add-on Store`.
   - Click the 3-dots in the upper right corner, select `Repositories`, and paste in this URL: [https://github.com/prankstr/hassio-husdata-exporter](https://github.com/prankstr/hassio-husdata-exporter)

2. **Install Husdata Exporter:**
   - Refresh the page
   - Find Husdata Exporter in the list of available add-ons, open it and click 'Install'.

## Configuration
Configure the add-on through the Home Assistant UI with the following options:
- `husdata_hostname` (required): Hostname or IP address of your Husdata gateway.
- `language` (optional): Language for metrics (options: "Swedish", "English"; default: "Swedish").
- `external_server` (optional): Enable/disable direct access to the add-on (default: false).
- `poll_interval` (optional): Time interval (in seconds) between data polls (default: 15).
- `log_level` (optional): Set the logging level (options: "Info", "Debug", "Warn", "Error"; default: "Info").

## Usage
Once the Husdata Exporter add-on is configured and running:
- The exporter will start polling data from your Husdata gateway immediately.
- For scraping within Home Assistant: Configure your monitoring tool to scrape metrics from `http://d5f5b367-husdata-exporter:8099/metrics`.
- For external scraping: Enable the external server in settings and configure your tool to scrape from `http://<your-server-ip-or-hostname>:9101/metrics`.

If the internal hostname isn't working - verify that it's correct by navigating to the add-on page and check the hostname. `Settings` -> `Add-ons` -> `Husdata Exporter`. Use the value underlined in red.

![Add-on Hostname Preview](images/hostname.png "Add-on Hostname")

## Available Metrics

| Register  | Metric Name                                             | Label Key                 | Label Value (English)    | Label Value (Swedish)    | Description                                               |
|-----------|---------------------------------------------------------|---------------------------|--------------------------|--------------------------|-----------------------------------------------------------|
| `0001`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Radiator return          | Returledning             | Radiator return temperature                               |
| `0004`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Heat carrier forward     | Framledning              | Internal heat carrier forward temperature                 |
| `0005`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Brine in / Condensor     | Köldbärare In            | Brine In / Condensor                                      |
| `0006`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Brine out / Evaporator   | Köldbärare Ut            | Brine Out / Evaporator                                    |
| `0003`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Heat carrier return      | Returledning intern      | Internal heat carrier return temperature                  |
| `0007`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Outdoor                  | Ute                      | Oudoor temperature                                        |
| `0008`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Indoor                   | Rum                      | Indoor temperature                                        |
| `0009`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Hot water top            | Varmvatten Topp          | Hot water top temperature                                 |
| `000A`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Hot water mid            | Varmvatten Laddning      | Hot water bottom/mid temperature                          |
| `000B`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Hot Gas                  | Hetgas                   | Hot Gas temperature                                       |
| `000C`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Suction Gas              | Suggas                   | Suction gas temperature                                   |
| `000D`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Liquid Flow              | Vätskeledning            | Liquid flow temperature                                   |
| `000E`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Air Intake               | Luftintag                | Air Intake temperature                                    |
| `0011`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Pool                     | Pool                     | Pool forward temperature                                  |
| `0020`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Radiator forward 2       | Framledning 2            | Internal heat carrier forward temperature for circuit 2   |
| `0022`    | `heatpump_sensor_temperature_celsius`                   | `sensor`                  | Radiator return 2        | Returledning 2           | Radiator return temperature for circuit 2                 |
| `4101`    | `heatpump_sensor_current_ampere`                        | `sensor`                  | L1                       | L1                       | Current draw L1                                           |
| `4102`    | `heatpump_sensor_current_ampere`                        | `sensor`                  | L2                       | L2                       | Current draw L2                                           |
| `4103`    | `heatpump_sensor_current_ampere`                        | `sensor`                  | L3                       | L3                       | Current draw L3                                           |
| `BC61`    | `heatpump_unit_starts_total`                            | `unit`                    | Compressor               | Kompressor               | Total amount of compressor starts                         |
| `6C60`    | `heatpump_unit_runtime_seconds_total`                   | `unit`                    | Compressor               | Kompressor               | Total compressor runtime in seconds                       |
| `6C63`    | `heatpump_unit_runtime_seconds_total`                   | `unit`                    | Add heat                 | Tillskott                | Total amount of added heat in seconds                     |
| `6C64`    | `heatpump_unit_runtime_seconds_total`                   | `unit`                    | Hot water                | Varmvattenladdning       | Total amount of hot water preparation in seconds          |
| `1A01`    | `heatpump_unit_on`                                      | `unit`                    | Compressor               | Kompressor               | Compressor on/of                                          |
| `1A02`    | `heatpump_unit_on`                                      | `unit`                    | Add Heat step 1          | Tillskott Steg I         | Add heat step 1 on/off                                    |
| `1A03`    | `heatpump_unit_on`                                      | `unit`                    | Add Heat step 2          | Tillskott Steg II        | Add heat step 2 on/off                                    |
| `1A04`    | `heatpump_unit_on`                                      | `unit`                    | Pump heat circuit        | Värmebärarpump           | Pump Heat Circuit on/off                                  |
| `1A05`    | `heatpump_unit_on`                                      | `unit`                    | Pump cold circuit        | Köldbärarpump            | Pump Cold Circuit on/off                                  |
| `1A06`    | `heatpump_unit_on`                                      | `unit`                    | Pump radiator            | Cirkulationspump         | Radiator pump on/off                                      |
| `1A07`    | `heatpump_unit_on`                                      | `unit`                    | Switch valve 1           | Växelventil 1            | Switch Valve 1 on/off                                     |
| `1A08`    | `heatpump_unit_on`                                      | `unit`                    | Switch valve 2           | Växelventil 2            | Switch Valve 2 on/off                                     |
| `1A0C`    | `heatpump_unit_on`                                      | `unit`                    | Heating cable            | Värmekabel               | Heating cable on/off                                      |
| `1A0D`    | `heatpump_unit_on`                                      | `unit`                    | Crank case heater        | Vevhusvärmare            | Crank case on/off                                         |
| `1A09`    | `heatpump_unit_on`                                      | `unit`                    | Fan                      | Fläkt                    | Fan on/off                                                |
| `1A0A`    | `heatpump_unit_on`                                      | `unit`                    | High pressostat          | Pressostat hög           | Pressostat High                                           |
| `1A0B`    | `heatpump_unit_on`                                      | `unit`                    | Low pressostat           | Pressostat låg           | Pressostat Low                                            |
| `2205`    | `heatpump_settings_curve_slope`                         | `supply`                  | Heat set 1 curveL        | Kurvlutning              | Slope of the heat curve                                   |
| `2222`    | `heatpump_settings_curve_slope`                         | `supply`                  | Heat set 1 curveL2       | Kurvlutning 2            | Slope of heat curve 2                                     |
| `2207`    | `heatpump_settings_curve_offset`                        | `supply`                  | Heat set 3 parallel      | Förskjutning värmekurva  | Offset of the heat curve                                  |
| `2224`    | `heatpump_settings_curve_offset`                        | `supply`                  | Heat set 3 parall2       | Förskjutning värmekurva 2| Offset for heat curve 2                                   |
| `0212`    | `heatpump_settings_hot_water_start_temperature_celsius` |                           |                          |                          | Temperature where hot water heating starts                |
| `0208`    | `heatpump_settings_hot_water_stop_temperature_celsius`  |                           |                          |                          | Temperature where hot water heating starts stops          |
| `2204`    | `heatpump_settings_room_compensation_celsius`           |                           |                          |                          | Room Sensor Influence                                     |
| `0203`    | `heatpump_settings_room_temperature_celsius`            |                           |                          |                          | Target room temperature                                   |
| `0217`    | `heatpump_settings_outdoor_temperature_offset_celsius`  |                           |                          |                          | Outdoor temperature offset                                |
| `8105`    | `heatpump_degree_minutes_celsius`                       |                           |                          |                          | Degree minutes/Integral/Gradminutes                       |
| `0107`    | `heatpump_heating_set_point_celsius`                    |                           |                          |                          | Heating setpoint/Börvärde                                 |
| `2201`    | `heatpump_settings_mode`                                |                           |                          |                          | Heatpump mode. Summer=0, Spring/Autum=1, Winter=2, Auto=3 |
| `2A20`    | `heatpump_alarm`                                        |                           |                          |                          | If any alarms are active                                  |
| `22F2`    | `heatpump_settings_alarm_reset`                         |                           |                          |                          | If an alarm reset has been triggered                      |
| `CFAA`    | `heatpump_sensor_power_watt`                            |                           |                          |                          | Draw in watt                                              |
| `3104`    | `heatpump_sensor_add_heat_percent`                      |                           |                          |                          | Add heat status percentage                                |

## Dashboard

I've included my `dashboard.json` for Grafana which is pre-configured in Swedish and can serve as a base for creating your custom dashboard to monitor your heatpump.

![Grafana Dashboard Preview](images/dashboard.png "Grafana Dashboard")

## License
This project is licensed under the [MIT License](LICENSE).