CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "public"."types" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("slug")
);

CREATE TABLE IF NOT EXISTS "public"."damage_class" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO damage_class (value)
  VALUES ('physical'), ('special'), ('status')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."moves" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "name" text NOT NULL,
  "accuracy" integer NOT NULL,
  "pp" integer NOT NULL,
  "power" integer NOT NULL,
  "damage_class_enum" text,
  "effect" text NOT NULL,
  "effect_chance" integer NOT NULL,
  "target" text NOT NULL,
  "type_id" uuid,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("damage_class_enum") REFERENCES "public"."damage_class" ("value") ON UPDATE CASCADE ON DELETE SET NULL,
  FOREIGN KEY ("type_id") REFERENCES "public"."types" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  UNIQUE ("slug")
);

CREATE TABLE IF NOT EXISTS "public"."pokemon" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "pokedex_id" integer NOT NULL,
  "slug" text NOT NULL,
  "name" text NOT NULL,
  "sprite" text NOT NULL,
  "hp" integer NOT NULL DEFAULT 0,
  "attack" integer NOT NULL DEFAULT 0,
  "defense" integer NOT NULL DEFAULT 0,
  "special_attack" integer NOT NULL DEFAULT 0,
  "special_defense" integer NOT NULL DEFAULT 0,
  "speed" integer NOT NULL DEFAULT 0,
  "is_baby" boolean NOT NULL DEFAULT FALSE,
  "is_legendary" boolean NOT NULL DEFAULT FALSE,
  "is_mythical" boolean NOT NULL DEFAULT FALSE,
  "description" text NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("slug")
);

CREATE TABLE IF NOT EXISTS "public"."pokemon_move" (
  "pokemon_id" uuid NOT NULL,
  "move_id" uuid NOT NULL,
  PRIMARY KEY ("pokemon_id", "move_id"),
  FOREIGN KEY ("move_id") REFERENCES "public"."moves" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "public"."pokemon_type" (
  "pokemon_id" uuid NOT NULL,
  "type_id" uuid NOT NULL,
  PRIMARY KEY ("pokemon_id", "type_id"),
  FOREIGN KEY ("pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("type_id") REFERENCES "public"."types" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "public"."stats" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("slug")
);

CREATE TABLE IF NOT EXISTS "public"."abilities" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "name" text NOT NULL,
  "effect" text NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("slug")
);

CREATE TABLE IF NOT EXISTS "public"."pokemon_ability" (
  "pokemon_id" uuid NOT NULL,
  "ability_id" uuid NOT NULL,
  PRIMARY KEY ("pokemon_id", "ability_id"),
  FOREIGN KEY ("pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("ability_id") REFERENCES "public"."abilities" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "public"."teams" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "public"."team_members" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "order" integer NOT NULL,
  "pokemon_id" uuid NOT NULL,
  "team_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("team_id") REFERENCES "public"."teams" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "public"."team_member_move" (
  "team_member_id" uuid NOT NULL,
  "move_id" uuid NOT NULL,
  "order" integer NOT NULL,
  PRIMARY KEY ("team_member_id", "move_id"),
  FOREIGN KEY ("move_id") REFERENCES "public"."moves" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("team_member_id") REFERENCES "public"."team_members" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

