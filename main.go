package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)   // по сути слайс(список) пользовательских структур данных 

type UserData struct {          // структуру можно сравнить с Классом в ООП
	firstName 		string
	lastName  		string
	email     		string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // создаем весовую группу для синхронизации легк. потоков. Делаем это для того, чтобы наша main горутина дала завершиться другим потокам входящим в нее.
// ф-цмя wg.Add() - увеличивает счетчик потоков, которых нашего основное приложение должно ожидать, а ф-ция wg.Done() уменьшает этот счетчик. Таким образом, когда счетчик равен нулю, что означает что у осноного потока нет потоков для ожидания - он может завершить работу приложения (так что ждать уже не нужно)

func main() {

	greetUsers()

	
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber :=  validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)   // добавляем ровно столько потоков, сколько нужно. В данном случае только один
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)


		if remainingTickets == 0 { 
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short. Try again")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign. Try again")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid. Try again")
		}
	}  
	wg.Wait()  // ожидаем пока все потоки, которые были добавлены в main выполнят свою работу, прежде чем приложение сможет заверщить работу. (то есть мы ждем пока поток go sendTicket() не будет завершен)
}



func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Println(conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
		for _, booking := range bookings {    // range возвращает индекс и значение для каждого элемента
			firstNames = append(firstNames, booking.firstName)
		}     
		return firstNames
}

func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Printf("Enter your first name: \n")
	fmt.Scan(&firstName)   // ввод данных присвоет это значение в память, где хранится переменная

	fmt.Printf("Enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: \n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: \n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v.\n", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############")
	wg.Done()  // ф-ция Done удаляет поток здесь из списка ожидания, по сути она сооющает группе ожидания(wg.Wait()- "что я закончил выполнение, так что основному потоку больше не нцжно меня ждать" )
}