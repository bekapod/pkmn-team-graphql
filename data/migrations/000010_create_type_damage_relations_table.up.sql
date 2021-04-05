CREATE TABLE IF NOT EXISTS "public"."damage_relation" (
  "value" text NOT NULL,
  PRIMARY KEY ("value")
);

INSERT INTO damage_relation (value)
  VALUES ('no-damage-to'), ('half-damage-to'), ('double-damage-to'), ('no-damage-from'), ('half-damage-from'), ('double-damage-from')
ON CONFLICT (value)
  DO NOTHING;

CREATE TABLE IF NOT EXISTS "public"."type_damage_relations" (
  "type_id" uuid NOT NULL,
  "related_type_id" uuid NOT NULL,
  "damage_relation_enum" text NOT NULL,
  "updated_at" timestamp without time zone NOT NULL DEFAULT now(),
  PRIMARY KEY ("type_id", "related_type_id", "damage_relation_enum"),
  FOREIGN KEY ("type_id") REFERENCES "public"."types" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("related_type_id") REFERENCES "public"."types" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  FOREIGN KEY ("damage_relation_enum") REFERENCES "public"."damage_relation" ("value") ON UPDATE CASCADE ON DELETE CASCADE
);

