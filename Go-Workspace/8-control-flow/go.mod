module main

go 1.15

replace example.com/while => ./while_loops

replace example.com/initCondPost => ./init_cond_post

replace example.com/ifStatements => ./if_control_flow

require (
	example.com/ifStatements v0.0.0-00010101000000-000000000000 // indirect
	example.com/initCondPost v0.0.0-00010101000000-000000000000
	example.com/while v0.0.0-00010101000000-000000000000
)
