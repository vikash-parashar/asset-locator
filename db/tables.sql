-- Create the Device Location Detail table
CREATE TABLE IF NOT EXISTS device_location (
    id SERIAL PRIMARY KEY, -- Add an 'id' field as primary key
    serial_number VARCHAR(255) NOT NULL, -- Keep 'serial_number' as unique
    device_make_model VARCHAR(255),
    model VARCHAR(255),
    device_type VARCHAR(255),
    data_center VARCHAR(255),
    region VARCHAR(255),
    dc_location VARCHAR(255),
    device_location VARCHAR(255),
    device_row_number INT,
    device_rack_number INT,
    device_ru_number VARCHAR(255)
);

-- Create the Device AMC/Owner Detail table
CREATE TABLE IF NOT EXISTS device_amc_owner (
    id SERIAL PRIMARY KEY, -- Add an 'id' field as primary key
    serial_number VARCHAR(255) NOT NULL, -- Keep 'serial_number' as unique
    device_make_model VARCHAR(255),
    model VARCHAR(255),
    po_number VARCHAR(255),
    po_order_date DATE,
    eosl_date DATE,
    amc_start_date DATE,
    amc_end_date DATE,
    device_owner VARCHAR(255)
);

-- Create the Device Power Detail table
CREATE TABLE IF NOT EXISTS device_power (
    id SERIAL PRIMARY KEY, -- Add an 'id' field as primary key
    serial_number VARCHAR(255) NOT NULL, -- Keep 'serial_number' as unique
    device_make_model VARCHAR(255),
    model VARCHAR(255),
    device_type VARCHAR(255),
    total_power_watt INT,
    total_btu DECIMAL(10, 2),
    total_power_cable INT,
    power_socket_type VARCHAR(255)
);

-- Create the Device Ethernet/Fiber Cable Detail table
CREATE TABLE IF NOT EXISTS device_ethernet_fiber (
    id SERIAL PRIMARY KEY, -- Add an 'id' field as primary key
    serial_number VARCHAR(255) NOT NULL, -- Keep 'serial_number' as unique
    device_make_model VARCHAR(255),
    model VARCHAR(255),
    device_type VARCHAR(255),
    device_physical_port VARCHAR(255),
    device_port_type VARCHAR(255),
    device_port_mac_address_wwn VARCHAR(255),
    connected_device_port VARCHAR(255)
);
