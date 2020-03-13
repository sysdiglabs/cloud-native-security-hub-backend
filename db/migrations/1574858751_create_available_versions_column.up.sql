ALTER TABLE IF EXISTS resources
  ADD COLUMN IF NOT EXISTS available_versions text[];

ALTER TABLE IF EXISTS latest_resources
  ADD COLUMN IF NOT EXISTS available_versions text[];
