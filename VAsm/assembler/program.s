.data
a: .float 0.0
float_val_a_2: .float 3.140000
b: .float 0.0
float_val_b_3: .float 1.000000
c: .quad 0
d: .quad 0
resultado_suma_0: .float 0.0
msg_0: .ascii "4.140000\n"
resultado_suma_1: .quad 0
msg_1: .ascii "12\n"
saludar: .quad 0
String_val_saludar_0: .ascii "Hola "
hi: .quad 0
String_val_hi_1: .ascii "Mundo"
resultado_suma_2: .quad 0
msg_2: .ascii "Hola Mundo\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #3
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #11
mov w8, #64
svc #0

    mov x0, #0
    mov x8, #93
    svc #0


// Suma 1
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
add x12, x10, x11
adr x13, resultado_suma_0
str x12, [x13]


// Suma 2
adr x10, c
ldr x10, [x10]
adr x11, d
ldr x11, [x11]
add x12, x10, x11
adr x13, resultado_suma_1
str x12, [x13]


// Suma 3
adr x10, saludar
ldr x10, [x10]
adr x11, hi
ldr x11, [x11]
add x12, x10, x11
adr x13, resultado_suma_2
str x12, [x13]