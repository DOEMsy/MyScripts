Set WshShell= WScript.CreateObject("WScript.Shell")

WshShell.AppActivate "’‘º““„"

for i=1 to 10000

WshShell.SendKeys "^v"

WshShell.SendKeys i

WshShell.SendKeys "%s"

Next