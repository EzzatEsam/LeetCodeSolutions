class Solution {
    public List<String> fullJustify(String[] words, int maxWidth) {
        int lineWidthUsed = 0;
        List<List<String>> linesArray = new ArrayList<> ();
        linesArray.add(new ArrayList<> ());
        List<Integer> lengthsArray = new ArrayList<> ();
        int spaces = 0;
        for (String word : words) {
            if (lineWidthUsed + word.length() + spaces <= maxWidth) {
                linesArray.get(linesArray.size() - 1).add(word);
                lineWidthUsed += word.length();
                spaces++;
            } else {
                lengthsArray.add(lineWidthUsed);
                linesArray.add(new ArrayList<> (Arrays.asList(word)));
                lineWidthUsed = word.length();
                spaces = 1;
            }  
        }
        lengthsArray.add(lineWidthUsed);
        return linesFill(linesArray, lengthsArray, maxWidth);
    }

    public static List<String> linesFill(List<List<String>> linesArray , List<Integer> lengthsArray, int maxWidth) {
        List<String> lines = new ArrayList<> ();
        for (int i = 0; i < linesArray.size() - 1; i++) {
            List<String> line = linesArray.get(i);
            int totalWords = line.size();
            int diff = maxWidth - lengthsArray.get(i);
            int innerSpaces = 0, extraSpaces = 0;
            if (totalWords > 1) {
                innerSpaces = diff / (totalWords - 1);
                extraSpaces = diff % (totalWords - 1);
            } else {
                innerSpaces = diff;
                extraSpaces = 0;
            }

            StringBuilder sb = new StringBuilder("");
            if (line.size() > 1) { 
                for (int j = 0; j < line.size() - 1; j++) {
                    sb.append(line.get(j));
                    if (extraSpaces > 0) {
                        sb.append(" ");
                        extraSpaces--;
                    }

                    for (int k = 0; k < innerSpaces; k++) {
                        sb.append(" ");
                    }
                }
                sb.append(line.get(line.size() - 1));
            } else {
                sb.append(line.get(0));
                for (int k = 0; k < innerSpaces; k++) {
                    sb.append(" ");
                }
            }
            lines.add(sb.toString());
        }

        // last line
        List<String> line = linesArray.get(linesArray.size() - 1);
        StringBuilder sb = new StringBuilder("");
        for (String word : line) {
            sb.append(word);
            if (sb.length() < maxWidth) {
                sb.append(" ");
            }
        }

        int rem = maxWidth - sb.length();
        for (int i = 0; i < rem; i++) {
            sb.append(" ");
        }
        lines.add(sb.toString());

        return lines;
    }
}