# TCC

## Overview

This CLI application is designed for finding RTSP paths using bio-inspired metaheuristics. It provides several commands for creating and managing point instances and graphs.

## Installation

(Provide installation instructions specific to your project)

## Global Flags

- `-s, --seed int64`: Set a seed for random number generation (default: 0)
  - Example: `-s 42`
- `-f, --folder string`: Directory to save maps and graphs (default: "./data")
  - Example: `-f ./my-data-folder`

## Commands

### 1. `create`

**Short Description**: Save resources into files

**Usage**: `tcc create`

**Behavior**:

- Runs both `map` and `graph` subcommands
- Generates point instances
- Creates corresponding graphs
- Saves files to the specified folder

### 2. `map`

**Short Description**: Save randomly generated point instances for Traveling Salesman Problem (TSP)

**Usage**: `tcc map`

**Features**:

- Generates point instances based on predefined frequencies
- Uses global seed for reproducibility
- Saves instances to a specified folder

**Current Configuration**:

- Generates 5 instances with 50 points

### 3. `graph`

**Short Description**: Generate graphs from point instances

**Usage**: `tcc graph`

**Behavior**:

- Loads point instances from the specified folder
- Creates graphs for each point instance
- Saves graph files in the same folder

## Example Usage

```bash
# Generate point instances and graphs with a specific seed
tcc create -s 42 -f ./experiment-data

# Generate only point instances
tcc map -s 123 -f ./point-instances

# Generate graphs from existing point instances
tcc graph -f ./point-instances
```
