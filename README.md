# Ventoux Scripting Language

## Usage

### Run Program

```shell
$ ventoux run hello.vx
```

## Examples

Hello World

```ventoux
"Hello World!"
```

Variables

```ventoux
greeting = "Hello"
greeting2 = "This really works!"
greeting3 = greeting2
pi = 3.141592
"Hello Variables!\n################"
greeting
greeting2
greeting3
pi
```

Expressions:

```ventoux
firstFib = 3
secondFib = 5
thirdFib = firstFib + secondFib
thirdFib
20 - 1
2 ^ 8
```