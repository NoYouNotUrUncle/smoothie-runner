module github.com/BayviewComputerClub/smoothie-runner/shared

go 1.14

replace github.com/BayviewComputerClub/smoothie-runner/protocol/runner => ../protocol/runner

replace github.com/BayviewComputerClub/smoothie-runner/util => ../util

require (
	github.com/BayviewComputerClub/smoothie-runner/protocol/runner v0.0.0-20200201204513-82f95cf7ffdf
	github.com/BayviewComputerClub/smoothie-runner/util v0.0.0-20191005014351-73e6f012bacd
	golang.org/x/net v0.0.0-20191003171128-d98b1b443823 // indirect
	google.golang.org/grpc v1.24.0 // indirect
)
