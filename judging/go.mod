module github.com/BayviewComputerClub/smoothie-runner/judging

go 1.14

require (
	github.com/BayviewComputerClub/smoothie-runner/adapters v0.0.0-20191004210318-697ce8368920
	github.com/BayviewComputerClub/smoothie-runner/cache v0.0.0-20191005014351-73e6f012bacd
	github.com/BayviewComputerClub/smoothie-runner/protocol/runner v0.0.0-20200201204513-82f95cf7ffdf
	github.com/BayviewComputerClub/smoothie-runner/sandbox v0.0.0-20200315011139-7b50a3d3e486
	github.com/BayviewComputerClub/smoothie-runner/shared v0.0.0-20200315011139-7b50a3d3e486
	github.com/BayviewComputerClub/smoothie-runner/util v0.0.0-20200315011139-7b50a3d3e486
	github.com/elastic/go-ucfg v0.8.1 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/xid v1.2.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20200128174031-69ecbb4d6d5d // indirect
	golang.org/x/mod v0.2.0 // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/tools v0.0.0-20200131211209-ecb101ed6550 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/BayviewComputerClub/smoothie-runner/protocol/runner => ../protocol/runner

replace github.com/BayviewComputerClub/smoothie-runner/protocol/test-data => ../protocol/test-data

replace github.com/BayviewComputerClub/smoothie-runner/util => ../util

replace github.com/BayviewComputerClub/smoothie-runner/sandbox => ../sandbox

replace github.com/BayviewComputerClub/smoothie-runner/cache => ../cache

replace github.com/BayviewComputerClub/smoothie-runner/shared => ../shared

replace github.com/BayviewComputerClub/smoothie-runner/adapters => ../adapters
