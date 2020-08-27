public class App {
    public static void main(String[] args) {
        System.out.println("Absolute value of -30: " + absolute(-30));
        System.out.println("Absolute value of -30.5: " + absolute(-30.5));
        System.out.println("Binomial Coefficient of 5,3: " + binomialCoefficient(5, 3));
        printGraph(5);
    }

    public static int absolute(int num) {
        if (num <= 0) // This will see if num is negative and flip to positive if so
            return (num * -1);
        else
            return (num);
    }

    public static double absolute(double num) { // overloaded method
        if (num <= 0)
            return (num * -1);
        else
            return (num);
    }

    public static int binomialCoefficient(int n, int r) {
        int i = 0;
        int result = 1;
        while (i < r) {
            result *= n - i;
            result /= i + 1;
            i++;
        }
        return (result);

    }

    public static void printGraph(int size) {
        if (size < 1) { // This is so any number less than 1 will return nothing
            return;
        } else {
            for (int e = -size; e <= size; e++) { // this is the counter for the columns
                if (e < 0 || e > 0) // Since the graph does same thing except for at 0
                    for (int t = -size; t <= 0; t++) { // This is for the individual row
                        if (t < 0 && t != -size) // More complex here since the amount of spaces varies
                            System.out.print("    ");
                        else if (t < 0 && t == -size)
                            System.out.print("      ");
                        else
                            System.out.print(" |\n\n");
                    }
                else
                    for (int x = -size; x <= size; x++) {
                        if (-size == x)
                            System.out.print(" - | -"); // At 0 first print needed an extra "-" at the beginning
                        else if (size == x)
                            System.out.print(" | -\n\n");
                        else
                            System.out.print(" | -");
                    }
            }
        }
    }
}