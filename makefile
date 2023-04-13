SHELL=cmd.exe
a :=
create model:
	go run -mod=mod entgo.io/ent/cmd/ent new $(M)
