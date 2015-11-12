-- Database init
CREATE USER hss WITH UNENCRYPTED PASSWORD 'hss';
CREATE DATABASE "hss";
GRANT ALL ON DATABASE "hss" TO "hss";

-- Switch to the parking db as the parking user.
\connect "hss";
set role "hss";

-- SensorValue

CREATE TABLE "sensor_value" (
    "sensor_id" text NOT NULL,
    "type" text NOT NULL,
    "time" timestamp with time zone DEFAULT now(),
    "value" DOUBLE PRECISION DEFAULT 0.0
);

CREATE UNIQUE INDEX "sensor_value_unique" ON "sensor_value" ("sensor_id", "time");
CREATE INDEX "sensor_value_type_time_index" ON "sensor_value" ("type", "time");    
CREATE INDEX "sensor_value_time_index" ON "sensor_value" ("time");    

 -- SystemEvent

CREATE TABLE "system_event" (
    "system_id" text NOT NULL,
    "time" timestamp with time zone DEFAULT now(),
    "message" text default ''
);

CREATE UNIQUE INDEX "system_event_unique" ON "system_event" ("system_id", "time");
