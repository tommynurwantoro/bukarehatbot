package text

import "strings"

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
	return "Nama group udah diganti menjadi `" + args + "` ya"
}

// ShowGroupName _
func ShowGroupName(name string) string {
	return "Nama group kalian " + name
}

// UnknownGroupName _
func UnknownGroupName() string {
	return "Admin belum ngubah nama group ya"
}

// SuccessAddMember _
func SuccessAddMember(usernames []string) string {
	return "Berhasil menambahkan " + strings.Join(usernames[:], ", ")
}

// InvalidParameter _
func InvalidParameter() string {
	return "Parameternya belum bener tuh, coba dicek lagi ya"
}

// AdminChanged _
func AdminChanged(username string) string {
	return "Admin sudah diubah ke @" + username + " ya"
}

// UnableToChangeAdmin _
func UnableToChangeAdmin(username string) string {
	return "Kamu gak bisa ubah admin, yang bisa ubah adminnya cuma admin yang sekarang. Coba tanya @" + username
}

// UserNotInGroup _
func UserNotInGroup(username string) string {
	return "@" + username + " gak ada di group kamu, coba cek lagi ya."
}

// Private //

func commands() string {
	return "/halo Cuma buat nyapa aja"
}
