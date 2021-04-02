CREATE TABLE IF NOT EXISTS "public"."habitat" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO habitat (value)
  VALUES ('cave'), ('forest'), ('grassland'), ('mountain'), ('rare'), ('rough-terrain'), ('sea'), ('urban'), ('waters-edge')
ON CONFLICT (value)
  DO NOTHING;

ALTER TABLE "public"."pokemon"
  ADD COLUMN IF NOT EXISTS habitat_enum text;

ALTER TABLE "public"."pokemon"
  ADD CONSTRAINT pokemon_habitat_enum_fkey FOREIGN KEY ("habitat_enum") REFERENCES "public"."habitat" ("value") ON UPDATE CASCADE ON DELETE SET NULL;

