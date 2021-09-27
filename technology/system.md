# System

**Color - stores the color information code**

| Field | Type | Information |
| :--- | :--- | ---: |
| rgba | number\[\] | Color information as RGBA |
| hsla | number\[\] | Color information as HSLA |
| hex | String | Color information as hex code |
| transparent | boolean | Hide the asset when true |

**ViewSpecs - visibility specifications of asset**

| Field | Type | Information |
| :--- | :--- | ---: |
| id | UUID | unique identifier |
| light | Color | color of the asset to be used in light mode |
| dark | Color | color of the asset to be used in dark mode |
| height | number | view height of the asset |
| width | number | view width of the asset |

**Component extends ViewSpecs - Preset stored for replicating asset properties**

| Field | Type | Information |
| :--- | :--- | ---: |
| key | String | key to reference the preset |
| name | String | title of the asset |

**ImageSpecs - specifications of image**

| Field | Type | Information |
| :--- | :--- | ---: |
| ref | String | image url |

**Item - Base class for all the puzzle components.**

| Field | Type | Information |
| :--- | :--- | ---: |
| id | UUID | unique identifier |
| name | String | title of the asset |
| text | String | long text for the asset |
| specs | ViewSpecs | visibility configuration |
| img | ImageSpecs | image information |

**Clue extends Item - contains clue's specifications**

| Field | Type | Information |
| :--- | :--- | ---: |
| rating | number | rating of clue for assigning priority |
| puzzleImgId | UUID | ref to the puzzle image to keep coordinates correct |
| xP | number | percent width from bottom left |
| yP | number | percent height from bottom left |

**Puzzle extends Item - contains puzzle's specifications**

| Field | Type | Information |
| :--- | :--- | ---: |
| note | String | user added text |
| solution | String | answer of the puzzle |
| clues | Clue\[\] | clues in the puzzle |

**Theme extends Item - contains theme's specifications**

| Field | Type | Information |
| :--- | :--- | ---: |
| puzzles | Puzzle\[\] | puzzles in the theme |
| solution |  | keyword to complete the scene |
|  |  |  |
|  |  |  |

**Data** **examples**

```text
v1: 16/07/2021
{
  "id": "e1e1718f-f03a-4429-b1e9-39b2705a542b",
  "name": "theme-name",
  "text": "theme-text",
  "specs": {
    "id": "c2082907-99a2-48b4-b291-2b2ae71a67a0",
    "light": {
      "color": {
        "r": 0.8117647,
        "g": 0.8745098,
        "b": 0.8117647,
        "a": 1
      },
      "transparent": false
    },
    "dark": {
      "color": {
        "r": 0.8117647,
        "g": 0.8745098,
        "b": 0.8117647,
        "a": 1
      },
      "transparent": false
    },
    "height": 100,
    "width": 100,
    "key": "theme"
  },
  "img": {
    "id": "ebcd6958-b29e-4740-88d2-0bf701e03ca5",
    "ref": "theme-image-ref"
  },
  "puzzles": [
    {
      "id": "5bad8836-f7c0-4fe6-8a2c-fbba0c326bcb",
      "name": "puzzle-name",
      "text": "puzzle-text",
      "specs": {
        "id": "d19b8cbc-861e-4220-9d1d-70332cc5e835",
        "light": {
          "color": {
            "r": 0.8117647,
            "g": 0.8745098,
            "b": 0.8117647,
            "a": 1
          },
          "transparent": false
        },
        "dark": {
          "color": {
            "r": 0.8117647,
            "g": 0.8745098,
            "b": 0.8117647,
            "a": 1
          },
          "transparent": false
        },
        "height": 100,
        "width": 100,
        "key": "puzzle"
      },
      "img": {
        "id": "b4126d02-1f66-4e02-aaf5-d4e049dbbbfc",
        "ref": "puzzle-image-ref"
      },
      "note": "puzzle-note",
      "solution": "puzzle-solution",
      "clues": [
        {
          "id": "162fd5b0-6e68-4622-8fe1-c1daef78bee5",
          "name": "clue-name",
          "text": "clue-text",
          "specs": {
            "id": "081f667a-68f2-49ff-a382-3d1d88eb4ebd",
            "light": {
              "color": {
                "r": 0.8117647,
                "g": 0.8745098,
                "b": 0.8117647,
                "a": 1
              },
              "transparent": false
            },
            "dark": {
              "color": {
                "r": 0.8117647,
                "g": 0.8745098,
                "b": 0.8117647,
                "a": 1
              },
              "transparent": false
            },
            "height": 10,
            "width": 10,
            "key": "clue"
          },
          "img": {
            "id": "612b22c0-5811-4999-9c06-62ac3784bfee",
            "ref": "clue-image-ref"
          },
          "rating": 100,
          "puzzleImgId": "b4126d02-1f66-4e02-aaf5-d4e049dbbbfc",
          "xP": 100,
          "yP": 100
        }
      ]
    }
  ]
}
```

