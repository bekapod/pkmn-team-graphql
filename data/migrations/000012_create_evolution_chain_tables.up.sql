CREATE TABLE IF NOT EXISTS "public"."evolution_trigger" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO evolution_trigger (value)
  VALUES ('level-up'), ('trade'), ('use-item'), ('shed'), ('other')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."time_of_day" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO time_of_day (value)
  VALUES ('day'), ('night'), ('any')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."gender" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO gender (value)
  VALUES ('male'), ('female'), ('unknown')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."item_attribute" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO item_attribute (value)
  VALUES ('countable'), ('consumable'), ('usable-overworld'), ('usable-in-battle'), ('holdable'), ('holdable-passive'), ('holdable-active'), ('underground')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."item_category" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO item_category (value)
  VALUES ('stat-boosts'), ('effort-drop'), ('medicine'), ('other'), ('in-a-pinch'), ('picky-healing'), ('type-protection'), ('baking-only'), ('collectibles'), ('evolution'), ('spelunking'), ('held-items'), ('choice'), ('effort-training'), ('bad-held-items'), ('training'), ('plates'), ('species-specific'), ('type-enhancement'), ('event-items'), ('gameplay'), ('plot-advancement'), ('unused'), ('loot'), ('all-mail'), ('vitamins'), ('healing'), ('pp-recovery'), ('revival'), ('status-cures'), ('mulch'), ('special-balls'), ('standard-balls'), ('dex-completion'), ('scarves'), ('all-machines'), ('flutes'), ('apricorn-balls'), ('apricorn-box'), ('data-cards'), ('jewels'), ('miracle-shooter'), ('mega-stones'), ('memories'), ('z-crystals')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."fling_effect" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO fling_effect (value)
  VALUES ('badly-poison'), ('burn'), ('berry-effect'), ('herb-effect'), ('paralyze'), ('poison'), ('flinch')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."regions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "name" text NOT NULL,
  updated_at timestamp without time zone NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  UNIQUE ("slug")
);

CREATE TABLE IF NOT EXISTS "public"."locations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "region_id" uuid NOT NULL,
  "name" text NOT NULL,
  updated_at timestamp without time zone NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  UNIQUE ("slug"),
  FOREIGN KEY ("region_id") REFERENCES "public"."regions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "public"."items" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid (),
  "slug" text NOT NULL,
  "name" text NOT NULL,
  "cost" integer NOT NULL,
  "fling_power" integer NOT NULL,
  "fling_effect" text,
  "effect" text NOT NULL,
  "sprite" text NOT NULL,
  "category_enum" text NOT NULL,
  "attribute_enums" text[] NOT NULL,
  updated_at timestamp without time zone NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  UNIQUE ("slug"),
  FOREIGN KEY ("category_enum") REFERENCES "public"."item_category" ("value") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "public"."pokemon_evolutions" (
  "from_pokemon_id" uuid NOT NULL,
  "to_pokemon_id" uuid NOT NULL,
  "trigger_enum" text NOT NULL,
  "item_id" uuid,
  "gender_enum" text,
  "held_item_id" uuid,
  "known_move_id" uuid,
  "known_move_type_id" uuid,
  "location_id" uuid,
  "min_level" integer,
  "min_happiness" integer,
  "min_beauty" integer,
  "min_affection" integer,
  "needs_overworld_rain" boolean NOT NULL DEFAULT FALSE,
  "party_species_pokemon_id" uuid,
  "party_type_id" uuid,
  "relative_physical_stats" integer,
  "time_of_day_enum" text,
  "trade_species_pokemon_id" uuid,
  "turn_upside_down" boolean NOT NULL DEFAULT FALSE,
  "spin" boolean NOT NULL DEFAULT FALSE,
  "take_damage" integer,
  "critical_hits" integer,
  updated_at timestamp without time zone NOT NULL DEFAULT now(),
  PRIMARY KEY ("from_pokemon_id", "to_pokemon_id", "time_of_day_enum", "gender_enum", "trigger_enum"),
  FOREIGN KEY ("from_pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("to_pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("trigger_enum") REFERENCES "public"."evolution_trigger" ("value") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("item_id") REFERENCES "public"."items" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("gender_enum") REFERENCES "public"."gender" ("value") ON UPDATE CASCADE ON DELETE SET NULL,
  FOREIGN KEY ("held_item_id") REFERENCES "public"."items" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("known_move_id") REFERENCES "public"."moves" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("known_move_type_id") REFERENCES "public"."types" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("location_id") REFERENCES "public"."locations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("party_species_pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("party_type_id") REFERENCES "public"."types" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("time_of_day_enum") REFERENCES "public"."time_of_day" ("value") ON UPDATE CASCADE ON DELETE SET NULL,
  FOREIGN KEY ("trade_species_pokemon_id") REFERENCES "public"."pokemon" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS region_idx ON regions (id);

CREATE INDEX IF NOT EXISTS location_idx ON locations (id);

CREATE INDEX IF NOT EXISTS item_idx ON items (id);

CREATE INDEX IF NOT EXISTS pokemon_evolution_from_pokemon_idx ON pokemon_evolutions (from_pokemon_id);

CREATE INDEX IF NOT EXISTS pokemon_evolution_to_pokemon_idx ON pokemon_evolutions (to_pokemon_id);

CREATE INDEX IF NOT EXISTS pokemon_evolution_trigger_idx ON pokemon_evolutions (trigger_enum);

CREATE INDEX IF NOT EXISTS pokemon_evolution_item_idx ON pokemon_evolutions (item_id);

CREATE INDEX IF NOT EXISTS pokemon_evolution_held_item_idx ON pokemon_evolutions (held_item_id);

CREATE INDEX IF NOT EXISTS pokemon_evolution_known_move_idx ON pokemon_evolutions (known_move_id);

CREATE INDEX IF NOT EXISTS pokemon_evolution_location_idx ON pokemon_evolutions (location_id);

