# Ventoux Scripting Language

## Usage

### Run Program

```shell
$ ventoux run hello.vx
```

## Examples

Hello World

```haskell
"Hello World!"
```

Variables

```haskell
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

```haskell
firstFib = 3
secondFib = 5
thirdFib = firstFib + secondFib
thirdFib
20 - 1
2 ^ 8
```

Functions:

```haskell
say s = s;
greeting = "Hello FnVentoux!"
say (greeting)
```