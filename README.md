# go-api-crowdfunding

#TEST NYARI EMAIL MENGGUNAKAN SERVICE

input := user.LoginInput {
Email: "yudistira@gmail.com",
Password: "yudistirar626",
}

user, err := userService.LoginUser(input)

if err != nil {
fmt.Println("Gagal Login")
fmt.Println(err.Error())
return
}
	
fmt.Println(user.Email)
fmt.Println(user.Name)

# TEST FIND BY EMAIL MENGGUNAKAN REPOSITORY
	
userByEmail, err := userRepository.FindByEmail("samsudin@gmail.com")

if err != nil {
fmt.Println(err.Error())
}

fmt.Println(userByEmail.Name)

# TEST CREATE USER MENGGUNAKAN SERVICE

userInput := user.RegisterUserInput{}
userInput.Name = "Pesulap merah"
userInput.Occupation = "Pesulap"
userInput.Email = "pesulapmerah@gmail.com"
userInput.Password = "12345"

userService.RegisterUser(userInput)


# TEST CREATE USER MENGGUNAKAN REPOSITORY


user := user.User {
Name : "Gus Samsudin",
Occupation: "Padepokna Nur Dzat",
Email: "samsudin@gmail.com",
}

userRepository.Save(user)
