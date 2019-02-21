# ubiquitous-barnacle

Go vs C++

## Challenge suites

- [Project Euler](https://projecteuler.net)

## Instructions

1. Creating the bare structure for a new Euler challenge, let's say challenge number 10, for hardcoded participants *maarten* and *harro*:

    ```sh
    ./new.sh go euler 10
    ```

    *Syntax: `./new.sh [go | cpp] [euler] [challenge #]`*

2. Building the go binary for participant *maarten*

    ```sh
    ./build.sh go euler maarten 10
    ```

    *Syntax: `./build.sh [go | cpp] [euler] [maarten | harro] [challenge #]`*

3. Running the go binary for participant *maarten*

    ```sh
    ./run.sh go euler maarten 10
    ```

    *Syntax: `./run.sh [go | cpp] [euler] [maarten | harro] [challenge #]`*
