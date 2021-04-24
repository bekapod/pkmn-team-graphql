DROP INDEX IF EXISTS region_idx;

DROP INDEX IF EXISTS location_idx;

DROP INDEX IF EXISTS item_idx;

DROP INDEX IF EXISTS pokemon_evolution_from_pokemon_idx;

DROP INDEX IF EXISTS pokemon_evolution_to_pokemon_idx;

DROP INDEX IF EXISTS pokemon_evolution_trigger_idx;

DROP INDEX IF EXISTS pokemon_evolution_item_idx;

DROP INDEX IF EXISTS pokemon_evolution_held_item_idx;

DROP INDEX IF EXISTS pokemon_evolution_known_move_idx;

DROP INDEX IF EXISTS pokemon_evolution_location_idx;

DROP TABLE IF EXISTS "public"."pokemon_evolutions";

DROP TABLE IF EXISTS "public"."items";

DROP TABLE IF EXISTS "public"."locations";

DROP TABLE IF EXISTS "public"."regions";

DROP TABLE IF EXISTS "public"."fling_effect";

DROP TABLE IF EXISTS "public"."item_category";

DROP TABLE IF EXISTS "public"."item_attribute";

DROP TABLE IF EXISTS "public"."gender";

DROP TABLE IF EXISTS "public"."time_of_day";

DROP TABLE IF EXISTS "public"."evolution_trigger";

