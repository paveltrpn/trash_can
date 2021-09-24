"use strict";
var global_cntnr = null;
function global_callback(src) {
    global_cntnr = src;
}
function loadResourceXHR_sync(url) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, false);
    xhr.onreadystatechange = function () {
        if (xhr.readyState != 4) {
            return;
        }
        if (xhr.status === 200) {
            global_callback(xhr.responseText);
        }
        else {
            global_callback(xhr.responseText);
        }
    };
    xhr.send();
}
function main() {
    console.log("call - loadResourceXHR_sync()");
    loadResourceXHR_sync('http://localhost:8080/assets/foo.html');
}
