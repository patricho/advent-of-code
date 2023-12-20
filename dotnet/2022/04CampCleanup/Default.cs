public static class Day04CampCleanup
{
    public static List<string> ReadTestInput()
    {
        return new List<string>
        {
            "2-4,6-8",
            "2-3,4-5",
            "5-7,7-9",
            "2-8,3-7",
            "6-6,4-6",
            "2-6,4-8"
        };
    }


    public static void Part1()
    {
        var file = Common.ReadInput("04CampCleanup");
        //file = ReadTestInput();

        var total = 0;

        foreach (var row in file)
        {
            var rowParts = row.Split(',');

            var first = rowParts[0].Split('-');
            var firstStart = int.Parse(first[0]);
            var firstEnd = int.Parse(first[1]);

            var second = rowParts[1].Split('-');
            var secondStart = int.Parse(second[0]);
            var secondEnd = int.Parse(second[1]);

            var contained = IsContained(firstStart, firstEnd, secondStart, secondEnd);

            if (contained)
            {
                Console.WriteLine("row: " + row);
                total++;
            }

            //Console.WriteLine("first: " + firstStart + " - " + firstEnd);
            //Console.WriteLine("second: " + secondStart + " - " + secondEnd);
            //Console.WriteLine("contained: " + contained);
        }

        Console.WriteLine("total: " + total);
    }


    public static void Part2()
    {
        var file = Common.ReadInput("04CampCleanup");
        //file = ReadTestInput();

        var total = 0;

        foreach (var row in file)
        {
            var rowParts = row.Split(',');

            var first = rowParts[0].Split('-');
            var firstStart = int.Parse(first[0]);
            var firstEnd = int.Parse(first[1]);

            var second = rowParts[1].Split('-');
            var secondStart = int.Parse(second[0]);
            var secondEnd = int.Parse(second[1]);

            var overlaps = IsOverlapping(firstStart, firstEnd, secondStart, secondEnd);


            if (overlaps)
            {
                Console.WriteLine("row: " + row);
                total++;
            }
        }

        Console.WriteLine("total: " + total);
    }


    private static bool IsContained(int firstStart, int firstEnd, int secondStart, int secondEnd)
    {
        if (firstStart >= secondStart && firstEnd <= secondEnd) return true;
        else if (secondStart >= firstStart && secondEnd <= firstEnd) return true;
        else return false;
    }


    private static bool IsOverlapping(int firstStart, int firstEnd, int secondStart, int secondEnd)
    {
        if (firstStart <= secondEnd && firstEnd >= secondStart) return true;
        else if (secondStart <= firstEnd && secondEnd >= firstStart) return true;
        else return false;
    }
}
