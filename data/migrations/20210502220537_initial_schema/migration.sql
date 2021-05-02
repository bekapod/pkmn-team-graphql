-- CreateEnum
CREATE TYPE "Color" AS ENUM ('BLACK', 'BLUE', 'BROWN', 'GRAY', 'GREEN', 'PINK', 'PURPLE', 'RED', 'WHITE', 'YELLOW');

-- CreateEnum
CREATE TYPE "DamageClass" AS ENUM ('PHYSICAL', 'SPECIAL', 'STATUS');

-- CreateEnum
CREATE TYPE "DamageRelation" AS ENUM ('NO_DAMAGE_TO', 'HALF_DAMAGE_TO', 'DOUBLE_DAMAGE_TO', 'NO_DAMAGE_FROM', 'HALF_DAMAGE_FROM', 'DOUBLE_DAMAGE_FROM');

-- CreateEnum
CREATE TYPE "EvolutionTrigger" AS ENUM ('LEVEL_UP', 'OTHER', 'SHED', 'TRADE', 'USE_ITEM');

-- CreateEnum
CREATE TYPE "Gender" AS ENUM ('MALE', 'FEMALE', 'ANY');

-- CreateEnum
CREATE TYPE "Habitat" AS ENUM ('CAVE', 'FOREST', 'GRASSLAND', 'MOUNTAIN', 'RARE', 'ROUGH_TERRAIN', 'SEA', 'URBAN', 'WATERS_EDGE');

-- CreateEnum
CREATE TYPE "ItemCategory" AS ENUM ('ALL_MACHINES', 'ALL_MAIL', 'APRICORN_BALLS', 'APRICORN_BOX', 'BAD_HELD_ITEMS', 'BAKING_ONLY', 'CHOICE', 'COLLECTIBLES', 'DATA_CARDS', 'DEX_COMPLETION', 'EFFORT_DROP', 'EFFORT_TRAINING', 'EVENT_ITEMS', 'EVOLUTION', 'FLUTES', 'GAMEPLAY', 'HEALING', 'HELD_ITEMS', 'IN_A_PINCH', 'JEWELS', 'LOOT', 'MEDICINE', 'MEGA_STONES', 'MEMORIES', 'MIRACLE_SHOOTER', 'MULCH', 'OTHER', 'PICKY_HEALING', 'PLATES', 'PLOT_ADVANCEMENT', 'PP_RECOVERY', 'REVIVAL', 'SCARVES', 'SPECIAL_BALLS', 'SPECIES_SPECIFIC', 'SPELUNKING', 'STANDARD_BALLS', 'STAT_BOOSTS', 'STATUS_CURES', 'TRAINING', 'TYPE_ENHANCEMENT', 'TYPE_PROTECTION', 'UNUSED', 'VITAMINS', 'Z_CRYSTALS');

-- CreateEnum
CREATE TYPE "ItemAttribute" AS ENUM ('CONSUMABLE', 'COUNTABLE', 'HOLDABLE', 'HOLDABLE_ACTIVE', 'HOLDABLE_PASSIVE', 'UNDERGROUND', 'USABLE_IN_BATTLE', 'USABLE_OVERWORLD');

-- CreateEnum
CREATE TYPE "MoveLearnMethod" AS ENUM ('LEVEL_UP', 'EGG', 'TUTOR', 'MACHINE', 'STADIUM_SURFING_PIKACHU', 'LIGHT_BALL_EGG', 'COLOSSEUM_PURIFICATION', 'XD_SHADOW', 'XD_PURIFICATION', 'FORM_CHANGE', 'RECORD', 'TRANSFER');

-- CreateEnum
CREATE TYPE "MoveTarget" AS ENUM ('SPECIFIC_MOVE', 'SELECTED_POKEMON_ME_FIRST', 'ALLY', 'USERS_FIELD', 'USER_OR_ALLY', 'OPPONENTS_FIELD', 'USER', 'RANDOM_OPPONENT', 'ALL_OTHER_POKEMON', 'SELECTED_POKEMON', 'ALL_OPPONENTS', 'ENTIRE_FIELD', 'USER_AND_ALLIES', 'ALL_POKEMON', 'ALL_ALLIES');

