# golox
Lox tree-walk interpreter written in go
Following the [Crafting Interpreters](http://craftinginterpreters.com/) book.
Also playing with compiling to wasm to run from the web.

**Notes to myself:**

Compile and run cli version with `go install ./...`

Compile the wasm version with `GOOS=js GOARCH=wasm go build -o ~/web/main.wasm`
from the `cmd/golox-wasm` directory 


Thoughts on go:
- So quick to learn!
- I'm so used to having a map() function in other languages!
- Capital letters denoting visibility is pretty odd. Was annoying to change from private to public. I assume there's a tool for this but I didn't figure it out quickly enough to just manually make that change
- Vim setup was very easy!
- IntelliJ has no go plugins?