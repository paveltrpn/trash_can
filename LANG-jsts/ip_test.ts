
/*  
    Напишите алгоритм валидации IPv4 адреса
    в Dot-десятичной нотации.
    IP является валидным, если состоит из
    четырёх октетов, ос значениями от 0 до 255
    включительно.
    На вход функции подаётся строка. на выходе
    ожидается булево значение true или false.
    
    Примеры:

    Валидные:
    1.2.3.4, 123.45.67.89

    Невалидные:
    1.2.3, 1.2.3.4.5, 123.456.78.90, 123.054.067.089

    Обратите внимание, что запись с ведущими нулями так же невалидна
*/

let numbers: string[] = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"];
let numbers_set = new Set(n);

function usage() {
    console.log("usage: node ip_test -ip <addr>");
}

/*
    function check_set_brute()

    Проверка вхождения символа в множество простым перебором.
    В Typescript нет типа чар, поэтому просто
    проверяем первый символ строки.
    Вообще хорошо бы вставить проеверку на 
    длину входящей строки, она должна быть 
    равна 1.
*/

function check_set_brute(key: string): boolean {
    for (let j = 0; j < numbers.length; j++) {
        if (key[0] == numbers[j]) {
            return true;
        }
    }
    return false;
}

/*
    function check_set_binary()

    Поиск значения в множестве бинарным поиском.
    Не реализовано.
*/

function check_set_binary(key: string): boolean {
    return false;
}

function check_set_native(key: string): boolean {
    return numbers_set.has(key[0]); 
}

function check_str(str: string): boolean {
    for (let i = 0; i < str.length; i++) {
        if (!check_set_native(str[i])) return false;
    }
    return true;
}


function check_ip(str: string): boolean {
    let ip: string[];
    let octet: number;

    ip = str.split(".");
    
    if (ip.length != 4) {
        console.log("addr is has wrong size!");
        return false;
    }

    for (let i = 0; i < 4; i++) {
        if (check_str(ip[i]) == false) {
            console.log("octet " + (i+1) + " has wrong character");
            return false;
        }

        //Проверяем ведущий ноль
        if (ip[i][0] == "0") {
            console.log("octet " + (i+1) + " has forward zero");
            return false;
        }

        octet = Number(ip[i]);

        if (octet > 255) {
            console.log("octet " + (i+1) + " greater than 255");
            return false;
        }
    }

    return true;
}

function main() {
    let l: number;

    l = process.argv.length;
    
    if (l != 4) {
        usage();
        process.exit(0);
    }

    if (process.argv[2] != "-ip") {
        usage();
        process.exit(0);
    }

    console.log(check_ip(process.argv[3]));
}

main()