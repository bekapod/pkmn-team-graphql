CREATE INDEX IF NOT EXISTS ability_idx ON abilities (id);

CREATE INDEX IF NOT EXISTS egg_group_idx ON egg_groups (id);

CREATE INDEX IF NOT EXISTS move_idx ON moves (id);

CREATE INDEX IF NOT EXISTS move_type_idx ON moves (type_id);

CREATE INDEX IF NOT EXISTS pokemon_idx ON pokemon (id);

CREATE INDEX IF NOT EXISTS type_idx ON types (id);

CREATE INDEX IF NOT EXISTS pokemon_ability_ability_idx ON pokemon_ability (ability_id);

CREATE INDEX IF NOT EXISTS pokemon_ability_pokemon_idx ON pokemon_ability (pokemon_id);

CREATE INDEX IF NOT EXISTS pokemon_egg_group_egg_group_idx ON pokemon_egg_group (egg_group_id);

CREATE INDEX IF NOT EXISTS pokemon_egg_group_pokemon_idx ON pokemon_egg_group (pokemon_id);

CREATE INDEX IF NOT EXISTS pokemon_move_move_idx ON pokemon_move (move_id);

CREATE INDEX IF NOT EXISTS pokemon_move_pokemon_idx ON pokemon_move (pokemon_id);

CREATE INDEX IF NOT EXISTS pokemon_type_type_idx ON pokemon_type (type_id);

CREATE INDEX IF NOT EXISTS pokemon_type_pokemon_idx ON pokemon_type (pokemon_id);

CREATE INDEX IF NOT EXISTS type_damage_relations_type_idx ON type_damage_relations (type_id);

CREATE INDEX IF NOT EXISTS type_damage_relations_related_type_idx ON type_damage_relations (related_type_id);

