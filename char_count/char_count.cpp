/*
Алгоритм :

1) Проверяете что Ваш файл есть, и открываете его.

2) Заводите std::map для подсчета вхождений символов и заполняете его.

3) Создаете csv файл и дампите туду содержимое map
*/

#include <iostream>
#include <fstream>
#include <filesystem>
#include <map>
#include <algorithm>
#include <iterator>

std::map<unsigned char, std::size_t> CalcSymbls(std::filesystem::path f_path) {

    if (!std::filesystem::exists(f_path)) {
        throw std::runtime_error{"File not found!"};
    }

    std::ifstream in_stream(f_path, std::ios::out);
    if (!in_stream.is_open()) {
        throw std::runtime_error{"Error open file!"};
    }

    std::map<unsigned char, std::size_t> symbls;        
    auto f_obj{[&](unsigned char ch){ symbls[ch]++;}};    
    std::for_each(std::istream_iterator<char>(in_stream), std::istream_iterator<char>(), f_obj);

    return symbls;
}


void WriteCSV(std::filesystem::path out_file_path, std::map<unsigned char, std::size_t> data) {

    std::ofstream out_csv_stream(out_file_path, std::ios::out);

    if (!out_csv_stream.is_open()) {
        throw std::runtime_error{"Error open file"};
    }

    auto f_obj{[&](std::pair<unsigned char, std::size_t> val){ out_csv_stream << val.first << ";" << val.second << "\n";}};
    std::for_each(std::begin(data), std::end(data), f_obj);
}


int main() {

    std::string file_name{"data"};
    auto file_path{std::filesystem::current_path() / file_name};
    auto symbls{CalcSymbls(file_path)};

    std::string out_name{"out.csv"};
    auto out_file_path{std::filesystem::current_path() / out_name};
    WriteCSV(out_file_path, symbls);

    return 0;
}