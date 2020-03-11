CREATE TABLE IF NOT EXISTS resources(
   id serial PRIMARY KEY,
   raw jsonb
);

CREATE INDEX IF NOT EXISTS resources_indexginp ON resources USING GIN (raw jsonb_path_ops);
