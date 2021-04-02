CREATE TABLE IF NOT EXISTS "public"."color" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO color (value)
  VALUES ('black'), ('blue'), ('brown'), ('gray'), ('green'), ('pink'), ('purple'), ('red'), ('white'), ('yellow')
ON CONFLICT (value)
  DO NOTHING;

ALTER TABLE "public"."pokemon"
  ADD COLUMN IF NOT EXISTS color_enum text;

ALTER TABLE "public"."pokemon"
  ADD CONSTRAINT pokemon_color_enum_fkey FOREIGN KEY ("color_enum") REFERENCES "public"."color" ("value") ON UPDATE CASCADE ON DELETE SET NULL;

