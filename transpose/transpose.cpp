
#include <iostream>
#include <vector>
#include <tuple>

std::vector<std::string> split_str(const std::string &string, char sep) {
    std::vector<std::string> list;
    std::string::size_type start = 0;
    std::string::size_type end;
    while ((end = string.find(sep, start)) != std::string::npos) {
        if (start != end)
            list.push_back(string.substr(start, end - start));
        start = end + 1;
    }
    if (start != string.size())
        list.push_back(string.substr(start));
    return list;
}

class transpose_c {
    private:
        const std::vector<std::tuple<std::string, std::string, int>>  scale =    {{{"A", "A", 0},
                                                                                   {"A#","Bb",1},
                                                                                   {"B", "B", 2},
                                                                                   {"C", "C", 3},
                                                                                   {"C#","Db",4},
                                                                                   {"D", "D", 5},
                                                                                   {"D#","Eb",6},
                                                                                   {"E", "E", 7},
                                                                                   {"F", "F", 8},
                                                                                   {"F#","Gb",9},
                                                                                   {"G", "G", 10},
                                                                                   {"G#","Ab",11}}};
    
        int             get_tone(const std::string &chord);
        std::string     get_chord(const int tone);

    public:
        transpose_c() {};
        ~transpose_c() {};
        std::string     transpose(const std::vector<std::string> &harm, const int step);
};

int transpose_c::get_tone(const std::string &chord) {
    for (auto i:scale) {
        if ((std::get<0>(i) == chord) || (std::get<1>(i) == chord)) {
            return std::get<2>(i);
        }
    }

    return 0;
}

std::string transpose_c::get_chord(const int tone) {
    std::string ret;

    if ((tone < 0) || (tone > 11)) {
        //therse no chord below 0 and upper 11
        return "none";
    }
    
    for (auto i:scale) {
        if (std::get<2>(i) == tone) {
            return std::get<0>(i) + "(" + std::get<1>(i) + ")";
        }
    }

    return ret;
}

std::string transpose_c::transpose(const std::vector<std::string> &harm, const int step) {
    int shift = 0;
    std::vector<int> tones;
    std::string ret;

    for (auto i:harm){
        tones.push_back(get_tone(i));
    }
    
    for (auto j:tones) {
        if (step >= 0) {
            shift = (j + step) % 12;
        } else if (step < 0) {
            shift = (12 + (j + step)) % 12;
        }

        ret = ret + get_chord(shift) + " ";
    }

    return ret;
}

void usage() {
    std::cout << "Program transpose chord progression \n" <<
				"Usage: main -c CHORD1:CHORD2: ... :CHORDn -t X \n" << 
                "where CHORD is A,B,C  ... \n" <<
                ": - separator \n" << 
                "X is a tone shift \n" << std::endl;
}

int main(int argc, char **argv) {
    transpose_c tsp;
    std::vector<std::string> harm;
    int step = 0;

    if (argc == 1) 
	{
		usage();
		return 0;
	}

    for (int i = 1; i < argc; i++) {
		if (std::string(argv[i]) == "-h") {
            usage();
            return 0;
        }

        if (std::string(argv[i]) == "-s") {
            step = std::stoi(argv[i+1]);
        }

        if (std::string(argv[i]) == "-c") {
            harm = split_str(argv[i+1],':');
        }
	}
    
    std::cout << "transposing chords: ";

    for (auto i :harm) std::cout << " " << i;

    std::cout << "\nto step:  " << step << "\n\n";

    std::cout << tsp.transpose(harm,step) << "\n";

    return 0;
}