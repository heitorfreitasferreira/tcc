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

### 4. `optimize <method>`

**Short Description**: Optimize one `.graph` instance with a selected method

**Methods**:

- `bruteforce`
- `ga`
- `aco`
- `pso`

**Shared Flags**:

- `--instance string`: Input graph file (default: `./data/10a.graph`)
- `-p, --population int`: Population size for population-based methods (default: `100`)
- `-i, --iterations int`: Number of iterations for population-based methods (default: `100`)
- `--results-dir string`: Base folder for structured outputs (default: `./data/results`)
- `--run-id string`: Optional explicit identifier for a run
- `--if-exists string`: Existing artifact policy: `skip|overwrite|error` (default: `skip`)
- `--progress bool`: Print improvement progress to stderr (default: `true`)

**Structured Outputs** (saved under `--results-dir`):

- `summary/<run_id>.json`: final best solution, metadata and params
- `evolution/<run_id>.jsonl`: only improvement events (one JSON record per improvement)
- `timing/<run_id>.json`: execution timing breakdown (`load_instance`, `optimize`, `serialize`, `total`)
- `logs/<run_id>.log`: recommended destination for stderr when running in batch

The `evolution` file only stores rows when the best makespan changes, including the iteration and evaluation count where it changed.

### 5. `serve`

**Short Description**: Start the RTSP results viewer web server

**Usage**: `tcc serve`

**Flags**:

- `--addr string`: Address where the web server listens (default: `:8080`)

**Behavior**:

- Renders `/` with a sidebar tree (`map -> method -> run`) and a canvas preview for map points
- Serves static files from `/css`, `/js`, and `/img`
- Uses HTMX endpoints for partial updates: `/ui/select-map`, `/ui/select-method`, `/ui/select-run`
- Uses embedded assets from `./data`: `*.graph`, `*.points`, `results/summary`, `results/evolution`, `results/timing`
- Does not embed `results/logs` files in the binary

## Example Usage

```bash
# Generate point instances and graphs with a specific seed
tcc create -s 42 -f ./experiment-data

# Generate only point instances
tcc map -s 123 -f ./point-instances

# Generate graphs from existing point instances
tcc graph -f ./point-instances

# Optimize with GA and save structured artifacts
tcc optimize ga --instance ./data/10a.graph --results-dir ./data/results

# Start the local RTSP viewer web server
tcc serve --addr :8080

# Batch execution with run_all.sh
./src/run_all.sh --method=ga --frequency=10:3,11:3 --results-dir=./src/data/results --if-exists=skip
```
