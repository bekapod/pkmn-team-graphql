ALTER TABLE "public"."pokemon"
  DROP CONSTRAINT IF EXISTS pokemon_color_enum_fkey;

ALTER TABLE "public"."pokemon"
  DROP COLUMN IF EXISTS color_enum;

DROP TABLE IF EXISTS "public"."color";

