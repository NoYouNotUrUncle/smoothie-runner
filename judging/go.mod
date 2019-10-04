module github.com/BayviewComputerClub/smoothie-runner/judging

go 1.13

require (
	github.com/BayviewComputerClub/smoothie-runner/adapters v0.0.0-20191004210318-697ce8368920
	github.com/BayviewComputerClub/smoothie-runner/protocol v0.0.0-00010101000000-000000000000
	github.com/BayviewComputerClub/smoothie-runner/shared v0.0.0-00010101000000-000000000000
	github.com/BayviewComputerClub/smoothie-runner/util v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20191003171128-d98b1b443823 // indirect
	golang.org/x/sys v0.0.0-20191003212358-c178f38b412c
	google.golang.org/grpc v1.24.0 // indirect
)

replace github.com/BayviewComputerClub/smoothie-runner/protocol => ../protocol

replace github.com/BayviewComputerClub/smoothie-runner/util => ../util

replace github.com/BayviewComputerClub/smoothie-runner/shared => ../shared

replace github.com/BayviewComputerClub/smoothie-runner/adapters => ../adapters
