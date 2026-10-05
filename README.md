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
for _, x := range y
we assume that y is a slice or string
when y is slice this for loop just goes on and on getting x the value of y[HowManyTimesRuned] it is quite usefull when dealing with slices or arrays or taking apart strings into characters
when using with strings it always gets the character instead of int just like in C

In the connect_user function we use the net package or library of go to use command conn and net.Dial which connect the potencial user to the server or in this case a port on localhost as specified in net.Dial for the UserID I have now noticed that it is completely useless and we use it just for debuging or loging when an error happens I might actually remove this sometimes I feel like its a bit too useless but I guess I will use it for now 

usage of net.Dial is pretty simple you have to just make sure you do
conn, _ := net.Dial()
you don't have to make the error checking if you don't want in the () you just specify the type of traffic you want to make or connection sorry and to where to connecto to in our case its "localhost:"+port also you wan to make sure to conn.Close() at the end of your session preferably with defer

you can run functions as goroutines also doing so with "go" keyword before executing the function like so
go function_name()

return keyword is for exiting the function not the whole code

you can define keywords as you like or exapmle when defining goroutine doing as in the root main.go "go func (c net.Conn)" now in the go function I can use "c" as the alias for "net.Conn" this can be pretty handy 

set your goroutines wg.Done() at the start of the function with defer to make it easier to use it wrongly

when error handeling we have to use "if err != nil" because we are taking it as an error when the error is not empty if the error was empty there was no error

the TCP servers or ports - the tcp servers are listening in 3 goroutines 1 for each port [8080, 8081, 8082] we do defer wg.Done() to mark the goroutine as done when it finishes, than we make it listen with net.Listen and error check with log.Printf("%s", err) and we return if any eror happens so it doesn't do it agian blindly, than we make sure the listener ends at the end of the program by doing defer litener.Close() than we make for loop to listen for all the actual connections in this for loop we made the listener accept connections and made a simple error loging as we did before but with continue this time since the error just means there is no client to satisfy so we continue the for loop, than we have a goroutine that defines c as net.Conn which acts kinda like an alias for it and we ensure the net.Conn ends with defer c.Close() and than we write on the page with c.Write([]byte("")) the []byte is a slice of byte and the byte is specified in ("") than we make sure the goroutine is using the newest accepted connection with the (conn) at the end of the goroutine defining and we do this code copy paste 3 times for 3 ports and we are done when we launch this script it listens for the TCP connections on those ports than we can run tho balance/main.go that will simulate the users splited into 3 slices connecting to 3 servers we use real connections but not real servers and users

flags flags are imported by "flags" in the import () they can be used on every type of variable in GO I think but you have to specify every of them before, for example if I would wanted to run my program with flag named "users" with value of 4 I would have do "go run main.go -users=4" how does it work under the hood is quite simple, you have to make a variable like users with the flag.Int() so full exapmle of usage is
users := flag.Int("users", 0, "int")
so in the () the "users" means the flag will be called users so it knows when I do -users=... it knows it has to asign ... to users the 0 means the default value if the user does not specify this number so I like to use 0 there, the "int" means that when user does flag -h it means I think -help so like it prints all the capable flags you can use when running this program so in this case it prints that the flag "users" wants an int input. Before you use the flags as the variables you have to do
flag.Parse()
so the flags are actually usable as the variables

in the function check_server we get 2 variables there server (or port) and timeout the server is defined as string variable which contains in this case one of the ports like "8080" the timeout is a time variable as you can see it is defined as type time.Duration which means it conains time inside this function we can see
conn, err := net.DialTimeout("tcp", "localhost:"+server, timeout)
as we could probably tell we call net function "DialTimeout" which probing or checking the server if it is up and responding in some time in our case it is defined as timeout variable (later specified as 500ms) than we can see the 1st thing we define in this function is the addres of this server which could be some ip addres or as in our case a localhost port, as the 1st thing to define is the type of connection the function tries to make in our case that is "TCP" as always we are than checking if the err variable was empty or not and returning 0 if it wasn't than making sure the connection shuts down after we made sure it was even up

when calling function server_check and using it in our program we made a for loop using i and servers variable servers is a slice with all the servers that could be used and i is new variable that holds the value of servers[HowManyTimesRuned] we use the keyword range there as we specified earlier this can be handy when making certain types of loops than we check if the function returns 0 or 1 with earch server if it returns 0 that means the server is down and not responding within 500ms as we specify when calling tis function if the function returns 1 which means the server is up and responding we add the server or "i" into healthy_servers which how the name hints holds all the responding servers

so as of right now we use slices to catch and spread all the users so the slices act as like if the users are already slplit into the groups for each server to handle than we just connect them to make it use the healthy servers only will be tricky I assume we will need to come up with different strategy for this one I don't want to make some basic thing like making 3 if statements for whatever server is down and shit like that I would want to make something fancier, we need to touch only the logic fo spliting the users so looking at it right now we have for loop that checks every user with selected_server variable and adds the user to each slice making it evenly splited out so we could make it so that we check the server if its healthy so we can make a function or a for loop inside each case and check if they are in the list of healthy servers if not they would pass the case to another server like it would just continue without adding we could probably make it with a function for better handeling and compactability

so I have came up and deployed my solution so we will start from the server runner so I have made a function for starting a server rather than making it hard-coded it does the same hting just wtith a variable instead of hardcoded port, than just called the functions as they should be which helps me when testing the balancer later on by commenting the servers easily. For the balancer I have made a function that checks if the server is in the healthy servers slice if yes return 1 and if not return 0 which is pretty self explaing since it helps me again for repeating a procces in one line rather than hardcoding a for loop 3 times, than I have called this function in the case statements which are in if statement that asks if the output of this function is 1 or 0 if 1 than it means that the server is responding and is healthy and adds it into its slice, if it returns 0 or any other value it is false and it subtracts 1 from i (i -= 1) since it means we did not add 1 user to its server and would left it hanging, other than this the logic stayed the same althougth I said earlier I don't want to make it with 3 if statements I have meant like some nested loops or something like that I think this is pretty clever solution and is pretty readable for other developers or even myself after a while, I have tested it on I think every case that could happen with 3 servers so I think it is well made


^^ SIDE NOTES ^^




This README is mainly for documentation of my notes and what I have learned by doing this project. Some parts of this are in Czech language since it is my main language I will probably not be translating it to english any time soon sorry about that.
Also you might think how am I learning when I probably code by AI I do some parts of the code by AI for example I don't know how I would make the connect_user function work really so I make the AI do the job of coding it but I try to understand it as much as I can I am asking questions and noting there what I have taken from it so I truely understand how it works. And any of the code here is not copy-pasted.
