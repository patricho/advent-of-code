static List<int> Part1()
{
    var sums = new List<int>();
    var file = File.ReadAllLines(Path.Combine(Environment.CurrentDirectory, "input.txt"));

    var idx = 0;
    sums.Add(0);

    foreach (var line in file)
    {
        if (string.IsNullOrEmpty(line))
        {
            idx++;
            sums.Add(0);
            continue;
        }

        sums[idx] += int.Parse(line);
    }

    sums = sums.OrderByDescending(i => i).ToList();
    Console.WriteLine($"elf with the most calories: {sums[0]}");

    return sums;
}

static void Part2()
{
    var sums = Part1();
    var top3sum = sums.Take(3).Sum();
    Console.WriteLine($"top 3 elves calorie sum: {top3sum}");
}

Part2();