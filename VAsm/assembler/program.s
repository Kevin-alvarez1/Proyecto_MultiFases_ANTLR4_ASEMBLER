mov x0, #42

.data
typeof_result_0: .ascii "int\n"
tipoNumero: .ascii "int"
float_val_decimal_1: .float 3.141600
typeof_result_1: .ascii "float\n"
tipoDecimal: .ascii "float"

.text
mov x0, #1
adr x1, typeof_result_0
mov x2, #4
mov w8, #64
svc #0

ldr s0, =float_val_decimal_1
mov x0, #1
adr x1, typeof_result_1
mov x2, #6
mov w8, #64
svc #0