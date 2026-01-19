# Anki flashcard import file format

This document describes the required structure and formatting rules for a plain text file that can be imported into Anki as flashcards. It is based on the official documentation for Text Files in Anki.

## Supported File Types

Anki accepts plain text files with `.txt` extension for import.

## Field Separators

Each line in the file represents one note (card).
Fields on each line must be separated by a pipe (`|`).

## Basic Structure

Each line corresponds to a single note. Each note/flashcard consists of two fields (e.g., Question, Answer).
Fields are ordered left-to-right:

<Field:Question><sep><Field:Answer>

Example with pipes:

Question 1|Answer 1
Question 2|Answer 2

The number of fields per line must be consistent; lines with missing fields will pad missing fields as empty, while extra fields beyond the expected count will be ignored.

## Comments

Lines beginning with # are ignored by the importer unless they are recognized as file headers.
Example: `# This is a comment and will be ignored`

## Multiline and Special Content

### Multiline Fields 
You must use HTML newline tags (`<br>`) to represent line breaks within fields.

Example with pipes:
```txt
Question 1|this is line 1.<br>this is line 2.
Question 2|Answer 2
```

### Including Delimiters Inside Fields

If the field text contains a delimiter character, you must escape it:
Use double quotes for CSV style (e.g., "value with | inside").

## Optional Header Directives

Anki supports special file header directives starting with a # at the top of the file:

#separator:Pipe
#html:true
#tags:Tag1 Tag2
#columns:Front|Back|Tags
#notetype:Basic
#deck:MyDeckName


Header keys and their meanings:

Directive	Meaning
separator:	Force a specific delimiter (Comma, Semicolon, Tab, Pipe, Colon, Space).
html:	true or false, controls HTML parsing in fields.
tags:	A space-separated list of tags added to every note.
columns:	Field names matching the number of columns present.
notetype:	Specify a default note type to use.
deck:	Specify the deck name where notes should be imported.

Example with headers:

#separator:Pipe
#html:false
#tags:Language Vocabulary
Question|Answer|Extra


Note: These header lines must appear before the first actual data line.

## Tags

Tags can be included in the imported fields and mapped in the import dialog.
If the file includes a separate “tags” field, you can assign it to Anki’s tags field during field mapping.
