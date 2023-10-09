

let process = require("process")

/**
 * Prints a line (if input is string) in stdout without a trailing new line.
 * If str is not a string we do coomon console.log() call with trailing new line.
 * @param {string | any} str 
 */
function println(str) {
    if (typeof str === "string") {
        process.stdout.write(str)
    } else {
        console.log(str)
    }
}

(function main() {
    println("foobar\n")

    let bar = [100, 101, 102, "one houndred", false]

    println(bar)

    for (let vr of bar) {
        println(vr+"\n")
    }
}) ()
