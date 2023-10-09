
/*  Программа для подготовки шпаркалки на тест по экологии.
    Принимает на вход два параметра - имена входного и выходного файлов.
    Входной файл - исходный код HTML страницы с тестом (там в абсолютно
    безхитростной форме перечислены вопросы и ответы на них в виде простого
    HTML, без кода javascript). На выходе текстовый файл вида:
        "Строка с вопросом"
        "Строка с ответом"
        ...
    Алгоритм простой:
        - Читаем файл построчно
        - Если строка содержит метку строки с вопросом, то
            - Читаем следующую строку
            - Добавляем её в выходной вектор
        - Если строка содержит метку с троки с ответом, то
            - Читаем следующую строку
            - Добавляем её в выходной вектор
        - Пишем выходной файл построчно
*/

#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <Windows.h>
#include <clocale>

using namespace std;

void usage() {
    cout << "test_parser --in <in file> --out <out file>\n";
}

int main(int argc, char **argv) {
    std::string in_fname, out_fname;
    std::string cur_string;
    vector<std::string> out_vector;
    std::ifstream in_file;
    std::ofstream out_file;
    
    /*  
        Это должно было помоч с выводом кириллицы в консоль,
        но не помогло 
    */
    SetConsoleCP(1251);// установка кодовой страницы win-cp 1251 в поток ввода
    SetConsoleOutputCP(1251); // установка кодовой страницы win-cp 1251 в поток вывода
    /*  
        Это второй способ выводить кириллицу в консоль.
        Тоже не помогло
    */
    setlocale (LC_ALL, "Rus");

    /* Аргументов должно быть 5 штук 
        См. void usage() */
    if (argc != 5) {
        usage();
        return 1;
    }

    if (string(argv[1]) != "--in") {
        usage();
        return 1;
    } else {
        in_fname = string(argv[2]);

        if (in_fname == "") {
            std::cout << "main(): input file name can't be empty!" << "\n";
            return 1;
        }

        in_file.open(in_fname);
        /* Читаем файл построчно */
        for (;getline(in_file,cur_string);) {
            /* Если строка содержит метку строки с вопросом */
            if (cur_string.find("<div class=\"page-header\">") != std::string::npos) {
                /* Читаем следующую строку */
                getline(in_file,cur_string);
                /* Добавляем её в выходной вектор */
                out_vector.push_back(cur_string);
            }

            /* Если строка содержит метку строки с ответом */
            if (cur_string.find("data-type=\"Правильный ответ\"") != std::string::npos) {
                /* Читаем следующую строку */
                getline(in_file,cur_string);
                /* Добавляем её в выходной вектор */
                out_vector.push_back(cur_string);
            }
        }
        in_file.close();
    }

    if (string(argv[3]) != "--out") {
        usage();
        return 1;
    } else {
        out_fname = string(argv[4]);

        if (out_fname == "") {
            std::cout << "main(): output file name can't be empty!" << "\n";
            return 1;
        }

        out_file.open(out_fname);
        for (auto &str : out_vector) {
            /* Пишем выходной файл построчно */
            out_file << str << "\n" << "\n";
        }
        out_file.close();
    }

    cout << in_fname << "\n" << out_fname << "\n";

    return 0;
}