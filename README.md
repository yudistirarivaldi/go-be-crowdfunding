# go-api-crowdfunding


# TEST JWT TOKEN
	
fmt.Println(authService.GenerateToken(1001))

# TEST UPLOAD AVATAR IN SERVICE
	
userService.SaveAvatar(6, "images/1-profile.png")

# TES SAVE AVATAR MENGGUNAKAN SERVICE

userService.SaveAvatar(5, "image-1-2-.jpg")
	
# CEK EMAIL TERSEDIA ATAU TIDAK MENGGUNAKAN SERVICE

input := user.CheckEmailInput {
Email: "pesulapmerah123@gmail.com",
}

bool, err := userService.IsEmailAvailable(input)
if err != nil {
fmt.Println("Gagal")
}
	
fmt.Println(bool)

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

# TEST LOGIN USER MENGGUNAKAN SERVICE

input := user.LoginInput {
		Email: "yudis123@gmail.com",
		Password: "password",
	}

	user, err := userService.LoginUser(input)
	if err != nil {
		fmt.Println("Gagal Login Brok")
	}

	fmt.Println(user.Email)
	fmt.Println(user.PasswordHash)

# TEST CREATE USER MENGGUNAKAN REPOSITORY


user := user.User {
Name : "Gus Samsudin",
Occupation: "Padepokna Nur Dzat",
Email: "samsudin@gmail.com",
}

userRepository.Save(user)
