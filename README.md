# nbasimulator
A simple app that simulates an NBA league with 30 teams, 450+ players and weekly schedule

## 1) Setup
You can choose one of the following two ways to run the application:

### Run with Docker
The application is fully dockerized and all dependencies are included. With docker compose the application can be run easily.
```
git clone https://github.com/aselimkaya/nbasimulator.git
cd nbasimulator
docker-compose up --build
```

### Run with bash script file
This file contains all dependencies and requirements. In order to run the project successfully, it will be sufficient to run the following commands in order.
```
git clone git clone https://github.com/aselimkaya/nbasimulator.git
cd nbasimulator
chmod +x run.sh
./run.sh
```

After successfully running these commands you should see this log in terminal:  ```http server started on [::]:1323```

## 2) API Reference

NBA simulator application supports two simple operations:
* Show real-time live game results
* Show assist leader after all games are completed

### 2.1) Real-Time Game Results

If you go to ```localhost:1323/results``` in the web browser, the game scores will be accessed in real time.

| Away | Away Score | Home Score | Home | Remaining |
|------|------------|------------|------|-----------|
| SAS  | 9          | 9          | CLE  | 203       |
| MIN  | 10         | 2          | MIL  | 203       |
| GSW  | 10         | 8          | ORL  | 203       |
| ATL  | 13         | 4          | CHA  | 203       |
| MEM  | 9          | 14         | NYK  | 203       |
| PHI  | 12         | 13         | POR  | 203       |
| UTA  | 4          | 9          | CHI  | 203       |
| BKN  | 2          | 2          | OKC  | 203       |
| BOS  | 11         | 3          | DAL  | 203       |
| WAS  | 5          | 4          | LAC  | 203       |
| DEN  | 11         | 8          | IND  | 203       |
| SAC  | 2          | 8          | PHX  | 203       |
| NOP  | 8          | 8          | TOR  | 203       |
| LAL  | 9          | 6          | MIA  | 203       |
| HOU  | 2          | 5          | DET  | 203       |

### 2.2) Leaderboard

Leaderboard page currently only support assist leader. If you go to ```localhost:1323/leaderboard``` in the web browser, assist leader information will be accessed.

| Assist Leader |                     |
|---------------|---------------------|
| Name          | Shareef Abdur-Rahim |
| Team          | MEM                 |
| Assists       | 6                   |

If the matches are not completed yet, an error message will appear on this page: ```results not yet finalized```

## 3) Database

Mongo DB is selected as database. The following four collections are created when the application runs successfully.
* Team
* Player
* PlayerGameInfo
* Game

### 3.1) Team Collection
The information of 30 teams in the league is written to the database thanks to [this free API](https://www.balldontlie.io/). An example team information is as follows.
```
{
  "_id": {
    "$oid": "6189b1b4769a3b93c2cce681"
  },
  "abbreviation": "BOS",
  "name": "Boston Celtics",
  "players": [
    {
      "name": "Jabari Bird",
      "team": "BOS"
    },
    {
      "name": "Michael Smith",
      "team": "BOS"
    },
    {
      "name": "John Bagley",
      "team": "BOS"
    },
    {
      "name": "Rick Fox",
      "team": "BOS"
    },
    {
      "name": "Marcus Webb",
      "team": "BOS"
    },
   ...
  ]
}
```

### 3.2) Player Collection
The information of players in the league is written to the database thanks to [this free API](https://www.balldontlie.io/). An example player information is as follows.
```
{
  "_id": {
    "$oid": "6189b1b4769a3b93c2cce554"
  },
  "name": "Hakeem Olajuwon",
  "team": "HOU"
}
```

### 3.3) Player Game Info Collection
Player stats are stored in this collection after the game is completed. An example player game information is as follows.
```
{
  "_id": {
    "$oid": "6189b2a4769a3b93c2cce6f1"
  },
  "game_id": "PHIvsPOR",
  "player": {
    "name": "Warren Kidd",
    "team": "PHI"
  },
  "player_stats": {
    "two_point_made": 2,
    "two_point_attempt": 3,
    "three_point_made": 1,
    "three_point_attempt": 1,
    "assist": 1
  }
}
```

### 3.4) Game Collection
The statistics of the two teams playing a game are stored in this collection. An example game is as follows.
```
{
  "_id": {
    "$oid": "6189b2a4769a3b93c2cce6da"
  },
  "game_id": "MINvsMIL",
  "away": {
    "game_id": "MINvsMIL",
    "team": {
      "abbreviation": "MIN",
      "name": "Minnesota Timberwolves",
      "players": [
        {
          "name": "Randy Breuer",
          "team": "MIN"
        },
        {
          "name": "Doug West",
          "team": "MIN"
        },
        {
          "name": "Tony Campbell",
          "team": "MIN"
        },
       ...
      ]
    },
    "players": [
      {
        "game_id": "MINvsMIL",
        "player": {
          "name": "Randy Breuer",
          "team": "MIN"
        },
        "player_stats": {
          "two_point_made": 1,
          "two_point_attempt": 4
        }
      },
      {
        "game_id": "MINvsMIL",
        "player": {
          "name": "Doug West",
          "team": "MIN"
        },
        "player_stats": {
          "two_point_made": 2,
          "two_point_attempt": 4,
          "three_point_attempt": 1,
          "assist": 1
        }
      },
      {
        "game_id": "MINvsMIL",
        "player": {
          "name": "Tony Campbell",
          "team": "MIN"
        },
        "player_stats": {
          "two_point_attempt": 3,
          "three_point_made": 1,
          "three_point_attempt": 2,
          "assist": 2
        }
      },
      ...
    ],
    "team_stats": {
      "game_id": "MINvsMIL",
      "score": 49
    }
  },
  "home": {
    "game_id": "MINvsMIL",
    "team": {
      "abbreviation": "MIL",
      "name": "Milwaukee Bucks",
      "players": [
        {
          "name": "Dan Schayes",
          "team": "MIL"
        },
        {
          "name": "Jay Humphries",
          "team": "MIL"
        },
        {
          "name": "Ricky Pierce",
          "team": "MIL"
        },
       ...
      ]
    },
    "players": [
      {
        "game_id": "MINvsMIL",
        "player": {
          "name": "Dan Schayes",
          "team": "MIL"
        },
        "player_stats": {
          "two_point_made": 2,
          "two_point_attempt": 2,
          "three_point_attempt": 1
        }
      },
      {
        "game_id": "MINvsMIL",
        "player": {
          "name": "Jay Humphries",
          "team": "MIL"
        },
        "player_stats": {
          "two_point_made": 1,
          "two_point_attempt": 4,
          "assist": 1
        }
      },
      {
        "game_id": "MINvsMIL",
        "player": {
          "name": "Ricky Pierce",
          "team": "MIL"
        },
        "player_stats": {
          "two_point_made": 2,
          "two_point_attempt": 3,
          "assist": 3
        }
      },
      ...
    ],
    "team_stats": {
      "game_id": "MINvsMIL",
      "score": 44
    }
  }
}
```

## 4) Simulator Working Mechanism
The game simulator consists of 4 sub-mechanisms
* Init
* Random
* Run
* Scheduler

### 4.1) Init
It is the first working sub-mechanism. It fetchs the information of basketball teams and players from the API and saves them in the database.

### 4.2) Random
It consists of random number generator functions provided for generating statistics in the simulator. A team's successful attack, whether the shot is 2 or 3 points, and whether the shot is accurate depends on the functions here.

### 4.3) Run
It is the main function that creates the services and starts the simulator.

### 4.4) Scheduler
It contains functions that match 30 teams in pairs and organize a game for each pair.