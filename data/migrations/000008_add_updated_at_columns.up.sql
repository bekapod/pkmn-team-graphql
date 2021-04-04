ALTER TABLE "public"."abilities"
  ADD COLUMN updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

ALTER TABLE "public"."egg_groups"
  ADD COLUMN updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

ALTER TABLE "public"."moves"
  ADD COLUMN updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

ALTER TABLE "public"."pokemon"
  ADD COLUMN updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

ALTER TABLE "public"."pokemon_type"
  ADD COLUMN updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

ALTER TABLE "public"."stats"
  ADD COLUMN updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

ALTER TABLE "public"."types"
  ADD COLUMN updated_at timestamp WITHOUT time zone NOT NULL DEFAULT now();

