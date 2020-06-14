package definitions

type Game struct {
	Id string `json:"id"`
	Timeout int `json:"timeout"`
}

type Board struct {
	Height int `json:"height"`
	Width int `json:"width"`
	Food []Point `json:"food"`
	Snakes []BattleSnake `json:"snakes"`
}

type BattleSnake struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Health int `json:"health"`
	Body []Point `json:"body"`
	Head Point `json:"head"`
	Length int `json:"length"`
	Shout string `json:"shout"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SnakeConfig struct {
	Apiversion string `json:"apiversion"`
    Author string `json:"author"`
    Color string `json:"color"`
    Head string `json:"head"`
    Tail string `json:"tail"`
}

type MainRequest struct {
	Game Game `json:"game"`
	Turn int `json:"turn"`
	Board Board `json:"board"`
	You BattleSnake `json:"you"`
}

type Move struct {
	Move string `json:"move"`
	Shout string `json:"shout"`
}