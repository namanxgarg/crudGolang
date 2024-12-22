//not connecting database, only on struct and slices

package main

type Movie struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Isbn string `json="isbn"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main(){

}