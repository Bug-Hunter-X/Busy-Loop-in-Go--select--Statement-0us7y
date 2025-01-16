# Busy Loop in Go `select` Statement

This repository demonstrates a common error in Go involving the `select` statement, leading to a busy loop. The program uses a `select` statement with a single case and a timer, resulting in excessive CPU consumption.

## Bug Description
The original code continuously prints an incrementing integer every second without a way to terminate the loop.  This is inefficient and consumes CPU resources unnecessarily.

## Solution
The solution demonstrates how to properly handle the `select` statement to avoid the busy-loop problem.  The solution includes a termination signal which can be sent to gracefully exit the program.
