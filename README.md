WireWorld - cellular automata

A Wireworld cell can be in one of four different states, usually numbered 0–3 in software, modeled by colors in the examples here:

empty (black),
electron head (blue),
electron tail (red),
conductor (yellow).
As in all cellular automata, time proceeds in discrete steps called generations (sometimes "gens" or "ticks"). Cells behave as follows:

empty → empty,
electron head → electron tail,
electron tail → conductor,
conductor → electron head if exactly one or two of the neighbouring cells are electron heads, otherwise remains conductor

-----------------------------------------------------------------------------------------------------------------------------

Installation

- GoLang 1.22.4 required
  
Run:

go run main.go - for executing

go build main.go - for creating .exe
