// Here’s the code for a Pipeline version of the word frequency program
// Перепиши на C++!

function read_words(file) {
  let words = read(file);
  return words;
}
function parse_words(words) {
  let wordList = [];
  let index = 0;
  let word = "";
  while (index < words.length) {
    let c = words[index];
    c = c.toLowerCase();
    if (c >= "a" && c <= "z") {
      word += c;
      index++;
      continue;
    }
    index++;
    if (word != "") {
      wordList.push(word);
    }
    word = "";
  }
  return wordList;
}

function remove_stopwords(wordList) {
  let stopWordsText = read("stopwords.txt");
  stopWordsText = stopWordsText.replace(/(\r\n|\n|\r)/gm, "");
  stopWords = stopWordsText.split(",");
  for (let i = 0; i < wordList.length; i++) {
    for (let j = 0; j < stopWords.length; j++) {
      if (wordList[i] == stopWords[j]) {
        for (let k = i; k < wordList.length-1; k++) {
          wordList[k] = wordList[k+1];
        }
        i--; // this handles two stopwords in a row
        wordList.pop();
        break;
      }
    }
  }
  return wordList;
}

function build_freqs(wordList) {
  let wordFreqs = [];
  wordFreqs.push([wordList[0],1]);
  for (let i = 1; i < wordList.length; i++) {
    for (let j = 0; j < wordFreqs.length; j++) {
      if (wordList[i] == wordFreqs[j][0]) {
        wordFreqs[j][1] += 1;
        break;
      }
    }
    wordFreqs.push([wordList[i], 1]);
  }
  return wordFreqs;
}

function sort_freqs(wordFreqs) {
  for (let i = wordFreqs.length-1; i > 0; i--) {
    for (let j = 0; j < i; j++) {
      if (wordFreqs[j][1] < wordFreqs[j+1][1]) {
        let temp = wordFreqs[j];
        wordFreqs[j] = wordFreqs[j+1];
        wordFreqs[j+1] = temp;
      }
    }
  }
  return wordFreqs;
}
\
function get_top_ten(wordFreqs) {
  let str = "";
  for (let i = 0; i < 10; i++) {
    str += wordFreqs[i][0] + ": " + wordFreqs[i][1] + "\n";
  }
  return str;
}
// main program
print(get_top_ten(sort_freqs(build_freqs(remove_stopwords
                (parse_words(read_words("text.txt")))))));