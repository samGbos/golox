# golox
Lox tree-walk interpreter written in go
Following the [Crafting Interpreters](http://craftinginterpreters.com/) book.
Also playing with compiling to wasm to run from the web.

**Notes to myself:**
Compile and run cli version with:
`go install ./...`

Compile the wasm version with:
`GOOS=js GOARCH=wasm go build -o ~/web/main.wasm`
from the `cmd/golox-wasm` directory 
