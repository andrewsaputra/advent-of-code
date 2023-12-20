import java.io.BufferedReader;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.function.BiFunction;

public class Day1 {
    private static void print(Object o) {
        System.out.println(o);
    }

    public static void main(String[] args) throws Exception {
        long startTime = System.currentTimeMillis();

        Path inputFile = Path.of("2023/java/inputs.txt");

        int res1 = solve(inputFile, false);
        print("Part 1: " + res1);

        int res2 = solve(inputFile, true);
        print("Part 2: " + res2);

        print(String.format("Duration: %dms", System.currentTimeMillis() - startTime));
    }

    private static int solve(Path inputFile, boolean isAlphanumeric) throws Exception {
        int sum = 0;
        try (BufferedReader bf = Files.newBufferedReader(inputFile)) {
            String line;
            while ((line = bf.readLine()) != null) {
                sum += findFirstDigit(line, isAlphanumeric, true) * 10;
                sum += findFirstDigit(line, isAlphanumeric, false);
            }
        }

        return sum;
    }

    private static int findFirstDigit(String line, boolean isAlphanumeric, boolean scanForward) {
        int length = line.length();
        int idx = 0;
        int modIdx = 1;
        if (!scanForward) {
            idx = length - 1;
            modIdx = -1;
        }

        while (idx >= 0 && idx < length) {
            char c = line.charAt(idx);
            if (c >= '0' && c <= '9') {
                return c - '0';
            }

            if (isAlphanumeric) {
                int digit;
                if ((digit = getAlphanumericDigit(line, idx)) != -1) {
                    return digit;
                }
            }

            idx += modIdx;
        }

        return -1;
    }

    private static int getAlphanumericDigit(String line, int idx) {
        int length = line.length();

        BiFunction<Integer, String, Boolean> doMatch = (startIdx, matcher) -> {
            int endIdx = startIdx + matcher.length();
            return endIdx <= length && matcher.equals(line.substring(startIdx, endIdx));
        };

        switch (line.charAt(idx)) {
            case 'o':
                if (doMatch.apply(idx, "one")) {
                    return 1;
                }
                break;
            case 't':
                if (doMatch.apply(idx, "two")) {
                    return 2;
                }
                if (doMatch.apply(idx, "three")) {
                    return 3;
                }
                break;
            case 'f':
                if (doMatch.apply(idx, "four")) {
                    return 4;
                }
                if (doMatch.apply(idx, "five")) {
                    return 5;
                }
                break;
            case 's':
                if (doMatch.apply(idx, "six")) {
                    return 6;
                }
                if (doMatch.apply(idx, "seven")) {
                    return 7;
                }
                break;
            case 'e':
                if (doMatch.apply(idx, "eight")) {
                    return 8;
                }
                break;
            case 'n':
                if (doMatch.apply(idx, "nine")) {
                    return 9;
                }
                break;
        }

        return -1;
    }
}
