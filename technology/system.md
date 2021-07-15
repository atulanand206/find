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

**Item - Base class for all the puzzle components.**

| Field | Type | Information |
| :--- | :--- | ---: |
| id | UUID | unique identifier |
| name | String | title of the asset |
| text | String | long text for the asset |
| specs | ViewSpecs | visibility configuration |
| img | String | reference to location of image |

**Clue extends Item - contains clue's specifications**

| Field | Type | Information |
| :--- | :--- | ---: |
| rating | number | rating of clue for assigning priority |

**Puzzle extends Item - contains puzzle's specifications**

| Field | Type | Information |
| :--- | :--- | ---: |
| note | String | user added text |
| solution | String | answer of the puzzle |
| clues | Clue\[\] | clues in the puzzle |

**Stage extends Item - contains stage's specifications**

| Field | Type | Information |
| :--- | :--- | ---: |
| puzzles | Puzzle\[\] | puzzles in the stage |

**Theme extends Item - contains theme's specifications**

| Field | Type | Information |
| :--- | :--- | ---: |
| stages | Stage\[\] | stages in the theme |
| solution |  | keyword to complete the scene  |

