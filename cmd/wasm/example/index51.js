// async function wasmBrowserInstantiate(wasmModuleUrl, importObject) {
//   let response = undefined;

//   if (!importObject) {
//     importObject = {
//       env: {
//         abort: () => console.log("Abort!"),
//       },
//     };
//   }

//   if (WebAssembly.instantiateStreaming) {
//     response = await WebAssembly.instantiateStreaming(fetch(wasmModuleUrl), importObject);
//   } else {
//     const fetchAndInstantiateTask = async () => {
//       const wasmArrayBuffer = await fetch(wasmModuleUrl).then((response) => response.arrayBuffer());
//       return WebAssembly.instantiate(wasmArrayBuffer, importObject);
//     };
//     response = await fetchAndInstantiateTask();
//   }

//   return response;
// }

// this code does not work
//
// WebAssembly.instantiateStreaming(fetch('my-file.wasm'))
//       .then(obj => {
//          // do something, say, print to console
//          console.log(obj.instance.exports.my_func());
//       });

// this code its works
//
// fetch('my-file.wasm').then(response =>
//   response.arrayBuffer()
// ).then(bytes =>
//   WebAssembly.instantiate(bytes)
// ).then(obj => {
//     console.log(obj.instance.exports.my_func());
// });

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

// const io = {
//   ...go.importObject,
//   env: {
//     ...go.importObject.env,
//     ...{
//       __memory_base: 0,
//       __table_base: 0,
//       memory: new WebAssembly.Memory({ initial: 1 }),
//     },
//   },
// };

const file = "51";

function run() {
  console.log(`BEFORE - FILE ${file}`);

  console.log("go.importObject", go.importObject);

  // WebAssembly.instantiateStreaming(fetch("main.wasm")).then((obj) => {
  //   console.log("here in promise");
  //   // do something, say, print to console
  //   console.log(obj.instance.exports.my_func());
  // });

  console.log("AFTER");
  fetch(`./${file}.wasm`)
    .then((response) => response.arrayBuffer())
    .then((bytes) => {
      console.log("bytes", bytes);
      return WebAssembly.instantiate(bytes, go.importObject);
    })
    .then((wasmModule) => {
      console.log("here in the last then of wasm promise");
      console.log(wasmModule);

      go.run(wasmModule.instance);
    })
    .catch((err) => {
      console.log("here in catch error", err);
    });

  // // Get the importObject from the go instance.
  // const importObject = go.importObject;

  // // Instantiate our wasm module
  // const wasmModule = await wasmBrowserInstantiate("main.wasm", importObject);

  // // Allow the wasm_exec go instance, bootstrap and execute our wasm module
  // go.run(wasmModule.instance);

  // // Call the Add function export from wasm, save the result
  // const addResult = wasmModule.instance.exports.add(24, 24);

  // // Set the result onto the body
  // document.body.textContent = `Hello World! addResult: ${addResult}`;
}

run();
