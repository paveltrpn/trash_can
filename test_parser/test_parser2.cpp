
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <Windows.h>
#include <clocale>

#include <boost/program_options.hpp>

namespace po = boost::program_options;

int main(int argc, char **argv) {
    po::options_description desc("Allowed options");
    po::variables_map vm;
    std::string in_fname, out_fname;

    desc.add_options()("help", "produce help message")
                      ("in", po::value<std::string>(), "input file name")
                      ("out", po::value<std::string>(), "outpus file name");

    
    po::store(po::parse_command_line(argc, argv, desc), vm);
    po::notify(vm);

    if (vm.count("help")) {
        std::cout << desc << "\n";
        return 1;
    }

    if (vm.count("in")) {
        in_fname = vm["in"].as<std::string>();

        if (in_fname == "") {
            std::cout << "main(): input file name can't be empty!" << "\n";
            return 1;
        }
    }

    std::cout << in_fname;

    return 0;
}