"use strict";
async function loadResourceFetch_allAsync(url) {
    let response = await fetch(url);
    return response.text();
}
async function main() {
    console.log("call - loadResourceFetch_allAsync()");
    var text = await loadResourceFetch_allAsync('http://localhost:8080/assets/foo.html');
    console.log(text);
}
main();
