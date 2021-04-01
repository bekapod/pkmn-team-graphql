ALTER TABLE "public"."pokemon_type"
  ADD COLUMN IF NOT EXISTS slot integer NOT NULL DEFAULT 0;

