import java.io.BufferedReader;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Function;
import java.util.function.Predicate;

public class Day2 {

    private static void print(Object o) {
        System.out.println(o);
    }

    public static void main(String[] args) throws Exception {
        long startTime = System.currentTimeMillis();

        Path inputFile = Path.of("2023/java/inputs.txt");

        int res1 = solvePart1(inputFile);
        print("Part 1: " + res1);

        int res2 = solvePart2(inputFile);
        print("Part 2: " + res2);

        print(String.format("Duration: %dms", System.currentTimeMillis() - startTime));
    }

    public record Game(
            int id,
            List<Map<String, Integer>> records) {
    }

    private static int solvePart1(Path inputFile) throws Exception {
        Predicate<Game> gameChecker = game -> {
            Map<String, Integer> maxBalls = Map.of("red", 12, "green", 13, "blue", 14);
            for (Map<String, Integer> record : game.records()) {
                for (Map.Entry<String, Integer> entry : record.entrySet()) {
                    String color = entry.getKey();
                    int amount = entry.getValue();
                    if (amount > maxBalls.get(color)) {
                        return false;
                    }
                }
            }

            return true;
        };

        int sum = 0;
        try (BufferedReader bf = Files.newBufferedReader(inputFile)) {
            String line;
            while ((line = bf.readLine()) != null) {
                Game game = lineToGame(line);
                if (gameChecker.test(game)) {
                    sum += game.id();
                }
            }
        }
        return sum;
    }

    private static int solvePart2(Path inputFile) throws Exception {
        Function<Game, Integer> powerCalculator = game -> {
            Map<String, Integer> maxBalls = new HashMap<>();
            for (Map<String, Integer> record : game.records()) {
                for (Map.Entry<String, Integer> entry : record.entrySet()) {
                    String color = entry.getKey();
                    int amount = entry.getValue();
                    Integer currMax;
                    if ((currMax = maxBalls.get(color)) == null || amount > currMax) {
                        maxBalls.put(color, amount);
                    }
                }
            }

            int total = 1;
            for (Map.Entry<String, Integer> entry : maxBalls.entrySet()) {
                total *= entry.getValue();
            }
            return total;
        };

        int sum = 0;
        try (BufferedReader bf = Files.newBufferedReader(inputFile)) {
            String line;
            while ((line = bf.readLine()) != null) {
                Game game = lineToGame(line);
                sum += powerCalculator.apply(game);
            }
        }
        return sum;
    }

    private static Game lineToGame(String line) {
        String[] data = line.split(": ");
        int gameId = Integer.parseInt(data[0].split(" ")[1]);

        List<Map<String, Integer>> records = new ArrayList<>();
        for (String roundStr : data[1].split("; ")) {
            Map<String, Integer> roundRecord = new HashMap<>();
            for (String ballStr : roundStr.split(", ")) {
                String[] balls = ballStr.split(" ");
                int amount = Integer.parseInt(balls[0]);
                String color = balls[1];
                roundRecord.put(color, amount);
            }

            records.add(roundRecord);
        }

        return new Game(gameId, records);
    }
}
