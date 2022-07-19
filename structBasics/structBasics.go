package main

import "fmt"

func main() {

	// struct  // a type in go // sequence of fields(named elements) of any type // this sequence is fixed and checked on compile time only

	//life without structs
	book1, author1, year1 := "boo1", "auth1", 1984
	book2, author2, year2 := "boo2", "auth2", 1948

	fmt.Println(book1, author1, year1)
	fmt.Println(book2, author2, year2)

	// creating structs // 1
	type book struct {
		title  string
		author string
		year   int
	}
	type book11 struct {
		title, author string
		year          int
	}
	b1 := book{"boo1", "auth1", 1984} // order has to be correct // if this is screwed, then eveything can go for a toss
	fmt.Println(b1)

	// creating structs // 2 //
	b1 = book{title: "boo1", author: "auth1", year: 1949}
	fmt.Println(b1)

	// creating structs // partial fields //
	b1 = book{title: "boo1"}
	fmt.Printf("%#v\n", b1) // main.book{title:"boo1", author:"", year:0}
	fmt.Println(b1.title)
	// fmt.Printf(b1.randomField) // compile error

	//updating structs
	b1.author = "auth22"
	fmt.Printf("%+v\n", b1) //  {title:boo1 author:auth22 year:0} // +v just prints fields:values

	//struct comparison
	// more of a compiler time ds // can clone to other variables and they are independent
	// we can compare structs using == and !=
	// TODO what happens is one field is derived from slice wagera // no comparable thing

	b2 := b1

	fmt.Println(b1 == b2)
	b2.author = "asdfa"
	fmt.Println(b1 == b2) // deep cloned with independent memory // hence now not equal

	// anonymous structs
	// one time structs // not defining a  named type for structs // the struct definition only becomes the names
	roger := struct {
		name  string
		place string
		age   int
	}{
		name:  "Roger",
		place: "Swiss",
		age:   10,
	}
	fmt.Printf("%#v\n", roger) // struct { name string; place string; age int }{name:"Roger", place:"Swiss", age:10}

	// anonymous fields
	// not defining a name for the fields // the datatype only becomes field names
	type Games struct {
		name string
		int
		bool
	}
	tennis := Games{"Tennis", 2, true}
	fmt.Printf("%#v\n", tennis)
	tennis.bool = false
	tennis.int = 10
	tennis.name = "TennisTable"
	fmt.Printf("%#v\n", tennis)

	// nested structs
	type Game struct {
		name        string
		playerCount int
	}
	type Player struct {
		name        string
		gameItPlays Game
	}

	kohli := Player{
		"kohli",
		Game{
			"kirkit",
			10,
		},
	}
	fmt.Printf("%#v\n", kohli)
	fmt.Println(kohli.gameItPlays.name)
	fmt.Println(kohli.name)
	fmt.Println(kohli.gameItPlays)

}
