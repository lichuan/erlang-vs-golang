# Scheduler Fairness and Real-Time Performance: Golang vs Erlang

## Overview

This repository contains a comparative analysis of the scheduling fairness and real-time performance between Golang and Erlang. The tests demonstrate that **Golang's goroutine scheduler performs poorly** - it's crude, unfair, and lacks good real-time characteristics. In contrast, Erlang's scheduler proves significantly superior in both fairness and real-time performance!

## Experimental Conclusions

Based on our rigorous testing:

- **Golang's scheduler** exhibits 粗糙 (crude) implementation, demonstrating clear unfairness and poor real-time performance
- **Erlang's scheduler** outperforms Golang dramatically, showing excellent fairness and real-time characteristics

## Test Methodology

### The Core Principle
The essence of scheduler evaluation lies in testing whether concurrent processes/goroutines receive fair treatment even when some routines are consuming substantial CPU resources. Other routines should enjoy equal execution opportunities rather than being forced into delay or blocking.

### Implementation Approach
We designed two identical functions (`work1` and `work2`) in both languages, then launched multiple resource-intensive routines (`work_busy`) to compete for system resources.

### Accounting for Performance Differences
While Golang (as a static language) naturally delivers higher performance, we adjusted the number of `work_busy` routines appropriately for each language to ensure:
- Both implementations maintain similar execution rates (approximately 60-80 executions per second)
- `work1` and `work2` achieve roughly the same execution counts within the same timeframe
- The comparison focuses specifically on scheduler behavior rather than raw language performance

## Results Analysis

After running approximately 10,000 iterations:

- **Golang results**: `work1` and `work2` showed a difference of **over 100 executions**
- **Erlang results**: `work1` and `work2` differed by only **1 execution** - an astonishing result!

## Key Insights

Erlang's scheduler demonstrates true soft real-time capabilities and implements genuinely fair preemptive scheduling - areas where Golang clearly falls short. The longer Golang runs under CPU-intensive conditions, the greater the disparity between `work1` and `work2` becomes.

## Code Structure

```
/
├── golang/          # Golang implementation
│   ├── main.go      # Primary test code
│   └── README.md    # Golang-specific instructions
├── erlang/          # Erlang implementation
│   ├── scheduler.erl # Primary test code
│   └── README.md    # Erlang-specific instructions
├── results/         # Test results and analysis
└── README.md        # This file
```

## Getting Started

### Prerequisites
- Golang (version 1.16 or higher)
- Erlang/OTP (version 23 or higher)

### Running the Tests
1. Clone this repository
2. Navigate to either `golang/` or `erlang/` directory
3. Follow language-specific instructions
4. Observe the execution count differences between work1 and work2

## Conclusion

This comparison reveals fundamental differences in scheduler design philosophy. Erlang's battle-tested process scheduler, designed for telecommunications-grade reliability, demonstrates superior fairness and real-time characteristics compared to Golang's more recent goroutine scheduler, particularly under sustained CPU pressure.

## Additional Resources

For further discussion on Erlang programming concepts and scheduler internals, please refer to the original WeChat article or join the Erlang communication group mentioned in the original content.

---
*Note: This research was conducted by Teacher Li based on extensive programming experience and deep Erlang expertise. The findings highlight important considerations for developers choosing between these languages for real-time systems.*