-- CreateEnum
CREATE TYPE "Shape" AS ENUM ('BALL', 'SQUIGGLE', 'FISH', 'ARMS', 'BLOB', 'UPRIGHT', 'LEGS', 'QUADRUPED', 'WINGS', 'TENTACLES', 'HEADS', 'HUMANOID', 'BUG_WINGS', 'ARMOR');

-- CreateEnum
CREATE TYPE "TimeOfDay" AS ENUM ('DAY', 'NIGHT', 'ANY');

-- CreateTable
CREATE TABLE "Pokemon" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "pokedexId" INTEGER NOT NULL,
    "slug" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "sprite" TEXT,
    "hp" INTEGER NOT NULL,
    "attack" INTEGER NOT NULL,
    "defense" INTEGER NOT NULL,
    "specialAttack" INTEGER NOT NULL,
    "specialDefense" INTEGER NOT NULL,
    "speed" INTEGER NOT NULL,
    "isBaby" BOOLEAN NOT NULL,
    "isLegendary" BOOLEAN NOT NULL,
    "isMythical" BOOLEAN NOT NULL,
    "description" TEXT,
    "color" "Color" NOT NULL,
    "shape" "Shape" NOT NULL,
    "habitat" "Habitat",
    "isDefaultVariant" BOOLEAN NOT NULL,
    "genus" TEXT NOT NULL,
    "height" INTEGER NOT NULL,
    "weight" INTEGER NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PokemonAbility" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "pokemonId" TEXT NOT NULL,
    "abilityId" TEXT NOT NULL,
    "slot" INTEGER NOT NULL,
    "isHidden" BOOLEAN NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PokemonEvolution" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "fromPokemonId" TEXT NOT NULL,
    "toPokemonId" TEXT NOT NULL,
    "trigger" "EvolutionTrigger" NOT NULL,
    "itemId" TEXT,
    "gender" "Gender" NOT NULL,
    "heldItemId" TEXT,
    "knownMoveId" TEXT,
    "knownMoveTypeId" TEXT,
    "minLevel" INTEGER,
    "minHappiness" INTEGER,
    "minBeauty" INTEGER,
    "minAffection" INTEGER,
    "needsOverworldRain" BOOLEAN NOT NULL,
    "partyPokemonId" TEXT,
    "partyTypeId" TEXT,
    "relativePhysicalStats" INTEGER,
    "timeOfDay" "TimeOfDay" NOT NULL,
    "tradeWithPokemonId" TEXT,
    "turnUpsideDown" BOOLEAN NOT NULL,
    "spin" BOOLEAN NOT NULL,
    "takeDamage" INTEGER,
    "criticalHits" INTEGER,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PokemonMove" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "pokemonId" TEXT NOT NULL,
    "moveId" TEXT NOT NULL,
    "learnMethod" "MoveLearnMethod" NOT NULL,
    "levelLearnedAt" INTEGER NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PokemonType" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "pokemonId" TEXT NOT NULL,
    "typeId" TEXT NOT NULL,
    "slot" INTEGER NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Ability" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "slug" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "effect" TEXT,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "EggGroup" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "slug" TEXT NOT NULL,
    "name" TEXT NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Item" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "slug" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "cost" INTEGER,
    "flingPower" INTEGER,
    "flingEffect" TEXT,
    "effect" TEXT,
    "sprite" TEXT,
    "category" "ItemCategory" NOT NULL,
    "attributes" "ItemAttribute"[],

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Move" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "slug" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "accuracy" INTEGER,
    "pp" INTEGER,
    "power" INTEGER,
    "damageClass" "DamageClass" NOT NULL,
    "effect" TEXT,
    "effectChance" INTEGER,
    "target" "MoveTarget" NOT NULL,
    "typeId" TEXT NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Type" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "slug" TEXT NOT NULL,
    "name" TEXT NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TypeDamageRelation" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "typeAId" TEXT NOT NULL,
    "typeBId" TEXT NOT NULL,
    "damageRelation" "DamageRelation" NOT NULL,

    PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "_EggGroupToPokemon" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "Pokemon.slug_unique" ON "Pokemon"("slug");

