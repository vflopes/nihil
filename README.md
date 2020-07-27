# nihil

**Temporal multidimensional space analytics.**

What is **nihil**?

- Think about **nihil** as a calculator
- And this calculator can [analyze](https://en.wikipedia.org/wiki/Dimensional_analysis) [multidimensional](https://en.wikipedia.org/wiki/Dimension) [spaces](https://en.wikipedia.org/wiki/Space_(mathematics))
- Each dimension is a group of [series](https://en.wikipedia.org/wiki/Series_(mathematics))
- These series have a [direction property](https://en.wikipedia.org/wiki/Orientation_(geometry)) defined by each [data point](https://en.wikipedia.org/wiki/Unit_of_observation)
- Each data point is an observation of a value for a given point in time

Why *"nihil"*?

["nihil"](https://en.wiktionary.org/wiki/nihil) is the Latin word for "nothing", there's a private and personal reason for this name.

What is this good for?

- Prediction, projection
- Behavior analysis
- Calculation parallelization
- Statistical inference
- Observability

How precise is **nihil**?

Well, nihil data points can be represented in two ways (all values are `float64` and all timestamps are `int64`):

- **Absolute values** - here the precision will be the precision of the timestamp, if you're representing timestamps in nanoseconds, then the value will be absolut at that point in time.
- [Candlesticks](https://en.wikipedia.org/wiki/Candlestick_chart) - as candlesticks can represent 4 value bounds (maximum/high, open, close, minimum/low) the precision depends on the error factor of the [measuring instrument](https://en.wikipedia.org/wiki/Measuring_instrument), so it can be theoretically infinite (if the measuring instrument is infinitely precise). This is really useful where you can't get absolute values but can infer an interval where there's a probability to contain the value.