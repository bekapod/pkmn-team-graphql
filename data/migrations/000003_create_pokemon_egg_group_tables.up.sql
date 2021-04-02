CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "public"."egg_groups" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("slug")
);

CREATE TABLE IF NOT EXISTS "public"."pokemon_egg_group" (
  "pokemon_id" uuid NOT NULL,
  "egg_group_id" uuid NOT NULL,
  PRIMARY KEY ("pokemon_id", "egg_group_id"),
  FOREIGN KEY ("egg_group_id") REFERENCES "public"."egg_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

