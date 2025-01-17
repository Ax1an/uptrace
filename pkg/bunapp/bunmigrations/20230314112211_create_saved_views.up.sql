CREATE TABLE saved_views (
  id int8 PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,

  project_id int4,

  name varchar(500),
  route varchar(100),
  params jsonb NOT NULL,
  query jsonb NOT NULL,
  pinned boolean NOT NULL,

  created_at timestamptz NOT NULL DEFAULT now()
);

--bun:split

CREATE INDEX saved_views_project_id_idx ON saved_views (project_id);
