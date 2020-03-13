CREATE TABLE IF NOT EXISTS apps(
   id serial PRIMARY KEY,
   raw jsonb
);

CREATE INDEX IF NOT EXISTS apps_indexginp ON apps USING GIN (raw jsonb_path_ops);
CREATE UNIQUE INDEX IF NOT EXISTS apps_name_unique ON apps(
  (raw ->> 'id'));
