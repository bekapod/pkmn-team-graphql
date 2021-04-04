ALTER TABLE "public"."pokemon_ability"
  DROP COLUMN IF EXISTS slot,
  DROP COLUMN IF EXISTS is_hidden,
  DROP COLUMN IF EXISTS updated_at;