-- CreateIndex
CREATE UNIQUE INDEX "PokemonAbility.pokemonId_abilityId_unique" ON "PokemonAbility"("pokemonId", "abilityId");

-- CreateIndex
CREATE UNIQUE INDEX "PokemonEvolution.fromPokemonId_toPokemonId_timeOfDay_gender_trigger_unique" ON "PokemonEvolution"("fromPokemonId", "toPokemonId", "timeOfDay", "gender", "trigger");

-- CreateIndex
CREATE UNIQUE INDEX "PokemonMove.pokemonId_moveId_learnMethod_unique" ON "PokemonMove"("pokemonId", "moveId", "learnMethod");

-- CreateIndex
CREATE UNIQUE INDEX "PokemonType.pokemonId_typeId_unique" ON "PokemonType"("pokemonId", "typeId");

-- CreateIndex
CREATE UNIQUE INDEX "Ability.slug_unique" ON "Ability"("slug");

-- CreateIndex
CREATE UNIQUE INDEX "EggGroup.slug_unique" ON "EggGroup"("slug");

-- CreateIndex
CREATE UNIQUE INDEX "Item.slug_unique" ON "Item"("slug");

-- CreateIndex
CREATE UNIQUE INDEX "Move.slug_unique" ON "Move"("slug");

-- CreateIndex
CREATE UNIQUE INDEX "Type.slug_unique" ON "Type"("slug");

-- CreateIndex
CREATE UNIQUE INDEX "TypeDamageRelation.typeAId_typeBId_damageRelation_unique" ON "TypeDamageRelation"("typeAId", "typeBId", "damageRelation");

-- CreateIndex
CREATE UNIQUE INDEX "_EggGroupToPokemon_AB_unique" ON "_EggGroupToPokemon"("A", "B");

-- CreateIndex
CREATE INDEX "_EggGroupToPokemon_B_index" ON "_EggGroupToPokemon"("B");

-- AddForeignKey
ALTER TABLE "PokemonAbility" ADD FOREIGN KEY ("pokemonId") REFERENCES "Pokemon"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonAbility" ADD FOREIGN KEY ("abilityId") REFERENCES "Ability"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("fromPokemonId") REFERENCES "Pokemon"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("toPokemonId") REFERENCES "Pokemon"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("itemId") REFERENCES "Item"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("heldItemId") REFERENCES "Item"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("knownMoveId") REFERENCES "Move"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("knownMoveTypeId") REFERENCES "Type"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("partyPokemonId") REFERENCES "Pokemon"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("partyTypeId") REFERENCES "Type"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonEvolution" ADD FOREIGN KEY ("tradeWithPokemonId") REFERENCES "Pokemon"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonMove" ADD FOREIGN KEY ("pokemonId") REFERENCES "Pokemon"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonMove" ADD FOREIGN KEY ("moveId") REFERENCES "Move"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonType" ADD FOREIGN KEY ("pokemonId") REFERENCES "Pokemon"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PokemonType" ADD FOREIGN KEY ("typeId") REFERENCES "Type"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Move" ADD FOREIGN KEY ("typeId") REFERENCES "Type"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TypeDamageRelation" ADD FOREIGN KEY ("typeAId") REFERENCES "Type"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TypeDamageRelation" ADD FOREIGN KEY ("typeBId") REFERENCES "Type"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_EggGroupToPokemon" ADD FOREIGN KEY ("A") REFERENCES "EggGroup"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_EggGroupToPokemon" ADD FOREIGN KEY ("B") REFERENCES "Pokemon"("id") ON DELETE CASCADE ON UPDATE CASCADE;
