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


Start to run golang:
![golang start](./res/golang_start.gif)
<br><br><br>
Start to run Erlang:
![erlang start](./res/erlang_start.gif)

## Results Analysis

After running approximately 10,000 iterations:

- **Golang results**: `work1` and `work2` showed a difference of **over 100 executions**
- **Erlang results**: `work1` and `work2` differed by only **1 execution** - an astonishing result!

Result of golang:  

![golang end](./res/golang_end.png)
<br><br><br>
Result of erlang:  

![erlang end](./res/erlang_end.png)

## Key Insights

Erlang's scheduler demonstrates true soft real-time capabilities and implements genuinely fair preemptive scheduling - areas where Golang clearly falls short. The longer Golang runs under CPU-intensive conditions, the greater the disparity between `work1` and `work2` becomes.

## Conclusion

This comparison reveals fundamental differences in scheduler design philosophy. Erlang's battle-tested process scheduler, designed for telecommunications-grade reliability, demonstrates superior fairness and real-time characteristics compared to Golang's more recent goroutine scheduler, particularly under sustained CPU pressure.
