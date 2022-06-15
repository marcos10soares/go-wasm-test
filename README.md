# go-wasm-test

Golang webassembly experiment, the project idea was to call golang functions from JS, the ideal test candidate for this project was the pkg `zxcvbn` which has different implementations with different results depending on the language that's implemented with ([issue: Score deviates from the coffee/javascript version](https://github.com/nbutton23/zxcvbn-go/issues/20)). As IMO the golang [zxcvbn-go](https://github.com/nbutton23/zxcvbn-go) pkg has the best implementation and results, the idea was to allow the frontend to use the exact pkg to avoid discrepancies between frontend and backend password strength score.

**Example: [demo](https://marcos10soares.github.io/go-wasm-test/assets/index.html)**

WASM Browser support: [https://caniuse.com/wasm](https://caniuse.com/wasm)


## Building
```bash
GOOS=js GOARCH=wasm go build -o assets/zxcvb.wasm cmd/zxcvb/main.go
```


## Serving the files

```bash
cd assets
python3 -m http.server 8080
```

and visit [http://localhost:8080](http://localhost:8080).

### Notes:

The file `wasm_exec.js` is copied directly from GO root by doing:
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets/
```

This is the code block responsible for fetching and exectuing the wasm binary:
```javascript
const go = new Go();
WebAssembly.instantiateStreaming(fetch("zxcvb.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
```

After running this code block, you can access the golang methods from JS.

