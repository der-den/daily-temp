# daily-temp
Remove files, older than n days. Perfect to keep your personal temp or download folder clean. 
Use windows task scheduler to run it daily.

Usage: `daily-temp <days> <directory>`

Sample: `daily-temp 5 c:\mytemp`
        `daily-temp 5 c:\mytemp <logfile.txt>`

- Delete all files/dirs/sub-dirs in the directory "mytemp" with a CREATION time older than 5 days.
- The logfile is optional.

Hint: code tested on windows only
 