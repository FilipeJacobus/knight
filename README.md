# knight API

An API that returns all possible knight positions (in chess) in two turns

## Install
### requires: 
- docker
- docker-compose

```
docker-compose up -d
```

## Endpoints

-  ### GET - /get-positions

    Return  all possible knight positions
    ### params:
    - player1 - Player 1 position 
    - player2 - Player 2 position 

## Example

#### Request:
```
curl --location --request GET 'localhost:9876/get-positions?player1=d4&player2=h1'
```

#### Response:
```
{
    "player1": [
        "f5",
        "e6",
        "c6",
        "b5",
        "b3",
        "c2",
        "e2",
        "f3"
    ],
    "player2": [
        "f2",
        "e3",
        "c3",
        "b2"
    ]
}
```