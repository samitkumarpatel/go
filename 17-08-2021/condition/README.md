### Decision
Go programming language provides the following types of decision making statements
- if
```
if(boolean_expression) {
   /* statement(s) will execute if the boolean expression is true */
}
```
- if...else
```
if(boolean_expression) {
   /* statement(s) will execute if the boolean expression is true */
} else {
   /* statement(s) will execute if the boolean expression is false */
}
```
- nested if
```
if( boolean_expression 1) {
   /* Executes when the boolean expression 1 is true */
   if(boolean_expression 2) {
      /* Executes when the boolean expression 2 is true */
   }
}
```
- switch
```
switch(boolean-expression or integral type){
   case boolean-expression or integral type :
      statement(s);      
   case boolean-expression or integral type :
      statement(s); 
   
   /* you can have any number of case statements */
   default : /* Optional */
      statement(s);
}
```
- select
```
select {
   case communication clause  :
      statement(s);      
   case communication clause  :
      statement(s); 
   /* you can have any number of case statements */
   default : /* Optional */
      statement(s);
}
```