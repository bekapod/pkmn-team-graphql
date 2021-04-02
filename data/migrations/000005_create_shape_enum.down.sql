ALTER TABLE "public"."pokemon"
  DROP CONSTRAINT IF EXISTS pokemon_shape_enum_fkey;

ALTER TABLE "public"."pokemon"
  DROP COLUMN IF EXISTS shape_enum;

DROP TABLE IF EXISTS "public"."shape";

