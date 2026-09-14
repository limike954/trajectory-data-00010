module github.com/limike954/trajectory-data-00010/examples

go 1.23.0

replace github.com/limike954/trajectory-data-00010/sqlite => ../sqlite

replace github.com/limike954/trajectory-data-00010 => ..

replace github.com/limike954/trajectory-data-00010/fsim => ../fsim

replace github.com/limike954/trajectory-data-00010/tpm => ../tpm

require (
	github.com/google/go-tpm v0.9.3
	github.com/google/go-tpm-tools v0.4.5
	github.com/limike954/trajectory-data-00010 v0.0.0-00010101000000-000000000000
	github.com/limike954/trajectory-data-00010/fsim v0.0.0-00010101000000-000000000000
	github.com/limike954/trajectory-data-00010/sqlite v0.0.0-00010101000000-000000000000
	github.com/limike954/trajectory-data-00010/tpm v0.0.0-00010101000000-000000000000
	github.com/syumai/workers v0.35.0
	hermannm.dev/devlog v0.5.0
)

require (
	github.com/ncruces/go-sqlite3 v0.24.1 // indirect
	github.com/ncruces/julianday v1.0.0 // indirect
	github.com/neilotoole/jsoncolor v0.7.1 // indirect
	github.com/syumai/workers-go v0.35.0 // indirect
	github.com/tetratelabs/wazero v1.9.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/term v0.30.0 // indirect
)
