package text

// Start _
func Start() string {
	return "Halo! Bot ini dibuat oleh Tommy (@tommynurwantoro).\nCoba gunakan /help untuk melihat perintah-perintah yang tersedia."
}

// Help _
func Help() string {
	return "Kamu bisa gunakan perintah-perintah ini loh:\n" + commands()
}

// Halo _
func Halo(username string) string {
	return "Halo, " + username + ". 👋🏻"
}

// InvalidCommand _
func InvalidCommand() string {
	return "Aku gak ngerti perintah itu, coba perintah yang lain ya."
}

// SuccessInsertMicrobreak _
func SuccessInsertMicrobreak(restTime string) string {
	return "Berhasil menambahkan microbreak pada " + restTime
}

// Private //

func commands() string {
	return "/halo Cuma buat nyapa aja"
}
