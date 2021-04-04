ALTER TABLE "public"."abilities"
  DROP COLUMN IF EXISTS updated_at;

ALTER TABLE "public"."egg_groups"
  DROP COLUMN IF EXISTS updated_at;

ALTER TABLE "public"."moves"
  DROP COLUMN IF EXISTS updated_at;

ALTER TABLE "public"."pokemon"
  DROP COLUMN IF EXISTS updated_at;

ALTER TABLE "public"."pokemon_type"
  DROP COLUMN IF EXISTS updated_at;

ALTER TABLE "public"."stats"
  DROP COLUMN IF EXISTS updated_at;

ALTER TABLE "public"."types"
  DROP COLUMN IF EXISTS updated_at;

