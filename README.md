# WireWorld - cellular automata

<p>A Wireworld cell can be in one of four different states, usually numbered 0–3 in software, modeled by colors in the examples here:</p>

<p>empty (black),</p>
<p>electron head (blue),</p>
<p>electron tail (red),</p>
<p>conductor (yellow).</p>

<p>As in all cellular automata, time proceeds in discrete steps called generations (sometimes "gens" or "ticks"). Cells behave as follows:</p>

<p>empty → empty,</p>
<p>electron head → electron tail,</p>
<p>electron tail → conductor,</p>
<p>conductor → electron head if exactly one or two of the neighbouring cells are electron heads, otherwise remains conductor</p>

-----------------------------------------------------------------------------------------------------------------------------

# Installation

- GoLang 1.22.4 required
  
Run:

 ``go run main.go`` - for executing

 ``go build main.go``- for creating .exe
