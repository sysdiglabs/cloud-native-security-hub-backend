CREATE TABLE IF NOT EXISTS latest_resources(
   id serial PRIMARY KEY,
   raw jsonb
);

CREATE INDEX IF NOT EXISTS latest_resources_indexginp ON latest_resources USING GIN (raw jsonb_path_ops);
CREATE UNIQUE INDEX IF NOT EXISTS latest_resources_type_name_unique ON latest_resources(
  (raw ->> 'kind'),
  (raw ->> 'app'),
  (raw ->> 'appVersion'));
