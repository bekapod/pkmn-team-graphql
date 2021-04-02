CREATE TABLE IF NOT EXISTS "public"."shape" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO shape (value)
  VALUES ('ball'), ('squiggle'), ('fish'), ('arms'), ('blob'), ('upright'), ('legs'), ('quadruped'), ('wings'), ('tentacles'), ('heads'), ('humanoid'), ('bug-wings'), ('armor')
ON CONFLICT (value)
  DO NOTHING;

ALTER TABLE "public"."pokemon"
  ADD COLUMN IF NOT EXISTS shape_enum text;

ALTER TABLE "public"."pokemon"
  ADD CONSTRAINT pokemon_shape_enum_fkey FOREIGN KEY ("shape_enum") REFERENCES "public"."shape" ("value") ON UPDATE CASCADE ON DELETE SET NULL;

