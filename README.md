# Go Reloaded

This repository contains my first independent project written in Golang, **Go Reloaded**. The project is a text processing tool that performs various modifications and formatting operations on an input text file and writes the processed content into an output file.  

## Features

The tool processes text with the following transformations:

1. **Hexadecimal Conversion**  
   Converts hexadecimal numbers followed by `(hex)` to their decimal equivalents.  
   - Example:  
     Input: `1E (hex)`  
     Output: `30`

2. **Binary Conversion**  
   Converts binary numbers followed by `(bin)` to their decimal equivalents.  
   - Example:  
     Input: `10 (bin)`  
     Output: `2`

3. **Case Transformations**  
   - `(up)`: Converts the previous word to uppercase.  
     Example:  
     Input: `go (up)`  
     Output: `GO`
   - `(low)`: Converts the previous word to lowercase.  
     Example:  
     Input: `SHOUTING (low)`  
     Output: `shouting`
   - `(cap)`: Capitalizes the previous word.  
     Example:  
     Input: `bridge (cap)`  
     Output: `Bridge`

   **Multi-Word Case Transformations:**  
   If followed by a number (e.g., `(low, 2)`), transforms the specified number of words.  
   - Example:  
     Input: `so EXCITED (low, 2)`  
     Output: `so excited`

4. **Punctuation Formatting**  
   Ensures correct spacing around punctuation such as `.`, `,`, `!`, `?`, `:` and `;`.  
   - Example:  
     Input: `hello ,world!`  
     Output: `hello, world!`

   Special handling for groups like `...` and `!?`.  
   - Example:  
     Input: `Wait ... here!?`  
     Output: `Wait... here!?`

5. **Quotation Marks**  
   Adjusts single quotes (`'`) to wrap words correctly.  
   - Example:  
     Input: `Elton John said: ' the best '`  
     Output: `Elton John said: 'the best'`

6. **Indefinite Article Correction**  
   Changes `a` to `an` if the following word begins with a vowel or `h`.  
   - Example:  
     Input: `a amazing story`  
     Output: `an amazing story`

## Usage

### Input and Output Files
The tool accepts two arguments:
1. **Input File**: Contains the text to be processed.  
2. **Output File**: Where the modified text will be saved.

### Running the Program
```bash
$ go run . <input_file> <output_file>
```

### Examples
1. **Case Conversion and Article Fix**  
   Input (`sample.txt`):  
   ```
   This is a example. It is a amazing experience.
   ```
   Command:  
   ```bash
   $ go run . sample.txt result.txt
   ```
   Output (`result.txt`):  
   ```
   This is an example. It is an amazing experience.
   ```

2. **Punctuation Handling**  
   Input:  
   ```
   hello ,world ! Is this ...good ?!
   ```
   Output:  
   ```
   hello, world! Is this... good?!
   ```

## What I Learned
- **File System Operations**: Reading from and writing to files in Go.  
- **String Manipulation**: Using Go's libraries to process and modify text efficiently.  
- **Problem Solving**: Implementing complex text transformations and handling edge cases.  
- **Best Practices**: Writing clean, maintainable Go code.  

## Future Improvements
- Add support for more transformation rules.  
- Enhance error handling for file operations and input validation.  
- Implement parallel processing for large files.  

## Acknowledgments
This was my first independent project in Golang, and I'm proud to have successfully implemented it!
