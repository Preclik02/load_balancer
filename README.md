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

So I am gonna start thinking about this new issue for dynamic servers I think this will be more of an issue in the server startup main.go not the logic itself although it will probably still asign the input but if not I mean no one ever specified if I had to do it in one script or 2 and even than if it is only with one I can asign the servers to grab the value from a file I will writhe the flag to and just loop the function with the port being ++ so yea ig bro and yea now I remembered, I need to split the users into x slices I didn't think aobut that, so we will definitely need some temporary files there to like send each script the server count but the startup script will still have the server number 1st that is not the question the main question here is how to make the slices be dynamic for the spliting right now I have hard coded 3 slices representing each server to make it dynamic I need to define x arrays or the easier way I can make a 2D array as in c it defines like so
char x[][]
but I don't know how to make that thing in GO since it is even now very different with defining the slices for int ors trings like []int{ } is very different from normal C syntax I will try to make it somehow and than come back to finish this idea okay I am back with the knowladge I need to make this happen althought I don't have time to finish this I will still write the basic idea of it now there 1st for the 2d arrays in general than the solution

so for the 2d arrays you can make them fixed like this
var x [][]int
or
x := [][]int
there the x is the 2d array and you can see it is int type so it fits our case we can add to the rows dynamically with append but for columns which in our case we can specify at the start of the program so we will don't have to make it multiple times so we literaly can just do
var x [][x]int
the x will be replaced with some new variable for the server and the x will be the server count since thre the values for columns and rows are the exact oposite from C we are in the 1st [] specifying the rows and int the 2nd [] columns and than we specify the type of the array or slice

for my solution I have come up to and end that we will make the startup server main.go to save the variable for the number of servers into a temp file and start the servers using for loop after that we are done with the server startups for the main.go logic that splits the users up ino groups we would have to get the value from the file than define a 2d array that would have fixed columns with the server count and rows would be left out like [] and they will be appended the users into them than for the spliting into groups we have now 3 cases we would work with the [][x] of the array and the server count so we would have to make a while loop that would work until user_asignd or something == users and for the users spliting we would make a variable i or sm that would count the column of the main slice or the port number and if statement at the end would check if it == to the server count if yes than reset it to 0 and start asigning again from server 0 till the users == users_asigned for the asigning make if statement if the server is healthy if yes than append this user and also make it += 1 and if not nothing happens for the connections we can keep the same logic again like make a while loop and we can make a function for the goroutines and as variables used there we can use the user wg and port before that we would need to make the port a string with method specified bellow and than call the function multiple times until it is the end of the column with the server so we will start scanning that server from 1 not 0 (skipping the server int) for the server checking it will be pretty much the same we will just need to convert the port into stirng and run it till the servers int or string and that is pretty much it that is the base idea behind that it will probably change along the way but this is what it will somehow look like and something I will look back on when making the actual function

now I am relising one "flaw" with this structure and that is that the server row is in int and not string as many functions want it to be I am gonna make a quick research on that if I can convert the int into string so importing strconv at the start is a great start and we can use 
x = strconv.Itoa(y[0])
which means x is eaqual to the y[0] but in string since "Itoa" means "Int to ASCII" and strconv is the keyword for calling that function the y[0] is just the value so before using the variable as a string we can make it a string from int using this 

reading files is pretty easy, or at least in our case, we have to import os and log and just do
data, err := os.ReadFile("filename")
if err != nil {
    log.Fatal(err)
}
and now the variable data is %s of the file so we can work with that. For writing data into a file in our case it is also pretty straight forward like so:
data := []byte("Hello world")
err := os.WriteFile("filename", data, 0644)
// error handling // 
in our case there is a proble where we have pointer int and we need it into a string we can do that by doing fmt.Sprintf("%d", data) into the byte slice

converting os.ReadFile into int is pretty complex process or at least I see it as too much complex than it should be since you have to read the file with os.ReadFile than convert it into string using strings.TrimSpace that will delete any new lines or spaces 

so the commit will happen now when I have in progress this update to dynamic servers - the program doesn't even compile now so that is yet to fix the reason for not getting it ready now or in the next commit is I have told myself I will do every day commits and the 2nd thing is I have to sleep for school now and I want good sleep since there are tests waiting for me tomorrow

