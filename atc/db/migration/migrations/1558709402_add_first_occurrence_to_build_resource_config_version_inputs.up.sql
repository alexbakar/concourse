BEGIN;
  ALTER TABLE build_resource_config_version_inputs ADD COLUMN first_occurrence boolean NOT NULL DEFAULT false;
COMMIT;
