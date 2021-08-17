### Operator
An operator is a symbol that tells the compiler to perform specific mathematical or logical manipulations. Go language is rich in built-in operators and provides the following types of operators −

- Arithmetic Operators
syntex
```
+, -, *, /, %, ++, --
```
example
```
Assume variable A holds 10 and variable B holds 20 then −
A + B gives 30
A - B gives -10
A * B gives 200
B / A gives 2
B % A gives 0
A++ gives 11
A-- gives 9
```

- Relational Operators
syntex
```
==, !=, >, <, >=, <=
```
example

```
Assume variable A holds 1 and variable B holds 0,
(A == B) is not true
(A != B) is true.
(A > B) is not true
(A < B) is true
(A >= B) is not true.
(A <= B) is true.
```
- Logical Operators
syntex
```
&&, ||, !
```

- Bitwise Operators

- Assignment Operators
syntex
```
=, +=, -=, *=, /=, %=, >>=, <<=, &=, ^=, |=
```
example
```
C = A + B will assign value of A + B into C
C += A is equivalent to C = C + A
C -= A is equivalent to C = C - A
C *= A is equivalent to C = C * A
C /= A is equivalent to C = C / A
C %= A is equivalent to C = C % A
C <<= 2 is same as C = C << 2
C >>= 2 is same as C = C >> 2
C &= 2 is same as C = C & 2
C ^= 2 is same as C = C ^ 2
C |= 2 is same as C = C | 2
```

- Miscellaneous Operators

syntex
```
&, *
```

### Operators Precedence in Go

Category|Operator|Associativity|
|-------|--------|-------------|
Postfix|() [] -> . ++ - -|Left to right|
Unary|	+ - ! ~ ++ - - (type)* & sizeof	|Right to left|
Multiplicative|	* / %	|Left to right|
|Additive|	+ -|	Left to right|
|Shift|	<< >>	|Left to right|
|Relational|	< <= > >=	|Left to right|
|Equality|	== !=	|Left to right|
