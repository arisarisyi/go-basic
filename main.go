package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)
const namaAwal = "Al"
const namaAkhir = "Okeoke"
const emailNama = "Al@gmail.com"
const ticketnya = 56

var bookingz = make([]UserData,0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {
	// mapping golang
	// var myslice []string
	// var mymap map[string]string
	// var userData = make(map[string]string);
	// userData["firstName"] = namaAwal;
	// userData["lastName"] = namaAkhir;
	// userData["email"] = emailNama;
	// userData["numberOfTicket"] = strconv.FormatUint(uint64(ticketnya),10);


	var userData = UserData{
		firstName: namaAwal,
		lastName: namaAkhir,
		email: emailNama,
		numberOfTicket: ticketnya,
	};


	fmt.Println(userData)

	greetUser("Al",23)
	conferenceName := "Go Conference";
	const conferenceTicket int = 50;
	var remainingTickets uint= 50;
	// uint tidak bisa negative

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T/n",conferenceName,remainingTickets,conferenceTicket);

	fmt.Printf("Welcome to %v app\n",conferenceName);
	fmt.Printf("We have total of %v ticketes and %v are still available\n",conferenceTicket,remainingTickets);
	fmt.Println("------++++--------");

	bookings :=[]string{};

	
	

	// ini infinite loop kalo mau di kasih kondisi sma kayak biasa
	// for remainingTickets>0 && len(booking)<50{}
	// bisa juga loop seperti biasa
	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }
for {
	firstName,lastName,email,userTickets	:= getUserInput()

	isValidName, isValidTicketNumber , isValidEmail :=	helper.ValidateUserInput(firstName,lastName,uint(userTickets),remainingTickets,email);

	// break untuk keluar loop continue lanjut ke iterasi selanjutnya

	if isValidTicketNumber && isValidEmail && isValidName {
			remainingTickets = remainingTickets - uint(userTickets)
			
			bookings = append(bookings,firstName+" " + lastName);

			fmt.Printf("The whole slice: %v\n",bookings);
			fmt.Printf("The first array: %v\n",bookings[0]);
			fmt.Printf("Slice type: %T\n",bookings);
			fmt.Printf("Slice lenght: %v\n",len(bookings));

			fmt.Printf("Thank you %v %v for booking %v tickets. You wiil receive a confirmation on %v \n",firstName,lastName,userTickets,email);
			fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName);

			firstNames := []string{};
			for _,booking := range bookings{
				names := strings.Fields(booking);
				firstNames = append(firstNames,names[0])
			}
			// aslinya ini kalo variablenya mau di cuekin ganti _ seperti di atas
			// for index,booking := range bookings{
			// 	names := strings.Fields(booking);
			//   firstNames = append(firstNames,names[index])
			// }
			fmt.Printf("These are all our bookings: %v\n",firstNames);

				// #####Menjalankan Go routine tinggal tambah go saja diawal####
				wg.Add(1)
				go sendTicket(userTickets,firstName,lastName,email)
				
				
				// sendTicket(userTickets,firstName,lastName,email)

			noTicketsRemaining := remainingTickets ==0;
			if noTicketsRemaining{
				fmt.Println("Our conference is booed out");
				break;
			}
	} else{

			if !isValidName{
				fmt.Println("First name or last name to short!")
			}

			if !isValidEmail{
				fmt.Println("Your email is not valid")
			}

			if !isValidTicketNumber{
				fmt.Println("Your ticket number is invalid")
			}
		
		}

		// contoh switch case sama kayak yang llain

		// city := "London";

		// switch city{
		// case "New York":
		// 	// 
		// 	break
		// case "Jombang":
		// 	break
		// default:
		// 	break
		// }
	}
	wg.Wait()
}

func greetUser(name string, umur int){
	fmt.Printf("Function di golang nama %v umur %v\n",name,umur)
}


func getUserInput()(string,string,string,int){
		var firstName string;
	var lastName string;
	var email string;
	var userTickets int;

	// & pointer biar di simpan di memory
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)
	
	fmt.Println("Enter your first name?")
	fmt.Scan(&firstName);
	
	fmt.Println("Enter your last name?")
	fmt.Scan(&lastName);
	
	fmt.Println("Enter your email address?")
	fmt.Scan(&email);

	fmt.Println("Enter your number tickets:")
	fmt.Scan(&userTickets);

	return firstName,lastName,email,userTickets
}

func sendTicket(userTickets int, firstName string,lastName string,email string){
	// setTimeotnya golang
	time.Sleep(10*time.Second)
	var ticket= fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName);
	fmt.Println("$$$$$$$$$$$$$$$$")
	fmt.Printf("Sending ticket %v to %v\n",ticket,email);
	wg.Done()
}