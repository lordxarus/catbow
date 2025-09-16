## TODO
- Create pull request for test/fix-gen-perf branch
- Approve + rebase pull request

- Create pull request for fix/rainbow-logic
- Rebase pull request

- change defaults to -freq .1 -spread 3
- add short option variants
## NICE TO TODO
- animated text 

- maybe an ability to have "profiles" of settings?
for example if you know that you'll only be seeing pretty short
messages you can use different settings to when you want to colorize
a wall of text
COLORS:

support for [various types](https://gist.github.com/kurahaupo/6ce0eaefe5e730841f03cb82b061daa2) of escape codes will be detected (TODO: how?) (--color-mode <truecolor | 256col | 16col> to override):
- 256 color palette ONLY
- 16 color palette ONLY 
  references:
  - [ANSI escape list](https://gist.github.com/JBlond/2fea43a3049b38287e5e9cefc87b2124)
  - [ANSI visualization](https://github.com/fidian/ansi)
CLI:
  - Cobra for argument parsing
  - automatically generate shell completions via build system 
  - package shell completions*
  - flags:
      -a, --animate
      -D, --duration (how long each segment animates for)

### lolcat Feature Parity Todo
- ability to interleave files and stdin: `catbow file0 - file1`* 
- when we are in a tty (our stdin is attached to a terminal) AND no files AND
  there is no stdin THEN we will allow the user to type into the terminal, buffer
  the text, and print rainbow text when they hit enter 

