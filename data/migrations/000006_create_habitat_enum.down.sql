ALTER TABLE "public"."pokemon"
  DROP CONSTRAINT IF EXISTS pokemon_habitat_enum_fkey;

ALTER TABLE "public"."pokemon"
  DROP COLUMN IF EXISTS habitat_enum;

DROP TABLE IF EXISTS "public"."habitat";

