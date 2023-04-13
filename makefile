SHELL=cmd.exe

create model:
	go run -mod=mod entgo.io/ent/cmd/ent new $(M)
