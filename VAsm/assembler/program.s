.data
a: .quad 0
b: .quad 0
f1: .float 0.0
float_val_f1_0: .float 2.500000
f2: .float 0.0
float_val_f2_1: .float 5.000000
text1: .quad 0
String_val_text1_0: .ascii "==== SUMA ===="
text2: .quad 0
String_val_text2_1: .ascii "==== RESTA ===="
text3: .quad 0
String_val_text3_2: .ascii "==== MULTIPLICACION ===="
text4: .quad 0
String_val_text4_3: .ascii "==== DIVISION ===="
text5: .quad 0
String_val_text5_4: .ascii "==== MODULO ===="
text6: .quad 0
String_val_text6_5: .ascii "==== MIXTOS ===="
msg_0: .ascii "==== SUMA ====\n"
resultado_suma_0: .quad 0
msg_1: .ascii "13\n"
resultado_suma_1: .float 0.0
msg_2: .ascii "12.500000\n"
resultado_suma_2: .float 0.0
msg_3: .ascii "7.500000\n"
msg_4: .ascii "==== RESTA ====\n"
resultado_resta_0: .quad 0
msg_5: .ascii "7\n"
resultado_resta_1: .float 0.0
resultado_resta_2: .float 0.0
msg_6: .ascii "2.000000\n"
msg_7: .ascii "==== MULTIPLICACION ====\n"
resultado_mul_0: .quad 0
msg_8: .ascii "30\n"
resultado_mul_1: .float 0.0
msg_9: .ascii "25.000000\n"
resultado_mul_2: .float 0.0
msg_10: .ascii "==== DIVISION ====\n"
resultado_div_0: .quad 0
msg_11: .ascii "3\n"
resultado_div_1: .float 0.0
msg_12: .ascii "4.000000\n"
resultado_div_2: .float 0.0
msg_13: .ascii "0.500000\n"
msg_14: .ascii "==== MODULO ====\n"
resultado_modulo_0: .quad 0
msg_15: .ascii "1\n"
msg_16: .ascii "==== MIXTOS ====\n"
resultado_suma_3: .float 0.0
resultado_resta_3: .float 0.0
msg_17: .ascii "-7.500000\n"
resultado_mul_3: .float 0.0
resultado_div_3: .float 0.0
msg_18: .ascii "1.666667\n"

.global _start
.text
_start:
mov x0, #1
adr x1, msg_0
mov x2, #15
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_1
mov x2, #3
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_4
mov x2, #16
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_5
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_3
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_6
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_7
mov x2, #25
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_8
mov x2, #3
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_9
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_10
mov x2, #19
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_11
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_12
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_13
mov x2, #9
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_14
mov x2, #17
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_15
mov x2, #2
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_16
mov x2, #17
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_2
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_17
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_9
mov x2, #10
mov w8, #64
svc #0

mov x0, #1
adr x1, msg_18
mov x2, #9
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
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fadd s2, s0, s1
adr x12, resultado_suma_1
str s2, [x12]


// Suma 3
adr x10, f1
ldr s0, [x10]
adr x11, f2
ldr s1, [x11]
fadd s2, s0, s1
adr x12, resultado_suma_2
str s2, [x12]


// Resta 1
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
sub x12, x10, x11
adr x13, resultado_resta_0
str x12, [x13]


// Resta 2
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fsub s2, s0, s1
adr x12, resultado_resta_1
str s2, [x12]


// Resta 3
adr x10, f2
ldr s0, [x10]
adr x11, b
ldr s1, [x11]
fsub s2, s0, s1
adr x12, resultado_resta_2
str s2, [x12]


// Multiplicación 1
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
mul x12, x10, x11
adr x13, resultado_mul_0
str x12, [x13]


// Multiplicación 2
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fmul s2, s0, s1
adr x12, resultado_mul_1
str s2, [x12]


// Multiplicación 3
adr x10, f1
ldr s0, [x10]
adr x11, f2
ldr s1, [x11]
fmul s2, s0, s1
adr x12, resultado_mul_2
str s2, [x12]


// División 1
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
sdiv x12, x10, x11
adr x13, resultado_div_0
str x12, [x13]


// División 2
adr x10, a
ldr s0, [x10]
adr x11, f1
ldr s1, [x11]
fdiv s2, s0, s1
adr x12, resultado_div_1
str s2, [x12]


// División 3
adr x10, f2
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fdiv s2, s0, s1
adr x12, resultado_div_2
str s2, [x12]


// Módulo 1
adr x10, a
ldr x10, [x10]
adr x11, b
ldr x11, [x11]
sdiv x12, x10, x11
mul x12, x12, x11
sub x12, x10, x12
adr x13, resultado_modulo_0
str x12, [x13]

// Suma 4
adr x10, f1
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fadd s2, s0, s1
adr x12, resultado_suma_3
str s2, [x12]


// Resta 4
adr x10, f1
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fsub s2, s0, s1
adr x12, resultado_resta_3
str s2, [x12]


// Multiplicación 4
adr x10, f1
ldr s0, [x10]
adr x11, a
ldr s1, [x11]
fmul s2, s0, s1
adr x12, resultado_mul_3
str s2, [x12]


// División 4
adr x10, f2
ldr s0, [x10]
adr x11, b
ldr s1, [x11]
fdiv s2, s0, s1
adr x12, resultado_div_3
str s2, [x12]