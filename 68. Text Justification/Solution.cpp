class Solution {
public:
    vector<string> fullJustify(vector<string>& words, int maxWidth) {
        int lineWidthUsed = 0;
        vector<vector<string>> linesArray;
        linesArray.push_back(vector<string> ());
        vector<int> lengthsArray;
        int spaces = 0;
        for (string word : words) {
            if (lineWidthUsed + word.length() + spaces <= maxWidth) {
                linesArray.back().push_back(word);
                lineWidthUsed += word.length();
                spaces++;
            } else {
                lengthsArray.push_back(lineWidthUsed);
                linesArray.push_back(vector<string> ({word}));
                lineWidthUsed = word.length();
                spaces = 1;
            }  
        }
        lengthsArray.push_back(lineWidthUsed);
        return linesFill(linesArray, lengthsArray, maxWidth);
    }

    static vector<string> linesFill(vector<vector<string>>& linesArray , vector<int>& lengthsArray, int maxWidth) {
        vector<string> lines;
        for (int i = 0; i < linesArray.size() - 1; i++) {
            vector<string> line = linesArray[i];
            int totalWords = line.size();
            int diff = maxWidth - lengthsArray[i];
            int innerSpaces = 0, extraSpaces = 0;
            if (totalWords > 1) {
                innerSpaces = diff / (totalWords - 1);
                extraSpaces = diff % (totalWords - 1);
            } else {
                innerSpaces = diff;
                extraSpaces = 0;
            }

            string sb = "";
            if (line.size() > 1) { 
                for (int j = 0; j < line.size() - 1; j++) {
                    sb += line[j];
                    if (extraSpaces > 0) {
                        sb += " ";
                        extraSpaces--;
                    }

                    for (int k = 0; k < innerSpaces; k++) {
                        sb += " ";
                    }
                }
                sb += line.back();
            } else {
                sb += line[0];
                for (int k = 0; k < innerSpaces; k++) {
                    sb += " ";
                }
            }
            lines.push_back(sb);
        }

        // last line
        vector<string> line = linesArray.back();
        string sb = "";
        for (string word : line) {
            sb += word;
            if (sb.length() < maxWidth) {
                sb += " ";
            }
        }

        int rem = maxWidth - sb.length();
        for (int i = 0; i < rem; i++) {
            sb += " ";
        }
        lines.push_back(sb);

        return lines;
    }
};