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

## Index

- [nihil](#nihil)
	- [Index](#index)
	- [Components](#components)
		- [Analytics](#analytics)
		- [LTP](#ltp)
	- [Data Model](#data-model)
		- [Metric Data Model (MDM)](#metric-data-model-mdm)
		- [Operational Data Model (ODM)](#operational-data-model-odm)
	- [Roadmap](#roadmap)
		- [Analysis](#analysis)


## Components

nihil has several components, each one responsible for doing just one thing. The overall focus is to provide high performance and scalability on UNIX-like operating systems.

### Analytics

A.k.a **nihil-analytics** is the calculator *per se*. The analytics takes input data models and performs computations to generate data enhanced by analysis routines and algorithms. Currently you can use [nihil-analytics](cmd/nihil-analytics) as a command line utility create your pipelines and process data.

### LTP

"LTP" comes from "long-term persistence" layer, it's responsible to manage the storage of the data model to be used in posterior analysis. Currently I'm planning how it'll be done, but certanly it'll be through an interface abstraction layer to make possible to implement any storage layer.

The default LTP will be a local disk storage driver, but I want to test something for intensive and low latency IO operations, something with [LSM](https://en.wikipedia.org/wiki/Log-structured_merge-tree).

## Data Model

The data model can be separated in two parts:

- Metric
- Operational

### Metric Data Model (MDM)

Everything starts with a **space**. The space has the global game rules and specifies zero or more axis. Each axis defines a dimension. A space having only one axis is a 1D space, 2 axis is a 2D space and so on. A 0 dimension space is a special case that we call singular space, it's a representation for undefined spaces.

Diving in axis we have series, series are temporal representation of values observations in time. An axis with 0 series is an empty axis, but it doesn't mean that the axis doesn't have a definition, it just don't have any oberved value yet. An axis can represent a context for a group of series, for example you can have an axis representing the **x** axis of your chart with two series:

- Velocity of a car by observer **A** (which can be a person beside the road)
- Velocity of a car by observer **B** (which can be another car)

This is useful for indicators too, if a series A represents an absolute oberved value we can have an indicator (like average) that can be a series B in the same axis.

An now we have the series which lead us to smallest data model unit: data points. Data points are single observations of a value or a value interval for a given point in time. We have two types of observations:

- Absolute values: as the name implies, they're absolute values for a point in time. An example of absolute values: the temperature of the day by seconds of the day.
- Candlesticks: candlesticks can represent intervals where there's a probability for the value to be contained within, and observation for the defined point in time if it's at the same time a point in time and a duration. Let's take our temperature example to help us understand candlesticks:

A daily observed temperature can have 4 values:

- The maximum (high) and minimum (low), which can be predicted by weather forecast.
  - Lets say our **high** value is 28°C and our **low** value is 15°C.
- The observed open temperature at the begining of the day and the close temperature at the end of day.
  - Lets say our **open** value is 26°C and our **close** value is 14°C.

See how this observation system can embed in a single point in time (the day) which is also an interval (the day duration) 4 values? This is the system that makes nihil an infinite precise temporal analytics, the measurement error rate is inversely proportional to the nihil's accuracy.

### Operational Data Model (ODM)

// TODO

## Roadmap

This project is currently in intense development and it's not mean to be used in production environment. This is a hobby project and I have plans to use it in some edge cases of financial market analysis and SRE observability subjects.

### Analysis

- [x] Implement *Financial Return*
- [x] Implement *Moving Average*
- [x] Implement *Standard Deviation*
- [x] Implement *Arithmetic Operations*
- [x] Implement *Bollinger Bands*
- [ ] Implement *EWMA*
- [ ] Implement CUDA engine