package text

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bot/bukarehatbot/entity"
)

// Start _
func Start() string {
	return "Halo! Bot ini dibuat oleh tim BukaRehat.\n" +
		"Jika ada pertanyaan silakan hubungi owner yang ada di https://github.com/tommynurwantoro/bukarehatbot\n" +
		"Coba gunakan /help untuk melihat perintah-perintah yang tersedia."
}

// Help _
func Help() string {
	return "Kamu bisa gunakan perintah-perintah ini loh:\n" + commands()
}

// HelpGroup _
func HelpGroup() string {
	return "Kamu bisa gunakan perintah-perintah ini loh:\n" + groupCommands()
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

// SuccessDeleteMicrobreak _
func SuccessDeleteMicrobreak(restTime string) string {
	return "Berhasil menghapus microbreak pada " + restTime
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

// NoLongerAnAdmin _
func NoLongerAnAdmin() string {
	return "Kamu sudah bukan admin BukaRehat lagi"
}

// AdminInitialized _
func AdminInitialized(username string) string {
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

// ReachMaxMicrobreak _
func ReachMaxMicrobreak() string {
	return "Microbreak kalian udah mencapai batas maksimal nih. Kamu udah gak bisa nambah lagi"
}

// NotFoundMicrobreak _
func NotFoundMicrobreak() string {
	return "Group kalian belum ada microbreak, buat pake command /micro {HH:mm} ya"
}

// NotFoundLeaderboard _
func NotFoundLeaderboard() string {
	return "Leaderboardnya belum ada nih, coba bikin event microbreak dulu ya"
}

// Private //

func commands() string {
	return "/halo Cuma buat nyapa aja"
}

func groupCommands() string {
	var b bytes.Buffer

	for _, c := range entity.GroupCommands() {
		b.WriteString(fmt.Sprintf("%s %s", c.Name, c.Description))
	}

	return b.String()
}
