# battlesnake

Experimenting with golang for https://play.battlesnake.com/

For now, the snake only tries to eat food following the shortest path available. The board is traversed with BFS

Some strategies worth exploring :

* Target closest food
* Target food with least enemy around
* Avoid moving to a square that will potentially have an enemy snake
* Maybe learn and try to implement some game theory algorithms

This program will listen on port 3000 and handle `/` `/start` `/move` `/end` HTTP requests

Using ngrok to open a tunnel to the local server


