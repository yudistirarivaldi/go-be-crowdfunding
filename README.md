# go-api-crowdfunding

====================================
TEST CREATE USER MENGGUNAKAN SERVICE
====================================

userInput := user.RegisterUserInput{}
userInput.Name = "Pesulap merah"
userInput.Occupation = "Pesulap"
userInput.Email = "pesulapmerah@gmail.com"
userInput.Password = "12345"

userService.RegisterUser(userInput)

=======================================
TEST CREATE USER MENGGUNAKAN REPOSITORY
=======================================

user := user.User {
Name : "Gus Samsudin",
Occupation: "Padepokna Nur Dzat",
Email: "samsudin@gmail.com",
}

userRepository.Save(user)
