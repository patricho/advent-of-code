Part2();

static void Part1()
{
    // The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors.
    // The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors.

    // The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
    // plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

    var plays = new Dictionary<string, Types.Play>
    {
        { "A", Types.Play.Rock },
        { "B", Types.Play.Paper },
        { "C", Types.Play.Scissors },

        { "X", Types.Play.Rock },
        { "Y", Types.Play.Paper },
        { "Z", Types.Play.Scissors },
    };

    var playScore = new Dictionary<Types.Play, int>
    {
        { Types.Play.Rock, 1 },
        { Types.Play.Paper, 2 },
        { Types.Play.Scissors, 3 },
    };

    var lines = File.ReadAllLines(Path.Combine(Environment.CurrentDirectory, "input.txt"));

    /*var lines = new List<string>
    {
        "A Y",
        "B X",
        "C Z"
    };*/

    var totalScore = 0;

    foreach (var line in lines)
    {
        var linePlay = line.Split(' ').Select(s => plays[s]);
        var them = linePlay.First();
        var me = linePlay.Last();

        var result = GetResult(them, me);

        var linePlayScore = playScore[me];

        var lineScore = (int)result + linePlayScore;
        totalScore += lineScore;

        Console.WriteLine($"{line}: {lineScore}");
    }

    Console.WriteLine($"totalScore: {totalScore}");
}

static void Part2()
{
    // "Anyway, the second column says how the round needs to end: X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
    // The total score is still calculated in the same way, but now you need to figure out what shape to choose so the round ends as indicated.

    var plays = new Dictionary<string, Types.Play>
    {
        { "A", Types.Play.Rock },
        { "B", Types.Play.Paper },
        { "C", Types.Play.Scissors }
    };

    var outcomes = new Dictionary<string, Types.Result>
    {
        { "X", Types.Result.Lose },
        { "Y", Types.Result.Draw },
        { "Z", Types.Result.Win }
    };

    var playScore = new Dictionary<Types.Play, int>
    {
        { Types.Play.Rock, 1 },
        { Types.Play.Paper, 2 },
        { Types.Play.Scissors, 3 },
    };

    var lines = File.ReadAllLines(Path.Combine(Environment.CurrentDirectory, "input.txt"));

    /*var lines = new List<string>
    {
        "A Y",
        "B X",
        "C Z"
    };*/

    var totalScore = 0;

    foreach (var line in lines)
    {
        var linePlay = line.Split(' ');
        var them = plays[linePlay[0]];
        var outcome = outcomes[linePlay[1]];

        var me = GetPlay(them, outcome);

        var result = GetResult(them, me);

        var linePlayScore = playScore[me];

        var lineScore = (int)result + linePlayScore;
        totalScore += lineScore;

        Console.WriteLine($"{line}: {lineScore}");
    }

    Console.WriteLine($"totalScore: {totalScore}");
}

static Types.Result GetResult(Types.Play them, Types.Play me)
{
    if (them == me) return Types.Result.Draw;

    if (them == Types.Play.Paper)
    {
        switch (me)
        {
            case Types.Play.Rock: return Types.Result.Lose;
            case Types.Play.Scissors: return Types.Result.Win;
        }
    }
    else if (them == Types.Play.Rock)
    {
        switch (me)
        {
            case Types.Play.Paper: return Types.Result.Win;
            case Types.Play.Scissors: return Types.Result.Lose;
        }
    }
    else if (them == Types.Play.Scissors)
    {
        switch (me)
        {
            case Types.Play.Paper: return Types.Result.Lose;
            case Types.Play.Rock: return Types.Result.Win;
        }
    }

    throw new ArgumentOutOfRangeException();
}

static Types.Play GetPlay(Types.Play them, Types.Result outcome)
{
    if (outcome == Types.Result.Draw) return them;

    if (them == Types.Play.Paper)
    {
        switch (outcome)
        {
            case Types.Result.Lose: return Types.Play.Rock;
            case Types.Result.Win: return Types.Play.Scissors;
        }
    }
    else if (them == Types.Play.Rock)
    {
        switch (outcome)
        {
            case Types.Result.Lose: return Types.Play.Scissors;
            case Types.Result.Win: return Types.Play.Paper;
        }
    }
    else if (them == Types.Play.Scissors)
    {
        switch (outcome)
        {
            case Types.Result.Lose: return Types.Play.Paper;
            case Types.Result.Win: return Types.Play.Rock;
        }
    }

    throw new ArgumentOutOfRangeException();
}

namespace Types
{
    enum Play
    {
        Rock,
        Paper,
        Scissors
    }

    enum Result
    {
        Lose = 0,
        Draw = 3,
        Win = 6
    }
}
