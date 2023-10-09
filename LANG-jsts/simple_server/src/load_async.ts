
function loadResourceXHR(url: string, callback: any) {
    var xhr = new XMLHttpRequest()
    xhr.open('GET', url)
    
    xhr.onreadystatechange = function() {
        if (xhr.readyState != 4) {
            return
        }

        if (xhr.status === 200) {
            callback(xhr.responseText);
        } else {
            callback(xhr.responseText);
        }
    }

    xhr.send()
}

async function loadResourceFetch(url: string) {
    let response = await fetch(url);
        
    let commits = await response.text();

    console.log("call - loadResourceFetch()");
    console.log(commits);
}

(function main() {
    let foo_callback = function(text: string) {
        console.log("call - loadResourceXHR()");
        console.log(text);
    };
    
    loadResourceXHR('http://localhost:8080/assets/foo.html', foo_callback);
    
    loadResourceFetch('http://localhost:8080/assets/foo.html')
})