
#include <iostream>
#include <fstream>
#include <vector>
#include <string>

using namespace std;

void usage() {
	cout << "Please, place usage here later" << "\n" << endl;
}

int32_t id_rw(int32_t i, int32_t j, int32_t n) {
	return (i*n + j);
};

vector<string> split_str(const string &str, char sep) {
    vector<string> list;
    string::size_type start = 0;
    string::size_type end;

    while ((end = str.find(sep, start)) != string::npos) {
        if (start != end) {
            list.push_back(str.substr(start, end - start));
		}
        start = end + 1;
    }

    if (start != str.size()) {
        list.push_back(str.substr(start));
	}

    return list;
}

string parse_str(const string &str) {
	string lstr = str;
	string rw_args;
	vector<string> sints; 
	vector<int> ints;
	int npos = 0;

	npos = str.find("id_rw");

	if (npos == -1) {
		return str;
	}

	rw_args = lstr.substr(npos+5, 9);

	for (auto &c: rw_args) {
		if ((c == '(') || (c == ')') || (c == ',')) {
			//&c - &rw_args[0] - индекс итерируемого элемента
			rw_args.erase(&c - &rw_args[0],1);
		} 
	}

	sints = split_str(rw_args, ' ');

	for (auto &i: sints) {
		ints.push_back(std::stoi(i));
	}
	
	lstr.erase(npos, 14);

	lstr.insert(npos, std::to_string(id_rw(ints[0], ints[1], ints[2])));

	lstr = parse_str(lstr);

	return lstr; 
}

int main(int argc, char **argv) {
	ifstream pfile;
	ofstream ofile;
	vector<string> file;
	string str;

	if ((argc != 2)) {
		cout << "wrong count of args!" << "\n" << endl;
		usage();
		return 0;
	}
	
	if (std::string(argv[1]) == "-h") {
		usage();
		return 0;
	}
	
	pfile.open(argv[1]);

	if (pfile.is_open()) {
		while (!pfile.eof()) {
			getline(pfile, str);
			file.push_back(parse_str(str));
		}

	} else {
		cout << "no file " << argv[1] << "\n";
		return 0;
	}

//	show file
	ofile.open("out.cpp");
	for (auto i: file) {
		cout << i << "\n";
		ofile << i << "\n";
	}

	pfile.close();
	ofile.close();
			
	return 0;
}