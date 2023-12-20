public static class Day03RucksackReorganization
{
    public static void Part2()
    {
        var file = Common.ReadInput("03RucksackReorganization");
        // file = ReadTestInput;

        var total = 0;

        for (var rowidx = 0; rowidx < file.Count; rowidx += 3)
        {
            Console.WriteLine("grp " + rowidx);
            Console.WriteLine("line 0 " + file[rowidx]);
            Console.WriteLine("line 1 " + file[rowidx + 1]);
            Console.WriteLine("line 2 " + file[rowidx + 2]);

            var common = FindCommon(file[rowidx], file[rowidx + 1], file[rowidx + 2]);
            var score = GetCharScore(common);

            total += score;

            Console.WriteLine("common: " + common);
            Console.WriteLine("score: " + score);
        }

        Console.WriteLine("total: " + total);
    }


    public static List<string> ReadTestInput()
    {
        return new List<string>
        {
            "vJrwpWtwJgWrhcsFMMfFFhFp",
            "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
            "PmmdzqPrVvPwwTWBwg",
            "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
            "ttgJtRGJQctTZtZT",
            "CrZsJsPPZsGzwwsLwLmpwMDw"
        };
    }


    public static void Part1()
    {
        var file = Common.ReadInput("03RucksackReorganization");
        // file = ReadTestInput;

        var total = 0;

        foreach (var row in file)
        {
            var rowParts = SplitRow(row);

            Console.WriteLine("input: " + row);
            Console.WriteLine("rowParts[0]: " + rowParts[0]);
            Console.WriteLine("rowParts[1]: " + rowParts[1]);

            var common = FindCommon(rowParts);
            var score = GetCharScore(common);

            total += score;

            Console.WriteLine("common: " + common);
            Console.WriteLine("score: " + score);
        }

        Console.WriteLine("total: " + total);
    }


    private static int GetCharScore(char common)
    {
        var lowercase = "abcdefghijklmnopqrstuvwxyz";
        var uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

        var idxl = lowercase.IndexOf(common);
        var idxu = uppercase.IndexOf(common);

        /*
        To help prioritize item rearrangement, every item type can be converted to a priority:

        Lowercase item types a through z have priorities 1 through 26.
        Uppercase item types A through Z have priorities 27 through 52.
        */

        if (idxl >= 0) return idxl + 1;
        if (idxu >= 0) return idxu + 27;

        return -1;
    }


    private static char FindCommon(string[] rowParts)
    {
        var first = rowParts[0];
        var second = rowParts[1];

        for (int i = 0; i < first.Length; i++)
        {
            for (int j = 0; j < second.Length; j++)
            {
                if (first[i] == second[j]) return first[i];
            }
        }

        return default;
    }


    private static char FindCommon(string first, string second, string third)
    {
        for (int i = 0; i < first.Length; i++)
        {
            for (int j = 0; j < second.Length; j++)
            {
                for (int k = 0; k < third.Length; k++)
                {
                    if (first[i] == second[j] && second[j] == third[k])
                    {
                        return first[i];
                    }
                }
            }
        }

        return default;
    }


    private static string[] SplitRow(string row)
    {
        var half = (int)(row.Length / 2m);

        var first = row[..half];
        var second = row[half..];

        return new[]
        {
            first,
            second
        };
    }
}
