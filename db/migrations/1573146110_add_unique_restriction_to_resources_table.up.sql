CREATE UNIQUE INDEX IF NOT EXISTS resources_kind_app_appversion_unique ON resources(
  (raw ->> 'kind'),
  (raw ->> 'app'),
  (raw ->> 'appVersion'),
  (raw ->> 'version'));
