package abstract_factory

import "fmt"

type AbstractFactory interface {
	CreateBook() Book
	CreateFilm() Film
	CreateGame() Game
}

type Book interface {
	Reading()
}

type Film interface {
	Watching()
}

type Game interface {
	Playing()
}

type HorrorFactory struct{}

func (f *HorrorFactory) CreateBook() Book {
	return &HorrorBook{}
}

func (f *HorrorFactory) CreateFilm() Film {
	return &HorrorFilm{}
}

func (f *HorrorFactory) CreateGame() Game {
	return &HorrorGame{}
}

type HorrorBook struct{}

func (b *HorrorBook) Reading() {
	fmt.Println("Reading the best horror book!")
}

type HorrorFilm struct{}

func (f *HorrorFilm) Watching() {
	fmt.Println("Watching the best horror film!")
}

type HorrorGame struct{}

func (f *HorrorGame) Playing() {
	fmt.Println("Playing the best horror game!")
}

type SciFiFactory struct{}

func (f *SciFiFactory) CreateBook() Book {
	return &SciFiBook{}
}

func (f *SciFiFactory) CreateFilm() Film {
	return &SciFiFilm{}
}

func (f *SciFiFactory) CreateGame() Game {
	return &SciFiGame{}
}

type SciFiBook struct{}

func (b *SciFiBook) Reading() {
	fmt.Println("Reading the best sci-fi book!")
}

type SciFiFilm struct{}

func (f *SciFiFilm) Watching() {
	fmt.Println("Watching the best sci-fi film!")
}

type SciFiGame struct{}

func (f *SciFiGame) Playing() {
	fmt.Println("Playing the best sci-fi game!")
}

type AdventureFactory struct{}

func (f *AdventureFactory) CreateBook() Book {
	return &AdventureBook{}
}

func (f *AdventureFactory) CreateFilm() Film {
	return &AdventureFilm{}
}

func (f *AdventureFactory) CreateGame() Game {
	return &AdventureGame{}
}

type AdventureBook struct{}

func (b *AdventureBook) Reading() {
	fmt.Println("Reading the best adventure book!")
}

type AdventureFilm struct{}

func (f *AdventureFilm) Watching() {
	fmt.Println("Watching the best adventure film!")
}

type AdventureGame struct{}

func (f *AdventureGame) Playing() {
	fmt.Println("Playing the best adventure game!")
}

type Activity struct {
	Book Book
	Film Film
	Game Game
}

func (a *Activity) HaveActivity(factory AbstractFactory) {
	a.Book = factory.CreateBook()
	a.Film = factory.CreateFilm()
	a.Game = factory.CreateGame()

	// Use the created objects
	fmt.Println("Having activities...:")
	a.Book.Reading()
	a.Film.Watching()
	a.Game.Playing()
}
