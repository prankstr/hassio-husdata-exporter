## Overview
Husdata Prometheus Exporter is a Home Assistant add-on designed to integrate with Husdata gateways. It polls and exports data from Husdata gateways(H60/H66), providing the Husdata heatpump metrics in a format compatible with Prometheus enhancing your ability to monitor and analyze the performance of your heat pump.

This has been developed towards a H66 connected and a Nibe heatpump with the Styr2002 interface but a lot of the metrics should work for other heat pumps as well, I'm not sure what works and what doesn't though.

## Features
- Periodic polling of data from Husdata systems.
- Metrics exposure in Prometheus format for easy integration and monitoring.
- Customizable polling intervals.
- Supports Swedish and English label values for sensors etc.
- Optional external server functionality for direct add-on access.

See the add-on documentation for more info!