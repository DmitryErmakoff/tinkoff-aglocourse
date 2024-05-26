using System;

class Program
{
    static int Gcd(int a, int b)
    {
        while (b != 0)
        {
            int temp = b;
            b = a % b;
            a = temp;
        }
        return a;
    }

    static int Lcm(int a, int b)
    {
        return (a * b) / Gcd(a, b);
    }

    static void Main()
    {
        string input = Console.ReadLine();
        string[] values = input.Split();
        int N = int.Parse(values[0]);
        int K = int.Parse(values[1]);

        int lcmValue = Lcm(N, K);

        Console.WriteLine(lcmValue);
    }
}