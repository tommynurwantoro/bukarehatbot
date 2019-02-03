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

// InvalidCommandForUser _
func InvalidCommandForUser(username string) string {
	return "Kamu gak boleh pakai perintah ini, ngomong dulu ke @" + username + " ya"
}

// AdminNotFound _
func AdminNotFound() string {
	return "Hanya admin yang boleh pakai perintah ini"
}

// ChangeGroupName _
func ChangeGroupName(args string) string {
	return "Nama group udah diganti menjadi " + args
}

// ShowGroupName _
func ShowGroupName(name string) string {
	return "Nama group adalah " + name
}

// UnknownGroupName _
func UnknownGroupName() string {
	return "Admin belum set nama group"
}

// Private //

func commands() string {
	return "/halo Cuma buat nyapa aja"
}
