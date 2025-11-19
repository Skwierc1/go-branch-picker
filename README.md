# Git Branch Picker

A simple interactive command-line tool for selecting and switching between Git branches.

## Description

Git Branch Picker (`gbp`) provides an interactive menu to quickly browse, search, and switch between Git branches. It displays all available branches in a searchable list and allows you to create new branches on the fly.

## Features

- Interactive branch selection with a searchable menu
- Quick branch switching
- Create new branches directly from the interface
- Search functionality to filter branches by name

## Usage

Run the program to see an interactive menu of all available Git branches:

```bash
./gbp
```

Use the arrow keys to navigate, type to search, and press Enter to select a branch or action.

## Installation

Build the program:

```bash
go build -o gbp main.go
```

## Preference

Copy binary to /usr/bin and run program directly.

```bash
cp gbp /usr/bin/
```