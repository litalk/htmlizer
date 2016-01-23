# Htmlizer #
Htmlizer is a tool which can translate Junit XML report to HTML.

## Output ##
The style of output HTML is based on Maven HTML report. It looks like this [sample_junit.html](https://github.com/wu8685/htmlizer/blob/master/sample_junit.html).

## Input ##
It has two parameters like following.
```
-i    The path of original Junit XML. It could point to a directory or a xml file.
-o    The path of output directory. The directory will be created if it doesn't exist.
```
The output HTML file names is almostly the same as the input XML file, and which would be followed by an extra "_1" if a file with the same name has existed. 
