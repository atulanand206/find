# System

**Color**

| Field | Type | Information |
| :--- | :--- | ---: |
| rgba | number\[\] | Color information as RGBA |
| hsla | number\[\] | Color information as HSLA |
| hex | String | Color information as hex code |
| transparent | boolean | Hide the asset when true |

**ViewSpecs**

| Field | Type | Information |
| :--- | :--- | ---: |
| id | UUID | unique identifier |
| light | Color | color of the asset to be used in light mode |
| dark | Color | color of the asset to be used in dark mode |
| height | number | view height of the asset |
| width | number | view width of the asset |

**Item**

| Field | Type | Information |
| :--- | :--- | ---: |
| id | UUID | unique identifier |
| name | String | title of the asset |
| text | String | long text for the asset |
| specs | ViewSpecs | visibility configuration |
| img | String | reference to location of image |

**Clue extends Item**

| Field | Type | Information |
| :--- | :--- | ---: |
| rating | number | rating of clue for assigning priority |

**Puzzle extends Item**

| Field | Type | Information |
| :--- | :--- | ---: |
| note | String | user added text |
| solution | String | answer of the puzzle |
| clues | Clue\[\] | clues in the puzzle |

**Stage extends Item**

| Field |  | Information |
| :--- | :--- | ---: |
| puzzles | Puzzle\[\] | puzzles in the stage |

**Theme extends Item**

| Field | Information |
| :--- | ---: |
| id | unique identifier |
| solution | keyword to complete the scene  |

| Field | Information |
| :--- | ---: |
| id | unique identifier |

