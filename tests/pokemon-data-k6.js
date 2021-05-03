import http from "k6/http";
import { sleep } from "k6";

const allPokemonQuery = `
  {
    pokemon {
      total
      pokemon {
        id
        name
        slug
        pokedexId
        sprite
        color
        shape
        habitat
        hp
        attack
        defense
        specialAttack
        specialDefense
        speed
        height
        weight
        isDefaultVariant
        isBaby
        isLegendary
        isMythical
        description
        abilities {
          total
          abilities {
            id
            slug
            name
            effect
          }
        }
        eggGroups {
          total
          eggGroups {
            id
            name
            slug
          }
        }
        moves {
          total
          moves {
            id
            slug
            name
            accuracy
            pp
            power
            damageClass
            effect
            effectChance
            target
            type {
                id
                name
                slug
            }
          }
        }
        types {
          total
          pokemonTypes {
            slot
            type {
                id
                name
                slug
            }
          }
        }
      }
    }
  }
`;

const headers = {
  "Content-Type": "application/json",
};

export const options = {
  stages: [
    { duration: "30s", target: 100 },
    { duration: "1m", target: 100 },
    { duration: "30s", target: 0 },
  ],
  discardResponseBodies: true,
  noUsageReport: true,
  thresholds: {
    http_req_failed: ["rate<0.01"],
    http_req_duration: ["p(100)<1000", "p(80)<100"],
  },
};

export default function () {
  const res = http.post(
    "http://localhost:5000/graphql",
    JSON.stringify({
      variables: {},
      query: allPokemonQuery,
    }),
    {
      headers,
    }
  );
  sleep(1);
}
