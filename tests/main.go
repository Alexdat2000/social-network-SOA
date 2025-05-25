package main

func main() {
	ClearTablePostgres("postgres_users", 5432, "users", "users")
	ClearTablePostgres("postgres_content", 5433, "content", "entries")
	ClearTablePostgres("postgres_content", 5433, "content", "comments")
	ClearTableClick("stats.views")
	ClearTableClick("stats.likes")
	ClearTableClick("stats.comments")
}
