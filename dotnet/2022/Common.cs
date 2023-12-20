public static class Common
{
    public static List<string> ReadInput(string folder, string file = "input.txt")
    {
        return File.ReadAllLines(Path.Combine(Environment.CurrentDirectory, folder, file)).ToList();
    }
}