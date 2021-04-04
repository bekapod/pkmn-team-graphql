ALTER TABLE "public"."pokemon_ability"
  ADD COLUMN IF NOT EXISTS slot integer NOT NULL DEFAULT 0,
  ADD COLUMN IF NOT EXISTS is_hidden bool NOT NULL DEFAULT FALSE,
  ADD COLUMN IF NOT EXISTS updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

