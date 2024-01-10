## Overview
Husdata Prometheus Exporter is a Home Assistant add-on that integrates with Husdata gateways(H60/H66). It polls and exports data from these Husdata gateways, presenting Husdata heat pump metrics in a format that is fully compatible with Prometheus, enhancing your ability to monitor and analyze your heat pump's performance effectively.

This add-on has been developed primarily for the Husdata H66 and Nibe ground source heat pumps equipped with the Styr2002 interface as well as the Bosch EHP AW pumps with the Rego800 interface. It offers mostly complete metrics coverage for these specific models, as supported by Husdata. While a lot of the metrics should work for other heat pumps as well the compatibility and functionality may vary.

## Features
- Periodic polling of data from Husdata gateways.
- Exports Husdata metrics in a prometheus format for easy integration and monitoring.
- Customizable polling intervals.
- Supports Swedish and English label values for sensors etc(see below).
- Supports Home Assistant Ingress and a toggleable external server for direct access.

See the add-on documentation for more info!