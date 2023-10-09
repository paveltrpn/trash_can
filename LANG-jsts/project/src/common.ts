
class cat_c {
    mew: string = "mew mew mew!";

    constructor () {}

    say_mew(times: number): void {
        let i: number = 0;

        for (i = 0; i < times; i++) {
            console.log("cat_c::say_mew() at " + i + "-th time!");
        }
    }
}

export {cat_c};
