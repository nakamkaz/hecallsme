using System;

public class Cscallsme
{
	public static void Main(string[] args)
	{
		String content = "";
		String CRLF = "\r\n";
		foreach (System.Collections.DictionaryEntry de in System.Environment.GetEnvironmentVariables())
		{
			content += de.Key+"="+de.Value + CRLF;
		}
		content += "==cs=="+CRLF ;
		for (int i=0; i < args.Length; i++)
		{
			content += "Arg" + i + ": " + args[i] + CRLF;
		}
		Console.WriteLine(content);
		System.IO.File.WriteAllText(@"C:\temp\cscallsme.txt", content);
	}
}
// how to build: csc.exe /target:exe /platform:anycpu cscallsme.cs

