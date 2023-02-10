# Ventoux Scripting Language
Ventoux is a functional scripting language powered by Golang, heavily impressed by Haskell.

## Usage

### Run Program

Just run `hello.vx` and print all final results to stdout
```shell
$ ventoux run hello.vx
```

Run and save the context to the filesystem:

```shell
$ ventoux run --export-virtual-machine-state=vm.vxcontext hello.vx
```

## Examples

### Hello World

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

### Expressions:

```haskell
firstFib = 3
secondFib = 5
thirdFib = firstFib + secondFib
thirdFib
20 - 1
2 ^ 8
```

### Functions:

```haskell
say s = s;
greeting = "Hello FnVentoux!"
say (greeting)
```

> Hello FnVentoux!

```haskell
add a b =  a + b;
ten = add (7 3)
add(ten ten)
```
> 20