so I have just merged my 1st PR which was essentially me being done in class of programming and I made a small bug fix with opening an issue and a PR on my 2nd account. Although it was me it is a great feeling gettin my 1st ever PR merge like me doing the merge action that is just sweet.

I have actually taken some notes on paper since sometimes I like it more when it is written on paper to even have like rought sketches on how the code will work which is quicker done on paper than in IDE or text editor, in short the notes just say that the asigning of users to servers should be done by the for loop that is and was there it will just refactor the use of selected_server and make if statements for the values limits so like it should be smaller than the servers_count and it will work like the i does at the end of 1 loop it will ++ it and that the asignment of it will just do the column of i and row of server_selected and append to that cordinate.

Well after making the asignment function of for loop my bad, it was too easy I thought that it will be harder on thinking I this feeling was probably coused by the size of the file since I thought that it will not comile a lot of times and stuff like that, it does compile BUT we don't know if it will work as intended we can only wish now, one thing I learned doing this is that when you appending to a slice like this you have to put the x[i] in the begining too like x[i] = append() so pretty cool I guess, than it is pretty straight forward like doing the ++ and checking for the limits of the server count.

So I just made another notes and a sketch on the connection for loops I am thinking of a nested loop I know this is not the most optimized solution to this but I am making this so it works not to be the quickes thing existing, also this is still the core of load_balancer the optimizatoin and structuring will come next after some fundemental changes. To the notes they just say something like that I am thinking about the nested loop and 1st one will check the server if it is valid like the number of server than the 2nd loop will connect all users using the goroutine that I made before also thinking about making this goroutine a function to save some time and 7lines of code but maybe some other time.

So I just did the full change with the server connections and server checks they both work I did in fact not touch the functions but I did just renamed the variables and made the logic with 2d slice rather than hardcoding a value the main challange was probably the check health logic since it is pretty much a lot of code putted close together and I have used a lot of converting there I know I could make it once and just use the variable but I just felt like this was better right now for some reason so maybe I will fix it sometimes soon I guess, also the chech health logic is now like less complex  Iremoved the range keyword since it was using the servers variable but now the servers is empty and the shortest solution is to just make a clasic for loop with servers_count so that was I guess the main change for this for the connecting I just made another clasic for loop with the goroutine and for loop in it the for loop with goroutine was left untouched really I guess only there is a strconv for the port to be in string and not int so yea that was it for this 2/2 update.

so about this time I am able to code while being in high school since this is my 1st year there I don't know if I can use personal notebooks without any restriction or not so I am half risking but as an IT major I would expect the notebooks to be compleatly free to use on breaks or sm like that, so I have made issue on free time in other classes from the 2nd github account about that I don't have good structure and code optimalization. For toady or like this commit I would like to prepare good functino uses like make even small repeated taks to be in a function not only for readability but also for the scaleability and update making, comming from C I know making functions can sometimes be really messy or not understandable well couse of the pointer and it is pretty much trail and error but whatever I guess.

So now on the high school break I don't have the brain capacity or time to make complex structure changing so I just went through the code and commented it all the way through with an exception of the server starting main.go I don't have the time to do that. Okay actually did all of the documenting and commenting but the break is over now. Going to push this code and comming back to it later on.

Okay so after hard days of schools I am absolutely tired but I will still make this commit because I don't want to lose my streak of commits, I have got distracted a lot of times while comming up with an idea for this since it is wednesday it is almost end of the week and my sleep schedule is not covering this also this week is a lot of academic pressure in school halfway through the week after like 5 test and at least 1 more is comming at me so w ig.

For this commit I have thought about making the project do the starting of each script at the same time improving the structure although the code will be now harder to test since before it was simplier kinda but it will be alright even now I assume, so for the code, the code will be pretty simple it wil accept the flags how the programs would do but they would do it one by one this one makes it into one call, after the flags are parsed than make a call on starting server program as goroutine since it will not finish

I will explain this code later on I have no time now ...


^^ SIDE NOTES ^^




This README is mainly for documentation of my notes and what I have learned by doing this project. Some parts of this are in Czech language since it is my main language I will probably not be translating it to english any time soon sorry about that.
Also you might think how am I learning when I probably code by AI I do some parts of the code by AI for example I don't know how I would make the connect_user function work really so I make the AI do the job of coding it but I try to understand it as much as I can I am asking questions and noting there what I have taken from it so I truely understand how it works. And any of the code here is not copy-pasted.
