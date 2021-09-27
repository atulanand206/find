# Communication

Web Sockets protocol will get connected whenever a browser instance is launched. 

Select a value from the WebSocketMessageAction list and send as action in a WebSocketMessage. Pass in the relevant field required by the action.

The message prepending `S_` will be sent by the server and others by the client.

**WebSocketMessageAction - enum of actions for the messages.**

| Action | Information |
| :--- | ---: |
| BEGIN | quizmaster; call SYNCED |
| JOIN | player, team, match; call SYNCED  |
| START | Void; If players filled; call SYNCED |
| ANSWER | questionId; call S\_ANSWER |
| SCORE | questionId, playerId, points; call S\_QUESTION & call S\_PLAYER |
| NEXT | gameId; call S\_QUESTION |
| OVER | gameId; call S\_OVER |
| EXTEND | Void; Increment rounds by 1; call SYNCED |
| S\_GAME | game |
| S\_PLAYER | player |
| S\_QUESTION | question |
| S\_ANSWER | answer |
| S\_OVER | conclusion |

**WebSocketMessage - stores the message delivered over the bridge.**

| Field | Type | Information |
| :--- | :--- | ---: |
| action | WebSocketMessageAction | action for the current message |
| quizmaster | Player | person conducting the game |
| player | Player | person joining the game |
| game | Game | game info for UI |
| conclusion | Conclusion | game info for completion |
| question | Question | question |
| answer | Answer | answer |
| gameId | string | id of the game |
| playerId | string | id of the player |
| questionId | string | id of the question |
| points | number  | points for the question |

**Game - specifications of game object**

| Field | Type | Information |
| :--- | :--- | ---: |
| Quizmaster | Player | quizmaster for the game |
| Tags | \[\]string | ids of genres of which questions are already asked |
| Players | \[\]Player | Participants of the game |

**Conclusion - specifications of result object**

| Field | Type | Information |
| :--- | :--- | ---: |
| Quizmaster | Player | quizmaster for the game |
| Players | \[\]Player | Participants of the game |
| Questions | \[\]Question | questions asked in the game |
| Answers | \[\]Answer | answers for the questions asked in the game |

**Player - specifications of result object**

| Field | Type | Information |
| :--- | :--- | ---: |
| Name | string | name of the player |
| Email | string | email of the player |
| Scores | Score | scores of the player |

**Question - specifications of result object**

| Field | Type | Information |
| :--- | :--- | ---: |
| Statements | \[\]string | lines of the question |
| Tag | string | ref to the genre, tag, index |

**Answer - specifications of result object**

| Field | Type | Information |
| :--- | :--- | ---: |
| QuestionId | string | ref to the question |
| Answer | string | answer of the question |

**Score - specifications of score; internal object**

| Field | Type | Information |
| :--- | :--- | ---: |
| current | number | score in the current match |
| overall | number | overall score of the player |

**Data** **examples**

```text
v1: 27/09/2021

Player
{
  "_id":"5b0471ce-5193-4707-b17b-9ad7f5628926",
  "name":"James",
  "email":"cat@gc.com"
  "scores":{
    "overall":16
  }
}

Game
{
  "_id":"fed69145-475d-4edb-8509-f7c91a7607b6",
  "players":[ // Can't be a quizmaster
    {
      "_id":"ef60f971-f5ff-4772-9e83-ea6e65af2061",
      "name":"James",
      "email":"cat@gp.com"
      "scores":{
        "current":1
      }
    },
    {
      "_id":"5b0471ce-5193-4707-b17b-9ad7f5628926",
      "name":"James",
      "email":"cat@gc.com"
      "scores":{
        "current":1
      }
    }
  ],
  "quizmaster":{ // Must not be in players
    "_id":"65ee1721-0f45-42b5-9656-b6b2451425a5",
    "name":"James",
    "email":"cat@gre.com"
    "scores":{
      "current":0 // Must be 0
    }
  },
  "tags":[
    "58546356-f7eb-4580-9b11-9553a2904268",
    "fad1f57b-5fa7-46cd-b7fd-d07c75c337e7",
    "c47cfd75-947d-48f5-92fb-35f2364f1fc4",
    "ba9de0e6-c3b1-4d06-a5c4-de6db35cb99b",
    "019c66d4-e606-43a2-9985-8af735c5b354",
    "6685b05d-9808-41d6-8a61-c21b7374070d"
  ]
}

Index
{
  "_id":"ef6f7b8f-4f37-4880-bf63-a428e5abcd5d",
  "tag":"marvel"
}

Question
{
  "_id":"0ac7e80f-eed1-42c1-9336-a5da3d28073d",
  "statements":[
    "X & Y have a lot in common. From staying drunk all the time to ignoring the glamorous film industry where these two had few successes. Not too many to get them a big break but not too less that they couldn’t live lavishly.",
    "One lives a walk away from the beach while the other a jump away from the pool. Both of them were in the hearts of people because they never considered media as a big deal and lived unapologetically.",
    "They both also hosted guys in need and provided them shelter and those became the lead’s sidekicks.",
    "ID X & Y."
  ],
  "tag":"035a789c-a6ab-45af-b1b9-75f09501fa1a"
}

Answer
{
  "_id":"7798ece6-2319-40af-b4bc-1239db305ba0",
  "question_id":"0ac7e80f-eed1-42c1-9336-a5da3d28073d",
  "answer":"X: Two and a Half Men, Y: BoJack Horseman"
}
```

 

