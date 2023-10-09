function readTextFile(file, callback) {
    var rawFile = new XMLHttpRequest();
    rawFile.overrideMimeType("application/text");
    rawFile.open("GET", file, true);
    rawFile.onreadystatechange = function() {
        if (rawFile.readyState === 4 && rawFile.status == "200") {
            callback(rawFile.responseText);
        }
    }
    rawFile.send(null);
}

readTextFile("/Users/Documents/workspace/test.json", function(text){
  var data = JSON.parse(text);
  console.log(data);
  
function fetchLocal(url) {
    return new Promise(function(resolve, reject) {
      var xhr = new XMLHttpRequest
      xhr.onload = function() {
        resolve(new Response(xhr.responseText, {status: xhr.status}))
      }
      xhr.onerror = function() {
        reject(new TypeError('Local request failed'))
      }
      xhr.open('GET', url)
      xhr.send(null)
    })
}

// Replace ./data.json with your JSON feed
fetch('./data.json').then(response => {
    return response.json();
  }).then(data => {
    // Work with JSON data here
    console.log(data);
  }).catch(err => {
    // Do something for an error here
  });
  
function webglBookLoad(url) {
    var request = new XMLHttpRequest();

    request.onreadystatechange = function() {
        if (request.readyState === 4 && request.status !== 404) {
            console.log(request.responseText);
        }
    }

    request.open('GET', url, true)
    request.send(null)
}
