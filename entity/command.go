package entity

// Command _
type Command struct {
	Name        string
	Description string
}

// GroupCommands _
func GroupCommands() []Command {
	return []Command{
		{"/halo", "Cuma buat nyapa aja\n"},
		{"/change_admin", "{username} Ngubah admin ke member lain di dalam tim\n"},
		{"/add_member", "{username} Nambahin user ke bukarehat\n"},
		{"/show_group_name", "Nunjukin nama tim kamu untuk leaderboard\n"},
		{"/change_group_name", "{name} Ngubah nama tim kamu untuk leaderboard\n"},
		{"/micro", "{HH:mm} Ngatur jam microbreak buat tim kamu\n"},
		{"/show_micros", "Nunjukin jam-jam microbreak yang udah diatur admin tim kamu\n"},
		{"/leaderboard", "Nunjukin poin-poin yang dimiliki tiap member\n"},
	}
}
