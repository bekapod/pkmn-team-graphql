ALTER TABLE "public"."pokemon"
  DROP COLUMN IF EXISTS is_default_variant,
  DROP COLUMN IF EXISTS genus,
  DROP COLUMN IF EXISTS height,
  DROP COLUMN IF EXISTS weight;

