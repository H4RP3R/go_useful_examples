package main

import (
	"fmt"
)

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %s article was written by %s", a.Title, a.Author)
}

// func typeCast(s fmt.Stringer) {
// 	article := s.(Article)
// 	fmt.Printf("%+v\n", article.Title)
// 	// fmt.Printf("%+v\n", s.Title)
// }

// func typeCast(s fmt.Stringer) {
// 	article, ok := s.(Article) // article == nil, if not ok
// 	if ok {
// 		fmt.Printf("%+v\n", article.Title)
// 	} else {
// 		fmt.Println("Invalid type assertion")
// 	}
// }

func typeCast(s fmt.Stringer) {
	if article, ok := s.(Article); ok {
		fmt.Printf("%+v\n", article.Title)
	} else {
		fmt.Println("Invalid type assertion")
	}
}

func main() {
	a := Article{
		Title:  "Go for beginners",
		Author: "James Bond",
	}

	typeCast(a)
}
