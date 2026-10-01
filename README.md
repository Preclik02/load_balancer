Load balancer - prvni udelam logiku pro to ze nejaky user patri nejakemu serveru takze
user bude vybran do 1 ze 3 arrayu tyto arraye budou reprezentovat servery napr server_8080 bude logicky server 1 server_8081 bude server 2 a server_8082 bude server 3 zatim to bude takto jednoduse hard coded cislo protoze logicky vis kolik mas k dispozici serveru pote se pridaji checky ze treba 1 server nefunguje tak se vyzuvaji jenom 2 a 3 nebo tak to uz vsak program bude zjistovat sam bez zadneho hard codidnug nebo inputu take se misto klasickeho inputu budou vyuzivat tzv. flags pri spousteni

^^ MAIN IDEA ^^

puzivame arraye nebo po novem slangu v GO "slicy" na to aby sme priradily serveru uzivatele takze finalni uzivatele pro kazdy server sou v slicu napr. server_8080

var wg sync.WaitGroup
wg.Add(1) - adds to counter like it tells it that some goroutine is running (state += 1)
wg.Done - makes the program think all goroutines are finished (state = -1)
wg.Wait - waits for all the goroutines to finish essentially just waiting for the state to go to 0
use this when connecting users with 3 for loops

go func() {}() - is a nameless function that is only there once you can't call it twice you just make it it executes itself kinda like another process and done there are a couple of things you should always do in it
always call defer wg.Done() it mars the goroutine or function as done when it ends
when you define this function it can have a couple of defited variables to it there is an example
go func(x int, y int) {
	fmt.Println("Hello, world")
	fmt.Println(x)
	fmt.Println(y)
}(10, 20)
This example makes the function have variables (int) x, y which are defined as 10, 20 at the end so we can use them like this

Keyword "Range" this keyword can be used like so
for _, x := Range y
we assume that y is a slice or string
when y is slice this for loop just goes on and on getting x the value of y[HowManyTimesRuned] it is quite usefull when dealing with slices or arrays or taking apart strings into characters
when using with strings it always gets the character instead of int just like in C




^^ SIDE NOTES ^^




This README is mainly for documentation of my notes and what I have learned by doing this project. Some parts of this are in Czech language since it is my main language I will probably not be translating it to english any time soon sorry about that.
Also you might think how am I learning when I probably code by AI I do some parts of the code by AI for example I don't know how I would make the connect_user function work really so I make the AI do the job of coding it but I try to understand it as much as I can I am asking questions and noting there what I have taken from it so I truely understand how it works. And any of the code here is not copy-pasted.
