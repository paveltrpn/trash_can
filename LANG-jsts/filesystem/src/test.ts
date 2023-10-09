
import * as fs from "fs"
import * as rng from "./range";
import * as smnp from "./string_mnp";

export function testfreadSync() {
    console.log("main(): fs test.")

    // let foo: freadSync = new freadSync();
    // foo.printFileStats("../assets/raven.txt");

    let fileContent = fs.readFileSync("../assets/raven.txt", "utf8")
    let words = smnp.splitBySpace(fileContent);

    for (let i = 0; i < 5; i++) {
        console.log(words[i])
    }

    console.log(smnp.removePunctuation(fileContent))
}

/* ============ ENTRY POINT ================ */
/* ============ function main() ============ */
(() => { 
    testfreadSync();
})();
/* ============ =============== ============ */