#!/usr/bin/with-contenv bashio

# Create main config
bashio::log.info "Setting up configuration"
export HUSDATA_HOSTNAME=$(bashio::config 'husdata_hostname')
export LANG=$(bashio::config 'language')
export LOG_LEVEL=$(bashio::config 'log_level')
export POLL_INTERVAL=$(bashio::config 'poll_interval')
export EXTERNAL_SERVER=$(bashio::config 'external_server')
export PORT=$(bashio::addon.port 9101)


# Start application
bashio::log.info "Starting husdata_exporter"
./husdata_exporter